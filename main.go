package main

import (
	"belajar_golang_restful_api/app"
	"belajar_golang_restful_api/controller"
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/middleware"
	"belajar_golang_restful_api/repository"
	"belajar_golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositroy()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	corsHandle := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},

		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: corsHandle.Handler(middleware.NewAuthMiddleware(router)),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	// client := &http.Client{}

	// req, _ := http.NewRequest("GET", "https://api.rajaongkir.com/starter/province", nil)
	// req.Header.Add("key", "a20d8602d25e7f5386647f9e9c660f44")
	// resp, err := client.Do(req)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer resp.Body.Close()
	// resp_body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(resp.Status)
	// fmt.Println(string(resp_body))

}
