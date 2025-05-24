package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Username string             `json:"username"`
	jwt.RegisteredClaims
}
