package users

import (
	"time"

	postgresGORM "rodrigues.igor.com/attic/internal/database"
)

type Repository interface {
	Insert(user *User) (*User, error)
	Update(user *User) (*User, error)
	Get(user *User) (*User, error)
	List() ([]User, error)
}

type repository struct {
	gorm *postgresGORM.PostgresGORM
}

func NewRepository(gorm *postgresGORM.PostgresGORM) Repository {
	return &repository{
		gorm: gorm,
	}
}

func (repo *repository) Insert(user *User) (*User, error) {
	user.SetCreatedAT(time.Now())
	if err := repo.gorm.GetGORM().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repository) Update(user *User) (*User, error) {
	user.SetUpdatedAT(time.Now())
	if err := repo.gorm.GetGORM().Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *repository) Get(user *User) (*User, error) {
	var fetchedUser User
	if err := repo.gorm.GetGORM().Find(&fetchedUser, []int64{user.ID}).Error; err != nil {
		return nil, err
	}

	return &fetchedUser, nil
}

func (repo *repository) List() ([]User, error) {
	var users []User
	if err := repo.gorm.GetGORM().Preload("users").Find(&users).Error; err != nil {
		return users, nil
	}

	return users, nil
}
