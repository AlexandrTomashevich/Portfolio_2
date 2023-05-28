package friend_service

import (
	"net/http"
)

type FriendService interface {
	// создаем нового пользователя
	CreateFriend(w http.ResponseWriter, r *http.Request)
	// Подружить двух пользователей
	MakeFriend(iw http.ResponseWriter, r *http.Request)
	// Удаляем пользователя
	DeleteFriend(w http.ResponseWriter, r *http.Request)
	// Показать всех друзей пользователя
	GetFriends(w http.ResponseWriter, r *http.Request)
	// Обновить возраст
	UpdateAge(w http.ResponseWriter, r *http.Request)
}
