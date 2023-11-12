package api

import (
	"context"
	"fmt"
	"ticketing_app/proto-gen/ticket"
)

// ModifyUserSeat updates a user's seat assignment.
func (s *TrainService) ModifyUserSeat(ctx context.Context, req *ticket.ModifyUserSeatRequest) (*ticket.ModifyUserSeatResponse, error) {
	userDetails, exists := s.receiptMap[req.UserId]
	if !exists {
		return nil, fmt.Errorf("user with ID '%s' not found", req.UserId)
	}

	if err := s.updateUserSection(userDetails, req.NewSection, req.NewSeatNo); err != nil {
		return nil, err
	}

	return &ticket.ModifyUserSeatResponse{Message: "User seat modified successfully"}, nil
}
func (s *TrainService) updateUserSection(userDetails *ticket.Ticket, newSection ticket.Section, newSeatNo string) error {
	// Remove user from old section
	s.removeFromSectionMap(userDetails.Section, userDetails.SeatNo, userDetails.UserId)

	// Check if the new seat is available before assigning
	sectionMap := s.getSectionMap(newSection)
	if _, taken := sectionMap[newSeatNo]; taken {
		return fmt.Errorf("seat %s in section %s is already taken", newSeatNo, newSection)
	}

	// Add user to the new section
	s.addToSectionMap(newSection, newSeatNo, userDetails.UserId)

	// Update the user's section and seat information
	userDetails.Section = newSection
	userDetails.SeatNo = newSeatNo

	return nil
}

func (s *TrainService) removeFromSectionMap(section ticket.Section, seatNo string, userId string) {
	sectionMap := s.getSectionMap(section)
	if users, ok := sectionMap[seatNo]; ok {
		for i, id := range users {
			if id == userId {
				sectionMap[seatNo] = append(users[:i], users[i+1:]...)
				break
			}
		}
	}
}

func (s *TrainService) addToSectionMap(section ticket.Section, seatNo string, userId string) {
	sectionMap := s.getSectionMap(section)
	sectionMap[seatNo] = append(sectionMap[seatNo], userId)
}
