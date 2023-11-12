package api

import (
	"ticketing_app/proto-gen/ticket"
)

func (s *TrainService) getSectionMap(section ticket.Section) map[string][]string {
	switch section {
	case ticket.Section_SECTION_A:
		return s.sectionA
	case ticket.Section_SECTION_B:
		return s.sectionB
	default:
		return nil
	}
}
