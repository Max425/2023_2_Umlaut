// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	multipart "mime/multipart"
	reflect "reflect"

	model "github.com/go-park-mail-ru/2023_2_Umlaut/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(ctx context.Context, user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), ctx, user)
}

// DeleteCookie mocks base method.
func (m *MockAuthorization) DeleteCookie(ctx context.Context, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCookie", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCookie indicates an expected call of DeleteCookie.
func (mr *MockAuthorizationMockRecorder) DeleteCookie(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCookie", reflect.TypeOf((*MockAuthorization)(nil).DeleteCookie), ctx, session)
}

// GenerateCookie mocks base method.
func (m *MockAuthorization) GenerateCookie(ctx context.Context, id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCookie", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCookie indicates an expected call of GenerateCookie.
func (mr *MockAuthorizationMockRecorder) GenerateCookie(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCookie", reflect.TypeOf((*MockAuthorization)(nil).GenerateCookie), ctx, id)
}

// GetDecodeUserId mocks base method.
func (m *MockAuthorization) GetDecodeUserId(ctx context.Context, message string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecodeUserId", ctx, message)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDecodeUserId indicates an expected call of GetDecodeUserId.
func (mr *MockAuthorizationMockRecorder) GetDecodeUserId(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecodeUserId", reflect.TypeOf((*MockAuthorization)(nil).GetDecodeUserId), ctx, message)
}

// GetSessionValue mocks base method.
func (m *MockAuthorization) GetSessionValue(ctx context.Context, session string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionValue", ctx, session)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionValue indicates an expected call of GetSessionValue.
func (mr *MockAuthorizationMockRecorder) GetSessionValue(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionValue", reflect.TypeOf((*MockAuthorization)(nil).GetSessionValue), ctx, session)
}

// GetUser mocks base method.
func (m *MockAuthorization) GetUser(ctx context.Context, mail, password string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, mail, password)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationMockRecorder) GetUser(ctx, mail, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorization)(nil).GetUser), ctx, mail, password)
}

// MockFeed is a mock of Feed interface.
type MockFeed struct {
	ctrl     *gomock.Controller
	recorder *MockFeedMockRecorder
}

// MockFeedMockRecorder is the mock recorder for MockFeed.
type MockFeedMockRecorder struct {
	mock *MockFeed
}

// NewMockFeed creates a new mock instance.
func NewMockFeed(ctrl *gomock.Controller) *MockFeed {
	mock := &MockFeed{ctrl: ctrl}
	mock.recorder = &MockFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeed) EXPECT() *MockFeedMockRecorder {
	return m.recorder
}

// GetNextUser mocks base method.
func (m *MockFeed) GetNextUser(ctx context.Context, params model.FilterParams) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextUser", ctx, params)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextUser indicates an expected call of GetNextUser.
func (mr *MockFeedMockRecorder) GetNextUser(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextUser", reflect.TypeOf((*MockFeed)(nil).GetNextUser), ctx, params)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateFile mocks base method.
func (m *MockUser) CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", ctx, userId, file, size)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockUserMockRecorder) CreateFile(ctx, userId, file, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockUser)(nil).CreateFile), ctx, userId, file, size)
}

// DeleteFile mocks base method.
func (m *MockUser) DeleteFile(ctx context.Context, userId int, link string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, userId, link)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockUserMockRecorder) DeleteFile(ctx, userId, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockUser)(nil).DeleteFile), ctx, userId, link)
}

// GetCurrentUser mocks base method.
func (m *MockUser) GetCurrentUser(ctx context.Context, userId int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", ctx, userId)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockUserMockRecorder) GetCurrentUser(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockUser)(nil).GetCurrentUser), ctx, userId)
}

// GetUserShareCridentials mocks base method.
func (m *MockUser) GetUserShareCridentials(ctx context.Context, userId int) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserShareCridentials", ctx, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserShareCridentials indicates an expected call of GetUserShareCridentials.
func (mr *MockUserMockRecorder) GetUserShareCridentials(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserShareCridentials", reflect.TypeOf((*MockUser)(nil).GetUserShareCridentials), ctx, userId)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), ctx, user)
}

// MockLike is a mock of Like interface.
type MockLike struct {
	ctrl     *gomock.Controller
	recorder *MockLikeMockRecorder
}

// MockLikeMockRecorder is the mock recorder for MockLike.
type MockLikeMockRecorder struct {
	mock *MockLike
}

// NewMockLike creates a new mock instance.
func NewMockLike(ctrl *gomock.Controller) *MockLike {
	mock := &MockLike{ctrl: ctrl}
	mock.recorder = &MockLikeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLike) EXPECT() *MockLikeMockRecorder {
	return m.recorder
}

// CreateLike mocks base method.
func (m *MockLike) CreateLike(ctx context.Context, like model.Like) (model.Dialog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", ctx, like)
	ret0, _ := ret[0].(model.Dialog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockLikeMockRecorder) CreateLike(ctx, like interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockLike)(nil).CreateLike), ctx, like)
}

// MockDialog is a mock of Dialog interface.
type MockDialog struct {
	ctrl     *gomock.Controller
	recorder *MockDialogMockRecorder
}

// MockDialogMockRecorder is the mock recorder for MockDialog.
type MockDialogMockRecorder struct {
	mock *MockDialog
}

// NewMockDialog creates a new mock instance.
func NewMockDialog(ctrl *gomock.Controller) *MockDialog {
	mock := &MockDialog{ctrl: ctrl}
	mock.recorder = &MockDialogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialog) EXPECT() *MockDialogMockRecorder {
	return m.recorder
}

// CreateDialog mocks base method.
func (m *MockDialog) CreateDialog(ctx context.Context, dialog model.Dialog) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDialog", ctx, dialog)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDialog indicates an expected call of CreateDialog.
func (mr *MockDialogMockRecorder) CreateDialog(ctx, dialog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDialog", reflect.TypeOf((*MockDialog)(nil).CreateDialog), ctx, dialog)
}

// GetDialog mocks base method.
func (m *MockDialog) GetDialog(ctx context.Context, id int) (model.Dialog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialog", ctx, id)
	ret0, _ := ret[0].(model.Dialog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialog indicates an expected call of GetDialog.
func (mr *MockDialogMockRecorder) GetDialog(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialog", reflect.TypeOf((*MockDialog)(nil).GetDialog), ctx, id)
}

// GetDialogs mocks base method.
func (m *MockDialog) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogs", ctx, userId)
	ret0, _ := ret[0].([]model.Dialog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogs indicates an expected call of GetDialogs.
func (mr *MockDialogMockRecorder) GetDialogs(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogs", reflect.TypeOf((*MockDialog)(nil).GetDialogs), ctx, userId)
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// GetDialogMessages mocks base method.
func (m *MockMessage) GetDialogMessages(ctx context.Context, userId, recipientId int) ([]model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogMessages", ctx, userId, recipientId)
	ret0, _ := ret[0].([]model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogMessages indicates an expected call of GetDialogMessages.
func (mr *MockMessageMockRecorder) GetDialogMessages(ctx, userId, recipientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogMessages", reflect.TypeOf((*MockMessage)(nil).GetDialogMessages), ctx, userId, recipientId)
}

// SaveOrUpdateMessage mocks base method.
func (m *MockMessage) SaveOrUpdateMessage(ctx context.Context, message model.Message) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrUpdateMessage", ctx, message)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveOrUpdateMessage indicates an expected call of SaveOrUpdateMessage.
func (mr *MockMessageMockRecorder) SaveOrUpdateMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrUpdateMessage", reflect.TypeOf((*MockMessage)(nil).SaveOrUpdateMessage), ctx, message)
}

// MockTag is a mock of Tag interface.
type MockTag struct {
	ctrl     *gomock.Controller
	recorder *MockTagMockRecorder
}

// MockTagMockRecorder is the mock recorder for MockTag.
type MockTagMockRecorder struct {
	mock *MockTag
}

// NewMockTag creates a new mock instance.
func NewMockTag(ctrl *gomock.Controller) *MockTag {
	mock := &MockTag{ctrl: ctrl}
	mock.recorder = &MockTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTag) EXPECT() *MockTagMockRecorder {
	return m.recorder
}

// GetAllTags mocks base method.
func (m *MockTag) GetAllTags(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTags", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTags indicates an expected call of GetAllTags.
func (mr *MockTagMockRecorder) GetAllTags(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTags", reflect.TypeOf((*MockTag)(nil).GetAllTags), ctx)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// CreateFeedFeedback mocks base method.
func (m *MockAdmin) CreateFeedFeedback(ctx context.Context, rec model.Recommendation) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeedFeedback", ctx, rec)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeedFeedback indicates an expected call of CreateFeedFeedback.
func (mr *MockAdminMockRecorder) CreateFeedFeedback(ctx, rec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeedFeedback", reflect.TypeOf((*MockAdmin)(nil).CreateFeedFeedback), ctx, rec)
}

// CreateFeedback mocks base method.
func (m *MockAdmin) CreateFeedback(ctx context.Context, stat model.Feedback) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeedback", ctx, stat)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeedback indicates an expected call of CreateFeedback.
func (mr *MockAdminMockRecorder) CreateFeedback(ctx, stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeedback", reflect.TypeOf((*MockAdmin)(nil).CreateFeedback), ctx, stat)
}

// CreateRecommendation mocks base method.
func (m *MockAdmin) CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecommendation", ctx, rec)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecommendation indicates an expected call of CreateRecommendation.
func (mr *MockAdminMockRecorder) CreateRecommendation(ctx, rec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecommendation", reflect.TypeOf((*MockAdmin)(nil).CreateRecommendation), ctx, rec)
}

// GetCSATType mocks base method.
func (m *MockAdmin) GetCSATType(ctx context.Context, userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCSATType", ctx, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCSATType indicates an expected call of GetCSATType.
func (mr *MockAdminMockRecorder) GetCSATType(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCSATType", reflect.TypeOf((*MockAdmin)(nil).GetCSATType), ctx, userId)
}

// GetFeedbackStatistics mocks base method.
func (m *MockAdmin) GetFeedbackStatistics(ctx context.Context) (model.FeedbackStatistic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedbackStatistics", ctx)
	ret0, _ := ret[0].(model.FeedbackStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedbackStatistics indicates an expected call of GetFeedbackStatistics.
func (mr *MockAdminMockRecorder) GetFeedbackStatistics(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedbackStatistics", reflect.TypeOf((*MockAdmin)(nil).GetFeedbackStatistics), ctx)
}

// GetRecommendationsStatistics mocks base method.
func (m *MockAdmin) GetRecommendationsStatistics(ctx context.Context) (model.RecommendationStatistic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendationsStatistics", ctx)
	ret0, _ := ret[0].(model.RecommendationStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationsStatistics indicates an expected call of GetRecommendationsStatistics.
func (mr *MockAdminMockRecorder) GetRecommendationsStatistics(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationsStatistics", reflect.TypeOf((*MockAdmin)(nil).GetRecommendationsStatistics), ctx)
}

// MockComplaint is a mock of Complaint interface.
type MockComplaint struct {
	ctrl     *gomock.Controller
	recorder *MockComplaintMockRecorder
}

// MockComplaintMockRecorder is the mock recorder for MockComplaint.
type MockComplaintMockRecorder struct {
	mock *MockComplaint
}

// NewMockComplaint creates a new mock instance.
func NewMockComplaint(ctrl *gomock.Controller) *MockComplaint {
	mock := &MockComplaint{ctrl: ctrl}
	mock.recorder = &MockComplaintMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplaint) EXPECT() *MockComplaintMockRecorder {
	return m.recorder
}

// CreateComplaint mocks base method.
func (m *MockComplaint) CreateComplaint(ctx context.Context, complaint model.Complaint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComplaint", ctx, complaint)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComplaint indicates an expected call of CreateComplaint.
func (mr *MockComplaintMockRecorder) CreateComplaint(ctx, complaint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComplaint", reflect.TypeOf((*MockComplaint)(nil).CreateComplaint), ctx, complaint)
}

// GetComplaintTypes mocks base method.
func (m *MockComplaint) GetComplaintTypes(ctx context.Context) ([]model.ComplaintType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplaintTypes", ctx)
	ret0, _ := ret[0].([]model.ComplaintType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplaintTypes indicates an expected call of GetComplaintTypes.
func (mr *MockComplaintMockRecorder) GetComplaintTypes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplaintTypes", reflect.TypeOf((*MockComplaint)(nil).GetComplaintTypes), ctx)
}

// GetNextComplaint mocks base method.
func (m *MockComplaint) GetNextComplaint(ctx context.Context) (model.Complaint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextComplaint", ctx)
	ret0, _ := ret[0].(model.Complaint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextComplaint indicates an expected call of GetNextComplaint.
func (mr *MockComplaintMockRecorder) GetNextComplaint(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextComplaint", reflect.TypeOf((*MockComplaint)(nil).GetNextComplaint), ctx)
}
