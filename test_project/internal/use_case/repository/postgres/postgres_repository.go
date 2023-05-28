package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/login/test_project/internal/entity"
)

type postgresRepository struct {
	// структура для работы с БД
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *postgresRepository {
	return &postgresRepository{
		db: db,
	}
}

// запрос на создания друга
func (r *postgresRepository) CreateFriend(in entity.Friend) (string, error) {
	// запрос к БД
	query := `insert into "user"(name, age) values ($1, $2) returning id;`
	// отправляем запрос к базе данные и получаем ответ структура Row
	row := r.db.QueryRow(query, in.Name, in.Age)
	id := ""
	// из байтов (010101) получаем наш id
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *postgresRepository) MakeFriend(in1 string, in2 string) (string, string, error) {
	// первый запрос вставка id друзей
	query := `insert into "friends"(to_friend, from_friend)
	values 
	($1, $2),
	($2, $1);`
	_, err := r.db.Exec(query, in1, in2)
	if err != nil {
		return "", "", err
	}

	// второй запрос имя первого друга
	var name1 entity.NameDTO
	query = `select name from "user" where id = $1;`
	if err := r.db.Get(&name1, query, in1); err != nil {
		return "", "", err
	}

	// третий имя второго друга
	var name2 entity.NameDTO
	query = `select name from "user" where id = $1;`
	if err := r.db.Get(&name2, query, in2); err != nil {
		return "", "", err
	}

	return name1.Name, name2.Name, nil
}

func (r *postgresRepository) DeleteFriend(in string) (string, error) {
	query := `delete from "friends" where to_friend = $1 or from_friend = $1;`
	_, err := r.db.Exec(query, in)
	if err != nil {
		return "", err
	}

	query = `delete from "user" where id = $1 returning name;`
	row := r.db.QueryRow(query, in)
	name := ""
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}

func (r *postgresRepository) GetFriends(in string) ([]entity.FriendDTO, error) {
	query := `select id, name, age
	from "friends"
	join "user" u on u.id = from_friend
	where to_friend=$1 `
	// взять выборку в слайс
	friends := []entity.FriendDTO{}
	if err := r.db.Select(&friends, query, in); err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *postgresRepository) UpdateAge(in string, age int) error {
	query := `update "user" set age=$1 where id=$2`
	_, err := r.db.Exec(query, age, in)
	if err != nil {
		return err
	}
	return nil
}
