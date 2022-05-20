package groups

import (
	"time"

	postgresGORM "rodrigues.igor.com/attic/internal/database"
)

type Repository interface {
	Insert(group *Group) (*Group, error)
	Update(group *Group) (*Group, error)
	Get(group *Group) (*Group, error)
	List() ([]Group, error)
}

type repository struct {
	gorm *postgresGORM.PostgresGORM
}

func NewRepository(gorm *postgresGORM.PostgresGORM) Repository {
	return &repository{
		gorm: gorm,
	}
}

func (repo *repository) Insert(group *Group) (*Group, error) {
	group.SetCreatedAT(time.Now())
	if err := repo.gorm.GetGORM().Create(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (repo *repository) Update(group *Group) (*Group, error) {
	group.SetUpdateAT(time.Now())
	if err := repo.gorm.GetGORM().Save(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (repo *repository) Get(group *Group) (*Group, error) {
	if err := repo.gorm.GetGORM().Find(group, []int64{int64(group.GetID())}).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (repo *repository) List() ([]Group, error) {
	var group []Group
	if err := repo.gorm.GetGORM().Preload("users").Find(&group).Error; err != nil {
		return group, nil
	}

	return group, nil
}
