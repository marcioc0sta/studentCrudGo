package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcioc0sta/gin-api-rest/db"
	"github.com/marcioc0sta/gin-api-rest/models"
)

func Students(c *gin.Context) {
	students := [] models.Student{}

	db.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func StudentById(c *gin.Context) {
	id := c.Param("id")

	student := models.Student{}

	db.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func Greet(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"message": "Hi " + name,
	})
}


func CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	foundStudent := db.DB.Where("cpf = ?", student.CPF).Find(&student)

	if foundStudent.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Student already exists",
		})

		return
	}
	
	db.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	student := models.Student{}

	db.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	db.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	student := models.Student{}

	db.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	db.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func FindStudentByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	student := models.Student{}


	db.DB.Where(models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}