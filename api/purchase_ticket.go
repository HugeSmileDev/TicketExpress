package api

import (
	"context"
	"errors"
	"ticketing_app/proto-gen/ticket"
)

// PurchaseTicket processes a ticket purchase request.
func (s *TrainService) PurchaseTicket(ctx context.Context, req *ticket.PurchaseTicketRequest) (*ticket.PurchaseTicketResponse, error) {
	if err := validatePurchaseTicketRequest(req); err != nil {
		return nil, err
	}

	reqTicket := req.GetTicket()

	if _, exists := s.receiptMap[reqTicket.UserId]; exists {
		return nil, errors.New("user with the same ID already exists")
	}

	if err := s.assignSeat(reqTicket.UserId, reqTicket.Section, reqTicket.SeatNo); err != nil {
		return nil, err
	}

	s.receiptMap[reqTicket.UserId] = reqTicket
	receipt := &ticket.Receipt{
		From:          reqTicket.From,
		To:            reqTicket.To,
		UserFirstName: reqTicket.UserFirstName,
		UserLastName:  reqTicket.UserLastName,
		UserEmail:     reqTicket.UserEmail,
		Price:         "$20", // This should ideally not be hard-coded.
		SeatNo:        reqTicket.SeatNo,
		Section:       reqTicket.Section,
	}

	return &ticket.PurchaseTicketResponse{Receipt: receipt}, nil
}

func validatePurchaseTicketRequest(req *ticket.PurchaseTicketRequest) error {
	reqTicket := req.GetTicket()
	if reqTicket.From == "" || reqTicket.To == "" || reqTicket.UserFirstName == "" || reqTicket.UserLastName == "" ||
		reqTicket.UserEmail == "" || reqTicket.UserId == "" || reqTicket.Section == ticket.Section_SECTION_UNSPECIFIED || reqTicket.SeatNo == "" {
		return errors.New("all fields are required11")
	}
	return nil
}

func (s *TrainService) assignSeat(userID string, section ticket.Section, seatNo string) error {
	sectionMap := s.getSectionMap(section)
	if sectionMap == nil {
		return errors.New("invalid section")
	}

	if seats, exists := sectionMap[seatNo]; exists && len(seats) > 0 {
		return errors.New("seat already taken")
	}

	sectionMap[seatNo] = []string{userID}
	return nil
}
