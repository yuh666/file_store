package handler

import (
	"encoding/json"
	"file_store/db"
	"file_store/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	SALT = "&&&&"
)

func RegisterUserHandler(writer http.ResponseWriter, request *http.Request) {
	//get 跳转
	//post 提交加盐
	if request.Method == "GET" {
		//读HTML文件返回
		var err error
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			ErrDeal(writer, err, "internal error")
			return
		}
		writer.Write(data)
	} else {
		request.ParseForm()
		username := request.Form.Get("username")
		password := request.Form.Get("password")
		newPasswd := util.Sha1([]byte(password + SALT))
		db.Insert(username, newPasswd)
		writer.Write([]byte("SUCCESS"))
	}
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	//教验
	//token
	//updatetoken
	//redirect home.html
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	entity, err := db.LoadByUsername(username)
	if err != nil {
		ErrDeal(writer, err, "internal error")
		return
	}
	newPasswd := util.Sha1([]byte(password + SALT))
	if entity.Passwd != newPasswd {
		writer.Write([]byte("Login Fail"))
		return
	}
	token := GenToken(username)
	db.UpdateToken(username, token)
	m := map[string]interface{}{}
	m["Token"] = token
	m["Username"] = username
	m["Location"] = "http://" + request.Host + "/static/view/home.html"
	r := Result{0, "ok", m}
	j, _ := json.Marshal(r)
	writer.Write(j)
}

func UserInfoHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	entity, err := db.LoadByUsername(username)
	if err != nil {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("LoginError"))
		return
	}
	m := map[string]interface{}{}
	m["Username"] = entity.Username
	m["SignupAt"] = entity.SignAt
	r := Result{0, "ok", m}
	j, _ := json.Marshal(r)
	writer.Write(j)
}

func GenToken(username string) string {
	//MD5(timestamp+passwd+salt)+ts[:8]
	ts := fmt.Sprintf("%d", time.Now().UnixNano())
	return util.MD5([]byte(ts+username+SALT)) + ts[:8]
}
