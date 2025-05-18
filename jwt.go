package mjwtauth

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSigner handles signing and verifying tokens.
type JWTSigner interface {
	Sign(claims map[string]interface{}) (string, error)
	Verify(tokenString string) (jwt.MapClaims, error)
}

// RS256Signer implements JWTSigner using RSA keys.
type RS256Signer struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRS256SignerFromPEM(privateKeyPEM, publicKeyPEM []byte) (*RS256Signer, error) {
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, err
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return nil, err
	}
	return &RS256Signer{
		privateKey: privKey,
		publicKey:  pubKey,
	}, nil
}

func (s *RS256Signer) Sign(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	for k, v := range claims {
		token.Claims.(jwt.MapClaims)[k] = v
	}
	return token.SignedString(s.privateKey)
}

func (s *RS256Signer) Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

type HS256Signer struct {
	secret []byte
}

func NewHS256Signer(secret []byte) *HS256Signer {
	return &HS256Signer{secret: secret}
}

func (s *HS256Signer) Sign(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	for k, v := range claims {
		token.Claims.(jwt.MapClaims)[k] = v
	}
	return token.SignedString(s.secret)
}

func (s *HS256Signer) Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
