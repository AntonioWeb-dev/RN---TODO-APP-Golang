package Task

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskModel struct {
	ObjectID    primitive.ObjectID `json:"-" bson:"_id"`
	ID          string             `json:"_id" bson:"-"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Done        bool               `json:"done" bson:"done,omitempty"`
	Priority    int                `json:"priority" bson:"priority,omitempty"`
	Estimate_at time.Time          `json:"estimate_at" bson:"estimate_at,omitempty"`
	Create_at   time.Time          `json:"create_at" bson:"create_at,omitempty"`
}
