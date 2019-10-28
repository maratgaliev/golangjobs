// Code generated by MockGen. DO NOT EDIT.
// Source: store/store.go

// Package mock is a generated GoMock package.
package mock

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/maratgaliev/golangjobs/model"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStore) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// Begin mocks base method
func (m *MockStore) Begin() (*sql.Tx, error) {
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin
func (mr *MockStoreMockRecorder) Begin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockStore)(nil).Begin))
}

// Commit mocks base method
func (m *MockStore) Commit(tx *sql.Tx) error {
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockStoreMockRecorder) Commit(tx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockStore)(nil).Commit), tx)
}

// Rollback mocks base method
func (m *MockStore) Rollback(tx *sql.Tx) error {
	ret := m.ctrl.Call(m, "Rollback", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockStoreMockRecorder) Rollback(tx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockStore)(nil).Rollback), tx)
}

// GetJob mocks base method
func (m *MockStore) GetJob(tx *sql.Tx, id int) (*model.Job, error) {
	ret := m.ctrl.Call(m, "GetJob", tx, id)
	ret0, _ := ret[0].(*model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob
func (mr *MockStoreMockRecorder) GetJob(tx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockStore)(nil).GetJob), tx, id)
}

// GetJobs mocks base method
func (m *MockStore) GetJobs(tx *sql.Tx) ([]*model.Job, error) {
	ret := m.ctrl.Call(m, "GetJobs", tx)
	ret0, _ := ret[0].([]*model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs
func (mr *MockStoreMockRecorder) GetJobs(tx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockStore)(nil).GetJobs), tx)
}

// CreateJob mocks base method
func (m *MockStore) CreateJob(tx *sql.Tx, job *model.Job) (*int, error) {
	ret := m.ctrl.Call(m, "CreateJob", tx, job)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob
func (mr *MockStoreMockRecorder) CreateJob(tx, job interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockStore)(nil).CreateJob), tx, job)
}

// UpdateJob mocks base method
func (m *MockStore) UpdateJob(tx *sql.Tx, job *model.Job) error {
	ret := m.ctrl.Call(m, "UpdateJob", tx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJob indicates an expected call of UpdateJob
func (mr *MockStoreMockRecorder) UpdateJob(tx, job interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockStore)(nil).UpdateJob), tx, job)
}

// DeleteJob mocks base method
func (m *MockStore) DeleteJob(tx *sql.Tx, id int) error {
	ret := m.ctrl.Call(m, "DeleteJob", tx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJob indicates an expected call of DeleteJob
func (mr *MockStoreMockRecorder) DeleteJob(tx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJob", reflect.TypeOf((*MockStore)(nil).DeleteJob), tx, id)
}