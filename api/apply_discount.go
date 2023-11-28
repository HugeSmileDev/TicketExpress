package api

import (
	"context"
	"fmt"
	ticketingapp "ticketing_app/proto-gen/ticket"
)

// ApplyDiscount is a function to apply a discount.
func (s *TrainService) ApplyDiscount(ctx context.Context, req *ticketingapp.DiscountRequest) (*ticketingapp.DiscountResponse, error) {
	updatedReceipt, found := s.receiptMap[req.ReceiptId]
	if !found {
		return nil, fmt.Errorf("receipt with ID '%s' not found", req.ReceiptId)
	}

	discountPrice := getDiscountPrice(req.DiscountCode)
	updatedReceipt.PricePaid -= discountPrice

	// Save changes
	s.receiptMap[req.ReceiptId] = updatedReceipt

	return &ticketingapp.DiscountResponse{UpdatedReceipt: updatedReceipt}, nil
}

// getDiscountPrice returns the discount amount based on the discount code.
func getDiscountPrice(discountCode string) float32 {
	switch discountCode {
	case "TBD123":
		return 1.0
	case "WOW1":
		return 2.0
	case "Test3":
		return 5.0
	default:
		return 0.0
	}
}
