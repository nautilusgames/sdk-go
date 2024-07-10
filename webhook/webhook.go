package webhook

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/nautilusgames/sdk-go/builder"
)

const (
	_xApiKey            = "x-api-key"
	_xTenantID          = "x-tenant-id"
	_xTenantPlayerToken = "x-tenant-player-token"
	_xGameID            = "x-game-id"

	_verifyPlayer = "/player/verify"
	_walletGet    = "/wallet/get"
	_walletCreate = "/wallet/create"
	_walletBet    = "/wallet/bet"
	_walletPayout = "/wallet/payout"
	_walletRefund = "/wallet/refund"
)

type (
	// player
	VerifyPlayer func(ctx context.Context, request *VerifyPlayerRequest) (*VerifyPlayerResponse, error)

	// wallet
	GetWallet func(ctx context.Context, request *GetWalletRequest) (*GetWalletResponse, error)
	Bet       func(ctx context.Context, request *BetRequest) (*WalletResponse, error)
	Payout    func(ctx context.Context, request *PayoutRequest) (*WalletResponse, error)
)

func HandleVerifyPlayer(mux *mux.Router, logger *zap.Logger, handler VerifyPlayer) {
	mux.HandleFunc(_verifyPlayer, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute verifyPlayer

		request := &VerifyPlayerRequest{}
		headerRequest := readHeader(r)
		request.Header = headerRequest
		response, err := handler(r.Context(), request)
		if err != nil {
			fmt.Printf("Route not found: %s %s\n", r.Method, r.URL)

			logger.Error(err.Error(), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		builder.SendResponse(w, response)
	})
}

func HandleGetWallet(mux *mux.Router, logger *zap.Logger, handler GetWallet) {
	mux.HandleFunc(_walletGet, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute getWallet

		request := &GetWalletRequest{}
		errorResponse := &GetWalletResponse{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			logger.Error(err.Error(), zap.Error(err))
			errorResponse.Error.Message = err.Error()
			builder.SendResponse(w, errorResponse)
			return
		}
		request.Header = headerRequest
		response, err := handler(r.Context(), request)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
			errorResponse.Error.Message = err.Error()
			builder.SendResponse(w, errorResponse)
			return
		}
		builder.SendResponse(w, response)
	})
}

func HandleBet(mux *mux.Router, logger *zap.Logger, handler Bet) {
	mux.HandleFunc(_walletBet, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute bet

		var err error
		request := &BetRequest{}
		response := &WalletResponse{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			logger.Error(err.Error(), zap.Error(err))
			response.Error.Message = err.Error()
			builder.SendResponse(w, response)
			return
		}
		request.Header = headerRequest
		response, err = handler(r.Context(), request)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
			response.Error.Message = err.Error()
			builder.SendResponse(w, response)
			return
		}
		builder.SendResponse(w, response)
	})
}

func HandlePayout(mux *mux.Router, logger *zap.Logger, handler Payout) {
	mux.HandleFunc(_walletPayout, func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute payout

		var err error
		request := &PayoutRequest{}
		response := &WalletResponse{
			Error: &Error{
				Code: http.StatusInternalServerError,
			},
		}
		headerRequest := readHeader(r)
		if err := builder.ToRequest(r.Body, request); err != nil {
			logger.Error(err.Error(), zap.Error(err))
			response.Error.Message = err.Error()
			builder.SendResponse(w, response)
			return
		}
		request.Header = headerRequest
		response, err = handler(r.Context(), request)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
			response.Error.Message = err.Error()
			builder.SendResponse(w, response)
			return
		}
		builder.SendResponse(w, response)
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
