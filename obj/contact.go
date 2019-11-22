package obj

import (
	"fmt"
	"time"
)

type Contact struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Contact) String() string {
	return fmt.Sprintf("ID=%d UserID=%d FirstName=%s LastName=%s Email=%s Phone=%s CreatedAt=%s UpdatedAt=%s", c.ID,
		c.UserID,
		c.FirstName,
		c.LastName, c.Email, c.Phone, c.CreatedAt, c.UpdatedAt)
}
