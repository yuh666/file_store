package db

import (
	"file_store/db/mysql"
)

func Insert(username, passwd string) bool {
	stmt, err := mysql.GetDB().Prepare("insert into tbl_user (user_name,user_pwd) value (?,?)")
	if err != nil {
		return false
	}
	_, err = stmt.Exec(username, passwd)
	if err != nil {
		return false
	}
	defer stmt.Close()
	return true
}

type UserEntity struct {
	Username string
	Passwd   string
	SignAt   string
}

type Token struct {
	Username string
	Token    string
}

func LoadTokenByUsername(token string) (*Token, error) {
	//stmt query
	//是否存在
	stmt, err := mysql.GetDB().Prepare("select user_name,token from tbl_token where token = ? limit 1")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(token)
	defer stmt.Close()
	fe := Token{}
	err = row.Scan(&fe.Username, &fe.Token)
	if err != nil {
		return nil, err
	}
	return &fe, nil
}

func LoadByUsername(username string) (*UserEntity, error) {
	//stmt query
	//是否存在
	stmt, err := mysql.GetDB().Prepare("select user_name,user_pwd,sign_at from tbl_user where user_name = ? limit 1")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(username)
	defer stmt.Close()
	fe := UserEntity{}
	err = row.Scan(&fe.Username, &fe.Passwd, &fe.SignAt)
	if err != nil {
		return nil, err
	}
	return &fe, nil
}

func UpdateToken(username, token string) bool {

	stmt, err := mysql.GetDB().Prepare("replace into tbl_token (user_name,token) value (?,?)")
	if err != nil {
		return false
	}
	_, err = stmt.Exec(username, token)
	if err != nil {
		return false
	}
	defer stmt.Close()
	return true
}
