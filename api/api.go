package api

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Uri     string
	Method  string
	In      interface{}
	Out     interface{}
	Comment string
	Handle  gin.HandlerFunc
}

func (a *Api) RegisterToGin(engine *gin.Engine) {
	engine.Handle(a.Method, a.Uri, a.Handle)
}

type RespWarp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (a *Api) ToMarkdown() {
	_ = faker.FakeData(&a.In)
	_ = faker.FakeData(&a.Out)
	if a.Uri[0] != '/' {
		a.Uri = "/" + a.Uri
	}
	fmt.Printf("### %s\n", a.Uri)
	fmt.Println(a.Comment, "   ")
	fmt.Printf("##### 请求示例\n")
	fmt.Println("```")
	fmt.Printf("%s http://localhost%s\n", a.Method, a.Uri)

	fmt.Println(Doc(a.In).ContentType())
	fmt.Println("```")
	fmt.Println()
	fmt.Println("```json")
	fmt.Println(IndentJson(a.In))
	fmt.Println("```")

	fmt.Printf("##### 返回示例\n")
	fmt.Println()
	fmt.Println("```json")
	fmt.Println(IndentJson(RespWarp{Code: 0, Data: a.Out}))
	fmt.Println("```")
	fmt.Println()
	fmt.Println("请求参数:")
	fmt.Println()
	fmt.Println(Doc(a.In).ToMarkDownIn())
	fmt.Println("返回参数:")
	fmt.Println()
	fmt.Println(Doc(a.Out).ToMarkDownOut())
}

func IndentJson(s interface{}) string {
	out, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return "{}"
	}
	return string(out)
}
