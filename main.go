package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanziDy/go_book/internal/repository"
	"github.com/xuanziDy/go_book/internal/repository/dao"
	"github.com/xuanziDy/go_book/internal/service"
	"github.com/xuanziDy/go_book/internal/web/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//初始化数据库连接
	db := initDB() //todo 应该弄成全局的？

	//注册路由
	r := gin.Default()
	initUserHandler(db).RegisterRouter(r)

	//运行服务
	if err := r.Run(":8080"); err != nil {
		panic("8080服务启动失败")
	}
}

// 临时放置
func initDB() *gorm.DB {
	//todo 待制作数据库
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	//todo 临时依赖gorm来建表
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

// todo 放在这里不能更奇怪了！！！！
func initUserHandler(db *gorm.DB) *handler.UserHandler {
	d := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(d)
	srv := service.NewUserService(repo)
	u := handler.NewUserHandler(srv)
	return u
}
