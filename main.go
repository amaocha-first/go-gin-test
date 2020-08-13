package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

//db初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("couldn't open db")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//追加
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("couldn't open db")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//全件取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("couldn't open db")
	}
	todos := []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//一件取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("couldn't open db")
	}
	var todo Todo
	db.First(&Todo, id)
	db.Close()
	return todo
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
