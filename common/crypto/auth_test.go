package crypto_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"testing"

	"v2ray.com/core/common/buf"
	. "v2ray.com/core/common/crypto"
	"v2ray.com/core/testing/assert"
)

func TestAuthenticationReaderWriter(t *testing.T) {
	assert := assert.On(t)

	key := make([]byte, 16)
	rand.Read(key)
	block, err := aes.NewCipher(key)
	assert.Error(err).IsNil()

	aead, err := cipher.NewGCM(block)
	assert.Error(err).IsNil()

	payload := make([]byte, 8*1024)
	rand.Read(payload)

	cache := buf.NewLocal(16 * 1024)
	iv := make([]byte, 12)
	rand.Read(iv)

	writer := NewAuthenticationWriter(&AEADAuthenticator{
		AEAD: aead,
		NonceGenerator: &StaticBytesGenerator{
			Content: iv,
		},
		AdditionalDataGenerator: &NoOpBytesGenerator{},
	}, cache)

	nBytes, err := writer.Write(payload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload))
	assert.Int(cache.Len()).GreaterThan(0)
	_, err = writer.Write([]byte{})
	assert.Error(err).IsNil()

	reader := NewAuthenticationReader(&AEADAuthenticator{
		AEAD: aead,
		NonceGenerator: &StaticBytesGenerator{
			Content: iv,
		},
		AdditionalDataGenerator: &NoOpBytesGenerator{},
	}, cache, false)

	actualPayload := make([]byte, 16*1024)
	nBytes, err = reader.Read(actualPayload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload))
	assert.Bytes(actualPayload[:nBytes]).Equals(payload)

	_, err = reader.Read(actualPayload)
	assert.Error(err).Equals(io.EOF)
}

func TestAuthenticationReaderWriterPartial(t *testing.T) {
	assert := assert.On(t)

	key := make([]byte, 16)
	rand.Read(key)
	block, err := aes.NewCipher(key)
	assert.Error(err).IsNil()

	aead, err := cipher.NewGCM(block)
	assert.Error(err).IsNil()

	payload := make([]byte, 8*1024)
	rand.Read(payload)

	cache := buf.NewLocal(16 * 1024)
	iv := make([]byte, 12)
	rand.Read(iv)

	writer := NewAuthenticationWriter(&AEADAuthenticator{
		AEAD: aead,
		NonceGenerator: &StaticBytesGenerator{
			Content: iv,
		},
		AdditionalDataGenerator: &NoOpBytesGenerator{},
	}, cache)

	nBytes, err := writer.Write(payload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload))
	assert.Int(cache.Len()).GreaterThan(0)
	_, err = writer.Write([]byte{})
	assert.Error(err).IsNil()

	reader := NewAuthenticationReader(&AEADAuthenticator{
		AEAD: aead,
		NonceGenerator: &StaticBytesGenerator{
			Content: iv,
		},
		AdditionalDataGenerator: &NoOpBytesGenerator{},
	}, cache, false)

	actualPayload := make([]byte, 7*1024)
	nBytes, err = reader.Read(actualPayload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(actualPayload))
	assert.Bytes(actualPayload[:nBytes]).Equals(payload[:nBytes])

	nBytes, err = reader.Read(actualPayload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload) - len(actualPayload))
	assert.Bytes(actualPayload[:nBytes]).Equals(payload[7*1024:])

	_, err = reader.Read(actualPayload)
	assert.Error(err).Equals(io.EOF)
}

func TestAuthenticationReaderWriterAggressive(t *testing.T) {
	assert := assert.On(t)

	key := make([]byte, 16)
	rand.Read(key)
	block, err := aes.NewCipher(key)
	assert.Error(err).IsNil()

	aead, err := cipher.NewGCM(block)
	assert.Error(err).IsNil()

	payload := make([]byte, 7*1024)
	rand.Read(payload)

	cache := buf.NewLocal(16 * 1024)
	iv := make([]byte, 12)
	rand.Read(iv)

	writer := NewAuthenticationWriter(&AEADAuthenticator{
		AEAD: aead,
		NonceGenerator: &StaticBytesGenerator{
			Content: iv,
		},
		AdditionalDataGenerator: &NoOpBytesGenerator{},
	}, cache)

	nBytes, err := writer.Write(payload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload))
	assert.Int(cache.Len()).GreaterThan(0)
	_, err = writer.Write(payload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload))
	_, err = writer.Write([]byte{})
	assert.Error(err).IsNil()

	reader := NewAuthenticationReader(&AEADAuthenticator{
		AEAD: aead,
		NonceGenerator: &StaticBytesGenerator{
			Content: iv,
		},
		AdditionalDataGenerator: &NoOpBytesGenerator{},
	}, cache, true)

	actualPayload := make([]byte, 16*1024)
	nBytes, err = reader.Read(actualPayload)
	assert.Error(err).IsNil()
	assert.Int(nBytes).Equals(len(payload) * 2)

	_, err = reader.Read(actualPayload)
	assert.Error(err).Equals(io.EOF)
}
