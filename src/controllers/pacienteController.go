package controllers

import (
	"GoPacientes/src/database"
	"GoPacientes/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePaciente(c *gin.Context) {
	db := database.GetDB()

	var pacientes models.Paciente

	err := c.ShouldBindJSON(&pacientes)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})
		return
	}

	err = db.Create(&pacientes).Error
	if err != nil {
		c.JSON(422, gin.H{
			"error": "Cannor create a paciente: " + err.Error(),
		})
		return
	}

	c.JSON(201, pacientes)
}

func GetPacientes(c *gin.Context) {
	db := database.GetDB()

	var pacientes []models.Paciente

	err := db.Find(&pacientes).Error
	if err != nil {
		c.JSON(422, gin.H{
			"error": "Cannot list a paciente: " + err.Error(),
		})
		return
	}

	c.JSON(200, pacientes)
}

func GetPacienteById(c *gin.Context) {
	db := database.GetDB()
	var paciente models.Paciente

	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to a integer: " + err.Error(),
		})
		return
	}

	err = db.Find(&paciente, newId).Error
	if err != nil {
		c.JSON(422, gin.H{
			"error": "ID has to a integer: " + err.Error(),
		})
		return
	}

	c.JSON(200, paciente)

}

func UpdatePaciente(c *gin.Context) {
	db := database.GetDB()

	var paciente models.Paciente

	id := c.Param("id")

	newID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to a integer: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&paciente)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update a paciente: " + err.Error(),
		})
		return
	}

	err = db.Save(&paciente).Where("ID = ?", newID).Error
	if err != nil {
		c.JSON(422, gin.H{
			"error": "Cannot save paciente: " + err.Error(),
		})
		return
	}

	c.JSON(200, paciente)

}

func DeletePaciente(c *gin.Context) {
	db := database.GetDB()
	var paciente models.Paciente

	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to a integer: " + err.Error(),
		})
		return
	}

	err = db.Delete(&paciente, newId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Cannot find paciente: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Paciente was deleted for success! ",
	})

}
