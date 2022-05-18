package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresGORM struct {
	dsn        string
	connection *gorm.DB
}

func NewPostgresGORM(dsn string) *PostgresGORM {
	return &PostgresGORM{
		dsn: dsn,
	}
}

func (p *PostgresGORM) Connect() (*PostgresGORM, error) {
	db, err := gorm.Open(postgres.Open(p.dsn), &gorm.Config{})
	if err != nil {
		return p, err
	}

	p.connection = db

	return p, nil
}

func (p *PostgresGORM) GetGORM() *gorm.DB {
	return p.connection
}
