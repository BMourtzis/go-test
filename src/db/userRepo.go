package db

import (
	"beautyproject/services/src/models"
	"fmt"
	"log"
)

func CreateUser(user models.User) (int, error) {
	var id int

	db := GetDbClient()
	queryString := fmt.Sprintf("INSERT INTO \"users\" (firstname, lastname, age) VALUES('%s', '%s', '%d') RETURNING id", user.FirstName, user.LastName, user.Age)

	log.Println(queryString)

	err := db.QueryRow(queryString).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// func ReadRecordById(t Table, id xid.ID) {
// 	db := getDbClient()
// }

func ReadUsers() ([]models.User, error) {
	db := GetDbClient()

	rows, err := db.Query("SELECT id, firstname, lastname, age FROM users")

	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Age)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// func UpdateRecord(t Table) {
// 	db := getDbClient()
// }

// func DeleteRecord(t Table) {
// 	db := getDbClient()
// }
