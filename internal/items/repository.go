package items

import (
	"time"

	postgresGORM "rodrigues.igor.com/attic/internal/database"
)

type Repository interface {
	Insert(item *Item) (*Item, error)
	Update(item *Item) (*Item, error)
	Get(item *Item) (*Item, error)
	List() ([]Item, error)
}

type repository struct {
	gorm *postgresGORM.PostgresGORM
}

func NewRepository(gorm *postgresGORM.PostgresGORM) Repository {
	return &repository{
		gorm: gorm,
	}
}

func (repo *repository) Insert(item *Item) (*Item, error) {
	item.SetCreatedAT(time.Now())
	if err := repo.gorm.GetGORM().Create(item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *repository) Update(item *Item) (*Item, error) {
	item.SetUpdateAT(time.Now())
	if err := repo.gorm.GetGORM().Save(item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *repository) Get(item *Item) (*Item, error) {
	if err := repo.gorm.GetGORM().Find(item, []int64{int64(item.GetID())}).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *repository) List() ([]Item, error) {
	var item []Item
	if err := repo.gorm.GetGORM().Preload("users").Find(&item).Error; err != nil {
		return item, nil
	}

	return item, nil
}
