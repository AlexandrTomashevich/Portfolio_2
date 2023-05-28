package repository

import "github.com/login/test_project/internal/entity"

type Repository interface {
	// создаем нового пользователя
	CreateFriend(in entity.Friend) (string, error)
	// Подружить двух пользователей
	MakeFriend(in1 string, in2 string) (string, string, error)
	// Удаляем пользователя
	DeleteFriend(in string) (string, error)
	// Показать всех друзей пользователя
	GetFriends(in string) ([]entity.FriendDTO, error)
	// Обновить возраст
	UpdateAge(in string, age int) error
}
