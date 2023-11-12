package api

import (
	"ticketing_app/proto-gen/ticket"
)

type TicketServiceServer interface {
	ticket.TicketServiceServer
}

// TrainService represents a service for managing train tickets.
type TrainService struct {
	ticket.UnimplementedTicketServiceServer
	receiptMap map[string]*ticket.Ticket
	sectionA   map[string][]string
	sectionB   map[string][]string
}

// NewTrainService initializes and returns a new TrainService.
func NewTicketServiceServer() TicketServiceServer {
	return &TrainService{
		receiptMap: make(map[string]*ticket.Ticket),
		sectionA:   make(map[string][]string),
		sectionB:   make(map[string][]string),
	}
}
