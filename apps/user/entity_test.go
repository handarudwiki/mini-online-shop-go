package user

import (
	"handarudwiki/mini-online-shop-go/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "hanndaru@gmail.com",
			Password: "password",
		}
		err := userEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("email is required", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "",
			Password: "password",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})

	t.Run("email is invalid", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "handaru.gmail.com",
			Password: "password",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	t.Run("password is required", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "handaru@gmail.com",
			Password: "",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("invalid password length", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "handaru@gmail.com",
			Password: "123",
		}
		err := userEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalidLength, err)
	})
}
