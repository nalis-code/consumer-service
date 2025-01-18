package internal

import (
	"context"
	"fmt"
	"time"
)

type PriceValue struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type Handler struct {
	timeGen string
}

func NewHandler() *Handler {
	h := &Handler{
		timeGen: time.Now().String(),
	}
	return h
}

func (h *Handler) Handle(ctx context.Context, price PriceValue) error {

	// PART 1 --- VALIDATION

	err := h.validatePrice(price)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	//PART 2 --- TRANSFORMATION

	//newModel, err := h.transform(ctx, p)
	//if err != nil {
	//	return fmt.Errorf("could not transform to new model: %w", err)
	//}

	// PART 3 --- POST

	//err = h.poster.Post(ctx, newModel)
	//if err != nil {
	//	return fmt.Errorf("could not send data: %w", err)
	//}

	return nil
}

func (h *Handler) validatePrice(price PriceValue) error {

	//Mandatory
	if price.Amount == 0 {
		return fmt.Errorf("amount can not be 0")
	}

	if price.Currency == "" {
		return fmt.Errorf("currency code can not be empty")
	}

	if !isValidCurrency(price.Currency) {
		return fmt.Errorf("wrong currency value, it can only be GBP or EUR or USD or AUD")
	}

	return nil
}

// Check if a given string is a valid Price enum value
func isValidCurrency(value string) bool {
	switch value {
	case "GBP", "AUD", "USD", "EUR":
		return true
	}
	return false
}
