// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source repository.go -destination repository_mock.go -package repository
//
// Package repository is a generated GoMock package.
package repository

import (
	model "project/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAllInRepo is a mock of AllInRepo interface.
type MockAllInRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAllInRepoMockRecorder
}

// MockAllInRepoMockRecorder is the mock recorder for MockAllInRepo.
type MockAllInRepoMockRecorder struct {
	mock *MockAllInRepo
}

// NewMockAllInRepo creates a new mock instance.
func NewMockAllInRepo(ctrl *gomock.Controller) *MockAllInRepo {
	mock := &MockAllInRepo{ctrl: ctrl}
	mock.recorder = &MockAllInRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllInRepo) EXPECT() *MockAllInRepoMockRecorder {
	return m.recorder
}

// ApplyJob_Repository mocks base method.
func (m *MockAllInRepo) ApplyJob_Repository(id uint64) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyJob_Repository", id)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyJob_Repository indicates an expected call of ApplyJob_Repository.
func (mr *MockAllInRepoMockRecorder) ApplyJob_Repository(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyJob_Repository", reflect.TypeOf((*MockAllInRepo)(nil).ApplyJob_Repository), id)
}

// CreateCompany mocks base method.
func (m *MockAllInRepo) CreateCompany(arg0 model.Company) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", arg0)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockAllInRepoMockRecorder) CreateCompany(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockAllInRepo)(nil).CreateCompany), arg0)
}

// CreateJob mocks base method.
func (m *MockAllInRepo) CreateJob(j model.Job) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", j)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockAllInRepoMockRecorder) CreateJob(j any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockAllInRepo)(nil).CreateJob), j)
}

// CreateUser mocks base method.
func (m *MockAllInRepo) CreateUser(arg0 model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAllInRepoMockRecorder) CreateUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAllInRepo)(nil).CreateUser), arg0)
}

// FetchUserByEmail mocks base method.
func (m *MockAllInRepo) FetchUserByEmail(arg0 string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserByEmail", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserByEmail indicates an expected call of FetchUserByEmail.
func (mr *MockAllInRepoMockRecorder) FetchUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserByEmail", reflect.TypeOf((*MockAllInRepo)(nil).FetchUserByEmail), arg0)
}

// GetAllCompany mocks base method.
func (m *MockAllInRepo) GetAllCompany() ([]model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCompany")
	ret0, _ := ret[0].([]model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCompany indicates an expected call of GetAllCompany.
func (mr *MockAllInRepoMockRecorder) GetAllCompany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCompany", reflect.TypeOf((*MockAllInRepo)(nil).GetAllCompany))
}

// GetAllJobs mocks base method.
func (m *MockAllInRepo) GetAllJobs() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJobs")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs.
func (mr *MockAllInRepoMockRecorder) GetAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockAllInRepo)(nil).GetAllJobs))
}

// GetCompany mocks base method.
func (m *MockAllInRepo) GetCompany(id int) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", id)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockAllInRepoMockRecorder) GetCompany(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockAllInRepo)(nil).GetCompany), id)
}

// GetJobId mocks base method.
func (m *MockAllInRepo) GetJobId(id uint64) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobId", id)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobId indicates an expected call of GetJobId.
func (mr *MockAllInRepoMockRecorder) GetJobId(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobId", reflect.TypeOf((*MockAllInRepo)(nil).GetJobId), id)
}

// GetJobs mocks base method.
func (m *MockAllInRepo) GetJobs(id int) ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", id)
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockAllInRepoMockRecorder) GetJobs(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockAllInRepo)(nil).GetJobs), id)
}
