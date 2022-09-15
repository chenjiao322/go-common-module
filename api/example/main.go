package main

import (
	"gitlab.geinc.cn/services/go-common-module/api"
	"mime/multipart"
)

type GetAppidMapListRequest struct {
	Limit  int                   `form:"limit"`
	Offset int                   `form:"offset"`
	Excel  *multipart.FileHeader `form:"excel" comment:"导入的excel文件" faker:"-"`
}

type GetAppidMapListResp struct {
	Count int `json:"count" comment:"总和"`
}

func main() {
	a := &api.Api{
		Uri:     "/hello/123",
		Method:  "GET",
		In:      GetAppidMapListRequest{},
		Out:     GetAppidMapListResp{},
		Comment: "test",
	}
	a.ToMarkdown()
}
