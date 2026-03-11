# 🔒 FlashLock 

A secure, fast, and memory-safe tool to encrypt your flash drives.

[![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go&logoColor=white)](https://golang.org/) 
![Build](https://github.com/AraVraelHalt/FlashLock/actions/workflows/go.yml/badge.svg)
---

## 🚀 Features

- AES‑256 encryption for data security  
- Argon2id for memory-hard password hashing  
- Secure memory handling and constant-time comparisons  
- Anti-brute force protections and tamper detection (Coming Soon)
- Terminal CLI 
- Cross-platform support  

---

## 🛠 Tech Stack

**Programming Language:** Go (fast, simple, memory-safe)  

**Cryptography:**  
- AES‑256 (data encryption)  
- Argon2 (password → key derivation) 

**Disk / File System Access:** 
- OS system APIs  

**Authentication Protection:**  
- Argon2id with salt 
- Memory-hard password hashing  

**Security Features:**  
- Secure memory handling  
- Constant-time comparisons  
- Anti-brute force protections (Coming soon)
- Tamper detection (Coming soon)

**Build & Packaging:**  
- Go compiler → single binary (coming soon)
- Git + GitHub 

## 📝 Usage
Parent folder should be at root of the flash drive and named `unlocked.container`, after the encryption flow is as follows:


**Scan for devices:**

```bash
scan
```

**Select flash drive from the list of scanned devices:**

```bash
select <index>
```

**Encrypt or Decrypt drive:**

```bash
encrypt <psswd>
decrypt <psswd>
```

**Eject drive before removing:**

```bash
eject
```


