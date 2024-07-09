package webhook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

type HttpServer interface {
	// add request or header to parameters/ handler response or error
	VerifyPlayer(ctx context.Context, request *model.VerifyPlayerRequest) (*model.VerifyPlayerResponse, error)
	GetWallet(ctx context.Context, request *model.GetWalletRequest) (*model.GetWalletResponse, error)
	// CreateWallet(ctx context.Context, request string) error // does not exist?
	Bet(ctx context.Context, request *model.BetRequest) (*model.WalletResponse, error)
	Payout(ctx context.Context, request *model.PayoutRequest) (*model.WalletResponse, error)
	// Refund(ctx context.Context, request string) error // TODO
}

type UnimplementedHttpServer struct{}

func (UnimplementedHttpServer) VerifyPlayer(_ context.Context, _ *model.VerifyPlayerRequest) (*model.VerifyPlayerResponse, error) {
	fmt.Println("method VerifyPlayer not implemented")
	return nil, nil
}

func (UnimplementedHttpServer) GetWallet(_ context.Context, _ *model.GetWalletRequest) (*model.GetWalletResponse, error) {
	fmt.Println("method GetWallet not implemented")
	return nil, nil
}

func (UnimplementedHttpServer) Bet(_ context.Context, _ *model.BetRequest) (*model.WalletResponse, error) {
	fmt.Println("method Bet not implemented")
	return nil, nil
}

func (UnimplementedHttpServer) Payout(_ context.Context, _ *model.PayoutRequest) (*model.WalletResponse, error) {
	fmt.Println("method Payout not implemented")
	return nil, nil
}

func RegisterHttpServer(server HttpServer, logger *zap.Logger, httpPathPrefix string) *mux.Router {
	router := mux.NewRouter()
	sub := router.PathPrefix(httpPathPrefix).Subrouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Error("Route not found: %s %s\n", zap.String("method", r.Method), zap.Any("url", r.URL))
	})

	sub.HandleFunc(_verifyPlayer, verifyPlayer(server, logger)).Methods(http.MethodPost)
	sub.HandleFunc(_walletGet, getWallet(server, logger)).Methods(http.MethodPost)
	sub.HandleFunc(_walletBet, bet(server, logger)).Methods(http.MethodPost)
	sub.HandleFunc(_walletPayout, payout(server, logger)).Methods(http.MethodPost)

	return router
}

func verifyPlayer(server HttpServer, logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// handler read request & call func execute verifyPlayer

		request := &model.VerifyPlayerRequest{}
		headerRequest := readHeader(r)
		request.Header = headerRequest
		response, err := server.VerifyPlayer(r.Context(), request)
		if err != nil {
			fmt.Printf("Route not found: %s %s\n", r.Method, r.URL)

			logger.Error(err.Error(), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		builder.SendResponse(w, response)
	}
}

func getWallet(server HttpServer, logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
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
		response, err := server.GetWallet(r.Context(), request)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
			errorResponse.Error.Message = err.Error()
			builder.SendResponse(w, errorResponse)
			return
		}
		builder.SendResponse(w, response)
	}
}

func bet(server HttpServer, logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
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
		response, err = server.Bet(r.Context(), request)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
			response.Error.Message = err.Error()
			builder.SendResponse(w, response)
			return
		}
		builder.SendResponse(w, response)
	}
}

func payout(server HttpServer, logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
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
		response, err = server.Payout(r.Context(), request)
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
