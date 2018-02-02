package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/jinzhu/gorm"
	"github.com/wangzitian0/golang-gin-starter-kit/articles"
	"github.com/wangzitian0/golang-gin-starter-kit/common"
	"github.com/wangzitian0/golang-gin-starter-kit/users"
	"github.com/gin-gonic/gin"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))

	articles.ArticlesRegister(v1.Group("/articles"))

	r.Run() // listen and serve on 0.0.0.0:8080
}
