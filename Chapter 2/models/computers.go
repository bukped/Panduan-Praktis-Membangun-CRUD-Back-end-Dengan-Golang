package models

import (
	"database/sql"
)

type Computer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ComputerCollection struct {
	Computers []Computer `json:"items"`
}

func GetComputers(db *sql.DB) ComputerCollection {
	sql := "SELECT * FROM computers"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := ComputerCollection{}
	for rows.Next() {
		computer := Computer{}
		err2 := rows.Scan(&computer.ID, &computer.Name, &computer.Price)

		if err2 != nil {
			panic(err2)
		}
		result.Computers = append(result.Computers, computer)
	}
	return result
}

func PutComputer(db *sql.DB, name string, price int) (int64, error) {
	sql := "INSERT INTO computers(name, price) VALUES(?,?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name, price)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditComputer(db *sql.DB, computerId int, name string, price int) (int64, error) {
	sql := "UPDATE computers set name = ?, price = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(name, price, computerId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteComputer(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM computers WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
