package webhook

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nautilusgames/sdk-go/builder"
)

var (
	_xApiKey            = "x-api-key"
	_xTenantID          = "x-tenant-id"
	_xTenantPlayerToken = "x-tenant-player-token"
	_xGameID            = "x-game-id"

	_verifyPlayer   = "/player/verify"
	_walletGet      = "/wallet/get"
	_walletBet      = "/wallet/bet"
	_walletPayout   = "/wallet/payout"
	_walletRefund   = "/wallet/refund"
	_walletRollback = "/wallet/rollback"
)

type (
	// player
	VerifyPlayer func(ctx context.Context, request *VerifyPlayerRequest) (*VerifyPlayerReply, error)

	// wallet
	GetWallet func(ctx context.Context, request *GetWalletRequest) (*GetWalletReply, error)
	Bet       func(ctx context.Context, request *TransactionRequest) (*TransactionReply, error)
	Payout    func(ctx context.Context, request *TransactionRequest) (*TransactionReply, error)
	Refund    func(ctx context.Context, request *TransactionRequest) (*TransactionReply, error)
	Rollback  func(ctx context.Context, request *TransactionRequest) (*TransactionReply, error)
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
		request := &TransactionRequest{}
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
		request := &TransactionRequest{}
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
		request := &TransactionRequest{}
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
		request := &TransactionRequest{}
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
		XApiKey:            r.Header.Get(_xApiKey),
		XTenantId:          r.Header.Get(_xTenantID),
		XTenantPlayerToken: r.Header.Get(_xTenantPlayerToken),
		XGameId:            r.Header.Get(_xGameID),
	}
}
