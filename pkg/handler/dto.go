package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type signInInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpInput struct {
	Name     string `json:"name" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type deleteLink struct {
	Link string `json:"link"`
}

type idResponse struct {
	Id int `json:"id"`
}

type ClientResponseDto[K comparable] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload K      `json:"payload"`
}

func NewClientResponseDto[K comparable](ctx context.Context, w http.ResponseWriter, statusCode int, message string, payload K) {
	response := ClientResponseDto[K]{
		Status:  statusCode,
		Message: message,
		Payload: payload,
	}
	sendData(ctx, w, response, statusCode, message)
}

func NewSuccessClientResponseDto[K comparable](ctx context.Context, w http.ResponseWriter, payload K) {
	NewClientResponseDto[K](ctx, w, 200, "success", payload)
}

func newErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, message string) {
	NewClientResponseDto[string](ctx, w, statusCode, message, "")
}

type ClientResponseArrayDto[K comparable] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload []K    `json:"payload"`
}

func NewClientResponseArrayDto[K comparable](ctx context.Context, w http.ResponseWriter, statusCode int, message string, payload []K) {
	response := ClientResponseArrayDto[K]{
		Status:  statusCode,
		Message: message,
		Payload: payload,
	}
	sendData(ctx, w, response, statusCode, message)
}

func NewSuccessClientResponseArrayDto[K comparable](ctx context.Context, w http.ResponseWriter, payload []K) {
	NewClientResponseArrayDto[K](ctx, w, 200, "success", payload)
}

func sendData(ctx context.Context, w http.ResponseWriter, response interface{}, statusCode int, message string) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	logger, ok := ctx.Value(keyLogger).(*zap.Logger)
	if !ok {
		log.Println("Logger not found in context")
	}
	*logger = *logger.With(zap.Any("Status", statusCode), zap.Any("Message", message))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
