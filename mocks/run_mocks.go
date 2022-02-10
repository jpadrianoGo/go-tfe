// Code generated by MockGen. DO NOT EDIT.
// Source: run.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockRuns is a mock of Runs interface.
type MockRuns struct {
	ctrl     *gomock.Controller
	recorder *MockRunsMockRecorder
}

// MockRunsMockRecorder is the mock recorder for MockRuns.
type MockRunsMockRecorder struct {
	mock *MockRuns
}

// NewMockRuns creates a new mock instance.
func NewMockRuns(ctrl *gomock.Controller) *MockRuns {
	mock := &MockRuns{ctrl: ctrl}
	mock.recorder = &MockRunsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuns) EXPECT() *MockRunsMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockRuns) Apply(ctx context.Context, runID string, options tfe.RunApplyOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockRunsMockRecorder) Apply(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockRuns)(nil).Apply), ctx, runID, options)
}

// Cancel mocks base method.
func (m *MockRuns) Cancel(ctx context.Context, runID string, options tfe.RunCancelOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockRunsMockRecorder) Cancel(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockRuns)(nil).Cancel), ctx, runID, options)
}

// Create mocks base method.
func (m *MockRuns) Create(ctx context.Context, options tfe.RunCreateOptions) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, options)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRunsMockRecorder) Create(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRuns)(nil).Create), ctx, options)
}

// Discard mocks base method.
func (m *MockRuns) Discard(ctx context.Context, runID string, options tfe.RunDiscardOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discard", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Discard indicates an expected call of Discard.
func (mr *MockRunsMockRecorder) Discard(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discard", reflect.TypeOf((*MockRuns)(nil).Discard), ctx, runID, options)
}

// ForceCancel mocks base method.
func (m *MockRuns) ForceCancel(ctx context.Context, runID string, options tfe.RunForceCancelOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceCancel", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceCancel indicates an expected call of ForceCancel.
func (mr *MockRunsMockRecorder) ForceCancel(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceCancel", reflect.TypeOf((*MockRuns)(nil).ForceCancel), ctx, runID, options)
}

// List mocks base method.
func (m *MockRuns) List(ctx context.Context, workspaceID string, options *tfe.RunListOptions) (*tfe.RunList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.RunList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRunsMockRecorder) List(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRuns)(nil).List), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockRuns) Read(ctx context.Context, runID string) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, runID)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRunsMockRecorder) Read(ctx, runID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRuns)(nil).Read), ctx, runID)
}

// ReadWithOptions mocks base method.
func (m *MockRuns) ReadWithOptions(ctx context.Context, runID string, options *tfe.RunReadOptions) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, runID, options)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockRunsMockRecorder) ReadWithOptions(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockRuns)(nil).ReadWithOptions), ctx, runID, options)
}
