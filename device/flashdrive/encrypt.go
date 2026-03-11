package flashdrive

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io/fs"
	"os"
	"path/filepath"
)

func (f *FlashDrive) EncryptContainer(password []byte) error {
	decryptedFolder := f.Path + "/unlocked.container"

	buf := new(bytes.Buffer)

	err := filepath.WalkDir(decryptedFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(decryptedFolder, path)
		
    if err != nil {
			return err
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := binary.Write(buf, binary.LittleEndian, uint16(len(relPath))); err != nil {
			return err
		}

		buf.Write([]byte(relPath))

		if err := binary.Write(buf, binary.LittleEndian, uint64(len(content))); err != nil {
			return err
		}

		buf.Write(content)

		return nil
	})

	if err != nil {
		return err
	}

	data := buf.Bytes()
	salt := make([]byte, 16)
	
  if _, err := rand.Read(salt); err != nil {
		return err
	}

	key := deriveKey(password, salt)
	block, err := aes.NewCipher(key)
	
  if err != nil {
		return err
	}
	
  gcm, err := cipher.NewGCM(block)
	
  if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	
  if _, err := rand.Read(nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	containerFile := f.Path + "/locked.container"

	if err := os.RemoveAll(decryptedFolder); err != nil {
		return err
	}

	outFile, err := os.Create(containerFile)
	
  if err != nil {
		return err
	}
	
  defer outFile.Close()

	if _, err := outFile.Write(salt); err != nil {
		return err
	}
	
  if _, err := outFile.Write(ciphertext); err != nil {
		return err
	}

	return nil
}
