package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "aplication/json")
	// writer.Header().Set("Access-Control-Allow-Origin", "*")
	// writer.Header().Set("Access-Control-Allow-Methods", "*")
	// writer.Header().Set("Access-Control-Allow-Headers", "*")
	//writer.Header().Add("Access-Control-Allow-Origin", "*")
	//writer.Header().Add("Vary", "Origin")
	//writer.Header().Add("Vary", "Access-Control-Request-Method")
	//writer.Header().Add("Vary", "Access-Control-Request-Headers")
	//writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	//writer.Header().Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
