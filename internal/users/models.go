package users

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	Active    bool      `json:"active,omitempty"`
	CreatedAT time.Time `json:"-"`
	UpdatedAT time.Time `json:"-"`
}

func NewUser() *User { return new(User) }

func (u *User) SetID(id int64) *User {
	u.ID = id
	return u
}
func (u *User) SetName(name string) *User {
	u.Name = name
	return u
}
func (u *User) SetRole(role string) *User {
	u.Role = role
	return u
}
func (u *User) SetActive(active bool) *User {
	u.Active = active
	return u
}

func (u *User) SetCreatedAT(now time.Time) *User {
	u.CreatedAT = now
	return u
}

func (u *User) SetUpdatedAT(now time.Time) *User {
	u.UpdatedAT = now
	return u
}
