package auth

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aripkur/go-learn-shop/external/database"
	"github.com/aripkur/go-learn-shop/infra/response"
	"github.com/aripkur/go-learn-shop/internal/config"
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

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "rahasia",
	}

	ctx := context.Background()
	err := svc.register(ctx, req)

	require.Nil(t, err)

}

func TestRegister_Failed(t *testing.T) {
	t.Run("email already used", func(t *testing.T) {
		req := RegisterRequestPayload{
			Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
			Password: "rahasia",
		}

		ctx := context.Background()
		_ = svc.register(ctx, req)

		err := svc.register(ctx, req)

		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlReadyUsed, err)
	})
}

func TestLogin_Success(t *testing.T) {
	email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
	pwd := "rahasia"

	reqRegis := RegisterRequestPayload{
		Email:    email,
		Password: pwd,
	}

	ctx := context.Background()
	err := svc.register(ctx, reqRegis)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email:    email,
		Password: pwd,
	}

	token, err := svc.login(ctx, reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)
	log.Println(token)
}
