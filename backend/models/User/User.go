package User

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title,omitempty"`
	Done        bool      `json:"done,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Estimate_at time.Time `json:"estimate_at,omitempty"`
	Create_at   time.Time `json:"create_at,omitempty"`
}

// User - user model
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Tasks    []Task `json:"tasks,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
