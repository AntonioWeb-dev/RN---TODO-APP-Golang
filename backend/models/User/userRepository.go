package User

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(user *UserModel) error
	FindAll() ([]*UserModel, error)
}

type repository struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

// func (repo *Repository) FindUserById(id string) (UserModel, error) {

// }
func InitRepo(ctx context.Context, userCollection *mongo.Collection) Repository {
	return &repository{userCollection, ctx}
}

// CreateUser - Store a new user in the database
func (repo *repository) CreateUser(userData *UserModel) error {
	ObjectID := primitive.NewObjectID()
	userData.ObjectID = ObjectID
	_, err := repo.userCollection.InsertOne(repo.ctx, userData)
	return err
}

func (repo *repository) FindAll() ([]*UserModel, error) {
	filter := bson.D{{}}
	var users []*UserModel
	cur, err := repo.userCollection.Find(repo.ctx, filter)
	if err != nil {
		return users, err
	}

	for cur.Next(repo.ctx) {
		var u UserModel
		err := cur.Decode(&u)
		u.ID = u.ObjectID.Hex()
		if err != nil {
			return users, err
		}
		users = append(users, &u)
	}

	if err = cur.Err(); err != nil {
		return users, err
	}

	cur.Close(repo.ctx)
	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}
