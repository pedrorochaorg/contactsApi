package obj

import (
	"fmt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Contacts  []Contact `json:"contacts,omitempty"`
}

func (c User) String() string {
	return fmt.Sprintf("ID=%s FirstName=%s LastName=%s CreatedAt=%s UpdatedAt=%s Contacts=%s", c.ID, c.FirstName,
		c.LastName, c.CreatedAt, c.UpdatedAt, c.Contacts)
}
