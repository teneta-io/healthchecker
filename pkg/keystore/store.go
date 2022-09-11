package keystore

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/teneta-io/healthchecker/pkg/logger"
)

const (
	PubKeyFileName     = "public.txt"
	PrivateKeyFileName = "private.txt"
)

type KeyStore interface {
	GetPublicKey() string
	SubscribeTransaction() bool // TODO in future add check keys
}

type Store struct {
	PublicKey  string
	PrivateKey string
}

func NewKeyStore() (*Store, error) {
	store := &Store{}

	err := store.LoadKeys()
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return nil, err
	}

	if strings.Contains(err.Error(), "no such file or directory") {
		err = store.GenerateKeys()
		if err != nil {
			return nil, err
		}
	}

	return store, nil
}

func (s *Store) GenerateKeys() error {
	// Create an account
	key, err := crypto.GenerateKey()
	if err != nil {
		return err
	}

	logger.Infof("crypto\n")
	// Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	logger.Infof(address)

	// Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	logger.Infof(privateKey)

	err = SaveFile(PubKeyFileName, address)
	if err != nil {
		logger.Errorf("Error save file: %s Error: %v", PubKeyFileName, err)
		return err
	}

	err = SaveFile(PrivateKeyFileName, privateKey)
	if err != nil {
		logger.Errorf("Error save file: %s Error: %v", PrivateKeyFileName, err)
		return err
	}

	s.PrivateKey = privateKey
	s.PublicKey = address

	return nil
}

func SaveFile(name, text string) error {
	f, err := os.Create(name)
	if err != nil {
		logger.Errorf("%v", err)
		return err
	}

	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		logger.Errorf("%v", err)
		return err
	}

	return nil
}

func (s *Store) LoadKeys() error {
	content, err := ioutil.ReadFile(PubKeyFileName)
	if err != nil {
		logger.Errorf("Error read file %s: %v", PubKeyFileName, err)
		return err
	}
	s.PublicKey = string(content)

	content, err = ioutil.ReadFile(PrivateKeyFileName)
	if err != nil {
		logger.Errorf("Error read file %s: %v", PrivateKeyFileName, err)
		return err
	}
	s.PrivateKey = string(content)

	return nil
}

func (s *Store) GetPublicKey() string {
	return s.PublicKey
}

func (s *Store) SubscribeTransaction() bool {
	return true
}
