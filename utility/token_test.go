package utility

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("generate_success", func(t *testing.T) {
		publicId := uuid.NewString()
		token, err := GenerateToken(publicId, "user", "rahasia")

		log.Println(publicId)
		log.Println(token)

		require.Nil(t, err)
		require.NotEmpty(t, token)

	})
	t.Run("verify_success", func(t *testing.T) {
		secret := "rtahasia"
		id := uuid.NewString()
		role := "user"

		token, err := GenerateToken(id, role, secret)
		require.Nil(t, err)

		idV, roleV, err := ValidateToken(token, secret)
		log.Println(idV)
		log.Println(roleV)

		require.Nil(t, err)
		require.Equal(t, id, idV)
		require.Equal(t, role, roleV)

	})
}
