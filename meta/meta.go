package meta

import (
	"sort"
	"file_store/db"
)

//文件元信息
type FileMeta struct {
	FileSha1 string
	FileName string
	Location string
	UploadAt string
	FileSize int64
}

var fileMetaMap map[string]*FileMeta

func init() {
	fileMetaMap = make(map[string]*FileMeta)
}

func UploadFileMeta(fileSha1 string, fileMeta *FileMeta) {
	fileMetaMap[fileSha1] = fileMeta
}

func UploadFileMetaDB(fileSha1 string, fileMeta *FileMeta) bool {
	return db.InsertFileMeta(fileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.Location)
}

func QueryFileMetaByHashDB(fileSha1 string) (*FileMeta, error) {
	entity, err := db.GetFileMeta(fileSha1)
	if err != nil {
		return nil, err
	}
	fm := FileMeta{
		entity.FileSha1.String,
		entity.FileName.String,
		entity.FileAddr.String,
		"",
		entity.FileSize.Int64,
	}
	return &fm, nil
}

func QueryFileMetaByHash(fileSha1 string) *FileMeta {
	return fileMetaMap[fileSha1]
}

func DeleteFileMetaByHash(fileSha1 string) {
	delete(fileMetaMap, fileSha1)
}

func QueryFileMetaLimit(limit int) []*FileMeta {
	l := len(fileMetaMap)
	if l == 0 {
		return []*FileMeta{}
	}
	fileMetas := make([]*FileMeta, 0, len(fileMetaMap))
	for _, v := range fileMetaMap {
		fileMetas = append(fileMetas, v)
	}
	sort.Sort(ByUploadTime(fileMetas))
	if limit == 0 {
		limit = l / 2
	}
	if limit == 0 {
		limit = 1
	}
	if limit > l {
		limit = l
	}
	return fileMetas[:limit]
}
