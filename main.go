package main

import (
	"context"
	"fmt"
	. "github.com/tigerbeetle/tigerbeetle-go"
	. "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
	"log"
	"time"
)

const (
	tigerbeetleServerAddress = "<TIGERBEETLE_SERVER_ADDRESS>"
	clusterID                = "tigerbeetle"
	accountTypeCode          = 718
	accountFlags             = 0
)

func main() {
	timestamp := uint64(time.Now().UnixMilli())
	// Assume these are constants you've defined for your application:

	client, err := NewClient(clusterID, []string{tigerbeetleServerAddress}, nil)
	if err != nil {
		panic(err)
	}

	// Create a context with a timeout or cancellation ability.
	ctx := context.Background()

	// Define accounts for users X and Y with an initial balance of 100 units each.
	accounts := []Account{
		{
			ID:            ToUint128(1),
			UserData128:   ToUint128(0),
			UserData64:    0,
			UserData32:    0,
			Code:          accountTypeCode,
			Flags:         accountFlags,
			DebitsPosted:  0,
			CreditsPosted: 100, // 100 units initial balance credited
			Timestamp:     timestamp,
			Reserved:      0,
			Ledger:        1,
		},
		{
			ID:            ToUint128(2),
			UserData128:   ToUint128(0),
			UserData64:    0,
			UserData32:    0,
			Code:          accountTypeCode,
			Flags:         accountFlags,
			DebitsPosted:  0,
			CreditsPosted: 100, // 100 units initial balance credited
			Timestamp:     timestamp,
			Reserved:      0,
			Ledger:        1,
		},
	}

	// Call the CreateAccounts method with the accounts array.
	result, err := client.CreateAccounts(accounts)
	if err != nil {
		log.Printf("Error creating accounts: %s", err)
		for _, err := range result {
			log.Printf("Error creating account %d: %s", err.Index, err.Result)
			return
		}
		panic(err)
	}

	fmt.Printf("Account creation results: %v\n", result)
}
