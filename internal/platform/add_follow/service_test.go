package add_follow

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	log_mock "github.com/uala-challenge/simple-toolkit/pkg/utilities/log/mock"
)

func TestAddFollow_Success(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	timestamp := float64(time.Now().Unix())

	mk.ExpectZScore("following:"+userID, followerID).SetErr(redis.Nil)

	mk.ExpectZAdd("following:"+userID, redis.Z{Score: timestamp, Member: followerID}).SetVal(1)

	mk.ExpectZAdd("followers:"+followerID, redis.Z{Score: timestamp, Member: userID}).SetVal(1)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.NoError(t, err)
	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAddFollow_AlreadyFollowing(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"

	mk.ExpectZScore("following:"+userID, followerID).SetVal(1.0)

	service := NewService(Dependencies{
		Client: mockRedis,
		Log:    mockLog,
	})

	err := service.Accept(context.TODO(), userID, followerID)

	assert.NoError(t, err)
	err = mk.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestAddFollow_ErrorCheckingFollowStatus(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	expectedErr := errors.New("error en Redis")

	mk.ExpectZScore("following:"+userID, followerID).SetErr(expectedErr)

	mockLog.On("WrapError", expectedErr, "Error verificando si el seguidor ya está en la lista").Return(expectedErr)

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

func TestAddFollow_ErrorAddingToFollowing(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	timestamp := float64(time.Now().Unix())
	expectedErr := errors.New("error en Redis")

	mk.ExpectZScore("following:"+userID, followerID).SetErr(redis.Nil)

	mk.ExpectZAdd("following:"+userID, redis.Z{Score: timestamp, Member: followerID}).SetErr(expectedErr)

	mockLog.On("WrapError", expectedErr, "Error agregando seguidor en following").Return(expectedErr)

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

func TestAddFollow_ErrorAddingToFollowers(t *testing.T) {
	mockRedis, mk := redismock.NewClientMock()
	mockLog := log_mock.NewService(t)

	userID := "user-123"
	followerID := "user-456"
	timestamp := float64(time.Now().Unix())
	expectedErr := errors.New("error en Redis")

	mk.ExpectZScore("following:"+userID, followerID).SetErr(redis.Nil)

	mk.ExpectZAdd("following:"+userID, redis.Z{Score: timestamp, Member: followerID}).SetVal(1)

	mk.ExpectZAdd("followers:"+followerID, redis.Z{Score: timestamp, Member: userID}).SetErr(expectedErr)

	mockLog.On("WrapError", expectedErr, "Error agregando seguidor en followers").Return(expectedErr)

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
