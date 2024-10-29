package friend

import (
	"github.com/muliCohenGini/go-api/internal/db"
)

type FriendRepository interface {
	addFriendRequest(friend Friend) (bool, error)
}

type Repository struct {
}

func NewFriendRepository() FriendRepository {
	return &Repository{}
}

func (r *Repository) addFriendRequest(friend Friend) (bool, error) {
	if friend.Status == "" {
		friend.Status = "pending"
	}
	query := `
		INSERT INTO "FriendRequests" ("senderId", "recipientId", "status", "community")
		VALUES ($1, $2, $3, $4);
	`
	_, err := db.DB.Exec(query, friend.SenderId, friend.RecipientId, friend.Status, friend.Community)
	if err != nil {
		return false, err
	}
	return true, nil
}
