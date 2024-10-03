package bitvora

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Network string

const (
	Mainnet Network = "https://api.bitvora.com"
	Signet  Network = "https://api.signet.bitvora.com"
)

type BitvoraClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func NewBitvoraClient(network Network, apiKey string) *BitvoraClient {
	return &BitvoraClient{
		BaseURL: string(network),
		APIKey:  apiKey,
		Client:  &http.Client{},
	}
}

type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError: StatusCode=%d, Body=%s", e.StatusCode, e.Body)
}

func (client *BitvoraClient) doPost(endpoint string, requestBody interface{}, responseBody interface{}) error {
	url := fmt.Sprintf("%s/%s", client.BaseURL, endpoint)
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &APIError{StatusCode: resp.StatusCode, Body: string(bodyBytes)}
	}

	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func (client *BitvoraClient) doGet(endpoint string, responseBody interface{}) error {
	url := fmt.Sprintf("%s/%s", client.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)

	resp, err := client.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &APIError{StatusCode: resp.StatusCode, Body: string(bodyBytes)}
	}

	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func (client *BitvoraClient) Withdraw(request WithdrawRequest) (*WithdrawResponse, error) {
	var response WithdrawResponse
	err := client.doPost("v1/bitcoin/withdraw/confirm", request, &response)
	return &response, err
}

func (client *BitvoraClient) EstimateWithdrawal(request EstimateWithdrawalRequest) (*EstimateWithdrawalResponse, error) {
	var response EstimateWithdrawalResponse
	err := client.doPost("v1/bitcoin/withdraw/estimate", request, &response)
	return &response, err
}

func (client *BitvoraClient) CreateLightningInvoice(request CreateLightningInvoiceRequest) (*CreateLightningInvoiceResponse, error) {
	var response CreateLightningInvoiceResponse
	err := client.doPost("v1/bitcoin/deposit/lightning-invoice", request, &response)
	return &response, err
}

func (client *BitvoraClient) CreateLightningAddress(request CreateLightningAddressRequest) (*CreateLightningAddressResponse, error) {
	var response CreateLightningAddressResponse
	err := client.doPost("v1/bitcoin/deposit/lightning-address", request, &response)
	return &response, err
}

func (client *BitvoraClient) CreateOnChainAddress(request CreateOnChainAddressRequest) (*CreateOnChainAddressResponse, error) {
	var response CreateOnChainAddressResponse
	err := client.doPost("v1/bitcoin/deposit/on-chain", request, &response)
	return &response, err
}

func (client *BitvoraClient) GetBalance() (*GetBalanceResponse, error) {
	var response GetBalanceResponse
	err := client.doGet("v1/transactions/balance", &response)
	return &response, err
}

func (client *BitvoraClient) GetTransactions() (*GetTransactionsResponse, error) {
	var response GetTransactionsResponse
	err := client.doGet("v1/transactions", &response)
	return &response, err
}

func (client *BitvoraClient) GetDeposit(id string) (*GetDepositResponse, error) {
	var response GetDepositResponse
	err := client.doGet(fmt.Sprintf("v1/transactions/deposits/%s", id), &response)
	return &response, err
}

func (client *BitvoraClient) GetWithdrawal(id string) (*WithdrawResponse, error) {
	var response WithdrawResponse
	err := client.doGet(fmt.Sprintf("v1/transactions/withdrawals/%s", id), &response)
	return &response, err
}
