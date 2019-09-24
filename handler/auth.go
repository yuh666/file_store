package handler

import (
	"file_store/db"
	"net/http"
)

func HttpInterceptor(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		token := request.Form.Get("token")
		tokenEntity, err := db.LoadTokenByUsername(token)
		if err != nil {
			writer.WriteHeader(http.StatusForbidden)
			writer.Write([]byte("LoginError"))
			return
		}
		username := tokenEntity.Username
		request.Form.Set("username", username)
		f(writer, request)
	}
}
