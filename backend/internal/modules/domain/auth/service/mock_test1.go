// Code generated by MockGen. DO NOT EDIT.
// Source: waterfall-backend/internal/modules/domain/auth/service (interfaces: IRoleRepo,IUserRepo,IEmailService,IFileStorageService)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	io "io"
	reflect "reflect"
	models "waterfall-backend/internal/models"
	roles "waterfall-backend/internal/models/roles"
	dto "waterfall-backend/internal/modules/domain/user/dto"
	dto0 "waterfall-backend/internal/modules/services/email/dto"
	dto1 "waterfall-backend/internal/modules/services/fs/dto"
	transaction "waterfall-backend/internal/pkg/transaction"

	gomock "github.com/golang/mock/gomock"
)

// MockIRoleRepo is a mock of IRoleRepo interface.
type MockIRoleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRoleRepoMockRecorder
}

// MockIRoleRepoMockRecorder is the mock recorder for MockIRoleRepo.
type MockIRoleRepoMockRecorder struct {
	mock *MockIRoleRepo
}

// NewMockIRoleRepo creates a new mock instance.
func NewMockIRoleRepo(ctrl *gomock.Controller) *MockIRoleRepo {
	mock := &MockIRoleRepo{ctrl: ctrl}
	mock.recorder = &MockIRoleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoleRepo) EXPECT() *MockIRoleRepoMockRecorder {
	return m.recorder
}

// GetWithPermissionsById mocks base method.
func (m *MockIRoleRepo) GetWithPermissionsById(arg0 context.Context, arg1 roles.Type) (*dto.Role, dto.Permissions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithPermissionsById", arg0, arg1)
	ret0, _ := ret[0].(*dto.Role)
	ret1, _ := ret[1].(dto.Permissions)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWithPermissionsById indicates an expected call of GetWithPermissionsById.
func (mr *MockIRoleRepoMockRecorder) GetWithPermissionsById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithPermissionsById", reflect.TypeOf((*MockIRoleRepo)(nil).GetWithPermissionsById), arg0, arg1)
}

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// CheckEmailExists mocks base method.
func (m *MockIUserRepo) CheckEmailExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailExists indicates an expected call of CheckEmailExists.
func (mr *MockIUserRepoMockRecorder) CheckEmailExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExists", reflect.TypeOf((*MockIUserRepo)(nil).CheckEmailExists), arg0, arg1)
}

// CreateWithPassword mocks base method.
func (m *MockIUserRepo) CreateWithPassword(arg0 context.Context, arg1 *dto.UserPasswordCreate) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithPassword", arg0, arg1)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithPassword indicates an expected call of CreateWithPassword.
func (mr *MockIUserRepoMockRecorder) CreateWithPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithPassword", reflect.TypeOf((*MockIUserRepo)(nil).CreateWithPassword), arg0, arg1)
}

// CreateWithService mocks base method.
func (m *MockIUserRepo) CreateWithService(arg0 context.Context, arg1 *dto.UserServiceCreate) (*dto.ExtendedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithService", arg0, arg1)
	ret0, _ := ret[0].(*dto.ExtendedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithService indicates an expected call of CreateWithService.
func (mr *MockIUserRepoMockRecorder) CreateWithService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithService", reflect.TypeOf((*MockIUserRepo)(nil).CreateWithService), arg0, arg1)
}

// GetExtendedByEmail mocks base method.
func (m *MockIUserRepo) GetExtendedByEmail(arg0 context.Context, arg1 string) (*dto.ExtendedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtendedByEmail", arg0, arg1)
	ret0, _ := ret[0].(*dto.ExtendedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtendedByEmail indicates an expected call of GetExtendedByEmail.
func (mr *MockIUserRepoMockRecorder) GetExtendedByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtendedByEmail", reflect.TypeOf((*MockIUserRepo)(nil).GetExtendedByEmail), arg0, arg1)
}

// GetExtendedByUuid mocks base method.
func (m *MockIUserRepo) GetExtendedByUuid(arg0 context.Context, arg1 string) (*dto.ExtendedUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtendedByUuid", arg0, arg1)
	ret0, _ := ret[0].(*dto.ExtendedUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtendedByUuid indicates an expected call of GetExtendedByUuid.
func (mr *MockIUserRepoMockRecorder) GetExtendedByUuid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtendedByUuid", reflect.TypeOf((*MockIUserRepo)(nil).GetExtendedByUuid), arg0, arg1)
}

// Tx mocks base method.
func (m *MockIUserRepo) Tx(arg0 context.Context) (transaction.TxRepo, transaction.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", arg0)
	ret0, _ := ret[0].(transaction.TxRepo)
	ret1, _ := ret[1].(transaction.Tx)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Tx indicates an expected call of Tx.
func (mr *MockIUserRepoMockRecorder) Tx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockIUserRepo)(nil).Tx), arg0)
}

// Update mocks base method.
func (m *MockIUserRepo) Update(arg0 context.Context, arg1 string, arg2 *dto.UserUpdate) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepo)(nil).Update), arg0, arg1, arg2)
}

// VerifyEmail mocks base method.
func (m *MockIUserRepo) VerifyEmail(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyEmail indicates an expected call of VerifyEmail.
func (mr *MockIUserRepoMockRecorder) VerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmail", reflect.TypeOf((*MockIUserRepo)(nil).VerifyEmail), arg0, arg1)
}

// MockIEmailService is a mock of IEmailService interface.
type MockIEmailService struct {
	ctrl     *gomock.Controller
	recorder *MockIEmailServiceMockRecorder
}

// MockIEmailServiceMockRecorder is the mock recorder for MockIEmailService.
type MockIEmailServiceMockRecorder struct {
	mock *MockIEmailService
}

// NewMockIEmailService creates a new mock instance.
func NewMockIEmailService(ctrl *gomock.Controller) *MockIEmailService {
	mock := &MockIEmailService{ctrl: ctrl}
	mock.recorder = &MockIEmailServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEmailService) EXPECT() *MockIEmailServiceMockRecorder {
	return m.recorder
}

// SendVerificationEmail mocks base method.
func (m *MockIEmailService) SendVerificationEmail(arg0 context.Context, arg1 string, arg2 *dto0.Verification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationEmail", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationEmail indicates an expected call of SendVerificationEmail.
func (mr *MockIEmailServiceMockRecorder) SendVerificationEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationEmail", reflect.TypeOf((*MockIEmailService)(nil).SendVerificationEmail), arg0, arg1, arg2)
}

// MockIFileStorageService is a mock of IFileStorageService interface.
type MockIFileStorageService struct {
	ctrl     *gomock.Controller
	recorder *MockIFileStorageServiceMockRecorder
}

// MockIFileStorageServiceMockRecorder is the mock recorder for MockIFileStorageService.
type MockIFileStorageServiceMockRecorder struct {
	mock *MockIFileStorageService
}

// NewMockIFileStorageService creates a new mock instance.
func NewMockIFileStorageService(ctrl *gomock.Controller) *MockIFileStorageService {
	mock := &MockIFileStorageService{ctrl: ctrl}
	mock.recorder = &MockIFileStorageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileStorageService) EXPECT() *MockIFileStorageServiceMockRecorder {
	return m.recorder
}

// CopyFile mocks base method.
func (m *MockIFileStorageService) CopyFile(arg0 context.Context, arg1 string, arg2 *models.ObjectRef, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyFile indicates an expected call of CopyFile.
func (mr *MockIFileStorageServiceMockRecorder) CopyFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFile", reflect.TypeOf((*MockIFileStorageService)(nil).CopyFile), arg0, arg1, arg2, arg3)
}

// CreateBucket mocks base method.
func (m *MockIFileStorageService) CreateBucket(arg0 context.Context, arg1 *models.ObjectRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockIFileStorageServiceMockRecorder) CreateBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockIFileStorageService)(nil).CreateBucket), arg0, arg1)
}

// CreateFile mocks base method.
func (m *MockIFileStorageService) CreateFile(arg0 context.Context, arg1 *dto1.NewFile, arg2 io.Reader) (*dto1.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto1.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockIFileStorageServiceMockRecorder) CreateFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockIFileStorageService)(nil).CreateFile), arg0, arg1, arg2)
}