package api

import (
	"context"
	"fmt"
	"ticketing_app/proto-gen/ticket"
)

// GetReceiptDetails retrieves receipt details for a given user.
func (s *TrainService) GetReceiptDetails(ctx context.Context, req *ticket.ReceiptDetailsRequest) (*ticket.ReceiptDetailsResponse, error) {
	userDetails, found := s.receiptMap[req.UserId]
	if !found {
		return nil, fmt.Errorf("user with ID '%s' not found", req.UserId)
	}

	receipt := &ticket.Receipt{
		From:          userDetails.From,
		To:            userDetails.To,
		UserFirstName: userDetails.UserFirstName,
		UserLastName:  userDetails.UserLastName,
		UserEmail:     userDetails.UserEmail,
		Price:         "$20", // This should ideally not be hard-coded.
		SeatNo:        userDetails.SeatNo,
		Section:       userDetails.Section,
	}

	return &ticket.ReceiptDetailsResponse{Receipt: receipt}, nil
}
