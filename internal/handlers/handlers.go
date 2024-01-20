package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/tigerbeetle"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/utils"
	"io"
	"net/http"
	"strconv"
)

type CreateAccountRequest struct {
	AccountID int `json:"account_id"`
	LedgerID  int `json:"ledger_id"`
}

//type GetAccountRequest struct {
//	AccountID int `json:"account_id"`
//	LedgerID  int `json:"ledger_id,omitempty"`
//}

type CreateTransactionRequest struct {
	Amount          float64 `json:"amount,omitempty"`
	DebitAccountID  int     `json:"debit_account_id,omitempty"`
	CreditAccountID int     `json:"credit_account_id,omitempty"`
	LedgerID        int     `json:"ledger_id,omitempty"`
}

//
//type GetTransactionRequest struct {
//	TransactionID int `json:"transaction_id"`
//	LedgerID      int `json:"ledger_id,omitempty"`
//}

func CreateAccount(c *gin.Context) {

	var createAccountRequest CreateAccountRequest
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to unmarshal AccountRequest Body",
			"message": err,
		})
		return
	}

	err = json.Unmarshal(jsonData, &createAccountRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to unmarshal AccountRequest Body",
			"message": err,
		})
		return
	}

	success, err := tigerbeetle.CreateAccount(uint64(createAccountRequest.AccountID), uint32(createAccountRequest.LedgerID))
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
			"message": err.Error(),
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
	accountId, present := c.GetQuery("account_id")
	if !present {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account_id not passed in query parameter",
		})
		return
	}

	accountIdInt, err := strconv.Atoi(accountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account_id is not passed as valid int",
		})
		return
	}

	account, err := tigerbeetle.GetAccount(uint64(accountIdInt))
	if account == nil || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get account",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, account)
	return
}

func GetTransaction(c *gin.Context) {
	transactionID, present := c.GetQuery("transaction_id")
	if !present {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "transaction_id not passed in query parameter",
		})
		return
	}

	transactionIDInt, err := strconv.Atoi(transactionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "account_id is not passed as valid int",
		})
		return
	}

	transaction, err := tigerbeetle.GetTransactionDetails(uint64(transactionIDInt))
	if transaction == nil || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to get transaction",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, transaction)
	return
}
