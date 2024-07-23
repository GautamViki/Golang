package handler

import (
	"fmt"

	r "main.go/repository"
)

// NewIdpRepo returns new IDP repository
func NewIdpRepo() r.StudentRepo {
	return &mysqlIdpRepo{}
}

// IDP repo mysql struct
type mysqlIdpRepo struct {
}

func (m *mysqlIdpRepo) Play() {
	fmt.Println("sdhvfudvbjkudj")
}

func (m *mysqlIdpRepo)Study(){
	fmt.Println("Study")
}
