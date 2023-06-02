package main

import (
	"github.com/astrophysics/webtest/controllers"
	"net/http"
)

func main() {
	//u := models.User{ID: 1, FirstName: "Vova", LastName: "I"}
	//fmt.Println(u)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
