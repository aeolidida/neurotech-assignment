package storage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"neurotech-assignment/backend/internal/errs"
	"os"
)

type EncryptedFileStorage struct {
	filePath string
	key      []byte
}

func NewEncryptedFileStorage(filePath string, keyStr string) (*EncryptedFileStorage, error) {
	op := "EncryptedFileStorage.NewEncryptedFileStorage"
	key, err := hex.DecodeString(keyStr)
	if err != nil {
		return nil, errs.WrapError(op, "error decoding key", nil)
	}
	return &EncryptedFileStorage{
		filePath: filePath,
		key:      key,
	}, nil
}

func (fs *EncryptedFileStorage) encrypt(data []byte) ([]byte, error) {
	op := "EncryptedFileStorage.encrypt"
	block, err := aes.NewCipher(fs.key)
	if err != nil {
		return nil, errs.WrapError(op, "error creating cipher", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errs.WrapError(op, "error creating GCM", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errs.WrapError(op, "error generating nonce", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (fs *EncryptedFileStorage) decrypt(data []byte) ([]byte, error) {
	op := "EncryptedFileStorage.decrypt"
	if len(data) == 0 {
		return []byte{}, nil
	}

	block, err := aes.NewCipher(fs.key)
	if err != nil {
		return nil, errs.WrapError(op, "error creating cipher", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errs.WrapError(op, "error creating GCM", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errs.WrapError(op, "invalid ciphertext", err)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errs.WrapError(op, "error decrypting data", err)
	}

	return plaintext, nil
}

func (fs *EncryptedFileStorage) Save(data []byte) error {
	op := "EncryptedFileStorage.Save"
	encryptedData, err := fs.encrypt(data)
	if err != nil {
		return errs.WrapError(op, "error encrypting data", err)
	}

	return os.WriteFile(fs.filePath, encryptedData, 0644)
}

func (fs *EncryptedFileStorage) Load() ([]byte, error) {
	op := "EncryptedFileStorage.Load"
	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, errs.WrapError(op, "error opening file", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, errs.WrapError(op, "error getting file info", err)
	}

	encryptedData := make([]byte, fileInfo.Size())
	_, err = file.Read(encryptedData)
	if err != nil {
		return nil, errs.WrapError(op, "error reading file", err)
	}

	if len(encryptedData) == 0 {
		return []byte{}, nil
	}

	data, err := fs.decrypt(encryptedData)
	if err != nil {
		return nil, errs.WrapError(op, "error decrypting data", err)
	}

	return data, nil
}
