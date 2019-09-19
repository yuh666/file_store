package db

import (
	"file_store/db/mysql"
	"database/sql"
)

func InsertFileMeta(fileHash string, fileName string, fileSize int64, fileAddr string) bool {
	stmt, err := mysql.GetDB().Prepare("insert into tbl_file (file_sha1, file_name, file_size, file_addr,status) value (?,?,?,?,?)")
	if err != nil {
		return false
	}
	_, err = stmt.Exec(fileHash, fileName, fileSize, fileAddr, 1)
	if err != nil {
		return false
	}
	defer stmt.Close()
	return true
}

type FileEntity struct {
	FileSha1 sql.NullString
	FileName sql.NullString
	FileAddr sql.NullString
	FileSize sql.NullInt64
}

func GetFileMeta(fileHash string) (*FileEntity, error) {
	stmt, err := mysql.GetDB().Prepare("select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1 = ? and status = 1 limit 1")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(fileHash)
	defer stmt.Close()
	fe := FileEntity{}
	err = row.Scan(&fe.FileSha1, &fe.FileName, &fe.FileSize, &fe.FileAddr)
	if err != nil {
		return nil, err
	}
	return &fe, nil
}
