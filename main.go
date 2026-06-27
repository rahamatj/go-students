package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func connectToDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "students.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database created successfully!")

	return db
}

func prompts() int {
	fmt.Println("Welcome to the Student Management System!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Add a student")
	fmt.Println("2. View all students")
	fmt.Println("3. Update a student")
	fmt.Println("4. Delete a student")
	fmt.Println("5. Exit")

	var choice int

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func addStudent() {
	db := connectToDatabase()
	defer db.Close()

	var name string
	var age int

	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter student age: ")
	fmt.Scanln(&age)

	result, err := db.Exec(
		"INSERT INTO students(name, age) VALUES(?, ?)",
		name,
		age,
	)
	if err != nil {
		log.Fatal(err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID:", id)
	fmt.Println("Student added successfully!")
}

func viewStudents() {
	db := connectToDatabase()
	defer db.Close()

	fmt.Println("All students: ")

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println()
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
		fmt.Println()
	}
}

func updateStudent() {
	db := connectToDatabase()
	defer db.Close()

	var id int

	fmt.Print("Enter student ID: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	var age int
	fmt.Print("Enter student age: ")
	fmt.Scanln(&age)

	_, err := db.Exec(
		"UPDATE students SET name = ?, age = ? WHERE id = ?",
		name,
		age,
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID:", id)
	fmt.Println("Student Updated successfully!")
}

func deleteStudent() {
	db := connectToDatabase()
	defer db.Close()

	var id int

	fmt.Print("Enter student ID: ")
	fmt.Scanln(&id)

	_, err := db.Exec(
		"DELETE FROM students WHERE id = ?",
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID:", id)
	fmt.Println("Student Deleted successfully!")
}

func options(option int) {
	switch option {
	case 1:
		addStudent()
	case 2:
		viewStudents()
	case 3:
		updateStudent()
	case 4:
		deleteStudent()
	case 5:
		fmt.Println("Exiting the program...")
		return
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func main() {
	choice := 0
	for choice != 5 {
		choice = prompts()
		options(choice)
	}
}