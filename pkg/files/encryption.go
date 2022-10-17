package files

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
	"golang.org/x/crypto/pbkdf2"
)

type AESKey struct {
	Cipher cipher.Block
	IV     []byte
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func GenerateEncryptionKey() string {
	random := make([]byte, config.CurrentConfig.EncryptionKeySize)
	rand.Read(random)
	return base64.RawURLEncoding.EncodeToString(random)
}

func GetAESKeyFromEncryptionKey(encryption_key string) (*AESKey, error) {
	salt, err := hex.DecodeString(config.CurrentConfig.Salt)
	if err != nil {
		fmt.Printf("%s\n%v\n", config.CurrentConfig.Salt, err)
		return nil, err
	}
	pdbkdf := pbkdf2.Key([]byte(encryption_key), salt, 4096, 48, sha512.New)
	cipher, err := aes.NewCipher(pdbkdf[:32])
	if err != nil {
		return nil, err
	}
	return &AESKey{Cipher: cipher, IV: pdbkdf[32:48]}, nil

}

func Encrypt(key *AESKey, message []byte) []byte {
	paddedMessage := PKCS5Padding(message, aes.BlockSize, len(message))
	ciphertext := make([]byte, len(paddedMessage))
	mode := cipher.NewCBCEncrypter(key.Cipher, key.IV)
	mode.CryptBlocks(ciphertext, paddedMessage)
	return ciphertext
}

func Decrypt(key *AESKey, ciphertext []byte) []byte {
	message := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(key.Cipher, key.IV)
	mode.CryptBlocks(message, ciphertext)
	return message
}
