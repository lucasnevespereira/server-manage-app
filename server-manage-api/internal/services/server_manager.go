package services

import (
	"context"
	"github.com/google/uuid"
	"server-manage-api/internal/models"
	"server-manage-api/internal/repository"
)

type ServerManagerService interface {
	Create(ctx context.Context, request models.CreateServerRequest) (*models.Server, error)
	List(ctx context.Context) ([]*models.Server, error)
	GetByID(ctx context.Context, id string) (*models.Server, error)
	DeleteByID(ctx context.Context, id string) error
}

type ServerManagerServiceImpl struct {
	repository repository.ServerRepository
}

// Enforces implementation of interface at compile time
var _ ServerManagerService = (*ServerManagerServiceImpl)(nil)

func NewServerManager(serverRepository repository.ServerRepository) *ServerManagerServiceImpl {
	return &ServerManagerServiceImpl{
		repository: serverRepository,
	}
}

func (s *ServerManagerServiceImpl) GetByID(ctx context.Context, id string) (*models.Server, error) {
	rowServer, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Server{
		ID:     rowServer.ID,
		Name:   rowServer.Name,
		Type:   rowServer.Type,
		Status: rowServer.Status,
	}, nil
}

func (s *ServerManagerServiceImpl) Create(ctx context.Context, request models.CreateServerRequest) (*models.Server, error) {
	createdServer, err := s.repository.Create(ctx, &repository.RowServer{
		ID:     uuid.New().String(),
		Name:   request.Name,
		Type:   request.Type,
		Status: request.Status,
	})
	if err != nil {
		return nil, err
	}

	return &models.Server{
		ID:     createdServer.ID,
		Name:   createdServer.Name,
		Type:   createdServer.Type,
		Status: createdServer.Status,
	}, nil
}

func (s *ServerManagerServiceImpl) DeleteByID(ctx context.Context, id string) error {
	err := s.repository.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServerManagerServiceImpl) List(ctx context.Context) ([]*models.Server, error) {
	rowServersList, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return toListServerModel(rowServersList), nil
}

func toListServerModel(rowServers []*repository.RowServer) []*models.Server {
	var serversList []*models.Server

	for _, rowServer := range rowServers {
		serversList = append(serversList, &models.Server{
			ID:     rowServer.ID,
			Name:   rowServer.Name,
			Type:   rowServer.Type,
			Status: rowServer.Status,
		})
	}

	return serversList
}
