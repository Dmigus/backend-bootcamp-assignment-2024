//go:build unit

package jwt

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/auth/usecases/login"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodec_Encode_Positive(t *testing.T) {
	t.Parallel()
	c := NewCodec([]byte("mykey"))
	token, err := c.Encode(models.AuthClaims{Role: models.Moderator, Name: "Sasha"})
	require.NoError(t, err)
	// ожидаемый токен, полученный на сайте jwt.io
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiU2FzaGEiLCJyb2xlIjoibW9kZXJhdG9yIn0.pBKOKGwuWpEvVOLRtMMHB3vP_jku84JKZEB5dZdBhkg"
	assert.Equal(t, expectedToken, token)
}

func TestCodec_Encode_Errors(t *testing.T) {
	t.Parallel()
	c := NewCodec([]byte("mykey"))
	_, err := c.Encode(models.AuthClaims{Role: models.UserRole(0), Name: "Sasha"})
	require.ErrorIs(t, err, login.ErrNoSuchUserType)
}

func TestCodec_Decode_Positive(t *testing.T) {
	t.Parallel()
	c := NewCodec([]byte("mykey"))
	claims, err := c.Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRGltYSIsInJvbGUiOiJjbGllbnQifQ.LjThmQ1aR0jGfF4z2JiOMdwvsTQcZSC76b6V1Jx_dBE")
	require.NoError(t, err)
	assert.Equal(t, claims.Role, models.Client)
	assert.Equal(t, claims.Name, "Dima")
}

func TestCodec_Decode_Errors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		token string
		want  error
	}{
		{
			"invalid token",
			"invalid",
			login.ErrInvalidToken,
		},
		{
			"wrong role",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRGltYSIsInJvbGUiOiJ3cm9uZyJ9.EINMeOLunVM6nK6ny_4TWlboRz2idlk5Zz3CsP58ujY",
			login.ErrInvalidToken,
		},
		{
			"wrong key",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRGltYSIsInJvbGUiOiJjbGllbnQifQ.GL_ROPKjphKyST0-jbPxvyU9y2rqq-4zgi2Hav9rSrI",
			login.ErrInvalidToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := NewCodec([]byte("mykey"))
			_, err := c.Decode(tt.token)
			require.ErrorIs(t, err, tt.want)
		})
	}
}
