package models

import "gorm.io/gorm"

type Paciente struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Cpf   int    `json:"cpf"`
}

func (p *Paciente) Expectativa() string {
	maxAge := 100
	expecAge := maxAge - p.Age
	res := ""

	if expecAge < 40 {
		res = "Expectativa de vida Baixa"
	}

	return res

}
