package lib

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func generateKeys(dirName, accountPassword string) (string, string, error) {
	nodeKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("couldn't generate key: %v", err)
	}

	nodePubKey := nodeKey.PublicKey
	address := crypto.PubkeyToAddress(nodeKey.PublicKey)

	accountKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("couldn't generate key: %v", err)
	}

	accountAddress := crypto.PubkeyToAddress(accountKey.PublicKey)

	if err := os.WriteFile(fmt.Sprintf("%s/nodekey", dirName), []byte(hexutil.Encode(crypto.FromECDSA(nodeKey))[2:]), 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save node key: %v", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/nodekey.pub", dirName), []byte(hexutil.Encode(crypto.FromECDSAPub(&nodePubKey))[4:]), 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save node public key: %v", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/address", dirName), []byte(address.Hex()[2:]), 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save node address: %v", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/accountPrivateKey", dirName), []byte(hexutil.Encode(crypto.FromECDSA(accountKey))), 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save account private key: %v", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/accountAddress", dirName), []byte(accountAddress.Hex()), 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save account address: %v", err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/accountPassword", dirName), []byte(accountPassword), 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save account password: %v", err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return "", "", fmt.Errorf("couldn't create random uuid: %v", err)
	}

	key := &keystore.Key{
		Id:         id,
		Address:    accountAddress,
		PrivateKey: accountKey,
	}

	keyJSON, err := keystore.EncryptKey(key, accountPassword, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return "", "", err
	}

	if err := os.WriteFile(fmt.Sprintf("%s/accountKeystore", dirName), keyJSON, 0644); err != nil {
		return "", "", fmt.Errorf("couldn't save account keystore: %v", err)
	}

	return hexutil.Encode(crypto.FromECDSAPub(&nodePubKey))[4:], accountAddress.Hex(), nil
}
