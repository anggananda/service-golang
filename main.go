package main

import (
	"service-golang/config"
	"service-golang/handlers"
	"service-golang/repositories"
	"service-golang/routes"
	"service-golang/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	db := config.InitDB()
	r := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	inventoryRepository := repositories.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepository)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)

	vehiclesRepository := repositories.NewVehiclesRepository(db)
	vehiclesService := services.NewVehiclesService(vehiclesRepository)
	vehiclesHandler := handlers.NewVehcilesHandler(vehiclesService)

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Mengizinkan semua origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
        AllowCredentials: true, // Setel ini ke true jika menggunakan cookie atau JWT
        MaxAge:           12 * time.Hour, // Cache preflight request untuk waktu yang ditentukan
    }))


	routes.SetUpRouter(r, userHandler, inventoryHandler, vehiclesHandler, taskHandler)

	if err := r.Run(":8080"); err != nil{
		panic("Server failed to start: " + err.Error())
	}
}