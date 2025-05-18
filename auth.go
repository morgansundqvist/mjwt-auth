package mjwtauth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimsCustomizer func(user AuthUser, claims map[string]interface{}) error

type AuthService struct {
	repo             UserRepository
	signer           JWTSigner
	claimsCustomizer ClaimsCustomizer
}

func NewAuthService(repo UserRepository, signer JWTSigner, customizer ClaimsCustomizer) *AuthService {
	return &AuthService{repo: repo, signer: signer, claimsCustomizer: customizer}
}

func (a *AuthService) Signup(username, password string) (AuthUser, error) {
	hash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	return a.repo.CreateUser(username, hash)
}

func (a *AuthService) Login(username, password string) (string, error) {
	user, err := a.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if err := CheckPasswordHash(user.GetPasswordHash(), password); err != nil {
		return "", errors.New("invalid password")
	}
	claims := map[string]interface{}{
		"sub":      user.GetID(),
		"username": user.GetUsername(),
	}
	if a.claimsCustomizer != nil {
		err = a.claimsCustomizer(user, claims)
		if err != nil {
			return "", err
		}
	}
	return a.signer.Sign(claims)
}

func (a *AuthService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	return a.signer.Verify(tokenString)
}
