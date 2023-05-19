package middleware

import (
	"encoding/json"
	"errors"
)

func RespondError(err error) []byte {
	if err == nil {
		err = errors.New("internal server error")
	}

	bts, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
	return bts
}

func RespondErrMsg(msg string) []byte {
	bts, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{
		Error: msg,
	})
	return bts
}

func Respond(data interface{}) []byte {
	bts, _ := json.Marshal(data)
	return bts
}

func RespondFind(data interface{}, count int64) []byte {
	bts, _ := json.Marshal(struct {
		Data  interface{} `json:"data"`
		Count int64       `json:"count"`
	}{
		Data:  data,
		Count: count,
	})
	return bts
}
