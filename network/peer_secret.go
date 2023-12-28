package p2p

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"

	"github.com/shaojianqing/smilebc/crypto/ecies"
)

const (
	AESKeyLength = 32
)

type AESEncryptKey []byte

type Secret struct {
	remoteNodeID    NodeID
	shareEncryptKey AESEncryptKey
	localPrivateKey *ecies.PrivateKey
	localPublicKey  *ecies.PublicKey
	remotePublicKey *ecies.PublicKey

	shareCipherBlock cipher.Block
}

func NewSecret(remoteNodeID NodeID, localPrivateKey *ecdsa.PrivateKey, remotePublicKey *ecdsa.PublicKey) (*Secret, error) {
	privateKey := ecies.ImportECDSA(localPrivateKey)
	publicKey := ecies.ImportECDSAPublic(remotePublicKey)
	shareEncryptKey, err := ecies.GenerateAESKey(AESKeyLength)
	if err != nil {
		return nil, err
	}
	shareCipherBlock, err := aes.NewCipher(shareEncryptKey)
	if err != nil {
		return nil, err
	}
	return &Secret{
		remoteNodeID:     remoteNodeID,
		localPrivateKey:  privateKey,
		localPublicKey:   &privateKey.PublicKey,
		remotePublicKey:  publicKey,
		shareEncryptKey:  shareEncryptKey,
		shareCipherBlock: shareCipherBlock,
	}, nil
}

func NewSecretWithAESKey(remoteNodeID NodeID, localPrivateKey *ecdsa.PrivateKey, remotePublicKey *ecdsa.PublicKey, aesEncryptKey AESEncryptKey) (*Secret, error) {
	privateKey := ecies.ImportECDSA(localPrivateKey)
	publicKey := ecies.ImportECDSAPublic(remotePublicKey)
	shareCipherBlock, err := aes.NewCipher(aesEncryptKey)
	if err != nil {
		return nil, err
	}
	return &Secret{
		remoteNodeID:     remoteNodeID,
		localPrivateKey:  privateKey,
		localPublicKey:   &privateKey.PublicKey,
		remotePublicKey:  publicKey,
		shareEncryptKey:  aesEncryptKey,
		shareCipherBlock: shareCipherBlock,
	}, nil
}

func (s *Secret) Encrypt(plainData []byte) []byte {

	blockSize := s.shareCipherBlock.BlockSize()
	plainData = s.pkcs7Padding(plainData, blockSize)

	cipherData := make([]byte, len(plainData))

	encrypt := cipher.NewCBCEncrypter(s.shareCipherBlock, s.shareEncryptKey[:blockSize])
	encrypt.CryptBlocks(cipherData, plainData)
	return cipherData
}

func (s *Secret) Decrypt(cipherData []byte) []byte {
	blockSize := s.shareCipherBlock.BlockSize()

	plainData := make([]byte, len(cipherData))
	mode := cipher.NewCBCDecrypter(s.shareCipherBlock, s.shareEncryptKey[:blockSize])
	mode.CryptBlocks(plainData, cipherData)

	plainData = s.pkcs7UnPadding(plainData)

	return plainData
}

func (s *Secret) pkcs7Padding(plainText []byte, blockSize int) []byte {
	paddingLength := blockSize - len(plainText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(plainText, paddingText...)
}

func (s *Secret) pkcs7UnPadding(cipherText []byte) []byte {
	length := len(cipherText)
	unPaddingLength := int(cipherText[length-1])
	return cipherText[:(length - unPaddingLength)]
}
