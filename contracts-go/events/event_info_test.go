package events_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/eluv-io/contracts/contracts-go/events"
	"github.com/eluv-io/contracts/contracts-go/tradable"
)

func TestEventInfo_Value(t *testing.T) {
	ev := &events.EventInfo{
		Name: "Transfer",
		ID:   common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		Types: []events.EventType{
			{
				Type:          reflect.TypeOf((*tradable.TransferERC20)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC20),
			}, {
				Type:          reflect.TypeOf((*tradable.TransferERC721)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC721),
			},
		},
	}

	log := types.Log{
		Address: common.HexToAddress("0xeda0813b6087443db2a5bc9f63f9e01f31943e71"),
		Topics: []common.Hash{
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			common.HexToHash("0x00000000000000000000000032ff9d67d7523cc9a56297c578f412d93a665667"),
		},
		Data:        common.FromHex("0x00000000000000000000000000000000000000000000d3c21bcecceda1000000"),
		BlockNumber: 1811496,
		TxHash:      common.HexToHash("0xedce33c6a79860b1608330d475b5a7b7c416e6ea43b6e5521aea50efda730da1"),
		TxIndex:     0,
		BlockHash:   common.HexToHash("0x4ff663bec7d4a87ace6d5b2fa6bbd39a5a7f9c60f13d107da9f0307b36421ae4"),
		Index:       0,
		Removed:     false,
	}
	var erc20 *tradable.TransferERC20
	{
		ok := false
		ierc20, err := ev.Value(log, reflect.TypeOf((*tradable.TransferERC20)(nil)))
		require.NoError(t, err)
		erc20, ok = ierc20.Interface().(*tradable.TransferERC20)
		require.True(t, ok)
	}
	require.Equal(t,
		common.Address{},
		erc20.From)
	require.Equal(t,
		common.HexToAddress("0x32ff9d67d7523cc9a56297c578f412d93a665667"),
		erc20.To)
	bi, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	require.Equal(t,
		bi,
		erc20.Value)

	//bb, err := json.MarshalIndent(erc20, "", "  ")
	//require.NoError(t, err)
	//fmt.Println(string(bb))

	_, err := ev.Value(log, reflect.TypeOf((*tradable.TransferERC721)(nil)))
	require.Error(t, err) //'topic/field count mismatch'

	//
	// reverse order of types
	//
	ev = &events.EventInfo{
		Name: "Transfer",
		ID:   common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		Types: []events.EventType{
			{
				Type:          reflect.TypeOf((*tradable.TransferERC721)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC721),
			},
			{
				Type:          reflect.TypeOf((*tradable.TransferERC20)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC20),
			},
		},
	}
	i, err := ev.Value(log)
	require.NoError(t, err)
	ok := false
	erc20, ok = i.Interface().(*tradable.TransferERC20)
	require.True(t, ok)

	type PseudoTransferERC20 struct {
		From  common.Address
		To    common.Address
		Value *big.Int
		Raw   types.Log // Blockchain specific contextual infos
	}
	ev = &events.EventInfo{
		Name: "Transfer",
		ID:   common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		Types: []events.EventType{
			{
				Type:          reflect.TypeOf((*PseudoTransferERC20)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC20),
			},
			{
				Type:          reflect.TypeOf((*tradable.TransferERC721)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC721),
			},
			{
				Type:          reflect.TypeOf((*tradable.TransferERC20)(nil)),
				BoundContract: tradable.BoundContract(tradable.K_ERC20),
			},
		},
	}
	i, err = ev.Value(log)
	require.NoError(t, err)
	ok = false
	_, ok = i.Interface().(*PseudoTransferERC20)
	require.True(t, ok)

	i, err = ev.Value(log, reflect.TypeOf((*tradable.TransferERC20)(nil)))
	require.NoError(t, err)
	ok = false
	_, ok = i.Interface().(*tradable.TransferERC20)
	require.True(t, ok)
}
