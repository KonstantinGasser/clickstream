package storage

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	nameDB   = "datalab_apptoken"
	nameColl = "apptoken"
)

// MongoClient implements the apptokens.Repository interface
type MongoClient struct {
	conn *mongo.Client
}

func NewMongoClient(conn *mongo.Client) (*MongoClient, error) {
	return &MongoClient{conn: conn}, nil
}

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Query must be any of bson.* or a struct with bson tags
// Returned data from the coll.InsertOne are ignored and will not be returned by the function
func (client MongoClient) Initialize(ctx context.Context, appToken apptokens.AppToken) error {

	data, err := bson.Marshal(appToken)
	if err != nil {
		return err
	}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// GetById looks up the app token behind the uuid and writes the result in the passed pointer
// // to the result. If none found returns mongo.ErrNoDocuments
func (client MongoClient) GetById(ctx context.Context, uuid string, result interface{}) error {
	filter := bson.M{"_id": uuid}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}

// Update updates the Jwt and Exp of the stored AppToken document
func (client MongoClient) Update(ctx context.Context, uuid, jwt string, exp int64, rfc int32) error {
	query := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"app_jwt":       jwt,
				"app_jwt_exp":   exp,
				"refresh_count": rfc,
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if _, err := coll.UpdateByID(ctx, uuid, query); err != nil {
		return err
	}
	return nil
}

func (client MongoClient) SetAppTokenLock(ctx context.Context, uuid string, lock bool) error {
	filter := bson.M{"_id": uuid}
	query := bson.D{
		{
			Key:   "$set",
			Value: bson.M{"locked": lock},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if _, err := coll.UpdateOne(ctx, filter, query); err != nil {
		return err
	}
	return nil
}

func (client MongoClient) AddMember(ctx context.Context, uuid, userUuid string) error {
	filter := bson.M{"_id": uuid}
	query := bson.D{
		{
			Key: "$addToSet",
			Value: bson.M{
				"member": userUuid,
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if _, err := coll.UpdateOne(ctx, filter, query); err != nil {
		return err
	}
	return nil
}

func (client MongoClient) RollbackAddMember(ctx context.Context, uuid, userUuid string) error {
	filter := bson.M{"_id": uuid}
	query := bson.D{
		{
			Key: "$pull",
			Value: bson.M{
				"member": userUuid,
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if _, err := coll.UpdateOne(ctx, filter, query); err != nil {
		return err
	}
	return nil
}
