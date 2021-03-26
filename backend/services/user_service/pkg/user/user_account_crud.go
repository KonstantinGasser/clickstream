package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// UserItem is a representation of a user document which lives in the
// mongoDB. Fields must be exported in order to serve as (de-)serialization for the mongoDB
type UserItem struct {
	UUID          string `bson:"_id"`
	Username      string `bson:"username"`
	Password      string `bson:"password"`
	FirstName     string `bson:"first_name"`
	LastName      string `bson:"last_name"`
	OrgnDomain    string `bson:"orgn_domain"`
	OrgnPosition  string `bson:"orgn_position"`
	ProfileImgURL string `bson:"profile_img_url"`
}

// UserItemUpdateable describes the fields of a user which can be updated by the user
type UserItemUpdateable struct {
	UUID          string `bson:"_id"`
	FirstName     string `bson:"first_name"`
	LastName      string `bson:"last_name"`
	OrgnPosition  string `bson:"orgn_position"`
	ProfileImgURL string `bson:"profile_img_url"`
}

// InsertNew inserts a new UserItem into the mongoDB. Checks before that the user name is not already taken if so returns
// an http.StatusBadeRequest, error
func (user user) InsertNew(ctx context.Context, storage storage.Storage, userItem UserItem) (int, error) {
	// check if username is taken
	taken, err := storage.Exists(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if taken { // username exists but must be unique
		return http.StatusBadGateway, errors.New("username already taken")
	}
	// inserts new user in storage
	if err := storage.InsertOne(ctx, userDatabase, userCollection, userItem); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Update updates the allowed fields of the user record
func (user user) Update(ctx context.Context, storage storage.Storage, userItem UserItemUpdateable) (int, error) {

	updateQuery := bson.D{
		{
			"$set", bson.D{
				{"first_name", userItem.FirstName},
				{"last_name", userItem.LastName},
				{"orgn_position", userItem.OrgnPosition},
			},
		},
	}
	if err := storage.UpdateOne(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID}, updateQuery); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// GetByIDs looks up all the records that match the provided UUIDs
func (user user) GetByIDs(ctx context.Context, storage storage.Storage, UUIDs []string) (int, []UserItem, error) {

	var userList []UserItem
	if err := storage.FindMany(ctx, userDatabase, userCollection, bson.M{"_id": bson.M{"$in": UUIDs}}, &userList); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, userList, nil
}

// GetByID collects all user details for the given user uuid
func (user user) GetByID(ctx context.Context, storage storage.Storage, UUID string) (int, UserItem, error) {
	var userItem UserItem
	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"_id": UUID}, &userItem); err != nil {
		return http.StatusInternalServerError, UserItem{}, nil
	}
	return http.StatusOK, userItem, nil
}