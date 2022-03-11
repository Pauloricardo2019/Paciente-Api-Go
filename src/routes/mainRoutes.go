package routes

import (
	"GoPacientes/src/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		pacientes := main.Group("pacientes")
		{
			pacientes.GET("/", controllers.GetPacientes)
			pacientes.GET("/:id", controllers.GetPacienteById)
			pacientes.POST("/", controllers.CreatePaciente)
			pacientes.PUT("/:id", controllers.UpdatePaciente)
			pacientes.DELETE("/:id", controllers.DeletePaciente)
		}
	}
	return router
}
