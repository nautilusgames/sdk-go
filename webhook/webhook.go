package webhook

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/nautilusgames/sdk-go/builder"
	"github.com/nautilusgames/sdk-go/model"
)

const (
	_x_api_key      = "x-api-key"
	_x_tenant_id    = "x-tenant-id"
	_x_tenant_token = "x-tenant-token"
	_x_game_id      = "x-game-id"

	_verifyPlayer = "/player/verify"
	_walletGet    = "/wallet/get"
	_walletCreate = "/wallet/create"
	_walletBet    = "/wallet/bet"
	_walletPayout = "/wallet/payout"
	_walletRefund = "/wallet/refund"
)

type (
	// player
	VerifyPlayer func(ctx context.Context, request *model.VerifyPlayerRequest) (*model.VerifyPlayerResponse, error)

	// wallet
	GetWallet func(ctx context.Context, request *model.GetWalletRequest) (*model.GetWalletResponse, error)
	Bet       func(ctx context.Context, request *model.BetRequest) (*model.WalletResponse, error)
	Payout    func(ctx context.Context, request *model.PayoutRequest) (*model.WalletResponse, error)
)

func HandleVerifyPlayer(logger *zap.Logger, handler VerifyPlayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute verifyPlayer

		request := &model.VerifyPlayerRequest{}
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
	}
}

func HandleGetWallet(logger *zap.Logger, handler GetWallet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute getWallet

		request := &model.GetWalletRequest{}
		errorResponse := &model.GetWalletResponse{
			Error: &model.Error{
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
	}
}

func HandleBet(logger *zap.Logger, handler Bet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute bet

		var err error
		request := &model.BetRequest{}
		response := &model.WalletResponse{
			Error: &model.Error{
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
	}
}

func HandlePayout(handler Payout, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute payout

		var err error
		request := &model.PayoutRequest{}
		response := &model.WalletResponse{
			Error: &model.Error{
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
	}
}

func readHeader(r *http.Request) *model.HookRequestHeader {
	header := &model.HookRequestHeader{}
	header.XApiKey = r.Header.Get(_x_api_key)
	header.XTenantId = r.Header.Get(_x_tenant_id)
	header.XTenantToken = r.Header.Get(_x_tenant_token)
	header.XGameId = r.Header.Get(_x_game_id)
	return header
}
