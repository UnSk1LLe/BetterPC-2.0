package errors

import "github.com/pkg/errors"

type OrderError struct {
	err error
}

func (o OrderError) Error() string {
	return o.err.Error()
}

var (
	ErrOrderCancelled         = OrderError{err: errors.New("order is cancelled")}
	ErrNotActiveOrder         = OrderError{err: errors.New("order is cancelled or closed")}
	ErrOrderNotFound          = OrderError{err: errors.New("order not found")}
	ErrOrderClosed            = OrderError{err: errors.New("order is closed")}
	ErrInvalidInput           = OrderError{err: errors.New("invalid input")}
	ErrActiveOrder            = OrderError{err: errors.New("active order")}
	ErrUnsupportedOrderStatus = OrderError{err: errors.New("unsupported order status")}
	ErrOrderAlreadyPaid       = OrderError{err: errors.New("order already paid")}
	ErrOrderOwnerMismatch     = OrderError{err: errors.New("order does not belong to the user")}
)
