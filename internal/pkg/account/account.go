package account

import (
	"github.com/leonardo3dp/binance-nft-buy/internal/domain/account"
	bapi "github.com/leonardo3dp/binance-nft-buy/pkg/binance-api"
)

type Account struct {
	Setting     account.Setting
	Auth        *bapi.Api
}

func InitAccount(setting account.Setting) (*Account, error) {
	auth, err := bapi.New(setting)
	if err != nil {
		return nil, err
	}
	return &Account{Setting: setting, Auth: auth}, nil
}
