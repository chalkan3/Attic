package physical_environment

import (
	postgresGORM "rodrigues.igor.com/attic/internal/database"
)

type Migrations struct {
	gorm *postgresGORM.PostgresGORM
}

func NewMigrations(gorm *postgresGORM.PostgresGORM) *Migrations {
	return &Migrations{
		gorm: gorm,
	}
}

func (m *Migrations) Migrate() {
	if ok := m.gorm.GetGORM().Migrator().HasTable(PhysicalEnvironment{}); !ok {
		m.gorm.GetGORM().Migrator().CreateTable(PhysicalEnvironment{})
	}
}
