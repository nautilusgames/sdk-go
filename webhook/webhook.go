package webhook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nautilusgames/sdk-go/builder"
)

const (
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
		headerRequest := readHeader(r)
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleGetWallet(mux *mux.Router, handler GetWallet) {
	mux.HandleFunc(_walletGet, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute getWallet

		request := &GetWalletRequest{}
		errorResponse := &GetWalletReply{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		request.Header = headerRequest
		reply, err := handler(r.Context(), request)
		if err != nil {
			errorResponse.Error.Message = err.Error()
			builder.SendReply(w, errorResponse)
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleBet(mux *mux.Router, handler Bet) {
	mux.HandleFunc(_walletBet, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute bet

		var err error
		request := &TransactionRequest{}
		reply := &TransactionReply{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		request.Header = headerRequest
		reply, err = handler(r.Context(), request)
		if err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandlePayout(mux *mux.Router, handler Payout) {
	mux.HandleFunc(_walletPayout, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute payout

		var err error
		request := &TransactionRequest{}
		reply := &TransactionReply{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		request.Header = headerRequest
		reply, err = handler(r.Context(), request)
		if err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleRefund(mux *mux.Router, handler Refund) {
	mux.HandleFunc(_walletRefund, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute payout

		var err error
		request := &TransactionRequest{}
		reply := &TransactionReply{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		request.Header = headerRequest
		reply, err = handler(r.Context(), request)
		if err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		builder.SendReply(w, reply)
	})
}

func HandleRollback(mux *mux.Router, handler Rollback) {
	mux.HandleFunc(_walletRollback, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute payout

		var err error
		request := &TransactionRequest{}
		reply := &TransactionReply{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		request.Header = headerRequest
		reply, err = handler(r.Context(), request)
		if err != nil {
			reply.Error.Message = err.Error()
			builder.SendReply(w, reply)
			return
		}
		builder.SendReply(w, reply)
	})
}

func readHeader(r *http.Request) *HookRequestHeader {
	header := &HookRequestHeader{}
	header.XApiKey = r.Header.Get(_xApiKey)
	header.XTenantId = r.Header.Get(_xTenantID)
	header.XTenantPlayerToken = r.Header.Get(_xTenantPlayerToken)
	header.XGameId = r.Header.Get(_xGameID)
	return header
}
