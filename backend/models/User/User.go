package User

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Task struct {
	ID          int64     `json:"_id" bson:"_id"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Done        bool      `json:"done" bson:"done,omitempty"`
	Priority    int       `json:"priority" bson:"priority,omitempty"`
	Estimate_at time.Time `json:"estimate_at" bson:"estimate_at,omitempty"`
	Create_at   time.Time `json:"create_at" bson:"create_at,omitempty"`
}

// User - user model
type UserModel struct {
	ObjectID primitive.ObjectID `json:"-" bson:"_id"`
	ID       string             `bson:"-"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Tasks    []Task             `json:"tasks" bson:"tasks,omitempty"`
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
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	return nil
}
