package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	Memory     uint32
	Time       uint32
	Threads    uint8
	KeyLength  uint32
	SaltLength uint32
}

var DefaultParams = &Argon2Params{
	Memory:     64 * 1024, // 64 MB
	Time:       3,
	Threads:    2,
	KeyLength:  32,
	SaltLength: 16,
}

func HashPassword(password string, p *Argon2Params) (string, error) {
	salt := make([]byte, p.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		p.Time,
		p.Memory,
		p.Threads,
		p.KeyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		p.Memory,
		p.Time,
		p.Threads,
		b64Salt,
		b64Hash,
	)

	return encoded, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	var memory uint32
	var time uint32
	var threads uint8
	var salt, hash []byte

	_, err := fmt.Sscanf(
		encodedHash,
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		&memory,
		&time,
		&threads,
		&salt,
		&hash,
	)
	if err != nil {
		return false, err
	}

	saltBytes, err := base64.RawStdEncoding.DecodeString(string(salt))
	if err != nil {
		return false, err
	}

	hashBytes, err := base64.RawStdEncoding.DecodeString(string(hash))
	if err != nil {
		return false, err
	}

	computedHash := argon2.IDKey(
		[]byte(password),
		saltBytes,
		time,
		memory,
		threads,
		uint32(len(hashBytes)),
	)

	return subtle.ConstantTimeCompare(hashBytes, computedHash) == 1, nil
}
