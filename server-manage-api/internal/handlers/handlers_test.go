package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"server-manage-api/internal/models"
	"server-manage-api/mocks"
	"testing"
)

var mockServerManager = &mocks.ServerManagerService{}
var mockRouter = gin.Default()
var mockHandler = NewHandler(mockServerManager)

func TestHealth(t *testing.T) {
	router := gin.New()
	router.GET("/health", Health)

	request, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedResponse := `{"status":"up"}`
	assert.JSONEq(t, expectedResponse, response.Body.String())
}

func TestNoRouteHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRouter.NoRoute(NoRoute)

	req, _ := http.NewRequest("GET", "/non-existing-route", nil)
	recorder := httptest.NewRecorder()

	mockRouter.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	expectedResponse := `{"status":404,"error":"Not Found"}`
	assert.JSONEq(t, expectedResponse, recorder.Body.String())
}

func TestCreateServerHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRouter.POST("/servers", mockHandler.CreateServer)

	testCases := []struct {
		name           string
		request        models.CreateServerRequest
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid request",
			request: models.CreateServerRequest{
				Name:   "Test Server",
				Type:   "medium",
				Status: "running",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"server Test Server was created with id 123456","status":200}`,
		},
		{
			name: "Invalid request with empty name",
			request: models.CreateServerRequest{
				Name:   "",
				Type:   "large",
				Status: "deployed",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"name cannot be empty","status":400}`,
		},
		{
			name: "Invalid request with invalid status",
			request: models.CreateServerRequest{
				Name:   "Nimbus",
				Type:   "large",
				Status: "deployed",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"server status, must be starting, running, stopping or stopped","status":400}`,
		},
		{
			name: "Invalid request with invalid type",
			request: models.CreateServerRequest{
				Name:   "Norther",
				Type:   "big",
				Status: "starting",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"server type must be small, medium or large","status":400}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestJSON, _ := json.Marshal(tc.request)
			requestBody := bytes.NewReader(requestJSON)
			response := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/servers", requestBody)
			req.Header.Set("Content-Type", "application/json")

			mockServerManager.On("Create", mock.Anything, tc.request).Return(&models.Server{
				ID:   "123456",
				Name: "Test Server",
			}, nil)

			mockRouter.ServeHTTP(response, req)

			assert.Equal(t, tc.expectedStatus, response.Code)
			assert.Equal(t, tc.expectedBody, response.Body.String())
		})
	}
}

func TestListServers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRouter.GET("/servers", mockHandler.ListServers)

	testCases := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Servers found",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"servers":[{"id":"1","name":"Server 1","type":"small","status":"stopped"},{"id":"2","name":"Server 2","type":"medium","status":"running"}],"total":2}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/servers", nil)

			var expectedServersList []*models.Server
			if tc.expectedStatus == http.StatusOK {
				expectedServersList = []*models.Server{
					{
						ID:     "1",
						Name:   "Server 1",
						Type:   "small",
						Status: "stopped",
					},
					{
						ID:     "2",
						Name:   "Server 2",
						Type:   "medium",
						Status: "running",
					},
				}
			}

			mockServerManager.On("List", mock.Anything).Return(expectedServersList, nil)

			mockRouter.ServeHTTP(response, req)

			assert.Equal(t, tc.expectedStatus, response.Code)
			assert.Equal(t, tc.expectedBody, response.Body.String())
		})
	}
}

func Test_validateRequest(t *testing.T) {
	type args struct {
		request *models.CreateServerRequest
	}
	tests := []struct {
		name           string
		args           args
		shouldThrowErr bool
	}{
		{
			name: "Valid server type and status",
			args: args{
				request: &models.CreateServerRequest{
					Name:   "Test Server",
					Type:   "small",
					Status: "running",
				},
			},
			shouldThrowErr: false,
		},
		{
			name: "Invalid server type",
			args: args{
				request: &models.CreateServerRequest{
					Name:   "Test Server",
					Type:   "big",
					Status: "running",
				},
			},
			shouldThrowErr: true,
		},
		{
			name: "Invalid server status",
			args: args{
				request: &models.CreateServerRequest{
					Name:   "Test Server",
					Type:   "small",
					Status: "deployed",
				},
			},
			shouldThrowErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRequest(tt.args.request)
			if tt.shouldThrowErr {
				assert.Error(t, err, "Expected error, but got nil")
			} else {
				assert.NoError(t, err, "Expected no error, but got %v", err)
			}
		})
	}
}
