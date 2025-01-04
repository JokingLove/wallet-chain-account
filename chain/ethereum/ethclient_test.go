package ethereum

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	url := "https://eth-holesky.g.alchemy.com/v2/BvSZ5ZfdIwB-5SDXMz8PfGcbICYQqwrl"
	dialEthClient, _ := DialEthClient(context.Background(), url)
	t.Run("Check Banlance And Account Type", func(t *testing.T) {
		// 2. 要检查的地址
		addressStr := "0xe833ED183B05c662b0eF1d617b8CBf956925E2Bd"
		address := common.HexToAddress(addressStr)

		getBalance, err := dialEthClient.GetBalance(address)
		assert.NoError(t, err)

		// 5. 打印详细信息
		t.Logf("adress: %s", address)
		t.Logf("Balance (wei) : %s", getBalance.String())
		t.Logf("Balance (ETH) : %v", WeiToEth(getBalance).Text('f', 18))
	})
}

// weiToEth 辅助函数，将  wei  转为 ETH

func WeiToEth(wei *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		new(big.Float).SetInt(big.NewInt(1e18)),
	)
}
