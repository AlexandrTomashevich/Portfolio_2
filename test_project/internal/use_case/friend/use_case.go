package friend

import (
	"github.com/login/test_project/internal/entity"
	"github.com/login/test_project/internal/use_case/repository"
)

type friendUseCase struct {
	repo repository.Repository
}

func New(repo repository.Repository) *friendUseCase {
	return &friendUseCase{
		repo: repo,
	}
}

func (u *friendUseCase) CreateFriend(in entity.Friend) (string, error) {
	id, err := u.repo.CreateFriend(in)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (u *friendUseCase) MakeFriend(in1 string, in2 string) (string, string, error) {
	name1, name2, err := u.repo.MakeFriend(in1, in2)
	if err != nil {
		return "", "", err
	}
	return name1, name2, nil
}

func (u *friendUseCase) DeleteFriend(in string) (string, error) {
	name, err := u.repo.DeleteFriend(in)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (u *friendUseCase) GetFriends(in string) ([]entity.FriendDTO, error) {
	friends, err := u.repo.GetFriends(in)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (u *friendUseCase) UpdateAge(in string, age int) error {
	if err := u.repo.UpdateAge(in, age); err != nil {
		return err
	}
	return nil
}
