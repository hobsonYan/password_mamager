package model

type AddForm struct {
	Id    string `form:"id"`
	Key   string `form:"key"`
	Value string `form:"value"`
}

type Password struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListForm struct {
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}

type ListVo struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}
