package main

import (
	"database/sql"
	"log"
	"test-project/controllers"
	"test-project/controllers/middlewares"
	"test-project/repositories"
	"test-project/services"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	db:=initDB()
    // приложение трехслойное и состоит: контроллеров,сервисов и репозитории
	// передача дб в репозиторииб, передача репозитории в сервис
	userRepository := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepository)

	phoneRepository := repositories.NewPhoneRepository(db)
    phoneService := services.NewPhoneService(phoneRepository)

	router:=gin.Default()
    // передача сервисов в контроллеры, 
	userController := controllers.NewUserController(userService)
	phoneController := controllers.NewPhoneController(phoneService)
	router.LoadHTMLGlob("templates/*")
     // паблик маршруты
	router.GET("/user/registration",userController.ShowRegistrationForm)
	router.POST("/user/register",userController.Registration)
	router.POST("/user/auth",userController.Login)
	// маршруты которые требуют ауторизации
	authorized := router.Group("/api")
    authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/user/:user_name",userController.GetByName)
		authorized.POST("/user/phone",phoneController.CreatePhone)
		authorized.GET("/user/phone/:phone_number",phoneController.GetPhone)
		authorized.PUT("/user/phone",phoneController.UpdatePhone)
		authorized.DELETE("/user/phone/:id",phoneController.DeletePhone)
	}
	
	router.Run(":8080")
}

func initDB() *sql.DB{
	db, err:=sql.Open("sqlite","project.db")
	if err != nil {
		log.Fatal(err)
	}
	
    _, err = db.Exec(`DROP TABLE IF EXISTS users;DROP TABLE IF EXISTS phones;`)
	if err != nil {
		log.Fatal(err)
	}
    // создание таблиц
	_, err = db.Exec(`
		CREATE TABLE users (
			user_id INTEGER PRIMARY KEY,
			login TEXT UNIQUE,
			password TEXT,
			name TEXT,
			age INTEGER
		);
		CREATE TABLE phones(
			phone_id INTEGER PRIMARY KEY,
			phone_number TEXT,
            is_fax INTEGER CHECK(is_fax BETWEEN 0 AND 1),
			description TEXT,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(user_id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}