package webhook

type HookRequestHeader struct {
	XApiKey            string `json:"x_api_key,omitempty"`             // x-api-key
	XTenantId          string `json:"x_tenant_id,omitempty"`           // x-tenant-id
	XTenantPlayerToken string `json:"x_tenant_player_token,omitempty"` // x-tenant-player-token
	XGameId            string `json:"x_game_id,omitempty"`             // x-game-id
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type VerifyPlayerRequest struct {
	Header *HookRequestHeader `json:"header"`
}

type VerifyPlayerReply struct {
	Data  *PlayerInfo `json:"data"`
	Error *Error      `json:"error"`
}

type PlayerInfo struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type GetWalletRequest struct {
	Header *HookRequestHeader `json:"header"`
}

type GetWalletReply struct {
	Data  *PlayerWallet `json:"data"`
	Error *Error        `json:"error"`
}

type BetRequest struct {
	Header *HookRequestHeader `json:"header,omitempty"`

	TxId      string  `json:"tx_id"`
	RefTxId   string  `json:"ref_tx_id"`
	SessionId string  `json:"session_id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
}

type RefundRequest struct {
	Header *HookRequestHeader `json:"header,omitempty"`

	TxId      string  `json:"tx_id"`
	RefTxId   string  `json:"ref_tx_id"`
	SessionId string  `json:"session_id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
}

type RollbackRequest struct {
	Header *HookRequestHeader `json:"header,omitempty"`

	TxId      string  `json:"tx_id"`
	RefTxId   string  `json:"ref_tx_id"`
	BetTxId   string  `json:"bet_tx_id"`
	SessionId string  `json:"session_id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
}

type PayoutRequest struct {
	Header *HookRequestHeader `json:"header,omitempty"`

	TxId      string  `json:"tx_id"`
	RefTxId   string  `json:"ref_tx_id"`
	SessionId string  `json:"session_id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
	// Possible values: [SETTLE, FEATURE, JACKPOT, BONUS_WITHDRAWAL, FREE_WITHDRAWAL]
	PayoutType string `json:"payout_type"`
}

type TransactionReply struct {
	Data  *TransactionData `json:"data"`
	Error *Error           `json:"error"`
}

type TransactionData struct {
	TenantTxId      string  `json:"tenant_tx_id"`
	TenantSessionId string  `json:"tenant_session_id"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	NewBalance      float64 `json:"new_balance"`
	CreatedAt       int64   `json:"created_at"`
}

type PlayerWallet struct {
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance"`
	UpdatedAt int64   `json:"updated_at"`
}
