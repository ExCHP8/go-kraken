package kraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
)

// NonceKey is the key used for the nonce value. A nonce simply stands for a Number used ONCE.
// It’s a unique token used to add a layer of security to the request and also to validate the intent
// of a user initiated action.
const NonceKey = "nonce"

// Signer represents a Kraken API signature.
type Signer struct {
	Secret Secret
}

// NewSigner returns a new Kraken API signer.
// Authenticated requests should be signed with the "API-Sign" header,
// using a signature generated with your private key, nonce, encoded payload, and URI path according to:
// HMAC-SHA512 of (URI path + SHA256(nonce + POST data)) and base64 decoded secret API key.
func NewSigner(s string) *Signer {
	secret, _ := base64.StdEncoding.DecodeString(s)

	return &Signer{
		Secret: Secret(secret),
	}
}

// Sign signs the Kraken API request.
// Docs: https://www.kraken.com/help/api#general-usage for more information
func (s *Signer) Sign(v url.Values, path string) string {
	sha := sha256.New()
	sha.Write([]byte(v.Get(NonceKey) + v.Encode()))

	mac := hmac.New(sha512.New, s.Secret)
	mac.Write(append([]byte(path), sha.Sum(nil)...))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
