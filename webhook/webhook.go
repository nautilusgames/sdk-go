package webhook

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nautilusgames/sdk-go/builder"
)

var (
	_xApiKey            = []string{"x-api-key", "x-tenant-secret", "t-secret"}
	_xTenantID          = []string{"x-tenant-id", "t-id"}
	_xTenantPlayerToken = []string{"x-tenant-player-token", "x-tenant-token", "t-token"}
	_xGameID            = []string{"x-game-id", "g-id"}

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
		headerRequest := readHeader(r)
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
		// handler read request & call func execute getWallet

		request := &GetWalletRequest{}
		headerRequest := readHeader(r)
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
		// handler read request & call func execute bet

		request := &TransactionRequest{}
		headerRequest := readHeader(r)
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
		// handler read request & call func execute payout

		request := &TransactionRequest{}
		headerRequest := readHeader(r)
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
		// handler read request & call func execute payout

		request := &TransactionRequest{}
		headerRequest := readHeader(r)
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
		// handler read request & call func execute payout

		request := &TransactionRequest{}
		headerRequest := readHeader(r)
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

func readHeader(r *http.Request) *HookRequestHeader {
	header := &HookRequestHeader{}
	header.XApiKey = extractHeader(r, _xApiKey...)
	header.XTenantId = extractHeader(r, _xTenantID...)
	header.XTenantPlayerToken = extractHeader(r, _xTenantPlayerToken...)
	header.XGameId = extractHeader(r, _xGameID...)
	return header
}

func extractHeader(r *http.Request, names ...string) string {
	for _, name := range names {
		value := r.Header.Get(name)
		if len(value) > 0 {
			return value
		}
	}

	return ""
}
