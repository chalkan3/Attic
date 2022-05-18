package physical_environment

import (
	"time"

	postgresGORM "rodrigues.igor.com/attic/internal/database"
)

type Repository interface {
	Insert(pe *PhysicalEnvironment) (*PhysicalEnvironment, error)
	Update(pe *PhysicalEnvironment) (*PhysicalEnvironment, error)
	Get(pe *PhysicalEnvironment) (*PhysicalEnvironment, error)
	List() ([]PhysicalEnvironment, error)
}

type repository struct {
	gorm *postgresGORM.PostgresGORM
}

func NewRepository(gorm *postgresGORM.PostgresGORM) Repository {
	return &repository{
		gorm: gorm,
	}
}

func (repo *repository) Insert(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	pe.SetCreatedAT(time.Now())
	if err := repo.gorm.GetGORM().Create(pe).Error; err != nil {
		return nil, err
	}

	return pe, nil
}

func (repo *repository) Update(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	pe.SetUpdateAT(time.Now())
	if err := repo.gorm.GetGORM().Save(pe).Error; err != nil {
		return nil, err
	}

	return pe, nil
}

func (repo *repository) Get(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	if err := repo.gorm.GetGORM().Find(pe, []int64{int64(pe.GetID())}).Error; err != nil {
		return nil, err
	}

	return pe, nil
}

func (repo *repository) List() ([]PhysicalEnvironment, error) {
	var pe []PhysicalEnvironment
	if err := repo.gorm.GetGORM().Preload("users").Find(&pe).Error; err != nil {
		return pe, nil
	}

	return pe, nil
}
