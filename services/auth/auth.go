package auth_services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIBOAIBAAJAc7DQV48nYshsJkVmuj0T1JSHlahAPiRY8z9r0DkFppzwO4zrx9uD
FHcm6j9Kh09XUSht2yDGC6xqWgOtTJnw2QIDAQABAkBwpUT3PIgZAxVq3kB8LmRU
pJqv+bczyqhhkOslP6Bk7TJ3AFXoPXaevm4EkRaLLANyJAI291Im/AnLfEvwPxXR
AiEAuJWAoZyFZCVaVLGtOgW9A3CPa0rNdx0cSFpgjgB7ObcCIQCgc6H5ShXTFBCf
TrQB+4Dn74YSreDKnS3fgEt1hDVp7wIgSDkfIp0myF+hL6Bx4lEaev0Q8O9M4719
MoZCX22qyZMCIH1spqsWbKUJxEyj2zbJgWTM6gNkBJqd76QMx+/fH1nlAiBCRhY8
Giv+YJmd+8r6TYJB+Ar1TuP93jmgoy8vKY43+A==
-----END RSA PRIVATE KEY-----
`

type Auth interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	privateKey string
}

func JWTAuthService() Auth {
	return &jwtServices{
		privateKey: privateKey,
	}
}

func (service *jwtServices) GenerateToken(email string, isUser bool) string {

	fmt.Println(service.privateKey)

	SigningKey := []byte(privateKey)
	key, err := jwt.ParseRSAPrivateKeyFromPEM(SigningKey)
	if err != nil {
		fmt.Println(err)
	}

	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		fmt.Println(err)
	}

	return token
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.privateKey), nil
	})

}
