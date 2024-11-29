package service

import (
	"fmt"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/paymentmethod"
	"github.com/stripe/stripe-go/v81/refund"
)

type StripeService struct {
	privateKey string
}

func NewStripeService(key string) *StripeService {
	stripe.Key = key
	return &StripeService{privateKey: key}
}

func (s *StripeService) CreateCustomer(email string, metadata map[string]string) (string, error) {
	params := &stripe.CustomerParams{
		Email:    stripe.String(email),
		Metadata: metadata,
	}

	result, err := customer.New(params)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

func (s *StripeService) AttachPaymentMethodToCustomer(customerId, paymentMethodId string) error {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerId),
	}

	_, err := paymentmethod.Attach(paymentMethodId, params)
	if err != nil {
		return err
	}

	return nil
}

func (s *StripeService) RemovePaymentMethod(paymentMethodId string) error {
	params := &stripe.PaymentMethodDetachParams{}

	_, err := paymentmethod.Detach(paymentMethodId, params)
	if err != nil {
		return err
	}

	return nil
}

func (s *StripeService) ListPaymentMethodsByCustomer(customerId string) ([]*stripe.PaymentMethod, error) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Customer: stripe.String(customerId),
	}

	params.Limit = stripe.Int64(3)

	cur := customer.ListPaymentMethods(params)
	list := cur.PaymentMethodList().Data

	return list, nil
}

func (s *StripeService) CreatePaymentIntent(amount int64, currency, paymentMethodID, returnUrl string, metadata map[string]string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Currency:      stripe.String(currency),
		PaymentMethod: stripe.String(paymentMethodID),
		Confirm:       stripe.Bool(true), // Automatically confirm the payment
		Metadata:      metadata,
		ReturnURL:     stripe.String(returnUrl),
	}

	// Create the Payment Intent
	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}

	return pi, nil
}

func (s *StripeService) GetPaymentIntent(paymentIntentId string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	if paymentIntentId == "" {
		return nil, fmt.Errorf("payment intent id cannot be empty")
	}
	if params == nil {
		params = &stripe.PaymentIntentParams{}
	}

	result, err := paymentintent.Get(paymentIntentId, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment intent: %w", err)
	}

	return result, nil
}

func (s *StripeService) RefundPayment(paymentIntentId string, amount int64) (*stripe.Refund, error) {
	// Refund the full charge if no amount is provided
	if amount == 0 {
		refundParams := &stripe.RefundParams{
			PaymentIntent: stripe.String(paymentIntentId),
		}
		return refund.New(refundParams)
	}

	// Refund a specific amount
	refundParams := &stripe.RefundParams{
		Charge: stripe.String(paymentIntentId),
		Amount: stripe.Int64(amount),
	}
	return refund.New(refundParams)
}
