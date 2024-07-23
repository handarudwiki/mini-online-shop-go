package user

import (
	"context"
	"fmt"
	"handarudwiki/mini-online-shop-go/external/database"
	"handarudwiki/mini-online-shop-go/infra/response"
	"handarudwiki/mini-online-shop-go/internal/config"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestResgisterSuccess(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "password",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}

func TestRegisterFail(t *testing.T) {
	t.Run("error email is already exist", func(t *testing.T) {
		req := RegisterRequestPayload{
			Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
			Password: "password",
		}
		err := svc.register(context.Background(), req)
		require.Nil(t, err)

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}

func TestLoginSuccess(t *testing.T) {
	reqRegister := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "password",
	}
	err := svc.register(context.Background(), reqRegister)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email:    reqRegister.Email,
		Password: reqRegister.Password,
	}

	token, err := svc.login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)
}
