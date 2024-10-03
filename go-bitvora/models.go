package bitvora

type WithdrawRequest struct {
	Amount      float64           `json:"amount"`
	Currency    string            `json:"currency"`
	Destination string            `json:"destination"`
	Metadata    map[string]string `json:"metadata"`
}

type WithdrawResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    WithdrawData `json:"data"`
}

type WithdrawData struct {
	ID               string                   `json:"id"`
	AmountSats       uint64                   `json:"amount_sats"`
	Recipient        string                   `json:"recipient"`
	FeeSats          float64                  `json:"fee_sats"`
	NetworkType      string                   `json:"network_type"`
	RailType         string                   `json:"rail_type"`
	Status           string                   `json:"status"`
	LightningPayment *LNDTrackPaymentResponse `json:"lightning_payment,omitempty"`
	ChainTxID        *string                  `json:"chain_tx_id,omitempty"`
	Metadata         *map[string]string       `json:"metadata,omitempty"`
	CreatedAt        string                   `json:"created_at"`
}

type LNDTrackPaymentResponse struct {
	PaymentHash     string           `json:"payment_hash"`
	Value           string           `json:"value"`
	CreationDate    string           `json:"creation_date"`
	Fee             string           `json:"fee"`
	PaymentPreimage string           `json:"payment_preimage"`
	ValueSat        string           `json:"value_sat"`
	ValueMsat       string           `json:"value_msat"`
	PaymentRequest  string           `json:"payment_request"`
	Status          string           `json:"status"`
	FeeSat          string           `json:"fee_sat"`
	FeeMsat         string           `json:"fee_msat"`
	CreationTimeNS  string           `json:"creation_time_ns"`
	HTLCs           []LNDHTLCAttempt `json:"htlcs"`
	PaymentIndex    string           `json:"payment_index"`
	FailureReason   string           `json:"failure_reason"`
}

type LNDHTLCAttempt struct {
	AttemptID     string            `json:"attempt_id"`
	Status        string            `json:"status"`
	Route         LNDPaymentRoute   `json:"route"`
	AttemptTimeNS string            `json:"attempt_time_ns"`
	ResolveTimeNS string            `json:"resolve_time_ns"`
	Failure       LNDPaymentFailure `json:"failure"`
	Preimage      string            `json:"preimage"`
}

type LNDPaymentRoute struct {
	TotalTimeLock int      `json:"total_time_lock"`
	TotalFees     string   `json:"total_fees"`
	TotalFeesMsat string   `json:"total_fees_msat"`
	TotalAmt      string   `json:"total_amt"`
	Hops          []LNDHop `json:"hops"`
}

type LNDHop struct {
	ChanID       string `json:"chan_id"`
	ChanCapacity string `json:"chan_capacity"`
	AmtToForward string `json:"amt_to_forward"`
	Expiry       int    `json:"expiry"`
}

type LNDPaymentFailure struct {
	Code               string           `json:"code"`
	ChannelUpdate      LNDChannelUpdate `json:"channel_update"`
	HTLCMsat           string           `json:"htlc_msat"`
	OnionSHA256        string           `json:"onion_sha_256"`
	CLTVExpiry         int              `json:"cltv_expiry"`
	Flags              int              `json:"flags"`
	FailureSourceIndex int              `json:"failure_source_index"`
	Height             int              `json:"height"`
}

type LNDChannelUpdate struct {
	Signature       string `json:"signature"`
	ChainHash       string `json:"chain_hash"`
	ChanID          string `json:"chan_id"`
	Timestamp       int    `json:"timestamp"`
	MessageFlags    int    `json:"message_flags"`
	ChannelFlags    int    `json:"channel_flags"`
	TimeLockDelta   int    `json:"time_lock_delta"`
	HTLCMinimumMsat string `json:"htlc_minimum_msat"`
	BaseFee         int    `json:"base_fee"`
	FeeRate         int    `json:"fee_rate"`
	HTLCMaximumMsat string `json:"htlc_maximum_msat"`
	ExtraOpaqueData string `json:"extra_opaque_data"`
}

type EstimateWithdrawalRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Destination string  `json:"destination"`
}

type EstimateWithdrawalResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    EstimateWithdrawalData `json:"data"`
}

type EstimateWithdrawalData struct {
	Recipient          string  `json:"recipient"`
	RecipientType      string  `json:"recipient_type"`
	AmountSats         uint64  `json:"amount_sats"`
	BitvoraFeeSats     float64 `json:"bitvora_fee_sats"`
	SuccessProbability float64 `json:"success_probability"`
}

type CreateLightningInvoiceRequest struct {
	Amount        float64           `json:"amount"`
	Currency      string            `json:"currency"`
	Description   string            `json:"description"`
	ExpirySeconds uint64            `json:"expiry_seconds"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

type CreateLightningInvoiceResponse struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    CreateLightningInvoiceData `json:"data"`
}

type CreateLightningInvoiceData struct {
	ID             string            `json:"id"`
	NodeID         string            `json:"node_id"`
	Memo           string            `json:"memo"`
	RPreimage      string            `json:"r_preimage"`
	RHash          string            `json:"r_hash"`
	AmountSats     uint64            `json:"amount_sats"`
	Settled        bool              `json:"settled"`
	PaymentRequest string            `json:"payment_request"`
	Metadata       map[string]string `json:"metadata,omitempty"`
}

type CreateLightningAddressRequest struct {
	Handle   string            `json:"handle"`
	Domain   string            `json:"domain"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateLightningAddressResponse struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    CreateLightningAddressData `json:"data"`
}

type CreateLightningAddressData struct {
	ID         string            `json:"id"`
	Handle     string            `json:"handle"`
	Domain     string            `json:"domain"`
	Address    string            `json:"address"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	CreatedAt  string            `json:"created_at"`
	LastUsedAt *string           `json:"last_used_at,omitempty"`
	DeletedAt  *string           `json:"deleted_at,omitempty"`
}

type CreateOnChainAddressRequest struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

type CreateOnChainAddressResponse struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    CreateOnChainAddressData `json:"data"`
}

type CreateOnChainAddressData struct {
	ID        string            `json:"id"`
	Address   string            `json:"address"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	CreatedAt string            `json:"created_at"`
}

type GetBalanceResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    GetBalanceData `json:"data"`
}

type GetBalanceData struct {
	Balance uint32 `json:"balance"`
}

type GetTransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

type Transaction struct {
	ID           string `json:"id"`
	CompanyID    string `json:"company_id"`
	AmountSats   uint64 `json:"amount_sats"`
	Recipient    string `json:"recipient"`
	RailType     string `json:"rail_type"`
	Type         string `json:"type"`
	FeeMicrosats uint64 `json:"fee_microsats"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
}

type GetDepositResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    GetDepositData `json:"data"`
}

type GetDepositData struct {
	ID                 string             `json:"id"`
	LedgerTxID         string             `json:"ledger_tx_id"`
	Recipient          string             `json:"recipient"`
	AmountSats         uint64             `json:"amount_sats"`
	FeeSats            float64            `json:"fee_sats"`
	ChainTxID          *string            `json:"chain_tx_id,omitempty"`
	RailType           string             `json:"rail_type"`
	NetworkType        string             `json:"network_type"`
	Status             string             `json:"status"`
	Metadata           *map[string]string `json:"metadata,omitempty"`
	LightningInvoiceID *string            `json:"lightning_invoice_id,omitempty"`
	CreatedAt          string             `json:"created_at"`
}
