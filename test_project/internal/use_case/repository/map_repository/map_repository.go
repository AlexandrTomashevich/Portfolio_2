package map_repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/login/test_project/internal/entity"
)

type mapRepository struct {
	storage map[string]*entity.Friend
}

func New() *mapRepository {
	return &mapRepository{
		storage: make(map[string]*entity.Friend),
	}
}

func (m *mapRepository) CreateFriend(in entity.Friend) (string, error) {
	id := uuid.New().String()
	if _, ok := m.storage[id]; ok {
		return "", fmt.Errorf("user exist id=%s", id)
	}
	in.Id = id
	m.storage[id] = &in
	return id, nil
}

func (m *mapRepository) MakeFriend(in1 string, in2 string) (string, string, error) {
	friend1, ok := m.storage[in1]
	if !ok {
		return "", "", fmt.Errorf("user don't exist id=%s", in1)
	}

	friend2, ok := m.storage[in2]
	if !ok {
		return "", "", fmt.Errorf("user don't exist id=%s", in2)
	}

	if ok := isFriends(*friend1, *friend2); ok {
		return "", "", fmt.Errorf("user1 id=%s is friend user2 id=%s", in1, in2)
	}
	friend1.Friends = append(friend1.Friends, *friend2)
	friend2.Friends = append(friend2.Friends, *friend1)
	m.storage[in1] = friend1
	m.storage[in2] = friend2

	return friend1.Name, friend2.Name, nil
}

func (m *mapRepository) DeleteFriend(in string) (string, error) {
	friend, ok := m.storage[in]
	if !ok {
		return "", fmt.Errorf("user don't exist id=%s", in)
	}
	delete(m.storage, in)

	return friend.Name, nil
}

func (m *mapRepository) GetFriends(in string) ([]entity.FriendDTO, error) {
	friend, ok := m.storage[in]
	if !ok {
		return nil, fmt.Errorf("user don't exist id=%s", in)
	}
	if len(friend.Friends) == 0 {
		return nil, fmt.Errorf("user don't have friends id=%s", in)
	}
	friends := []entity.FriendDTO{}
	for _, friend := range friend.Friends {
		data := entity.FriendDTO{
			Name: friend.Name,
		}
		friends = append(friends, data)
	}
	return friends, nil
}

func (m *mapRepository) UpdateAge(in string, age int) error {
	friend, ok := m.storage[in]
	if !ok {
		return fmt.Errorf("user don't exist id=%s", in)
	}
	friend.Age = age
	m.storage[in] = friend
	return nil
}

func isFriends(friend1, friend2 entity.Friend) bool {
	check := 0
	for _, v1 := range friend1.Friends {
		if v1.Id == friend2.Id {
			check++
		}
	}
	for _, v2 := range friend2.Friends {
		if v2.Id == friend1.Id {
			check++
		}
	}
	return check == 2
}
