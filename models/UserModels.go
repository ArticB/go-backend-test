package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id            primitive.ObjectID `bson:"_id"`
	first_name    *string            `json:"first_name" min="2" max="100"`
	last_name     *string            `json:"last_name" validate:"required" min="2" max="100"`
	email         *string            `json:"email" validate:"email, required"`
	password      *string            `json:"password" validate:"required" min="8"`
	phone         *string            `json:"phone" validate:"required"`
	user_type     *string            `json:"user_type" validate:"required" eq=ADMIN|eq=CREATOR|eq=USER`
	token         *string            `json:"token" `
	refresh_token *string            `json:"refresh_token"`
	created_at    time.Time          `json:"created_at"`
	updated_at    time.Time          `json:"updated_at"`
	user_id       *string            `json:"user_id"`
}
