package api

import (
	"fmt"
	"reflect"
	"strings"
)

type ArgumentDocTable struct {
	Name string
	rows []*ArgumentDocRow
}

func (r ArgumentDocTable) ContentType() string {
	if len(r.rows) == 0 {
		return ""
	}
	switch r.rows[0].Protocol {
	case "json":
		return "Content-Type: application/json"
	case "form":
		return "Content-Type: multipart/form-data"
	default:
		panic("unknown protocol")
	}
}

func (r ArgumentDocTable) ToMarkDownIn() string {
	if len(r.rows) == 0 {
		return ""
	}
	ans := ""
	ans += "|请求字段|类型|校验规则|必填|说明|\n"
	ans += "|----|----|----|----|----|\n"
	for _, v := range r.rows {
		ans += v.ToMarkDownIn()
	}
	ans += "\n"
	return ans
}

func (r ArgumentDocTable) ToMarkDownOut() string {
	if len(r.rows) == 0 {
		return ""
	}
	ans := ""
	ans += "|返回字段|类型|说明|\n"
	ans += "|----|----|----|\n"
	for _, v := range r.rows {
		ans += v.ToMarkDownOut()
	}
	ans += "\n"
	return ans
}

type ArgumentDocRow struct {
	Key      string
	Type     string
	Required bool
	Rule     string
	Comment  string
	Protocol string
}

func (r *ArgumentDocRow) ToMarkDownIn() string {
	key := strings.TrimLeft(r.Key, ".")
	return fmt.Sprintf("| %s | %s | %s | %v | %s |\n", key, r.Type, r.Rule, r.Required, r.Comment)
}

func (r *ArgumentDocRow) ToMarkDownOut() string {
	key := strings.TrimLeft(r.Key, ".")
	return fmt.Sprintf("| %s | %s | %s |\n", key, r.Type, r.Comment)
}

func Doc(ptr interface{}) ArgumentDocTable {
	reType := reflect.TypeOf(ptr)
	if reType.Kind() != reflect.Struct {
		panic("参数必须是结构体值")
	}
	table := ArgumentDocTable{rows: make([]*ArgumentDocRow, 0)}
	table.Name = reType.String()
	requestDoc(reType, &table, "")
	return table
}

func requestDoc(v reflect.Type, rt *ArgumentDocTable, prefix string) {
	for i := 0; i < v.NumField(); i++ {
		row := &ArgumentDocRow{}
		structField := v.Field(i)
		tag := structField.Tag
		rule := tag.Get("validate")
		if rule == "" {
			rule = tag.Get("binding")
		}
		if tag.Get("json") != "" {
			row.Protocol = "json"
			row.Key = prefix + "." + tag.Get("json")
		} else if tag.Get("form") != "" {
			row.Protocol = "form"
			row.Key = prefix + "." + tag.Get("form")
		}
		row.Type = structField.Type.String()
		if row.Type == "*multipart.FileHeader" {
			row.Type = "file_object"
		}
		row.Comment = tag.Get("comment")
		row.Required = strings.Index(rule, "required") != -1
		row.Rule = rule
		rt.rows = append(rt.rows, row)
		if structField.Type.Kind() == reflect.Struct {
			row.Type = "object"
			requestDoc(structField.Type, rt, row.Key)
		} else if structField.Type.Kind() == reflect.Slice {
			if structField.Type.Elem().Kind() == reflect.Struct {
				row.Type = "array"
				requestDoc(structField.Type.Elem(), rt, row.Key)
			} else {
				row.Type = structField.Type.Elem().Kind().String() + " array"
			}
		}
	}
}
