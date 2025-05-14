package license

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"os"
	"time"
)

type License struct {
	Fingerprint string    `json:"fingerprint"`
	ValidUntil  time.Time `json:"valid_until"`
	Signature   []byte    `json:"signature,omitempty"`
}

// 签名
func SignLicense(lic *License, privKey *rsa.PrivateKey) error {
	lic.Signature = nil
	data, _ := json.Marshal(lic)
	hash := sha256.Sum256(data)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
	if err != nil {
		return err
	}
	lic.Signature = sig
	return nil
}

// 验签
func VerifyLicense(lic *License, pubKey *rsa.PublicKey) error {
	sig := lic.Signature
	lic.Signature = nil
	data, _ := json.Marshal(lic)
	hash := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], sig)
}

// 兼容加载 PKCS#1 和 PKCS#8 的 RSA 私钥
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	// 先尝试作为 PKCS#1 解析
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}

	// 再尝试作为 PKCS#8 解析
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("parsed key is not RSA private key")
	}
	return rsaKey, nil
}

// 加载公钥（X.509 格式）
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}
	return key, nil
}
