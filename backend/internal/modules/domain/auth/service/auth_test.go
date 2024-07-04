package service_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/models/session"
	user "waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/utils/ptr"
)

type txMock struct {
}

func (t *txMock) Commit() error {
	return nil
}

func (t *txMock) Rollback() error {
	return nil
}

func TestAuthService_GenerateGuestSession_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIRoleRepo(ctrl, nil)
	service := NewAuthService(repo, nil, nil, nil, nil, config.Config{})

	expectedRole := &user.Role{Id: 1, Name: "Guest", ReleaseDelay: time.Second}
	expectedPerms := []permissions.Type{permissions.Read, permissions.Write}

	repo.EXPECT().GetWithPermissionsById(gomock.Any(), roles.Guest).Return(expectedRole, expectedPerms, nil)

	// Создание ожидаемого extendedUser
	extendedUser := &user.ExtendedUser{
		User: user.User{
			Uuid:     "",
			Email:    ptr.String("guest"),
			Username: "guest",
		},
		Roles:       []*user.Role{expectedRole},
		Permissions: expectedPerms,
	}

	expectedSession := &session.Session{
		SID:          "session-id",
		UserUuid:     "",
		Email:        ptr.String("guest"),
		Username:     "guest",
		Roles:        []roles.Type{expectedRole.Id},
		Permissions:  []permissions.Type{permissions.Read, permissions.Write},
		ReleaseDelay: time.Second,
	}

	repo.EXPECT().GetWithPermissionsById(gomock.Any(), roles.Guest).Return(expectedRole, expectedPerms, nil)

	session, err := service.GenerateGuestSession(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, expectedSession, session)
	}
}
