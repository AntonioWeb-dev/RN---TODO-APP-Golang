package User

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(user *UserModel) error
	FindAll() ([]*UserModel, error)
	FindByID(id string) (UserModel, error)
	CreateTask(userId string, taskData *Task) error
	FindByEmail(email string) (UserModel, error)
}

type repository struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

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

// CreateTask - Store a new user in the database
func (repo *repository) CreateTask(userId string, taskData *Task) error {
	ObjectID := primitive.NewObjectID()
	UserId, _ := primitive.ObjectIDFromHex(userId)
	user, err := repo.FindByID(userId)
	if err != nil {
		return err
	}

	taskData.ObjectID = ObjectID
	taskData.Create_at = time.Now()
	var tasks []Task
	if user.Tasks != nil {
		tasks = append(user.Tasks, *taskData)
	} else {
		tasks = append(tasks, *taskData)
	}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "tasks", Value: tasks},
	}}}
	_, err = repo.userCollection.UpdateByID(repo.ctx, UserId, update)
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
		if u.Tasks != nil {
			for index, task := range u.Tasks {
				u.Tasks[index].ID = task.ObjectID.Hex()
			}
		}
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

func (repo *repository) FindByID(id string) (UserModel, error) {
	primitiveID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": primitiveID}
	var user UserModel
	err := repo.userCollection.FindOne(repo.ctx, filter).Decode(&user)
	if err == mongo.ErrNilDocument {
		return UserModel{}, errors.New("User not found")
	} else if err != nil {
		return UserModel{}, err
	}
	user.ID = user.ObjectID.Hex()
	if user.Tasks != nil {
		for index, task := range user.Tasks {
			user.Tasks[index].ID = task.ObjectID.Hex()
		}
	}

	return user, nil
}

func (repo *repository) FindByEmail(email string) (UserModel, error) {
	filter := bson.M{"email": email}
	var user UserModel
	err := repo.userCollection.FindOne(repo.ctx, filter).Decode(&user)
	if err == mongo.ErrNilDocument {
		return UserModel{}, errors.New("User not found")
	} else if err != nil {
		return UserModel{}, err
	}
	user.ID = user.ObjectID.Hex()
	if user.Tasks != nil {
		for index, task := range user.Tasks {
			user.Tasks[index].ID = task.ObjectID.Hex()
		}
	}

	return user, nil
}
