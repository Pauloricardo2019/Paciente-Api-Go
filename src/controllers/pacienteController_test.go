package controllers

import (
	"GoPacientes/src/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreate_Paciente(t *testing.T) {
	createUserOk := func(c *gin.Context) {
		paciente := models.Paciente{
			Name:  "Paulo Ricardo",
			Email: "pauloric@outlook.com",
			Age:   18,
			Cpf:   01234567,
		}
		newPaciente, _ := json.Marshal(&paciente)

		response, _ := http.Post("127.0.0.1:8081", "aplication/json", newPaciente)

		assert.Equal(response, ioutil.ReadAll(response.Body))
	}

	createUser102 := func(c *gin.Context) {
		paciente := models.Paciente{
			Name:  "Paulo Ricardo",
			Email: "pauloric@outlook.com",
			Age:   101,
			Cpf:   01234567,
		}
		newPaciente, _ := json.Marshal(&paciente)

		response, _ := http.Post("127.0.0.1:8081", "aplication/json", newPaciente)

		assert.Equal(response, string("error: Age max 100 years: "))
	}

	createUserEmail := func(c *gin.Context) (*models.Paciente, error) {
		paciente := models.Paciente{
			Name:  "Paulo Ricardo",
			Email: "pauloriasaoionjifióaifjsafjasfáísfjiaijoasjoiajiofasjfassiajfaśfaaishfhafhahiaic@outlook.com",
			Age:   18,
			Cpf:   01234567,
		}
		newPaciente, _ := json.Marshal(&paciente)

		response, _ := http.POST("127.0.0.1:8081", "aplication/json", newPaciente)

		assert.Equal(response, string("error: Max length for email is 50"))

	}

	CreatePaciente()
}
