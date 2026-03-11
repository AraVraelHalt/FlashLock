package flashdrive

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
)

func (f *FlashDrive) DecryptContainer(password []byte) error {
	containerFile := f.Path + "/locked.container"
	data, err := os.ReadFile(containerFile)
	
  if err != nil {
		return err
	}

	if len(data) < 16 {
		return errors.New("container file too short")
	}

	salt := data[:16]
	ciphertext := data[16:]
	key := deriveKey(password, salt)
	block, err := aes.NewCipher(key)

	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	
  if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	
  if len(ciphertext) < nonceSize {
		return errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	
  if err != nil {
		return err
	}

	tempDir := f.Path + "/.temp_decrypt"
	
  if err := os.MkdirAll(tempDir, 0755); err != nil {
		return err
	}

	buf := bytes.NewReader(plaintext)
	
  for {
		var nameLen uint16
		
    if err := binary.Read(buf, binary.LittleEndian, &nameLen); err != nil {
			break // EOF
		}

		nameBytes := make([]byte, nameLen)
		
    if _, err := buf.Read(nameBytes); err != nil {
			return err
		}
  
		relPath := string(nameBytes)

		var contentLen uint64
		
    if err := binary.Read(buf, binary.LittleEndian, &contentLen); err != nil {
			return err
		}
		
    content := make([]byte, contentLen)
		
    if _, err := buf.Read(content); err != nil {
			return err
		}

		outPath := filepath.Join(tempDir, relPath)
		
    if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}
		
    if err := os.WriteFile(outPath, content, 0644); err != nil {
			return err
		}
	}

	if err := os.Remove(containerFile); err != nil {
		return err
	}

  decryptedFolder := f.Path + "/unlocked.container"
	
  if err := os.Rename(tempDir, decryptedFolder); err != nil {
		return err
	}

	return nil
}
