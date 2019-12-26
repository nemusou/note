package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
  userId int `json:userId`
  userName string `json:userName`
}

func main() {
  db := gormConnect()
  defer db.Close()

  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  e.GET("/user_test", func(c echo.Context) error {
    users := []User{}
    db.Find(&users)
    return c.JSON(http.StatusOK, users)
  })
  e.Logger.Fatal(e.Start(":1323"))
}

func gormConnect() *gorm.DB {
	db,err := gorm.Open("mysql", "popo:popo@tcp(database:3306)/note")
	if err != nil {
    panic(err.Error())
  }
  return db
}
