package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"server-manage-api/internal/models"
	"server-manage-api/internal/repository"
	"server-manage-api/mocks"
	"testing"
)

func TestServerManagerServiceImpl_GetByID(t *testing.T) {
	testCases := []struct {
		name           string
		serverID       string
		mockReturn     *repository.RowServer
		mockErr        error
		shouldThrowErr bool
		expectedServer *models.Server
	}{
		{
			name:     "Server found",
			serverID: "123",
			mockReturn: &repository.RowServer{
				ID:     "123",
				Name:   "Test Server",
				Type:   "medium",
				Status: "running",
			},
			mockErr:        nil,
			shouldThrowErr: false,
			expectedServer: &models.Server{
				ID:     "123",
				Name:   "Test Server",
				Type:   "medium",
				Status: "running",
			},
		},
		{
			name:           "Server not found",
			serverID:       "123",
			mockReturn:     nil,
			mockErr:        errors.New("server not found"),
			shouldThrowErr: true,
			expectedServer: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mocks.ServerRepository{}
			mockRepo.On("GetByID", mock.Anything, tc.serverID).Return(tc.mockReturn, tc.mockErr)
			serverManager := NewServerManager(mockRepo)

			server, err := serverManager.GetByID(context.Background(), tc.serverID)

			if tc.shouldThrowErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedServer, server)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestServerManagerServiceImpl_Create(t *testing.T) {
	mockRepo := &mocks.ServerRepository{}
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*repository.RowServer")).Return(&repository.RowServer{
		ID:     "123",
		Name:   "Test Server",
		Type:   "medium",
		Status: "running",
	}, nil)

	serverManager := NewServerManager(mockRepo)
	request := models.CreateServerRequest{
		Name:   "Test Server",
		Type:   "medium",
		Status: "running",
	}

	server, err := serverManager.Create(context.Background(), request)

	expectedServer := &models.Server{
		ID:     "123",
		Name:   "Test Server",
		Type:   "medium",
		Status: "running",
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedServer, server)

	mockRepo.AssertExpectations(t)
}

func TestServerManagerServiceImpl_List(t *testing.T) {
	mockRepo := &mocks.ServerRepository{}
	mockRepo.On("List", mock.Anything).Return([]*repository.RowServer{
		{
			ID:     "123",
			Name:   "Test Server 1",
			Type:   "small",
			Status: "stopped",
		},
		{
			ID:     "456",
			Name:   "Test Server 2",
			Type:   "medium",
			Status: "running",
		},
	}, nil)

	serverManager := NewServerManager(mockRepo)

	servers, err := serverManager.List(context.Background())

	expectedServers := []*models.Server{
		{
			ID:     "123",
			Name:   "Test Server 1",
			Type:   "small",
			Status: "stopped",
		},
		{
			ID:     "456",
			Name:   "Test Server 2",
			Type:   "medium",
			Status: "running",
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedServers, servers)

	mockRepo.AssertExpectations(t)
}
