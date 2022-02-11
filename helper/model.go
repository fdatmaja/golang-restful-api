package helper

import (
	"belajar_golang_restful_api/model/domain"
	"belajar_golang_restful_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesRensponses []web.CategoryResponse
	for _, category := range categories {
		categoriesRensponses = append(categoriesRensponses, ToCategoryResponse(category))
	}

	return categoriesRensponses
}
