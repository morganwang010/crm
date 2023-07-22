package main

import (
	"fmt"
	"gogin/model"
	"gogin/routes"
)

func main() {

	// first := models.DB.Table("dr.users").First(1)
	// result := models.DB.Table("dr.user").Select(`*`).Where("employee_number = ?", "wangjw30").First(1)
	// result := models.DB.First(1)
	model.DataBase()
	r := routes.NewRouter()
	fmt.Println("result")
	r.Run()
}
