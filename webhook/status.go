package webhook

const (
	ErrUnknown                 = 1000
	ErrInvalidRequest          = 1034
	ErrTenantFailed            = 1035
	ErrInternalServerError     = 1200
	ErrInvalidTenantID         = 1204
	ErrInvalidTenantApiKey     = 1205
	ErrInvalidPlayerSession    = 1300
	ErrEmptyPlayerSession      = 1301
	ErrServerErrorOccurs       = 1303
	ErrInvalidPlayer           = 1305
	ErrPlayerBlocked           = 1306
	ErrExpiredPlayerSession    = 1308
	ErrPlayerInactive          = 1309
	ErrVerifyFailed            = 1310
	ErrPlayerInProgress        = 1315
	ErrGameIsUnderMaintenance  = 1400
	ErrGameIsInactive          = 1401
	ErrInvalidGameID           = 1402
	ErrValueCannotBeNull       = 3001
	ErrPlayerDoesNotExist      = 3004
	ErrNoBetExists             = 3021
	ErrBetAlreadyPayout        = 3022
	ErrBetAlreadyExisted       = 3032
	ErrBetFailed               = 3033
	ErrPayoutFailed            = 3034
	ErrTransactionDoesNotExist = 3040
	ErrInvalidCurrency         = 3206
	ErrInsufficient            = 3202
)
