// Code generated by MockGen. DO NOT EDIT.
// Source: policy.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockPolicies is a mock of Policies interface.
type MockPolicies struct {
	ctrl     *gomock.Controller
	recorder *MockPoliciesMockRecorder
}

// MockPoliciesMockRecorder is the mock recorder for MockPolicies.
type MockPoliciesMockRecorder struct {
	mock *MockPolicies
}

// NewMockPolicies creates a new mock instance.
func NewMockPolicies(ctrl *gomock.Controller) *MockPolicies {
	mock := &MockPolicies{ctrl: ctrl}
	mock.recorder = &MockPoliciesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicies) EXPECT() *MockPoliciesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPolicies) Create(ctx context.Context, organization string, options tfe.PolicyCreateOptions) (*tfe.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPoliciesMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPolicies)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockPolicies) Delete(ctx context.Context, policyID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, policyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPoliciesMockRecorder) Delete(ctx, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPolicies)(nil).Delete), ctx, policyID)
}

// Download mocks base method.
func (m *MockPolicies) Download(ctx context.Context, policyID string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, policyID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockPoliciesMockRecorder) Download(ctx, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockPolicies)(nil).Download), ctx, policyID)
}

// List mocks base method.
func (m *MockPolicies) List(ctx context.Context, organization string, options *tfe.PolicyListOptions) (*tfe.PolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.PolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPoliciesMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicies)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockPolicies) Read(ctx context.Context, policyID string) (*tfe.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, policyID)
	ret0, _ := ret[0].(*tfe.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockPoliciesMockRecorder) Read(ctx, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPolicies)(nil).Read), ctx, policyID)
}

// Update mocks base method.
func (m *MockPolicies) Update(ctx context.Context, policyID string, options tfe.PolicyUpdateOptions) (*tfe.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, policyID, options)
	ret0, _ := ret[0].(*tfe.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPoliciesMockRecorder) Update(ctx, policyID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPolicies)(nil).Update), ctx, policyID, options)
}

// Upload mocks base method.
func (m *MockPolicies) Upload(ctx context.Context, policyID string, content []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, policyID, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockPoliciesMockRecorder) Upload(ctx, policyID, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockPolicies)(nil).Upload), ctx, policyID, content)
}
