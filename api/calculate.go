package api

import (
	"io"
	"net/http"
)

type requestBodyStruct struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type responseBodyStruct struct {
	Result int `json:"result"`
}

type API struct {
	ReadAll   func(reader io.Reader) ([]byte, error)
	Unmarshal func(data []byte, v interface{}) error
	Marshal   func(v interface{}) ([]byte, error)
}

func (api API) CalculateAPI(responseWriter http.ResponseWriter, request *http.Request) {
	requestBody, err := api.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := requestBodyStruct{}
	err = api.Unmarshal(requestBody, &body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	response := responseBodyStruct{
		Result: body.Number1 + body.Number2,
	}
	responseBodyJSON, err := api.Marshal(response)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.Write(responseBodyJSON)
}
