package state

type OpError struct{ msg string }

func (err OpError) Error() string {
	return err.msg
}

var (
	ErrInvalidFeeToken     = &OpError{"feeToken is not supported"}
	ErrInvalidToken        = &OpError{"token is not supported"}
	ErrAccountNotFound     = &OpError{"account does not exist"}
	ErrInvalidAccount      = &OpError{"invalid account id"}
	ErrInvalidAuthData     = &OpError{"l1 auth data(signature) is incorrect"}
	ErrInvalidSignature    = &OpError{"signature is incorrect"}
	ErrNonceMismatch       = &OpError{"nonce mismatch"}
	ErrInsufficientBalance = &OpError{"not enough balance"}
	ErrFromAccountLocked   = &OpError{"account is locked"}
	ErrAccountIncorrect    = &OpError{"account id is incorrect"}
	ErrAccountIdTooBig     = &OpError{"account id is bigger than max limit"}
)
