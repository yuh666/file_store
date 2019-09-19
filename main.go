package main

import (
	"net/http"
	"file_store/handler"
	"log"
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

	log.Fatalln(http.ListenAndServe(":7771", nil))

}
