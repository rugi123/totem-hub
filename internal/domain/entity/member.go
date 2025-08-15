package entity

import "github.com/google/uuid"

type Member struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	ChatID  uuid.UUID
	Role    string
	IsMuted bool
}

func ExtractMemberIDs(members []Member) []uuid.UUID {
	var IDs []uuid.UUID
	for _, member := range members {
		IDs = append(IDs, member.ID)
	}
	return IDs
}
