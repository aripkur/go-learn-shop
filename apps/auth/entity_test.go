package auth

import (
	"log"
	"testing"

	"github.com/aripkur/go-learn-shop/infra/response"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestValidateAuthEntity(t *testing.T) {
	t.Run("email success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "arif@gmail.com",
			Password: "1234567",
		}
		err := authEntity.validate()

		require.Nil(t, err)
	})
	t.Run("email required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "1234567",
		}
		err := authEntity.validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})
	t.Run("email invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "arifgmail.com",
			Password: "1234567",
		}
		err := authEntity.validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})
	t.Run("password success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "arif@gmail.com",
			Password: "1234567",
		}
		err := authEntity.validate()

		require.Nil(t, err)
	})
	t.Run("password required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "arif@gmail.com",
			Password: "",
		}
		err := authEntity.validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})
	t.Run("password invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "arif@gmail.com",
			Password: "12345",
		}
		err := authEntity.validate()

		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalid, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "arip@gmail.com",
			Password: "rahasia",
		}

		err := authEntity.encryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)

		log.Printf("%+v\n", authEntity)
	})
}
