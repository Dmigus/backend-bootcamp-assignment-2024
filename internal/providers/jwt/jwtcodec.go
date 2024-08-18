package jwt

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/login"
	"github.com/golang-jwt/jwt/v5"
)

type Codec struct {
	key []byte
}

func NewCodec(key []byte) *Codec {
	return &Codec{key: key}
}

func (J *Codec) Encode(claims models.AuthClaims) (string, error) {
	jwtClaims := jwt.MapClaims{}
	roleStr := claims.Role.String()
	if roleStr == "" {
		return "", login.ErrNoSuchUserType
	}
	jwtClaims["role"] = roleStr
	jwtClaims["name"] = claims.Name
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return t.SignedString(J.key)
}

func (J *Codec) Decode(tokenStr string) (*models.AuthClaims, error) {
	claims, err := J.claimsFromToken(tokenStr)
	if err != nil {
		return nil, err
	}
	role, err := roleFromClaims(claims)
	if err != nil {
		return nil, err
	}
	name, err := nameFromClaims(claims)
	if err != nil {
		return nil, err
	}
	return &models.AuthClaims{Role: role, Name: name}, nil
}

func (J *Codec) claimsFromToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return J.key, nil
	})
	if err != nil {
		return nil, login.ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, login.ErrInvalidToken
	}
	return claims, nil
}

func roleFromClaims(claims jwt.MapClaims) (models.UserRole, error) {
	switch claims["role"] {
	case models.ClientRoleName:
		return models.Client, nil
	case models.ModeratorRoleName:
		return models.Moderator, nil
	default:
		return models.UserRole(0), login.ErrInvalidToken
	}
}

func nameFromClaims(claims jwt.MapClaims) (string, error) {
	if claims["name"] == nil {
		return "", nil
	}
	name, ok := claims["name"].(string)
	if !ok {
		return "", login.ErrInvalidToken
	}
	return name, nil
}
