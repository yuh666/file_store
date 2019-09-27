package handler

import (
	"encoding/json"
	"file_store/const"
	"file_store/db"
	"file_store/meta"
	"file_store/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultPath = "/tmp"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	if request.Method == "GET" {
		//读HTML文件返回
		var err error
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			ErrDeal(writer, err, "internal error")
			return
		}
		writer.Write(data)
	} else if request.Method == "POST" {
		path := defaultPath
		//上传文件逻辑
		var err error
		file, header, err := request.FormFile("file")
		if err != nil {
			ErrDeal(writer, err, "internal error")
			return
		}
		defer file.Close()
		if temp := request.Form.Get("path"); temp != "" {
			path = temp
		}
		newFile, err := os.Create(path + "/" + header.Filename)
		if err != nil {
			ErrDeal(writer, err, "internal error")
			return
		}
		defer newFile.Close()
		bytesLen, err := io.Copy(newFile, file)
		if err != nil {
			ErrDeal(writer, err, "internal error")
			return
		}
		log.Println("上传文件完成：" + newFile.Name())
		newFile.Seek(0, 0)
		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: newFile.Name(),
			UploadAt: time.Now().Format(mixall.YYYYMMddHHmmss),
			FileSize: bytesLen,
			FileSha1: util.FileSha1(newFile),
		}
		//meta.UploadFileMeta(fileMeta.FileSha1, &fileMeta)
		_ = meta.UploadFileMetaDB(fileMeta.FileSha1, &fileMeta)
		log.Println("文件元信息更新成功,sha1：" + fileMeta.FileSha1)
		//写入用户文件表
		db.InsertFileUser(request.Form.Get("username"), fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
		//redirect
		http.Redirect(writer, request, "/static/view/home.html", http.StatusPermanentRedirect)
	}
}

func UploadSucHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("upload success"))
}

func QueryFileMetaByHash(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fileHash := request.Form.Get("filehash")
	fileMeta := meta.QueryFileMetaByHash(fileHash)
	fileMeta, err := meta.QueryFileMetaByHashDB(fileHash)
	if err != nil {
		ErrDeal(writer, err, "internal error")
		return
	}
	bytes, _ := json.Marshal(fileMeta)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(bytes)
}

func QueryFileMetaLimit(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	entity, _ := db.LoadFileUserByUserName(request.Form.Get("username"))
	bytes, _ := json.Marshal(entity)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(bytes)
}

func DownloadFileHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	filehash := request.Form.Get("filehash")
	fileMeta := meta.QueryFileMetaByHash(filehash)
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		ErrDeal(writer, err, "internal error")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		ErrDeal(writer, err, "internal error")
	}
	log.Println("开始下载文件：" + fileMeta.FileName)
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", "attachment;filename=\""+fileMeta.FileName+"\"")
	writer.Write(data)
}

func ModifyFileHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	opType := request.Form.Get("op")
	fileHash := request.Form.Get("filehash")
	fileName := request.Form.Get("filename")
	if opType == "0" {
		//改名
		fileMeta := meta.QueryFileMetaByHash(fileHash)
		if fileMeta != nil {
			s := fileMeta.FileName
			fileMeta.FileName = fileName
			log.Println("文件名修改成功：" + s + " -> " + fileName)
		}
	}
	writer.WriteHeader(http.StatusOK)
}

func DeleteFileHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fileHash := request.Form.Get("filehash")
	fileMeta := meta.QueryFileMetaByHash(fileHash)
	if fileMeta == nil {
		writer.WriteHeader(http.StatusOK)
		return
	}
	err := os.Remove(fileMeta.Location)
	if err != nil {
		ErrDeal(writer, err, "internal error")
		return
	}
	meta.DeleteFileMetaByHash(fileHash)
	log.Println("文件删除完成：" + fileMeta.FileName)
	writer.WriteHeader(http.StatusOK)
}

func FastUpload(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	filename := request.Form.Get("filename")
	filehash := request.Form.Get("filehash")
	entity, err := db.GetFileMeta(filehash)
	if err != nil {
		ErrDeal(writer, err, "internal error")
		return
	}
	if entity == nil {
		writer.Write([]byte("秒传失败"))
		return
	}
	//写入用户文件表
	db.InsertFileUser(request.Form.Get("username"), filehash, filename, entity.FileSize.Int64)
	writer.Write([]byte("成功"))
	return
}

func ErrDeal(writer http.ResponseWriter, err error, msg string) {
	log.Println(err)
	writer.WriteHeader(http.StatusInternalServerError)
	io.WriteString(writer, msg)
}
