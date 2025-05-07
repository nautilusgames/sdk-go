package webhook

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nautilusgames/sdk-go/builder"
)

const (
	Authorization      = "Authorization"
	XApiKey            = "x-api-key"
	XTenantId          = "x-tenant-id"
	XTenantPlayerToken = "x-tenant-player-token"
	XTenantPlayerId    = "x-tenant-player-id"
	XGameId            = "x-game-id"
	XRequestId         = "x-request-id"
)

const (
	_verifyPlayer   = "/player/verify"
	_walletGet      = "/wallet/get"
	_walletBet      = "/wallet/bet"
	_walletPayout   = "/wallet/payout"
	_walletRefund   = "/wallet/refund"
	_walletRollback = "/wallet/rollback"
	_walletCancel   = "/wallet/cancel"
)

type (
	// player
	VerifyPlayer func(ctx context.Context, request *VerifyPlayerRequest) (*VerifyPlayerReply, error)

	// wallet
	GetWallet func(ctx context.Context, request *GetWalletRequest) (*GetWalletReply, error)
	Bet       func(ctx context.Context, request *BetRequest) (*TransactionReply, error)
	Payout    func(ctx context.Context, request *PayoutRequest) (*TransactionReply, error)
	Refund    func(ctx context.Context, request *RefundRequest) (*TransactionReply, error)
	Rollback  func(ctx context.Context, request *RollbackRequest) (*TransactionReply, error)
	Cancel    func(ctx context.Context, request *CancelRequest) (*TransactionReply, error)
)

func HandleVerifyPlayer(mux *mux.Router, handler VerifyPlayer) {
	mux.HandleFunc(_verifyPlayer, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute verifyPlayer

		request := &VerifyPlayerRequest{}
		headerRequest := buildRequestHeaders(r)
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleGetWallet(mux *mux.Router, handler GetWallet) {
	mux.HandleFunc(_walletGet, func(w http.ResponseWriter, r *http.Request) {
		request := &GetWalletRequest{}
		headerRequest := buildRequestHeaders(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleBet(mux *mux.Router, handler Bet) {
	mux.HandleFunc(_walletBet, func(w http.ResponseWriter, r *http.Request) {
		request := &BetRequest{}
		headerRequest := buildRequestHeaders(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandlePayout(mux *mux.Router, handler Payout) {
	mux.HandleFunc(_walletPayout, func(w http.ResponseWriter, r *http.Request) {
		request := &PayoutRequest{}
		headerRequest := buildRequestHeaders(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleRefund(mux *mux.Router, handler Refund) {
	mux.HandleFunc(_walletRefund, func(w http.ResponseWriter, r *http.Request) {
		request := &RefundRequest{}
		headerRequest := buildRequestHeaders(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleRollback(mux *mux.Router, handler Rollback) {
	mux.HandleFunc(_walletRollback, func(w http.ResponseWriter, r *http.Request) {
		request := &RollbackRequest{}
		headerRequest := buildRequestHeaders(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleCancel(mux *mux.Router, handler Cancel) {
	mux.HandleFunc(_walletCancel, func(w http.ResponseWriter, r *http.Request) {
		request := &CancelRequest{}
		headerRequest := buildRequestHeaders(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			var httpErr *HTTPError
			if errors.As(err, &httpErr) {
				http.Error(w, httpErr.Message, httpErr.StatusCode)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		builder.SendReply(w, reply)
	})
}

func buildRequestHeaders(r *http.Request) *HookRequestHeader {
	return &HookRequestHeader{
		XApiKey:            r.Header.Get(XApiKey),
		XTenantId:          r.Header.Get(XTenantId),
		XTenantPlayerToken: r.Header.Get(XTenantPlayerToken),
		XTenantPlayerId:    r.Header.Get(XTenantPlayerId),
		XGameId:            r.Header.Get(XGameId),
		XRequestId:         r.Header.Get(XRequestId),
	}
}
