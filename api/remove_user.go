package api

import (
	"context"
	"errors"
	"ticketing_app/proto-gen/ticket"
)

// RemoveUser removes a user from the system.
func (s *TrainService) RemoveUser(ctx context.Context, req *ticket.RemoveUserRequest) (*ticket.RemoveUserResponse, error) {
	userDetails, exists := s.receiptMap[req.UserId]
	if !exists {
		return nil, errors.New("user not found")
	}

	// Remove the user from the section map
	if err := s.removeFromSection(userDetails); err != nil {
		return nil, err
	}

	// Remove the user from the receipt map
	delete(s.receiptMap, req.UserId)

	return &ticket.RemoveUserResponse{Message: "User removed successfully"}, nil
}

func (s *TrainService) removeFromSection(userDetails *ticket.Ticket) error {
	sectionMap := s.getSectionMap(userDetails.Section)
	if sectionMap == nil {
		return errors.New("invalid section")
	}

	if _, ok := sectionMap[userDetails.SeatNo]; ok {
		delete(sectionMap, userDetails.SeatNo)
		return nil
	}

	return errors.New("user seat not found in section")
}
