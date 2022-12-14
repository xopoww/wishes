// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xopoww/wishes/internal/service/repository (interfaces: Repository,Transaction)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/xopoww/wishes/internal/models"
	repository "github.com/xopoww/wishes/internal/service/repository"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddList mocks base method.
func (m *MockRepository) AddList(arg0 context.Context, arg1 *models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddList", arg0, arg1)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddList indicates an expected call of AddList.
func (mr *MockRepositoryMockRecorder) AddList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddList", reflect.TypeOf((*MockRepository)(nil).AddList), arg0, arg1)
}

// AddListItems mocks base method.
func (m *MockRepository) AddListItems(arg0 context.Context, arg1 *models.List, arg2 []models.ListItem) ([]models.ListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListItems", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.ListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddListItems indicates an expected call of AddListItems.
func (mr *MockRepositoryMockRecorder) AddListItems(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListItems", reflect.TypeOf((*MockRepository)(nil).AddListItems), arg0, arg1, arg2)
}

// AddOAuth mocks base method.
func (m *MockRepository) AddOAuth(arg0 context.Context, arg1, arg2 string, arg3 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOAuth", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOAuth indicates an expected call of AddOAuth.
func (mr *MockRepositoryMockRecorder) AddOAuth(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOAuth", reflect.TypeOf((*MockRepository)(nil).AddOAuth), arg0, arg1, arg2, arg3)
}

// AddUser mocks base method.
func (m *MockRepository) AddUser(arg0 context.Context, arg1 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockRepositoryMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepository)(nil).AddUser), arg0, arg1)
}

// Begin mocks base method.
func (m *MockRepository) Begin() (repository.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(repository.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockRepositoryMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockRepository)(nil).Begin))
}

// CheckOAuth mocks base method.
func (m *MockRepository) CheckOAuth(arg0 context.Context, arg1, arg2 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOAuth indicates an expected call of CheckOAuth.
func (mr *MockRepositoryMockRecorder) CheckOAuth(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOAuth", reflect.TypeOf((*MockRepository)(nil).CheckOAuth), arg0, arg1, arg2)
}

// CheckUsername mocks base method.
func (m *MockRepository) CheckUsername(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsername", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsername indicates an expected call of CheckUsername.
func (mr *MockRepositoryMockRecorder) CheckUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsername", reflect.TypeOf((*MockRepository)(nil).CheckUsername), arg0, arg1)
}

// Close mocks base method.
func (m *MockRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepository)(nil).Close))
}

// DeleteList mocks base method.
func (m *MockRepository) DeleteList(arg0 context.Context, arg1 *models.List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockRepositoryMockRecorder) DeleteList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockRepository)(nil).DeleteList), arg0, arg1)
}

// DeleteListItems mocks base method.
func (m *MockRepository) DeleteListItems(arg0 context.Context, arg1 *models.List, arg2 []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListItems", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteListItems indicates an expected call of DeleteListItems.
func (mr *MockRepositoryMockRecorder) DeleteListItems(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListItems", reflect.TypeOf((*MockRepository)(nil).DeleteListItems), arg0, arg1, arg2)
}

// EditList mocks base method.
func (m *MockRepository) EditList(arg0 context.Context, arg1 *models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditList", arg0, arg1)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditList indicates an expected call of EditList.
func (mr *MockRepositoryMockRecorder) EditList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditList", reflect.TypeOf((*MockRepository)(nil).EditList), arg0, arg1)
}

// EditUser mocks base method.
func (m *MockRepository) EditUser(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser.
func (mr *MockRepositoryMockRecorder) EditUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockRepository)(nil).EditUser), arg0, arg1)
}

// GetItemTaken mocks base method.
func (m *MockRepository) GetItemTaken(arg0 context.Context, arg1, arg2 int64) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemTaken", arg0, arg1, arg2)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemTaken indicates an expected call of GetItemTaken.
func (mr *MockRepositoryMockRecorder) GetItemTaken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemTaken", reflect.TypeOf((*MockRepository)(nil).GetItemTaken), arg0, arg1, arg2)
}

// GetList mocks base method.
func (m *MockRepository) GetList(arg0 context.Context, arg1 int64) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", arg0, arg1)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockRepositoryMockRecorder) GetList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockRepository)(nil).GetList), arg0, arg1)
}

// GetListItems mocks base method.
func (m *MockRepository) GetListItems(arg0 context.Context, arg1 *models.List) ([]models.ListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListItems", arg0, arg1)
	ret0, _ := ret[0].([]models.ListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListItems indicates an expected call of GetListItems.
func (mr *MockRepositoryMockRecorder) GetListItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListItems", reflect.TypeOf((*MockRepository)(nil).GetListItems), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(arg0 context.Context, arg1 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), arg0, arg1)
}

// GetUserLists mocks base method.
func (m *MockRepository) GetUserLists(arg0 context.Context, arg1 int64, arg2 bool) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLists", arg0, arg1, arg2)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLists indicates an expected call of GetUserLists.
func (mr *MockRepositoryMockRecorder) GetUserLists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLists", reflect.TypeOf((*MockRepository)(nil).GetUserLists), arg0, arg1, arg2)
}

// SetItemTaken mocks base method.
func (m *MockRepository) SetItemTaken(arg0 context.Context, arg1, arg2 int64, arg3 *int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetItemTaken", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetItemTaken indicates an expected call of SetItemTaken.
func (mr *MockRepositoryMockRecorder) SetItemTaken(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetItemTaken", reflect.TypeOf((*MockRepository)(nil).SetItemTaken), arg0, arg1, arg2, arg3)
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// AddList mocks base method.
func (m *MockTransaction) AddList(arg0 context.Context, arg1 *models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddList", arg0, arg1)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddList indicates an expected call of AddList.
func (mr *MockTransactionMockRecorder) AddList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddList", reflect.TypeOf((*MockTransaction)(nil).AddList), arg0, arg1)
}

// AddListItems mocks base method.
func (m *MockTransaction) AddListItems(arg0 context.Context, arg1 *models.List, arg2 []models.ListItem) ([]models.ListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListItems", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.ListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddListItems indicates an expected call of AddListItems.
func (mr *MockTransactionMockRecorder) AddListItems(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListItems", reflect.TypeOf((*MockTransaction)(nil).AddListItems), arg0, arg1, arg2)
}

// AddOAuth mocks base method.
func (m *MockTransaction) AddOAuth(arg0 context.Context, arg1, arg2 string, arg3 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOAuth", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOAuth indicates an expected call of AddOAuth.
func (mr *MockTransactionMockRecorder) AddOAuth(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOAuth", reflect.TypeOf((*MockTransaction)(nil).AddOAuth), arg0, arg1, arg2, arg3)
}

// AddUser mocks base method.
func (m *MockTransaction) AddUser(arg0 context.Context, arg1 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockTransactionMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockTransaction)(nil).AddUser), arg0, arg1)
}

// CheckOAuth mocks base method.
func (m *MockTransaction) CheckOAuth(arg0 context.Context, arg1, arg2 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOAuth indicates an expected call of CheckOAuth.
func (mr *MockTransactionMockRecorder) CheckOAuth(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOAuth", reflect.TypeOf((*MockTransaction)(nil).CheckOAuth), arg0, arg1, arg2)
}

// CheckUsername mocks base method.
func (m *MockTransaction) CheckUsername(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsername", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsername indicates an expected call of CheckUsername.
func (mr *MockTransactionMockRecorder) CheckUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsername", reflect.TypeOf((*MockTransaction)(nil).CheckUsername), arg0, arg1)
}

// Commit mocks base method.
func (m *MockTransaction) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit))
}

// DeleteList mocks base method.
func (m *MockTransaction) DeleteList(arg0 context.Context, arg1 *models.List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockTransactionMockRecorder) DeleteList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockTransaction)(nil).DeleteList), arg0, arg1)
}

// DeleteListItems mocks base method.
func (m *MockTransaction) DeleteListItems(arg0 context.Context, arg1 *models.List, arg2 []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListItems", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteListItems indicates an expected call of DeleteListItems.
func (mr *MockTransactionMockRecorder) DeleteListItems(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListItems", reflect.TypeOf((*MockTransaction)(nil).DeleteListItems), arg0, arg1, arg2)
}

// EditList mocks base method.
func (m *MockTransaction) EditList(arg0 context.Context, arg1 *models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditList", arg0, arg1)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditList indicates an expected call of EditList.
func (mr *MockTransactionMockRecorder) EditList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditList", reflect.TypeOf((*MockTransaction)(nil).EditList), arg0, arg1)
}

// EditUser mocks base method.
func (m *MockTransaction) EditUser(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser.
func (mr *MockTransactionMockRecorder) EditUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockTransaction)(nil).EditUser), arg0, arg1)
}

// GetItemTaken mocks base method.
func (m *MockTransaction) GetItemTaken(arg0 context.Context, arg1, arg2 int64) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemTaken", arg0, arg1, arg2)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemTaken indicates an expected call of GetItemTaken.
func (mr *MockTransactionMockRecorder) GetItemTaken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemTaken", reflect.TypeOf((*MockTransaction)(nil).GetItemTaken), arg0, arg1, arg2)
}

// GetList mocks base method.
func (m *MockTransaction) GetList(arg0 context.Context, arg1 int64) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", arg0, arg1)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockTransactionMockRecorder) GetList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockTransaction)(nil).GetList), arg0, arg1)
}

// GetListItems mocks base method.
func (m *MockTransaction) GetListItems(arg0 context.Context, arg1 *models.List) ([]models.ListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListItems", arg0, arg1)
	ret0, _ := ret[0].([]models.ListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListItems indicates an expected call of GetListItems.
func (mr *MockTransactionMockRecorder) GetListItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListItems", reflect.TypeOf((*MockTransaction)(nil).GetListItems), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockTransaction) GetUser(arg0 context.Context, arg1 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockTransactionMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockTransaction)(nil).GetUser), arg0, arg1)
}

// GetUserLists mocks base method.
func (m *MockTransaction) GetUserLists(arg0 context.Context, arg1 int64, arg2 bool) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLists", arg0, arg1, arg2)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLists indicates an expected call of GetUserLists.
func (mr *MockTransactionMockRecorder) GetUserLists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLists", reflect.TypeOf((*MockTransaction)(nil).GetUserLists), arg0, arg1, arg2)
}

// Rollback mocks base method.
func (m *MockTransaction) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransaction)(nil).Rollback))
}

// SetItemTaken mocks base method.
func (m *MockTransaction) SetItemTaken(arg0 context.Context, arg1, arg2 int64, arg3 *int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetItemTaken", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetItemTaken indicates an expected call of SetItemTaken.
func (mr *MockTransactionMockRecorder) SetItemTaken(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetItemTaken", reflect.TypeOf((*MockTransaction)(nil).SetItemTaken), arg0, arg1, arg2, arg3)
}
