package handler

import "net/http"

func RegisterUserHandler(writer http.ResponseWriter, request *http.Request) {
	//get 跳转
	//post 提交加盐
}

func RegisterInHandler(writer http.ResponseWriter, request *http.Request) {
	//教验
	//token
	//updatetoken
	//redirect home.html
}

func GenToken(username string) string {
	//MD5(timestamp+passwd+salt)+ts[:8]
}

func UpdateToken(username, token string) error {

}
