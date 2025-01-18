//go:build unit

package internal

import (
	"context"
	"fmt"
	"testing"
)

func TestIsValidCurrency(t *testing.T) {
	tests := []struct {
		currency string
		expected bool
	}{
		{"GBP", true},
		{"USD", true},
		{"EUR", true},
		{"AUD", true},
		{"INR", false},
		{"", false},
	}

	for _, test := range tests {
		result := isValidCurrency(test.currency)
		if result != test.expected {
			t.Errorf("isValidCurrency(%s) = %v; want %v", test.currency, result, test.expected)
		}
	}
}

func TestHandle(t *testing.T) {
	handler := NewHandler()

	tests := []struct {
		price    PriceValue
		expected error
	}{
		{PriceValue{Amount: 100, Currency: "GBP"}, nil},
		{PriceValue{Amount: 200, Currency: "USD"}, nil},
		{PriceValue{Amount: 300, Currency: "EUR"}, nil},
		{PriceValue{Amount: 400, Currency: "AUD"}, nil},
		{PriceValue{Amount: 500, Currency: "INR"}, fmt.Errorf("validation failed: wrong currency value, it can only be GBP or EUR or USD or AUD")},
	}

	for _, test := range tests {
		err := handler.Handle(context.Background(), test.price)
		if err != nil && err.Error() != test.expected.Error() {
			t.Errorf("Handle(%v) = %v; want %v", test.price, err, test.expected)
		}
	}
}
