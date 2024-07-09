package model

type HookRequestHeader struct {
	XApiKey      string `json:"x_api_key,omitempty"`      // x-api-key
	XTenantId    string `json:"x_tenant_id,omitempty"`    // x-tenant-id
	XTenantToken string `json:"x_tenant_token,omitempty"` // x-tenant-token
	XGameId      string `json:"x_game_id,omitempty"`      // x-game-id
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type VerifyPlayerRequest struct {
	Header *HookRequestHeader `json:"header"`
}

type VerifyPlayerResponse struct {
	Data *PlayerInfo `json:"data"`
}

type PlayerInfo struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type BetRequest struct {
	SessionId int64              `json:"session_id"`
	Amount    int64              `json:"amount"`
	Header    *HookRequestHeader `json:"header,omitempty"`
}

type PayoutRequest struct {
	SessionId int64              `json:"session_id"`
	Amount    int64              `json:"amount"`
	Header    *HookRequestHeader `json:"header,omitempty"`
}

type GetWalletRequest struct {
	Header *HookRequestHeader `json:"header"`
}

type GetWalletResponse struct {
	Data  *PlayerWallet `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
}

type PlayerWallet struct {
	Balance  int64 `json:"balance"`
	LastTxId int64 `son:"last_tx_id"`
}

type WalletResponse struct {
	Data  *WalletTransaction `json:"data"`
	Error *Error             `json:"error"`
}

type WalletTransaction struct {
	Id         int64 `json:"id"`
	SessionId  int64 `json:"session_id"`
	Amount     int64 `json:"amount"`
	NewBalance int64 `json:"new_balance"`
	CreatedAt  int64 `json:"created_at"`
}
