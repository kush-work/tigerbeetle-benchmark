package tigerbeetle

import (
	"errors"
	"fmt"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/constants"
	nonce "github.com/kushthedude/tigerbeetle-benchmark/internal/nonce"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/security"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/utils"
	. "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func GetAccount(accountID uint64) (*Account, error) {
	accounts, err := constants.TigerBeetleClient.LookupAccounts([]Uint128{ToUint128(accountID)})

	if err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		return nil, errors.New("account not found")
	}

	return &accounts[0], nil
}

func GetTransactionDetails(txnID uint64) (*Transfer, error) {
	transfers, err := constants.TigerBeetleClient.LookupTransfers([]Uint128{ToUint128(txnID)})

	if err != nil {
		return nil, err
	}

	if len(transfers) == 0 {
		return nil, errors.New("account not found")
	}

	return &transfers[0], nil
}
func CreateAccount(accountID uint64, ledgerID uint32) (bool, error) {
	accountToCreate := []Account{{
		ID:            ToUint128(accountID),
		UserData128:   ToUint128(0),
		UserData64:    uint64(0),
		UserData32:    uint32(0),
		Code:          constants.AccountTypeCode,
		Flags:         constants.AccountFlags,
		DebitsPosted:  ToUint128(0),
		CreditsPosted: ToUint128(0), // 100 units initial balance credited
		Reserved:      uint32(0),
		Ledger:        ledgerID,
	}}
	result, err := constants.TigerBeetleClient.CreateAccounts(accountToCreate)
	if err != nil {
		fmt.Println(result)
		fmt.Println(err)
		return false, err
	}
	if len(result) != 0 {
		return false, errors.New("account creation result size is not zero")
	}
	return true, nil
}

func PostCredits(accountIdToBeCredited uint64, accountIdToBeDebited uint64, amountToBeCredited uint64, ledgerID uint64) (bool, int, error) {
	txnID := utils.RandomNumberGenerator().Int()
	privateKey := constants.GetPrivateKey()
	publicKey := constants.GetPublicKey()
	transferObj := Transfer{
		ID:              ToUint128(uint64(txnID)),
		DebitAccountID:  ToUint128(accountIdToBeDebited),
		CreditAccountID: ToUint128(accountIdToBeCredited),
		Amount:          ToUint128(amountToBeCredited),
		PendingID:       ToUint128(0),
		UserData128:     ToUint128(2),
		UserData64:      0,
		UserData32:      0,
		Timeout:         0,
		Ledger:          uint32(ledgerID),
		Code:            1,
		Flags:           0,
		Timestamp:       0,
	}

	signedTxn, err := security.SignTransaction(&transferObj, privateKey)
	if err != nil {
		panic("failed to sign txn")
	}

	verified, err := security.VerifySignature(&transferObj, signedTxn, &publicKey)
	if err != nil {
		fmt.Println("failed at verifysignatute")
		return false, 0, err
	}

	if !verified {
		panic("fake txn")
	}

	transferObject := []Transfer{{
		ID:              ToUint128(uint64(txnID)),
		DebitAccountID:  ToUint128(accountIdToBeDebited),
		CreditAccountID: ToUint128(accountIdToBeCredited),
		Amount:          ToUint128(amountToBeCredited),
		PendingID:       ToUint128(0),
		UserData128:     ToUint128(2),
		UserData64:      0,
		UserData32:      0,
		Timeout:         0,
		Ledger:          uint32(ledgerID),
		Code:            1,
		Flags:           0,
		Timestamp:       0,
	}}

	nonceObj, err := nonce.NonceInstance.GenerateNonce()
	if err != nil {
		fmt.Printf("Error generating nonce: %v\n", err)
		return false, 0, err
	}

	fmt.Printf("Generated nonce: %s\n", nonceObj)

	valid, err := nonce.NonceInstance.VerifyNonce(nonceObj)
	if err != nil {
		return false, 0, err
	}

	if !valid {
		fmt.Println("failed at nonce invalid")
		return false, 0, errors.New("nonce is invalid")
	}

	result, err := constants.TigerBeetleClient.CreateTransfers(transferObject)
	if err != nil {
		fmt.Println("failed at here")
		fmt.Println(result)
		return false, 0, err
	}
	if len(result) != 0 {
		fmt.Println("failed at sdsds")
		fmt.Println(result)
		return false, 0, errors.New("transaction result size is not zero")
	}
	return true, txnID, nil
}
