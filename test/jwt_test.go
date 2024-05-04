package test

import (
	"github.com/Catlordx/CampusTrade/internal/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	userId := 1
	role := "admin"
	expectedExp := time.Now().Add(time.Hour * 24).Unix()
	token, err := utils.GenerateToken(uint(userId), role)
	require.NoError(t, err, "Generating JWT should not return an error")
	require.NotEmpty(t, token, "Generated JWT should not be empty")
	parsedToken, err := utils.VerifyToken(token)
	require.NoError(t, err, "JWT should be valid")
	require.Equal(t, uint(userId), parsedToken.UserID, "UserID claim mismatch")
	require.Equal(t, role, parsedToken.Role, "Role claim mismatch")
	require.Equal(t, expectedExp, parsedToken.ExpiresAt.Unix(), "Expiration time claim mismatch")
}
