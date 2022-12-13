package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcioc0sta/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.Students)
	r.GET("/students/:id", controllers.StudentById)
	r.GET("/:name", controllers.Greet)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.Run()
}