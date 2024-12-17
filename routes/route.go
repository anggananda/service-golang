package routes

import (
	"service-golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, userHandler *handlers.UserHandler, inventoryHandler *handlers.InventoryHandler, vehiclesHandler *handlers.VehiclesHandler, taskHandler *handlers.TaskHandler) {
	api := r.Group("/api/v1")
	{
		api.POST("/user", userHandler.CreateUser)
		api.GET("/users", userHandler.GetAllUser)
		api.GET("/users/:id", userHandler.GetUserByID)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

		// Inventory
		api.GET("/inventory", inventoryHandler.GetAllInventory)
		api.POST("/inventory", inventoryHandler.CreateInventory)
		api.PUT("/inventory/:id", inventoryHandler.UpdateInventory)
		api.DELETE("/inventory/:id", inventoryHandler.DeleteInventory)

		// Vehicles
		api.GET("/vehicles", vehiclesHandler.GetAllVehicles)
		api.POST("vehicles", vehiclesHandler.CreateVehicles)
		api.PUT("/vehicles/:id", vehiclesHandler.UpdateVehicles)
		api.DELETE("/vehicles/:id", vehiclesHandler.DeleteVehicles)

		// Tasks
		api.GET("/task", taskHandler.GetAllTask)
		api.POST("task", taskHandler.CreateTask)
		api.PUT("/task/:id", taskHandler.UpdateTask)
		api.DELETE("/task/:id", taskHandler.DeleteTask)
	}
}
