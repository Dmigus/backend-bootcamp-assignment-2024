//go:build unit

package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthService_DummyLogin_Positive(t *testing.T) {
	t.Parallel()
	a := &AuthService{
		key: []byte("mykey"),
	}
	token, err := a.DummyLogin(Client)
	require.NoError(t, err)
	// ожидаемый токен, полученный на сайте jwt.io
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiY2xpZW50In0.-pYcuqSQ9CPJcjKmy4E_fu7yugYXlaSxOQKBHZHsBdY"
	assert.Equal(t, expectedToken, token)
}

func TestAuthService_DummyLogin_Errors(t *testing.T) {
	t.Parallel()
	a := &AuthService{
		key: []byte("mykey"),
	}
	_, err := a.DummyLogin(UserType(0))
	require.ErrorIs(t, err, ErrNoSuchUserType)
}

func TestAuthService_GetRole_Client(t *testing.T) {
	t.Parallel()
	a := &AuthService{
		key: []byte("mykey"),
	}
	clientToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiY2xpZW50In0.-pYcuqSQ9CPJcjKmy4E_fu7yugYXlaSxOQKBHZHsBdY"
	role, err := a.GetRole(clientToken)
	require.NoError(t, err)
	assert.Equal(t, Client, role)
}

func TestAuthService_GetRole_Moderator(t *testing.T) {
	t.Parallel()
	a := &AuthService{
		key: []byte("mykey"),
	}
	moderatorToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoibW9kZXJhdG9yIn0.g2DHwsP5mL3Oxql152N0k8V-aZ2K9YqgO-gLTVKwyUg"
	role, err := a.GetRole(moderatorToken)
	require.NoError(t, err)
	assert.Equal(t, Moderator, role)
}

func TestAuthService_GetRole_Errors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token string
		want  error
	}{
		{
			"invalid token",
			"invalid",
			ErrInvalidToken,
		},
		{
			"wrong role",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoid3Jvbmdyb2xlIn0.M5aJiMqHOQ3zpvy-7ahhgUEHrHBE-qjDbzaCHRffCC8",
			ErrInvalidToken,
		},
		{
			"wrong key",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiY2xpZW50In0.T6DBdTL0P05-v_t6Tx3o-RizKzXGWvEoyAUxJNSfe9s",
			ErrInvalidToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &AuthService{
				key: []byte("mykey"),
			}
			_, err := a.GetRole(tt.token)
			require.ErrorIs(t, err, tt.want)
		})
	}
}
