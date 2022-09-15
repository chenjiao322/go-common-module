package api

import (
	"fmt"
	"mime/multipart"
	"testing"
)

type GetWordReq struct {
	TenantId     string    `json:"tenant_id"`
	Offset       int       `json:"offset"`
	Limit        int       `json:"limit" validate:"required,max=20"`
	Order        int       `json:"order"  comment:"0: 按更新时间正序, 1: 按更新时间倒序"`
	WordLike     string    `json:"word_like" comment:"模糊匹配敏感词"`
	UserNameLike string    `json:"user_name_like" comment:"模糊匹配用户名"`
	UpdateTime   TimeRange `json:"update_time" comment:"查询的更新时间范围"`
	Names        []string  `json:"names" comment:"名字"`
}

type TimeRange struct {
	Lower int64 `json:"lower"`
	Upper int64 `json:"upper"`
}

type GetAppidMapListRequest struct {
	Limit  int                   `form:"limit"`
	Offset int                   `form:"offset"`
	Excel  *multipart.FileHeader `form:"excel" comment:"导入的excel文件" faker:"-"`
}

type GetAppidMapListResp struct {
	Count int `json:"count" comment:"总和"`
}

func Test_main(t *testing.T) {
	table := Doc(GetWordReq{})
	fmt.Println(table.ToMarkDownIn())
	a := &Api{
		Uri:     "/hello/123",
		Method:  "GET",
		In:      GetAppidMapListRequest{},
		Out:     GetAppidMapListResp{},
		Comment: "test",
	}
	a.ToMarkdown()
}
