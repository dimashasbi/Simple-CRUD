package setvalue

import (
	"M-GateDBConfig/dbsetup"
	"M-GateDBConfig/model"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"io"
)

// SetValue to input value to DB
func SetValue(db *gorm.DB) {
	backsetting := &model.BackSettingsConfig{
		URLJavaMPAY: "localhost:27000",
		URLJavaSVA:  "localhost:27001",
	}
	res1B, err := json.Marshal(backsetting)
	if err != nil {
		panic(err)
	}
	key := []byte("passwordnyaharus") // should 16

	encryptedval, err := encrypt(key, res1B)
	if err != nil {
		panic(err)
	}
	_, ans := dbsetup.DBsetSettings(db, base64.StdEncoding.EncodeToString(encryptedval))
	decodestring, err := base64.StdEncoding.DecodeString(ans)
	if err != nil {
		panic(err)
	}
	decr, err := decrypt(key, decodestring)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON VALUE =  %v\n", string(decr))
}
func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
