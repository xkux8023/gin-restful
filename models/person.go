package models

import (
	db "gin-restful/database"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Age       int    `json:"age" form:"age"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name, age) VALUES (?, ?, ?)", p.FirstName, p.LastName, p.Age)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name, age FROM person")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *Person) GetPerson() (person Person, err error) {
	err = db.SqlDB.QueryRow("SELECT id, first_name, last_name, age FROM person WHERE id=?", p.Id).Scan(&person.Id, &person.FirstName, &person.LastName, &person.Age)
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE person SET first_name=?, last_name=?, age=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id, p.Age)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM person WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
