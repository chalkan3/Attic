package groups

import (
	"time"
)

type Group struct {
	ID                    int    `gorm:"primaryKey" json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	Active                bool   `json:"active,omitempty"`
	UserID                int    `json:"user_id,omitempty"`
	PhysicalEnvironmentID int
	CreatedAT             time.Time `json:"created_at,omitempty"`
	UpdatedAT             time.Time `json:"updated_at,omitempty"`
}

func NewGroup() *Group {
	return new(Group)
}

func (pe *Group) SetID(id int) *Group {
	pe.ID = id
	return pe
}

func (pe *Group) SetUserID(id int) *Group {
	pe.UserID = id
	return pe
}

func (pe *Group) SetPhysicalEnvironmentID(id int) *Group {
	pe.PhysicalEnvironmentID = id
	return pe
}

func (pe *Group) SetName(name string) *Group {
	pe.Name = name
	return pe
}

func (pe *Group) SetActive(active bool) *Group {
	pe.Active = active
	return pe
}

func (pe *Group) SetCreatedAT(createdAT time.Time) *Group {
	pe.CreatedAT = createdAT
	return pe
}

func (pe *Group) SetUpdateAT(updateAT time.Time) *Group {
	pe.UpdatedAT = updateAT
	return pe
}
func (pe *Group) GetID() int                    { return pe.ID }
func (pe *Group) GetUserID() int                { return pe.UserID }
func (pe *Group) GetPhysicalEnvironmentID() int { return pe.PhysicalEnvironmentID }

func (pe *Group) GetName() string         { return pe.Name }
func (pe *Group) GetActive() bool         { return pe.Active }
func (pe *Group) GetCreatedAT() time.Time { return pe.CreatedAT }
func (pe *Group) GetUpdatedAT() time.Time { return pe.UpdatedAT }
