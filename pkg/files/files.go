package files

import (
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/apierrors"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type FileInfos struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	EncryptionKey string `json:"encryption_key"`
}

func GenerateFileID() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetIDPath(id string) string {
	return path.Join(config.CurrentConfig.FileServePath, id)
}

func GetFilePath(id string, name string) string {
	return path.Join(GetIDPath(id), name)
}

func ReadFile(id string, name string, encryption_key string) ([]byte, error) {
	if strings.Contains(id, "/") || strings.Contains(name, "/") || strings.Contains(encryption_key, "/") {
		return nil, apierrors.FileDoesNotExist
	}
	filepath := GetFilePath(id, name)
	content, err := os.ReadFile(filepath)
	if os.IsNotExist(err) {
		return nil, apierrors.FileDoesNotExist
	} else if err != nil {
		return nil, err
	}
	key, err := GetAESKeyFromEncryptionKey(encryption_key)
	if err != nil {
		return nil, err
	}
	return Decrypt(key, content), nil
}

func WriteFile(name string, content []byte) (*FileInfos, error) {
	fileInfos := FileInfos{Name: name, ID: GenerateFileID(), EncryptionKey: GenerateEncryptionKey()}
	idPath := GetIDPath(fileInfos.ID)
	filepath := GetFilePath(fileInfos.ID, name)
	key, err := GetAESKeyFromEncryptionKey(fileInfos.EncryptionKey)
	if err != nil {
		return nil, err
	}
	cipher := Encrypt(key, content)
	os.Mkdir(idPath, os.ModePerm)
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}
	file.Write(cipher)
	file.Close()
	return &fileInfos, nil
}
