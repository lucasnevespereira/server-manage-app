package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbSsl      string
}

type ServerRepositoryImpl struct {
	db *gorm.DB
}

// Enforces implementation of interface at compile time
var _ ServerRepository = (*ServerRepositoryImpl)(nil)

func NewServerRepository(config Config) (*ServerRepositoryImpl, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
		config.DbSsl,
	)

	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, errors.Wrapf(err, "could not create postgres client")
	}

	internalDB, errInternalDB := database.DB()
	if errInternalDB != nil {
		return nil, errors.Wrapf(errInternalDB, "could not get internal db")
	}

	if errPing := internalDB.Ping(); errPing != nil {
		return nil, errors.Wrapf(errPing, "could not ping database")
	}

	log.Println("Server Repository started")
	return &ServerRepositoryImpl{db: database}, nil

}

func (r *ServerRepositoryImpl) AutoMigrate() error {
	return r.db.AutoMigrate(&RowServer{})
}

func (r *ServerRepositoryImpl) GetByID(ctx context.Context, id string) (*RowServer, error) {
	var rowServer *RowServer
	result := r.db.WithContext(ctx).First(&rowServer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return rowServer, nil
}

func (r *ServerRepositoryImpl) Create(ctx context.Context, rowServer *RowServer) (*RowServer, error) {
	result := r.db.WithContext(ctx).Create(rowServer)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create server: %v", result.Error)
	}

	return rowServer, nil
}

func (r *ServerRepositoryImpl) List(ctx context.Context) ([]*RowServer, error) {
	var rowServers []*RowServer
	result := r.db.WithContext(ctx).Find(&rowServers)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list servers: %v", result.Error)
	}

	return rowServers, nil
}

func (r *ServerRepositoryImpl) DeleteByID(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(RowServer{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
