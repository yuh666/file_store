package main

import (
	"file_store/handler"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//upload
	http.HandleFunc("/file/upload", handler.UploadHandler)
	//uploadSuc
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	//query by hash
	http.HandleFunc("/file/meta", handler.QueryFileMetaByHash)
	//query limit
	http.HandleFunc("/file/query", handler.QueryFileMetaLimit)
	//download
	http.HandleFunc("/file/download", handler.DownloadFileHandler)
	//modify
	http.HandleFunc("/file/modify", handler.ModifyFileHandler)
	//delete
	http.HandleFunc("/file/delete", handler.DeleteFileHandler)

	//user/register
	http.HandleFunc("/user/signup", handler.RegisterUserHandler)
	http.HandleFunc("/user/signin", handler.LoginHandler)
	http.HandleFunc("/user/info", handler.HttpInterceptor(handler.UserInfoHandler))

	//file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Fatalln(http.ListenAndServe(":7771", nil))

}
