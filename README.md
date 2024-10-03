# Bitvora Go SDK

This SDK provides a convenient way to interact with the Bitvora API using the Go programming language. It allows you to manage Bitcoin deposits and withdrawals, create lightning invoices and addresses, and retrieve balance and transaction information.

## Installation

To install the Bitvora Go SDK, use the following command:

```bash
go get github.com/bitvora/go-bitvora
```

## Usage

### 1. Initialization

First, you need to create a `BitvoraClient` instance. You'll need your Bitvora API key and choose either the mainnet or signet network.

```go
package main

import (
	"fmt"
	"log"

	"github.com/bitvora/go-bitvora"
)

func main() {
	apiKey := "YOUR_BITVORA_API_KEY" // Replace with your actual API key
	client := bitvora.NewBitvoraClient(bitvora.Mainnet, apiKey) // Use bitvora.Signet for testnet

	// ... your Bitvora API calls here ...
}
```

### 2. API Calls

The SDK provides methods for various Bitvora API endpoints:

#### a) Get Balance

Retrieves your Bitvora account balance.

```go
balance, err := client.GetBalance()
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Balance: %d sats\n", balance.Data.Balance)
```

#### b) Get Transactions

Retrieves a list of your transactions.

```go
transactions, err := client.GetTransactions()
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Transactions: %+v\n", transactions.Data)
```

#### c) Withdraw Bitcoin

Initiates a Bitcoin withdrawal. Requires amount, currency, destination address, and optional metadata.

```go
amount := 100.50
currency := bitvora.USD
destination := "bc1qxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // chain address, lightning invoice, or lightning address
metadata := map[string]string{"user_id": "123"}

withdrawal, err := client.Withdraw(amount, currency, destination, metadata)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Withdrawal successful: %+v\n", withdrawal)
```

#### d) Estimate Withdrawal Fee

Estimates the fee for a Bitcoin withdrawal before actually initiating it.

```go
amount := 1000.0 // Amount in Bitcoin
currency := "btc"
destination := "bc1qxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

estimate, err := client.EstimateWithdrawal(amount, currency, destination)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Withdrawal estimate: %+v\n", estimate)
```

#### e) Create Lightning Invoice

Creates a Lightning invoice.

```go
amount := 1000.0   // Amount in sats
currency := "sats"
description := "Payment for goods"
expirySeconds := 3600 // 1 hour
metadata := map[string]string{"order_id": "12345"}

invoice, err := client.CreateLightningInvoice(amount, currency, description, expirySeconds, metadata)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Lightning Invoice: %+v\n", invoice)
```

#### f) Create Lightning Address

Creates a Lightning address.

```go
handle := "my-store"
domain := "bitvora.com"
metadata := map[string]string{"description": "My Lightning Address"}

address, err := client.CreateLightningAddress(handle, domain, metadata)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Lightning Address: %+v\n", address)
```

#### g) Create On-Chain Address

Creates a new Bitcoin on-chain deposit address.

```go
metadata := map[string]string{"description": "On-chain deposit"}

address, err := client.CreateOnChainAddress(metadata)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("On-Chain Address: %+v\n", address)
```

#### h) Get Deposit

Retrieves details of a specific deposit using its ID.

```go
depositID := "YOUR_DEPOSIT_ID" // Replace with the actual deposit ID
deposit, err := client.GetDeposit(depositID)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Deposit details: %+v\n", deposit)
```

#### i) Get Withdrawal

Retrieves details of a specific withdrawal using its ID.

```go
withdrawalID := "YOUR_WITHDRAWAL_ID" // Replace with the actual withdrawal ID
withdrawal, err := client.GetWithdrawal(withdrawalID)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Withdrawal details: %+v\n", withdrawal)
```

### 3. Error Handling

The SDK uses custom `APIError` struct to represent errors returned by the Bitvora API. This struct contains the HTTP status code and the response body.

```go
if err != nil {
	apiErr, ok := err.(*bitvora.APIError)
	if ok {
		fmt.Printf("Bitvora API error: Status Code=%d, Body=%s\n", apiErr.StatusCode, apiErr.Body)
	} else {
		log.Fatal(err) // Handle other errors
	}
}
```

## Data Structures

The SDK defines structs corresponding to the Bitvora API response structures. Refer to the code for detailed information on the fields within each response struct (e.g., `WithdrawResponse`, `GetBalanceResponse`, etc.).

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This SDK is released under the MIT License.
