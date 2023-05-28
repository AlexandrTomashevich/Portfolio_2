package friend_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/login/test_project/internal/entity"
	"github.com/login/test_project/internal/use_case/friend"
)

type friendService struct {
	friendUseCase friend.FriendUseCase
}

// конструктоп для создания сервиса
func New(friendUseCase friend.FriendUseCase) *friendService {
	return &friendService{
		friendUseCase: friendUseCase,
	}
}

// Создать пользователя
func (s *friendService) CreateFriend(w http.ResponseWriter, r *http.Request) {
	// получаем данные от пользователя
	w.Header().Set("Content-Type", "application/json")
	friend := &entity.Friend{}
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// проверяем данные
	if len(friend.Name) == 0 || friend.Age < 1 {
		writeError(w, "incorrect values", http.StatusBadRequest)
		return
	}
	// обращаемся к useCase
	id, err := s.friendUseCase.CreateFriend(*friend)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// отправляем ответ
	body := entity.ResponseCreate{
		Status: http.StatusOK,
		Id:     id,
	}
	resp, err := jsoniter.Marshal(body)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// Подружить двух пользователей
func (s *friendService) MakeFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	makeFriend := &entity.MakeFriend{}
	if err := json.NewDecoder(r.Body).Decode(&makeFriend); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := uuid.Parse(makeFriend.SourceId); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := uuid.Parse(makeFriend.TargetId); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	nameSource, nameTarget, err := s.friendUseCase.MakeFriend(makeFriend.SourceId, makeFriend.TargetId)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := fmt.Sprintf("%s и %s теперь друзья", nameSource, nameTarget)
	body := entity.ResponseMakeFriend{
		Status: http.StatusOK,
		Result: result,
	}
	resp, err := jsoniter.Marshal(body)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

// Удаляем пользователя
func (s *friendService) DeleteFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "user")
	if _, err := uuid.Parse(id); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	name, err := s.friendUseCase.DeleteFriend(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	body := entity.ResponseMakeFriend{
		Status: http.StatusOK,
		Result: name,
	}
	resp, err := jsoniter.Marshal(body)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

// Показать всех друзей пользователя
func (s *friendService) GetFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "user_id")
	if _, err := uuid.Parse(id); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	friends, err := s.friendUseCase.GetFriends(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := jsoniter.Marshal(friends)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

// Обновить возраст
func (s *friendService) UpdateAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// парсим возраст из запроса
	newAge := &entity.NewAge{}
	if err := json.NewDecoder(r.Body).Decode(&newAge); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// парсим id из запроса
	id := chi.URLParam(r, "user_id")
	if _, err := uuid.Parse(id); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// вызываем useCase
	err := s.friendUseCase.UpdateAge(id, newAge.Age)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
