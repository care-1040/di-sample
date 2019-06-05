package dao

import (
	. "github.com/ha-t2/di-sample/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gormを通してmysqlにアクセスし、todoを取ってくる
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
	DBMS := "mysql"
	USER := "todo"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "app"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	todoDao := &TodoDao{db}
	return todoDao
}
