package e2e

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/vishnushankarsg/metachain/e2e/framework"
	"github.com/vishnushankarsg/metachain/types"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/wallet"
)

var (
	one = big.NewInt(1)
)

func TestJsonRPC(t *testing.T) {
	//nolint:godox
	// TODO: Reuse the same tests for websockets and IPC. (to be fixed in EVM-521)
	fund, err := wallet.GenerateKey()
	require.NoError(t, err)

	bytecode, _ := hex.DecodeString(sampleByteCode)

	ibftManager := framework.NewIBFTServersManager(
		t,
		1,
		IBFTDirPrefix,
		func(i int, config *framework.TestServerConfig) {
			config.Premine(types.Address(fund.Address()), framework.EthToWei(10))
			config.SetBlockTime(1)
		},
	)

	ibftManager.StartServers(context.Background())
	defer ibftManager.StopServers()

	srv := ibftManager.GetServer(0)
	client := srv.JSONRPC().Eth()

	t.Run("eth_getBalance", func(t *testing.T) {
		key1, _ := wallet.GenerateKey()

		// Test. return zero if the account does not exists
		balance1, err := client.GetBalance(key1.Address(), ethgo.Latest)
		require.NoError(t, err)
		require.Equal(t, balance1, big.NewInt(0))

		// Test. return the balance of an account
		newBalance := big.NewInt(22000)
		txn, err := srv.Txn(fund).Transfer(key1.Address(), newBalance).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		balance1, err = client.GetBalance(key1.Address(), ethgo.Latest)
		require.NoError(t, err)
		require.Equal(t, balance1, newBalance)

		//nolint:godox
		// Test. return 0 if the balance of an existing account is empty
		// TODO (to be fixed in EVM-521)
	})

	t.Run("eth_getTransactionCount", func(t *testing.T) {
		key1, _ := wallet.GenerateKey()

		//nolint:godox
		// TODO. return zero if the account does not exists in the state (Does not work) (to be fixed in EVM-521)
		_, err := client.GetNonce(key1.Address(), ethgo.Latest)
		require.NoError(t, err)

		txn, err := srv.Txn(fund).Transfer(key1.Address(), big.NewInt(10000000000000000)).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		// Test. increase the nonce with new transactions
		txn = srv.Txn(key1)
		txn, err = txn.Send()
		require.NoError(t, err)
		txn.NoFail(t)

		nonce1, err := client.GetNonce(key1.Address(), ethgo.Latest)
		require.NoError(t, err)
		require.Equal(t, nonce1, uint64(1))

		// Test. you can query the nonce at any block number in time
		nonce1, err = client.GetNonce(key1.Address(), ethgo.BlockNumber(txn.Receipt().BlockNumber-1))
		require.NoError(t, err)
		require.Equal(t, nonce1, uint64(0))

		// Test. TODO. you can query the nonce at any block hash in time (to be fixed in EVM-521) //nolint:godox
		block, err := client.GetBlockByNumber(ethgo.BlockNumber(txn.Receipt().BlockNumber)-1, false)
		require.NoError(t, err)

		_, err = client.GetNonce(key1.Address(), ethgo.BlockNumber(block.Number))
		require.NoError(t, err)
	})

	t.Run("eth_getStorage", func(t *testing.T) {
		t.Skip("TODO")
	})

	t.Run("eth_getCode", func(t *testing.T) {
		// we use a predefined private key so that the deployed contract address is deterministic.
		// Note that in order to work, this private key should only be used for this test.
		priv, _ := hex.DecodeString("2c15bd0dc992a47ca660983ae4b611f4ffb6178e14e04e2b34d153f3a74ce741")
		key1, _ := wallet.NewWalletFromPrivKey(priv)

		// fund the account so that it can deploy a contract
		txn, err := srv.Txn(fund).Transfer(key1.Address(), big.NewInt(10000000000000000)).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		codeAddr := ethgo.HexToAddress("0xDBfca0c43cA12759256a7Dd587Dc4c6EEC1D89A5")

		// Test. We get empty code from an empty contract
		code, err := client.GetCode(codeAddr, ethgo.Latest)
		require.NoError(t, err)
		require.Equal(t, code, "0x")

		txn = srv.Txn(key1)
		txn, err = txn.Deploy(bytecode).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		receipt := txn.Receipt()

		// Test. The deployed address is the one expected
		require.Equal(t, codeAddr, receipt.ContractAddress)

		// Test. We can retrieve the code by address on latest, block number and block hash
		cases := []ethgo.BlockNumberOrHash{
			ethgo.Latest,
			ethgo.BlockNumber(receipt.BlockNumber),
			// receipt.BlockHash, TODO: It does not work (to be fixed in EVM-521) //nolint:godox
		}
		for _, c := range cases {
			code, err = client.GetCode(codeAddr, c)
			require.NoError(t, err)
			require.NotEqual(t, code, "0x")
		}

		// Test. We can query in past state (when the code was empty)
		code, err = client.GetCode(codeAddr, ethgo.BlockNumber(receipt.BlockNumber-1))
		require.NoError(t, err)
		require.Equal(t, code, "0x")

		// Test. Query by pending should default to latest
		code, err = client.GetCode(codeAddr, ethgo.Pending)
		require.NoError(t, err)
		require.NotEqual(t, code, "0x")
	})

	t.Run("eth_getBlockByX", func(t *testing.T) {
		t.Skip("TODO")
	})

	t.Run("eth_getTransactionReceipt", func(t *testing.T) {
		key1, _ := wallet.GenerateKey()

		txn := srv.Txn(fund)
		txn, err := txn.Transfer(key1.Address(), one).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		// Test. We cannot retrieve a receipt of an empty hash
		emptyReceipt, err := client.GetTransactionReceipt(ethgo.ZeroHash)
		require.NoError(t, err) // Note that ethgo does not return an error when the item does not exists
		require.Nil(t, emptyReceipt)

		// Test. We can retrieve the receipt by the hash
		receipt, err := client.GetTransactionReceipt(txn.Receipt().TransactionHash)
		require.NoError(t, err)

		// Test. The populated fields match with the block
		block, err := client.GetBlockByHash(receipt.BlockHash, false)
		require.NoError(t, err)

		require.Equal(t, receipt.TransactionHash, txn.Receipt().TransactionHash)
		require.Equal(t, receipt.BlockNumber, block.Number)
		require.Equal(t, receipt.BlockHash, block.Hash)

		// Test. The receipt of a deployed contract has the 'ContractAddress' field.
		txn = srv.Txn(fund)
		txn, err = txn.Deploy(bytecode).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		require.NotEqual(t, txn.Receipt().ContractAddress, ethgo.ZeroAddress)
	})

	t.Run("eth_getTransactionByHash", func(t *testing.T) {
		key1, _ := wallet.GenerateKey()

		// Test. We should be able to query the transaction by its hash
		txn := srv.Txn(fund)
		txn, err := txn.Transfer(key1.Address(), one).Send()
		require.NoError(t, err)
		txn.NoFail(t)

		ethTxn, err := client.GetTransactionByHash(txn.Receipt().TransactionHash)
		require.NoError(t, err)

		// Test. The dynamic 'from' field is populated
		require.NotEqual(t, ethTxn.From, ethgo.ZeroAddress)
	})
}
