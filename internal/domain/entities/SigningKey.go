package entities

import "time"

type KeyStatus string

const (
	KeyPending KeyStatus = "PENDING"
	KeyActive  KeyStatus = "ACTIVE"
	KeyRetired KeyStatus = "RETIRED"
)

type SigningKey struct {
	ID        string
	Kid       string
	Algorithm string
	Status    KeyStatus

	PublicJWK           []byte
	EncryptedPrivateKey []byte
	EncryptedDEK        []byte
	Nonce               []byte
	KMSKeyID            string

	CreatedAt   time.Time
	PublishedAt *time.Time
	ActivatedAt *time.Time
	RetiredAt   *time.Time
	ExpiresAt   *time.Time
}
