package repository

import (
	"database/sql"
	"github.com/ipan97/go-dig-sample/model"
	_ "github.com/mattn/go-sqlite3"
)

type PersonRepository struct {
	database *sql.DB
}

func (repository *PersonRepository) FindAll() []*model.Person {
	rows, _ := repository.database.Query(`SELECT id, name, age FROM people;`)
	defer rows.Close()

	var people []*model.Person

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		_ = rows.Scan(&id, &name, &age)

		people = append(people, &model.Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}
