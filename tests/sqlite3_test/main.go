package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type People struct {
	id   int
	name string
	age  int
}

type appContext struct {
	db *sql.DB
}

func connectDB(driverName string, dbName string) (*appContext, error) {
	db, err := sql.Open(driverName, dbName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &appContext{db}, nil
}

// Create
func (c *appContext) Create() {
	stmt, err := c.db.Prepare("INSERT INTO users(name,age) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec("Jack", 1)
	if err != nil {
		fmt.Printf("add error: %v", err)
		return
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted id is ", lastID)
}

// Read
func (c *appContext) Read() {
	rows, err := c.db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := new(People)
		err := rows.Scan(&p.id, &p.name, &p.age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p.id, p.name, p.age)
	}
}

// UPDATE
func (c *appContext) Update() {
	stmt, err := c.db.Prepare("UPDATE users SET age = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(10, 1)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update affect rows is ", affectNum)
}

// DELETE
func (c *appContext) Delete() {
	stmt, err := c.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete affect rows is ", affectNum)
}

func main() {
	c, err := connectDB("sqlite3", "abc.db")
	if err != nil {
		print(err)
	}

	c.Create()
	fmt.Println("add action done!")

	c.Read()
	fmt.Println("get action done!")

	c.Update()
	fmt.Println("update action done!")

	c.Delete()
	fmt.Println("delete action done!")
}
