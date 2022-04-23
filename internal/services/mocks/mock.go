// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/services/service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	core "github.com/arthurshafikov/tg-notebot/internal/core"
	gomock "github.com/golang/mock/gomock"
)

// MockCategories is a mock of Categories interface.
type MockCategories struct {
	ctrl     *gomock.Controller
	recorder *MockCategoriesMockRecorder
}

// MockCategoriesMockRecorder is the mock recorder for MockCategories.
type MockCategoriesMockRecorder struct {
	mock *MockCategories
}

// NewMockCategories creates a new mock instance.
func NewMockCategories(ctrl *gomock.Controller) *MockCategories {
	mock := &MockCategories{ctrl: ctrl}
	mock.recorder = &MockCategoriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategories) EXPECT() *MockCategoriesMockRecorder {
	return m.recorder
}

// AddCategory mocks base method.
func (m *MockCategories) AddCategory(ctx context.Context, telegramChatID int64, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCategory", ctx, telegramChatID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCategory indicates an expected call of AddCategory.
func (mr *MockCategoriesMockRecorder) AddCategory(ctx, telegramChatID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCategory", reflect.TypeOf((*MockCategories)(nil).AddCategory), ctx, telegramChatID, name)
}

// ListCategories mocks base method.
func (m *MockCategories) ListCategories(ctx context.Context, telegramChatID int64) ([]core.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", ctx, telegramChatID)
	ret0, _ := ret[0].([]core.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockCategoriesMockRecorder) ListCategories(ctx, telegramChatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockCategories)(nil).ListCategories), ctx, telegramChatID)
}

// RemoveCategory mocks base method.
func (m *MockCategories) RemoveCategory(ctx context.Context, telegramChatID int64, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCategory", ctx, telegramChatID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCategory indicates an expected call of RemoveCategory.
func (mr *MockCategoriesMockRecorder) RemoveCategory(ctx, telegramChatID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCategory", reflect.TypeOf((*MockCategories)(nil).RemoveCategory), ctx, telegramChatID, name)
}

// RenameCategory mocks base method.
func (m *MockCategories) RenameCategory(ctx context.Context, telegramChatID int64, name, newName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameCategory", ctx, telegramChatID, name, newName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameCategory indicates an expected call of RenameCategory.
func (mr *MockCategoriesMockRecorder) RenameCategory(ctx, telegramChatID, name, newName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameCategory", reflect.TypeOf((*MockCategories)(nil).RenameCategory), ctx, telegramChatID, name, newName)
}

// MockNotes is a mock of Notes interface.
type MockNotes struct {
	ctrl     *gomock.Controller
	recorder *MockNotesMockRecorder
}

// MockNotesMockRecorder is the mock recorder for MockNotes.
type MockNotesMockRecorder struct {
	mock *MockNotes
}

// NewMockNotes creates a new mock instance.
func NewMockNotes(ctrl *gomock.Controller) *MockNotes {
	mock := &MockNotes{ctrl: ctrl}
	mock.recorder = &MockNotesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotes) EXPECT() *MockNotesMockRecorder {
	return m.recorder
}

// AddNote mocks base method.
func (m *MockNotes) AddNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", ctx, telegramChatID, categoryName, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockNotesMockRecorder) AddNote(ctx, telegramChatID, categoryName, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockNotes)(nil).AddNote), ctx, telegramChatID, categoryName, content)
}

// ListNotesFromCategory mocks base method.
func (m *MockNotes) ListNotesFromCategory(ctx context.Context, telegramChatID int64, categoryName string) ([]core.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotesFromCategory", ctx, telegramChatID, categoryName)
	ret0, _ := ret[0].([]core.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNotesFromCategory indicates an expected call of ListNotesFromCategory.
func (mr *MockNotesMockRecorder) ListNotesFromCategory(ctx, telegramChatID, categoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotesFromCategory", reflect.TypeOf((*MockNotes)(nil).ListNotesFromCategory), ctx, telegramChatID, categoryName)
}

// RemoveNote mocks base method.
func (m *MockNotes) RemoveNote(ctx context.Context, telegramChatID int64, categoryName, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNote", ctx, telegramChatID, categoryName, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNote indicates an expected call of RemoveNote.
func (mr *MockNotesMockRecorder) RemoveNote(ctx, telegramChatID, categoryName, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNote", reflect.TypeOf((*MockNotes)(nil).RemoveNote), ctx, telegramChatID, categoryName, content)
}

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// CheckChatIDExists mocks base method.
func (m *MockUsers) CheckChatIDExists(ctx context.Context, telegramChatID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckChatIDExists", ctx, telegramChatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckChatIDExists indicates an expected call of CheckChatIDExists.
func (mr *MockUsersMockRecorder) CheckChatIDExists(ctx, telegramChatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckChatIDExists", reflect.TypeOf((*MockUsers)(nil).CheckChatIDExists), ctx, telegramChatID)
}

// CreateIfNotExists mocks base method.
func (m *MockUsers) CreateIfNotExists(ctx context.Context, telegramChatID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIfNotExists", ctx, telegramChatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIfNotExists indicates an expected call of CreateIfNotExists.
func (mr *MockUsersMockRecorder) CreateIfNotExists(ctx, telegramChatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIfNotExists", reflect.TypeOf((*MockUsers)(nil).CreateIfNotExists), ctx, telegramChatID)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockLogger) Error(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", err)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), err)
}
