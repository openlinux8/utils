package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type User struct {
	Id       uint8  `gorm:"primaryKey"`
	Username string `json:"username" gorm:"type:varchar(25)"`
	Password string `json:"password" gorm:"type:varchar(25)"`
}

type Product struct {
	Id   uint8  `gorm:"primaryKey"`
	Type string `gorm:"index;type:varchar(25)"`
	Name string `gorm:"type:varchar(25)"`
}

func init() {
	dsn := "golang:golang@tcp(192.168.2.83:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "golang_",
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	model := db.Migrator()
	if !model.HasTable(&User{}) {
		model.AutoMigrate(&User{})
		var users = []User{{Username: "admin", Password: "admin"}, {Username: "root", Password: "root"}, {Username: "golang", Password: "88888888"}}
		db.Create(&users)
	}
	if !model.HasTable(&Product{}) {
		model.AutoMigrate(&Product{})
		var products = []Product{{Type: "汽车", Name: "宝马"}, {Type: "汽车", Name: "奔驰"}, {Type: "手机", Name: "小米"}, {Type: "手机", Name: "华为"}}
		db.Create(&products)
	}
}

func main() {
	e := echo.New()
	e.GET("/api/product", getProduct)
	e.POST("/api/login", postLogin)
	e.Logger.Fatal(e.Start(":8080"))
}

func getProduct(c echo.Context) error {
	product := new(Product)
	id := c.QueryParam("id")
	fmt.Println(id)
	return c.JSON(200, product)
}

func postLogin(c echo.Context) error {
	var loginInfo map[string]string
	user := new(User)
	c.Bind(user)
	if user.Username == "admin" && user.Password == "admin" {
		loginInfo = map[string]string{
			"message":  "login success",
			"user":     user.Username,
			"password": user.Password,
		}
	} else {
		loginInfo = map[string]string{
			"message":  "login failure",
			"user":     user.Username,
			"password": user.Password,
		}
	}
	return c.JSON(200, loginInfo)
}
