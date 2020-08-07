package model

//Pagination 分页参数
type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"pagesize"`
	Total    uint `json:"total"`
}

//DBCount 用于gin mysql 原生SQL 的计数方法
type DBCount struct {
	Num int `gorm:"column:num"`
}
