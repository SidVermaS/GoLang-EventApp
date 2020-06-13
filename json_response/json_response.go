package json_response

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string, err string)	{
	RespondWithJson(w, code, map[string]string{"message": message, "error": err})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{})	{
	response, _:=json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}