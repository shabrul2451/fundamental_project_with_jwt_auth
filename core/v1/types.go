package v1

import (
	"awesomeProject/enums"
	"crypto/rsa"
	"time"
)

// User holds users info.
type User struct {
	Id          string          `json:"id" bson:"id"`
	FirstName   string          `json:"first_name" bson:"first_name" `
	LastName    string          `json:"last_name" bson:"last_name"`
	Email       string          `json:"email" bson:"email" `
	Phone       string          `json:"phone" bson:"phone" `
	Password    string          `json:"password" bson:"password" `
	Status      enums.STATUS    `json:"status" bson:"status"`
	CreatedDate time.Time       `json:"created_date" bson:"created_date"`
	UpdatedDate time.Time       `json:"updated_date" bson:"updated_date"`
	CreatedBy   string          `json:"created_by" bson:"created_by"`
	Role        enums.ROLE      `json:"role" bson:"role"`
	AuthType    enums.AUTH_TYPE `json:"auth_type" bson:"auth_type"`
}

// RsaKeys contains RSA keys.
type RsaKeys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// LoginDto contains user login info.
type LoginDto struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserResourcePermissionDto struct {
	UserId string     `json:"user_id" bson:"user_id"`
	Email  string     `json:"email" bson:"email"`
	Role   enums.ROLE `json:"role" bson:"role"`
}

// JWTPayLoad contains payload of JWT token.
type JWTPayLoad struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}
