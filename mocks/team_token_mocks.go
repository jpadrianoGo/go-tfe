// Code generated by MockGen. DO NOT EDIT.
// Source: team_token.go
//
// Generated by this command:
//
//	mockgen -source=team_token.go -destination=mocks/team_token_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockTeamTokens is a mock of TeamTokens interface.
type MockTeamTokens struct {
	ctrl     *gomock.Controller
	recorder *MockTeamTokensMockRecorder
}

// MockTeamTokensMockRecorder is the mock recorder for MockTeamTokens.
type MockTeamTokensMockRecorder struct {
	mock *MockTeamTokens
}

// NewMockTeamTokens creates a new mock instance.
func NewMockTeamTokens(ctrl *gomock.Controller) *MockTeamTokens {
	mock := &MockTeamTokens{ctrl: ctrl}
	mock.recorder = &MockTeamTokensMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamTokens) EXPECT() *MockTeamTokensMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTeamTokens) Create(ctx context.Context, teamID string) (*tfe.TeamToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, teamID)
	ret0, _ := ret[0].(*tfe.TeamToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTeamTokensMockRecorder) Create(ctx, teamID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTeamTokens)(nil).Create), ctx, teamID)
}

// CreateWithOptions mocks base method.
func (m *MockTeamTokens) CreateWithOptions(ctx context.Context, teamID string, options tfe.TeamTokenCreateOptions) (*tfe.TeamToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithOptions", ctx, teamID, options)
	ret0, _ := ret[0].(*tfe.TeamToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithOptions indicates an expected call of CreateWithOptions.
func (mr *MockTeamTokensMockRecorder) CreateWithOptions(ctx, teamID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithOptions", reflect.TypeOf((*MockTeamTokens)(nil).CreateWithOptions), ctx, teamID, options)
}

// Delete mocks base method.
func (m *MockTeamTokens) Delete(ctx context.Context, teamID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, teamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTeamTokensMockRecorder) Delete(ctx, teamID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTeamTokens)(nil).Delete), ctx, teamID)
}

// DeleteByID mocks base method.
func (m *MockTeamTokens) DeleteByID(ctx context.Context, tokenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, tokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockTeamTokensMockRecorder) DeleteByID(ctx, tokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTeamTokens)(nil).DeleteByID), ctx, tokenID)
}

// List mocks base method.
func (m *MockTeamTokens) List(ctx context.Context, organizationID string, options *tfe.TeamTokenListOptions) (*tfe.TeamTokenList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organizationID, options)
	ret0, _ := ret[0].(*tfe.TeamTokenList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTeamTokensMockRecorder) List(ctx, organizationID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTeamTokens)(nil).List), ctx, organizationID, options)
}

// Read mocks base method.
func (m *MockTeamTokens) Read(ctx context.Context, teamID string) (*tfe.TeamToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, teamID)
	ret0, _ := ret[0].(*tfe.TeamToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTeamTokensMockRecorder) Read(ctx, teamID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTeamTokens)(nil).Read), ctx, teamID)
}

// ReadByID mocks base method.
func (m *MockTeamTokens) ReadByID(ctx context.Context, teamID string) (*tfe.TeamToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByID", ctx, teamID)
	ret0, _ := ret[0].(*tfe.TeamToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByID indicates an expected call of ReadByID.
func (mr *MockTeamTokensMockRecorder) ReadByID(ctx, teamID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByID", reflect.TypeOf((*MockTeamTokens)(nil).ReadByID), ctx, teamID)
}
