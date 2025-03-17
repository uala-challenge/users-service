package remove_follow

import (
	"context"
	"errors"
	"testing"

	"github.com/redis/go-redis/v9"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	log_mock "github.com/uala-challenge/simple-toolkit/pkg/utilities/log/mock"
)

func TestRemoveFollow_Success(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"

	mk.ExpectZScore("following:"+userID, followerID).SetVal(1.0)

	mk.ExpectZRem("following:"+userID, followerID).SetVal(1)

	mk.ExpectZRem("followers:"+followerID, userID).SetVal(1)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.NoError(t, err)
	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRemoveFollow_NotFollowing(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"

	mk.ExpectZScore("following:"+userID, followerID).SetErr(redis.Nil)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.NoError(t, err)
	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRemoveFollow_ErrorCheckingFollowStatus(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	expectedErr := errors.New("error en Redis")

	mk.ExpectZScore("following:"+userID, followerID).SetErr(expectedErr)

	mockLog.On("WrapError", expectedErr, "Error verificando si el seguidor existe en la lista").Return(expectedErr)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockLog.AssertExpectations(t)

	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRemoveFollow_ErrorRemovingFromFollowing(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	expectedErr := errors.New("error en Redis")

	mk.ExpectZScore("following:"+userID, followerID).SetVal(1.0)

	mk.ExpectZRem("following:"+userID, followerID).SetErr(expectedErr)

	mockLog.On("WrapError", expectedErr, "Error eliminando seguidor en following").Return(expectedErr)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockLog.AssertExpectations(t)

	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRemoveFollow_ErrorRemovingFromFollowers(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	expectedErr := errors.New("error en Redis")

	mk.ExpectZScore("following:"+userID, followerID).SetVal(1.0)

	mk.ExpectZRem("following:"+userID, followerID).SetVal(1)

	mk.ExpectZRem("followers:"+followerID, userID).SetErr(expectedErr)

	mockLog.On("WrapError", expectedErr, "Error eliminando seguidor en followers").Return(expectedErr)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockLog.AssertExpectations(t)

	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}
