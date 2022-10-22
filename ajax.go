package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

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

func main() {
	db, _ = sql.Open("mysql", "golang:golang@tcp(192.168.2.83:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local")
	e := echo.New()
	e.GET("/api/product", getProduct)
	e.POST("/api/login", postLogin)
	e.Logger.Fatal(e.Start(":8080"))
}

func getProduct(c echo.Context) error {
	product := new(Product)
	id := c.QueryParam("id")
	err := db.QueryRow("select type,name from golang_products where id="+id).Scan(&product.Type, &product.Name)
	//err := db.QueryRow("select type,name from golang_products where id=?", id).Scan(&product.Type, &product.Name)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(200, product)
}

func postLogin(c echo.Context) error {
	var username, password string
	var loginInfo map[string]string
	user := new(User)
	c.Bind(user)
	err := db.QueryRow("select username,password from golang_users where username=\""+user.Username+"\""+" and password=\""+user.Password+"\"").Scan(&username, &password)
	//err := db.QueryRow("select username,password from golang_users where username=? and password=?", user.Username, user.Password).Scan(&username, &password)
	fmt.Println(username, password)
	if err != nil {
		fmt.Println(err)
	}
	if username != "" && password != "" {
		loginInfo = map[string]string{
			"message":   "login success",
			"user":      user.Username,
			"password":  user.Password,
			"quser":     username,
			"qpassword": password,
		}
	} else {
		loginInfo = map[string]string{
			"message":   "login failure",
			"user":      user.Username,
			"password":  user.Password,
			"quser":     username,
			"qpassword": password,
		}
	}
	return c.JSON(200, loginInfo)
}
