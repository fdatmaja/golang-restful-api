package app

import (
	"belajar_golang_restful_api/controller"
	"belajar_golang_restful_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	// router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Header.Get("Access-Control-Request-Method") != "" {
	// 		// Set CORS headers
	// 		header := w.Header()
	// 		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	// 		header.Set("Access-Control-Allow-Origin", "*")
	// 	}

	// 	// Adjust status code to 204
	// 	w.WriteHeader(http.StatusNoContent)
	// })

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
