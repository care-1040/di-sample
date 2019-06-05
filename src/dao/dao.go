package dao

import (
	. "github.com/care-1040/di-sample/src"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

// gormを通してmysqlにアクセスし、todoを取ってくる
type TodoDaoInterface interface {
	Get(id int) Todo
}
type TodoDao struct {
	db *gorm.DB
}
func (t *TodoDao) Get(id int) Todo {
	todo := Todo{}
	t.db.First(&todo, id)
	return todo
}

// TodoDaoのコンストラクタ
func NewTodoDao() *TodoDao {
	DBMS     := "mysql"
	USER     := "todo"
	PASS     := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME   := "app"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=true"
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	todoDao := &TodoDao{db}
	return todoDao
}