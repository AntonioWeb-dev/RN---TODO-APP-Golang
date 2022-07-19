package User

import (
	"api/helpers/hash"
	"api/models/Task"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User - user model
type UserModel struct {
	ObjectID primitive.ObjectID `json:"-" bson:"_id"`
	ID       string             `json:"_id" bson:"-"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Tasks    []Task.TaskModel   `json:"tasks" bson:"tasks,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

func (user *UserModel) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}
	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *UserModel) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name: invalid arguments")
	}
	if user.Username == "" {
		return errors.New("Username: invalid arguments")
	}
	if user.Email == "" {
		return errors.New("Email: invalid arguments")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email invalid")
	}
	if user.Password == "" && step == "create" {
		return errors.New("Password: invalid arguments")
	}
	return nil
}

func (user *UserModel) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if step == "create" {
		passwordHash, err := hash.GenerateHash(user.Password)
		if err != nil {
			return err
		}
		user.Password = passwordHash
	}
	return nil
}
