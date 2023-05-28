// создание таблицы(CREATE TABLE)
// первичный ключ по которому соединяются таблицы(PRIMARY KEY)

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name text not null,
    age int not null
);

// внешний ключ, с какой таблицой соединиться(REFERENCES)
CREATE TABLE friends (
    to_friend uuid REFERENCES "user" (id),
    from_friend uuid REFERENCES "user" (id)
);

// добавить запись (insert into)
// что добавить, значения (values)
insert into "user"(id, name, age) values ('803c50d2-d114-11ed-afa1-0242ac120002', 'Alex', 29);
insert into "user"(id, name, age) values ('803c50d2-d114-11ed-afa1-0242ac121002', 'Misha', 32);

insert into friends (to_friend, from_friend) values ('803c50d2-d114-11ed-afa1-0242ac120002','803c50d2-d114-11ed-afa1-0242ac121002');
insert into friends (to_friend, from_friend) values ('803c50d2-d114-11ed-afa1-0242ac121002','803c50d2-d114-11ed-afa1-0242ac120002');

// атрибуты, поля таблицы какие нужно выбрать  (select)
// имя таблицы из которой нужно выбрать (from)
// условия выбора данных (where)

select * from "user"

select * from "user" where age < 30

select * from friends where to_friend = '803c50d2-d114-11ed-afa1-0242ac120002'