package mysql

//create table tbl_user_file (
//id          int(11)     NOT NULL PRIMARY KEY AUTO_INCREMENT,
//user_name   varchar(64) NOT NULL             DEFAULT ''
//COMMENT '用户名',
//file_sha1   varchar(64) NOT NULL             DEFAULT ''
//COMMENT '文件hash',
//file_size   bigint(20)  NOT NULL             DEFAULT '0'
//COMMENT '文件大小',
//upload_at   datetime    NOT NULL             DEFAULT current_timestamp
//COMMENT '上传时间',
//last_update datetime    NOT NULL             DEFAULT current_timestamp
//ON UPDATE CURRENT_TIMESTAMP
//COMMENT '最近更新时间',
//status      int(11)     NOT NULL             DEFAULT '0'
//COMMENT '文件状态',
//unique key uni_user_file(user_name, file_sha1),
//key idx_status(status),
//key idx_user_id (user_name)
//)

func InsertFileUser(username, fileSha1, fileName string, fileSize int64) bool {
	stmt, err := db.Prepare("insert into tbl_user_file(user_name,file_sha1,file_size,file_name) value (?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, e := stmt.Exec(username, fileSha1, fileSize, fileName)
	if e != nil {
		return false
	}
	return true
}

type FileUser struct {
	UserName    string
	FileName    string
	FileHash    string
	FileSize    int
	UploadAt    string
	LastUpdated string
}

func LoadFileUserByUserName(username string) ([]*FileUser, error) {
	stmt, err := db.Prepare("select user_name,file_sha1,file_size,upload_at,file_name,last_update from tbl_user_file where user_name=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	fileusers := make([]*FileUser, 0)
	for rows.Next() {
		f := FileUser{}
		rows.Scan(&f.UserName, &f.FileHash, &f.FileSize, &f.UploadAt, &f.FileName, &f.LastUpdated)
		fileusers = append(fileusers, &f)
	}
	return fileusers, nil
}
