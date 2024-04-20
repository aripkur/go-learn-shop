package database

import (
	"testing"

	"github.com/aripkur/go-learn-shop/internal/config"
	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectMysql(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectMysql(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})

	t.Run("invalid password", func(t *testing.T) {
		cfg := config.Cfg.DB
		cfg.Password = "invalid password"
		db, err := ConnectMysql(cfg)
		require.Nil(t, db)
		require.NotNil(t, err)
	})
}
