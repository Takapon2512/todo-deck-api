package main

import (
	"log"
	"todo-deck-api/backend/db"
	"todo-deck-api/backend/model"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	users := []model.User{
		{ Email: "user1@test.com", Password: hashPassword("password123") },
		{ Email: "user2@test.com", Password: hashPassword("password456") },
	}

	for _, user := range users {
		result := dbConn.Create(&users)
		if result.Error != nil {
			log.Fatalf("Failed to seed user %s: %v", user.Email, result.Error)
		}
	}
	log.Println("Seeding completed successfully")
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hashedPassword)
}