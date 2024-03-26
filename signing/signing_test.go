package signing

import "testing"

func TestGenerateRSAPair(t *testing.T) {
	_, err := GenerateAndSaveRSAPair("./", "test")
	if err != nil {
		t.Errorf("Error generating RSA pair: %s", err)
	}
}

func TestLoadPrivateKey(t *testing.T) {
	_, err := LoadPrivateKey("test")

	if err != nil {
		t.Errorf("Error loading private key: %s", err)
	}
}

func TestLoadPublicKey(t *testing.T) {
	_, err := LoadPublicKey("test_pub")

	if err != nil {
		t.Errorf("Error loading public key: %s", err)
	}
}

func TestSigning(t *testing.T) {
	key, err := LoadPrivateKey("test")

	if err != nil {
		t.Errorf("Error loading private key: %s", err)
	}

	message := []byte("test message")

	signature, err := SignMessage(key, message)

	if err != nil {
		t.Errorf("Error signing message: %s", err)
	}

	publicKey, err := LoadPublicKey("test_pub")

	if err != nil {
		t.Errorf("Error loading public key: %s", err)
	}

	err = VerifySignature(publicKey, message, signature)

	if err != nil {
		t.Errorf("Error verifying signature: %s", err)
	}
}

func TestSigningBigMessage(t *testing.T) {
	key, err := LoadPrivateKey("test")

	if err != nil {
		t.Errorf("Error loading private key: %s", err)
	}

	msg := make([]byte, 0, 1024)

	for i := 0; i < 1024; i++ {
		msg = append(msg, 'a')
	}

	signature, err := SignMessage(key, msg)

	if err != nil {
		t.Errorf("Error signing message: %s", err)
	}

	publicKey, err := LoadPublicKey("test_pub")

	if err != nil {
		t.Errorf("Error loading public key: %s", err)
	}

	err = VerifySignature(publicKey, msg, signature)

	if err != nil {
		t.Errorf("Error verifying signature: %s", err)
	}

}
