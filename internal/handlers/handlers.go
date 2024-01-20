package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/tigerbeetle"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/utils"
	"io"
	"net/http"
)

type CreateAccountRequest struct {
	AccountID int `json:"account_id"`
	LedgerID  int `json:"ledger_id"`
}

type GetAccountRequest struct {
	AccountID int `json:"account_id"`
	LedgerID  int `json:"ledger_id,omitempty"`
}

type CreateTransactionRequest struct {
	Amount          float64 `json:"amount,omitempty"`
	DebitAccountID  int     `json:"debit_account_id,omitempty"`
	CreditAccountID int     `json:"credit_account_id,omitempty"`
	LedgerID        int     `json:"ledger_id,omitempty"`
}

type GetTransactionRequest struct {
	TransactionID int `json:"transaction_id"`
	LedgerID      int `json:"ledger_id,omitempty"`
}

func CreateAccount(c *gin.Context) {

	var getAccountRequest GetAccountRequest
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to unmarshal AccountRequest Body",
			"message": err,
		})
		return
	}

	err = json.Unmarshal(jsonData, &getAccountRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to unmarshal AccountRequest Body",
			"message": err,
		})
		return
	}

	success, err := tigerbeetle.CreateAccount(uint64(getAccountRequest.AccountID), uint32(getAccountRequest.LedgerID))
	if !success || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to create account",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
	return
}

func CreateTransaction(c *gin.Context) {
	var createTransactionRequest CreateTransactionRequest
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read createTransaction Body",
		})
		return
	}

	err = json.Unmarshal(jsonData, &createTransactionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to unmarshal createTransaction Body",
		})
		return
	}

	amountInLowest := utils.GetAmountInLowestForm(createTransactionRequest.Amount)

	success, txnID, err := tigerbeetle.PostCredits(uint64(createTransactionRequest.CreditAccountID), uint64(createTransactionRequest.DebitAccountID), uint64(amountInLowest), uint64(createTransactionRequest.LedgerID))
	if !success || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to create transaction",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":       "true",
		"transactionID": txnID,
	})
	return
}

func GetAccount(c *gin.Context) {
	var getAccountRequest GetAccountRequest
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read GetAccountRequest Body",
		})
		return
	}

	err = json.Unmarshal(jsonData, &getAccountRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to unmarshal GetAccountRequest Body",
		})
		return
	}

	account, err := tigerbeetle.GetAccount(uint64(getAccountRequest.AccountID))
	if account == nil || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get account",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, account)
	return
}

func GetTransaction(c *gin.Context) {
	var getTransactionRequest GetTransactionRequest
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read GetTransaction Body",
		})
		return
	}

	err = json.Unmarshal(jsonData, &getTransactionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to unmarshal GetTransactionRequest Body",
		})
		return
	}

	transaction, err := tigerbeetle.GetTransactionDetails(uint64(getTransactionRequest.TransactionID))
	if transaction == nil || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get transaction",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, transaction)
	return
}
