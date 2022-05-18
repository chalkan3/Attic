package physical_environment

import (
	"time"
)

type PhysicalEnvironment struct {
	ID        int       `gorm:"primaryKey" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Active    bool      `json:"active,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	CreatedAT time.Time `json:"created_at,omitempty"`
	UpdatedAT time.Time `json:"updated_at,omitempty"`
}

func NewPhysicalEnvironment() *PhysicalEnvironment {
	return new(PhysicalEnvironment)
}

func (pe *PhysicalEnvironment) SetID(id int) *PhysicalEnvironment {
	pe.ID = id
	return pe
}

func (pe *PhysicalEnvironment) SetUserID(id int) *PhysicalEnvironment {
	pe.UserID = id
	return pe
}

func (pe *PhysicalEnvironment) SetName(name string) *PhysicalEnvironment {
	pe.Name = name
	return pe
}

func (pe *PhysicalEnvironment) SetActive(active bool) *PhysicalEnvironment {
	pe.Active = active
	return pe
}

func (pe *PhysicalEnvironment) SetCreatedAT(createdAT time.Time) *PhysicalEnvironment {
	pe.CreatedAT = createdAT
	return pe
}

func (pe *PhysicalEnvironment) SetUpdateAT(updateAT time.Time) *PhysicalEnvironment {
	pe.UpdatedAT = updateAT
	return pe
}

func (pe *PhysicalEnvironment) GetID() int              { return pe.ID }
func (pe *PhysicalEnvironment) GetUserID() int          { return pe.UserID }
func (pe *PhysicalEnvironment) GetName() string         { return pe.Name }
func (pe *PhysicalEnvironment) GetActive() bool         { return pe.Active }
func (pe *PhysicalEnvironment) GetCreatedAT() time.Time { return pe.CreatedAT }
func (pe *PhysicalEnvironment) GetUpdatedAT() time.Time { return pe.UpdatedAT }
