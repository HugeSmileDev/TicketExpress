package api

import (
	"context"
	"fmt"
	"ticketing_app/proto-gen/ticket"
)

// GetUserSeats retrieves user seat assignments for a specified section.
func (s *TrainService) GetUserSeats(ctx context.Context, req *ticket.UserSeatsRequest) (*ticket.UserSeatsResponse, error) {
	var sectionMap map[string][]string
	switch req.Section {
	case ticket.Section_SECTION_A:
		sectionMap = s.sectionA
	case ticket.Section_SECTION_B:
		sectionMap = s.sectionB
	default:
		return nil, fmt.Errorf("invalid section: %v", req.Section)
	}

	seatAssignments := make([]*ticket.UserSeatsResponse_SeatAssignment, 0)
	for seatNo, users := range sectionMap {
		for _, userId := range users {
			if userDetails, ok := s.receiptMap[userId]; ok && userDetails.SeatNo == seatNo {
				seatAssignments = append(seatAssignments, &ticket.UserSeatsResponse_SeatAssignment{
					UserId:   userId,
					UserSeat: seatNo,
				})
			}
		}
	}

	return &ticket.UserSeatsResponse{SeatAssignments: seatAssignments}, nil
}
