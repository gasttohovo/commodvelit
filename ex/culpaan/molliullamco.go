import (
	"context"
	"fmt"
	"math/big"

	"example.com/ethereum/go-ethereum/common"
	"example.com/ethereum/go-ethereum/core/types"
	"example.com/ethereum/go-ethereum/crypto"
	"example.com/ethereum/go-ethereum/ethclient"
)

// web3Stake performs a stake transaction on the VST contract.
func web3Stake(privateKey *ecdsa.PrivateKey, client *ethclient.Client, gasLimit uint64) (*types.Transaction, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	value := big.NewInt(0)
	data := []byte{}

	tx := types.NewTransaction(nonce, common.HexToAddress(vstContractAddress), value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("Stake transaction sent: %s\n", signedTx.Hash().Hex())

	return signedTx, nil
}
  
