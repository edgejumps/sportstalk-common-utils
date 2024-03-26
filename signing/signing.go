package signing

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	SportstalkPrivateKey = "SPORTSTALK PRIVATE KEY"
	SportstalkPublicKey  = "SPORTSTALK PUBLIC KEY"
)

var (
	ErrInvalidPublicKey = fmt.Errorf("invalid public key")
)

func GenerateAndSaveRSAPair(dir string, filename string) (*rsa.PrivateKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, err
	}

	err = SavePrivateKey(private, dir, filename)

	if err != nil {
		return nil, err
	}

	pubFileName := fmt.Sprintf("%s_pub", filename)

	err = SavePublicKey(&private.PublicKey, dir, pubFileName)

	if err != nil {
		return nil, err
	}

	return private, nil
}

func GenerateRSAPair() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func SavePrivateKey(private *rsa.PrivateKey, dir string, filename string) error {
	err := EnsureDirectoryExists(dir)

	if err != nil {
		return err
	}

	path := filepath.Join(dir, filename)

	fmt.Printf("Saving private key to %s\n", path)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes := MarshalPrivateKey(private)

	block := ConvertToPemBlock(bytes, SportstalkPrivateKey)

	return pem.Encode(file, block)
}

func SavePublicKey(public *rsa.PublicKey, dir string, filename string) error {
	err := EnsureDirectoryExists(dir)

	if err != nil {
		return err
	}

	path := filepath.Join(dir, filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := MarshalPublicKey(public)

	if err != nil {
		return err
	}

	block := ConvertToPemBlock(bytes, SportstalkPublicKey)

	return pem.Encode(file, block)
}

func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pemData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return RestorePrivateKeyFromPemDate(pemData)
}

func LoadPublicKey(filename string) (*rsa.PublicKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pemData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return RestorePublicKeyFromPemData(pemData)
}

func RestorePrivateKeyFromPemDate(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func RestorePublicKeyFromPemData(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey.(*rsa.PublicKey), nil
}

func MarshalPublicKey(public *rsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(public)
}

func MarshalPrivateKey(private *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(private)
}

func ConvertToPemBlock(data []byte, blockType string) *pem.Block {
	return &pem.Block{
		Type:  blockType,
		Bytes: data,
	}
}

func EnsureDirectoryExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0700)
	}

	return nil
}

func StringifyPublicKey(public *rsa.PublicKey) (string, error) {
	bytes, err := x509.MarshalPKIXPublicKey(public)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func ParsePublicKey(data string) (*rsa.PublicKey, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return nil, err
	}

	public, err := x509.ParsePKIXPublicKey(bytes)

	if err != nil {
		return nil, err
	}

	return public.(*rsa.PublicKey), nil
}

func StringifyPrivateKey(private *rsa.PrivateKey) string {
	bytes := x509.MarshalPKCS1PrivateKey(private)

	return base64.StdEncoding.EncodeToString(bytes)
}

func ParsePrivateKey(data string) (*rsa.PrivateKey, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(bytes)
}

func HashPublicKey(public interface{}) (string, error) {
	var encoded string

	switch t := public.(type) {
	case *rsa.PublicKey:
		str, err := StringifyPublicKey(t)

		if err != nil {
			return "", err
		}
		encoded = str
	case string:
		_, err := ParsePublicKey(t)

		if err != nil {
			return "", err
		}

		encoded = t
	default:
		return "", ErrInvalidPublicKey
	}

	hasher := sha256.New()

	_, err := hasher.Write([]byte(encoded))

	if err != nil {
		return "", err
	}

	encoded = base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	return encoded, nil
}

func SignMessage(private *rsa.PrivateKey, message interface{}) (string, error) {
	var data []byte

	switch t := message.(type) {
	case []byte:
		data = truncate(t)
	case string:
		data = truncate([]byte(t))
	default:
		bytes, err := json.Marshal(t)

		if err != nil {
			return "", err
		}
		data = truncate(bytes)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, private, 0, data)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func VerifySignature(public interface{}, message interface{}, signature string) error {

	var key *rsa.PublicKey

	switch t := public.(type) {
	case *rsa.PublicKey:
		key = t
	case string:
		parsedKey, err := ParsePublicKey(t)

		if err != nil {
			return err
		}

		key = parsedKey
	default:
		return ErrInvalidPublicKey
	}

	var data []byte

	switch t := message.(type) {
	case []byte:
		data = truncate(t)
	case string:
		data = truncate([]byte(t))
	default:
		bytes, err := json.Marshal(t)

		if err != nil {
			return err
		}
		data = truncate(bytes)
	}

	bytes, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(key, 0, data, bytes)
}

func truncate(data []byte) []byte {

	// the max data size that can be signed by RSA is 256-11=245 bytes for 2048-bit keys
	if len(data) > 245 {
		return data[:245]
	}
	return data
}
