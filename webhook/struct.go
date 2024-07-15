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
	Data *PlayerInfo `json:"data"`
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

type TransactionRequest struct {
	Currency  string             `json:"currency"`
	SessionId int64              `json:"session_id"`
	Amount    float64            `json:"amount"`
	Header    *HookRequestHeader `json:"header,omitempty"`
}

type TransactionReply struct {
	Data  *TransactionData `json:"data"`
	Error *Error           `json:"error"`
}

type TransactionData struct {
	Id         int64   `json:"id"`
	Currency   string  `json:"currency"`
	SessionId  int64   `json:"session_id"`
	Amount     float64 `json:"amount"`
	NewBalance float64 `json:"new_balance"`
	CreatedAt  int64   `json:"created_at"`
}

type PlayerWallet struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
	LastTxId int64   `json:"last_tx_id"`
}
