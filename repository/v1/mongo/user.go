package mongo

import (
	v1 "awesomeProject/core/v1"
	"awesomeProject/core/v1/repository"
	"awesomeProject/enums"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// UserCollection collection name
var (
	UserCollection = "user_collection"
)

type userRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (u userRepository) GetByEmail(email string) v1.User {
	var res v1.User
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"email": email}}
	query["$and"] = and
	coll := u.manager.Db.Collection(UserCollection)
	result, err := coll.Find(u.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
		return v1.User{}
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.User)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

func (u userRepository) Store(user v1.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	user.Password = string(hashedPassword)
	user.Status = enums.ACTIVE
	coll := u.manager.Db.Collection(UserCollection)
	_, err = coll.InsertOne(u.manager.Ctx, user)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (u userRepository) Get() []v1.User {
	var results []v1.User
	query := bson.M{
		"$or": []bson.M{
			{"status": enums.ACTIVE},
			{"status": enums.INACTIVE},
		},
	}
	coll := u.manager.Db.Collection(UserCollection)
	result, err := coll.Find(u.manager.Ctx, query, &options.FindOptions{
		Sort: bson.M{"created_date": -1},
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.User)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func NewUserRepository(timeout int) repository.User {
	return &userRepository{
		manager: GetDmManager(),
		timeout: 0,
	}
}
