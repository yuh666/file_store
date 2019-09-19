package meta

import (
	"time"
	"file_store/const"
)

type ByUploadTime []*FileMeta

func (b ByUploadTime) Len() int {
	return len(b)
}

func (b ByUploadTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByUploadTime) Less(i, j int) bool {
	iTime, _ := time.Parse(mixall.YYYYMMddHHmmss, b[i].UploadAt)
	jTime, _ := time.Parse(mixall.YYYYMMddHHmmss, b[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()
}
