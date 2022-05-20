package items

import (
	"time"
)

type Item struct {
	ID                    int       `gorm:"primaryKey" json:"id,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Active                bool      `json:"active,omitempty"`
	Quantity              int64     `json:"quantity,omitempty"`
	GroupID               int       `json:"group_id,omitempty"`
	UserID                int       `json:"user_id,omitempty"`
	PhysicalEnvironmentID int       `json:"physical_environment_id,omitempty"`
	CreatedAT             time.Time `json:"created_at,omitempty"`
	UpdatedAT             time.Time `json:"updated_at,omitempty"`
}

func NewItem() *Item {
	return new(Item)
}

func (pe *Item) SetID(id int) *Item {
	pe.ID = id
	return pe
}

func (pe *Item) SetUserID(id int) *Item {
	pe.UserID = id
	return pe
}

func (pe *Item) SetPhysicalEnvironmentID(id int) *Item {
	pe.PhysicalEnvironmentID = id
	return pe
}

func (pe *Item) SetName(name string) *Item {
	pe.Name = name
	return pe
}

func (pe *Item) SetActive(active bool) *Item {
	pe.Active = active
	return pe
}

func (pe *Item) SetCreatedAT(createdAT time.Time) *Item {
	pe.CreatedAT = createdAT
	return pe
}

func (pe *Item) SetUpdateAT(updateAT time.Time) *Item {
	pe.UpdatedAT = updateAT
	return pe
}
func (pe *Item) GetID() int                    { return pe.ID }
func (pe *Item) GetUserID() int                { return pe.UserID }
func (pe *Item) GetPhysicalEnvironmentID() int { return pe.PhysicalEnvironmentID }

func (pe *Item) GetName() string         { return pe.Name }
func (pe *Item) GetActive() bool         { return pe.Active }
func (pe *Item) GetCreatedAT() time.Time { return pe.CreatedAT }
func (pe *Item) GetUpdatedAT() time.Time { return pe.UpdatedAT }
