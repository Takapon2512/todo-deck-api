package main

import (
	"fmt"
	"log"
	"todo-deck-api/backend/db"
	"todo-deck-api/backend/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	clearData(dbConn)

	users := []model.User{
		{ Email: "user1@test.com", Password: hashPassword("password123") },
		{ Email: "user2@test.com", Password: hashPassword("password456") },
	}

	for _, user := range users {
		result := dbConn.Create(&user)
		if result.Error != nil {
			log.Fatalf("Failed to seed user %s: %v", user.Email, result.Error)
		}
	}
	log.Println("User seeding completed successfully")

	tasks := []model.Task{
		{ Title: "Task1", Description: "Description1", Completed: 0, UserId: 1 },
		{ Title: "Task2", Description: "Description2", Completed: 1, UserId: 1 },
		{ Title: "Task3", Description: "Description3", Completed: 2, UserId: 2 },
		{ Title: "Task4", Description: "Description4", Completed: 1, UserId: 2 },
	}

	for _, task := range tasks {
		result := dbConn.Create(&task)
		if result.Error != nil {
			log.Fatalf("Failed to seed task %s: %v", fmt.Sprint(task.ID), result.Error)
		}
	}
	log.Println("Task seeding completed successfully")
}

// データを初期化
func clearData(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}


func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hashedPassword)
}