package update_timeline

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	rest_mock "github.com/uala-challenge/simple-toolkit/pkg/client/rest/mock"
	log_mock "github.com/uala-challenge/simple-toolkit/pkg/utilities/log/mock"
)

func TestUpdateTimeline_Success(t *testing.T) {
	mockRest := rest_mock.NewService(t)
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"

	// Simular respuesta exitosa de la API
	mockResponse := &resty.Response{
		RawResponse: nil,
	}

	headers := map[string]string{"Aapply": "application/json"}
	body := map[string]string{"follower_id": followerID}
	requestBody, _ := json.Marshal(body)

	mockRest.On("Patch", mock.Anything, fmt.Sprintf("/timeline/%s", userID), requestBody, headers).
		Return(mockResponse, nil)

	service := NewService(Dependencies{
		Client: mockRest,
		Log:    mockLog,
	})

	rsp, err := service.Apply(context.TODO(), userID, followerID)

	assert.NoError(t, err)
	assert.NotNil(t, rsp)

	mockRest.AssertExpectations(t)
}

func TestUpdateTimeline_ErrorRequest(t *testing.T) {
	mockRest := rest_mock.NewService(t)
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	expectedErr := errors.New("API error")

	headers := map[string]string{"Aapply": "application/json"}
	body := map[string]string{"follower_id": followerID}
	requestBody, _ := json.Marshal(body)

	mockRest.On("Patch", mock.Anything, fmt.Sprintf("/timeline/%s", userID), requestBody, headers).
		Return(nil, expectedErr)

	// Configuración correcta del mock del log
	mockLog.On("Error", mock.Anything, expectedErr, "Error getting value", mock.Anything).
		Return()

	service := NewService(Dependencies{
		Client: mockRest,
		Log:    mockLog,
	})

	rsp, err := service.Apply(context.TODO(), userID, followerID)

	assert.Nil(t, rsp)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)

	mockRest.AssertExpectations(t)
	mockLog.AssertExpectations(t)
}
