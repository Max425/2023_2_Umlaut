package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMessageService_GetDialogMessages(t *testing.T) {
	messageText := "Message 1"
	mockMessages := []model.Message{
		{Id: ptrToInt(1), Text: &messageText},
		{Id: ptrToInt(2), Text: &messageText},
	}

	tests := []struct {
		name           string
		mockBehavior   func(r *mock_repository.MockMessage)
		expectedResult []model.Message
		expectedError  error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockMessage) {
				r.EXPECT().GetDialogMessages(gomock.Any(), 1).Return(mockMessages, nil)
			},
			expectedResult: mockMessages,
			expectedError:  nil,
		},
		{
			name: "Error",
			mockBehavior: func(r *mock_repository.MockMessage) {
				r.EXPECT().GetDialogMessages(gomock.Any(), 1).Return(nil, errors.New("error getting messages"))
			},
			expectedResult: nil,
			expectedError:  errors.New("error getting messages"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMessage := mock_repository.NewMockMessage(ctrl)
			test.mockBehavior(repoMessage)

			service := &MessageService{repoMessage: repoMessage}
			messages, err := service.GetDialogMessages(context.Background(), 1)

			assert.Equal(t, test.expectedResult, messages)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestMessageService_SaveOrUpdateMessage(t *testing.T) {
	messageText := "Message 1"
	mockMessage := model.Message{
		Id:   ptrToInt(1),
		Text: &messageText,
	}

	tests := []struct {
		name           string
		inputMessage   model.Message
		mockBehavior   func(r *mock_repository.MockMessage)
		expectedResult model.Message
		expectedError  error
	}{
		{
			name:         "Success Create",
			inputMessage: model.Message{Text: &messageText},
			mockBehavior: func(r *mock_repository.MockMessage) {
				r.EXPECT().CreateMessage(gomock.Any(), gomock.Any()).Return(mockMessage, nil)
			},
			expectedResult: mockMessage,
			expectedError:  nil,
		},
		{
			name:         "Success Update",
			inputMessage: model.Message{Id: ptrToInt(1), Text: &messageText},
			mockBehavior: func(r *mock_repository.MockMessage) {
				r.EXPECT().UpdateMessage(gomock.Any(), gomock.Any()).Return(mockMessage, nil)
			},
			expectedResult: mockMessage,
			expectedError:  nil,
		},
		{
			name:         "Error",
			inputMessage: model.Message{},
			mockBehavior: func(r *mock_repository.MockMessage) {
				r.EXPECT().CreateMessage(gomock.Any(), gomock.Any()).Return(model.Message{}, errors.New("error creating message"))
			},
			expectedResult: model.Message{},
			expectedError:  errors.New("error creating message"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMessage := mock_repository.NewMockMessage(ctrl)
			test.mockBehavior(repoMessage)

			service := &MessageService{repoMessage: repoMessage}
			resultMessage, err := service.SaveOrUpdateMessage(context.Background(), test.inputMessage)

			assert.Equal(t, test.expectedResult, resultMessage)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func ptrToInt(i int) *int {
	return &i
}