package utility

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicID := uuid.NewString()
		tokenString, err := GenerateToken(publicID, "user", "secretBanget")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		secret := "iniSecret"
		role := "user"
		publicID := uuid.NewString()
		tokenString, err := GenerateToken(publicID, role, secret)
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		jwtID, jwtRole, err := ValidateToken(tokenString, secret)
		require.Nil(t, err)
		require.Equal(t, role, jwtRole)
		require.Equal(t, publicID, jwtID)
	})
}
