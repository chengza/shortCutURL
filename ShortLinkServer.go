package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)
import "github.com/google/uuid"

func main() {
	db, _ := sql.Open("sqlite3", "./data")
	db.Exec(`create table if not exists record(
    uuid varchar(64) primary key not null ,
    url varchar(2014) not null ,
    createdTime timestamp not null 
)`)
	statement, _ := db.Prepare(`insert into record (uuid,url,createdTime)`)

	server := gin.Default()
	server.GET("/new", func(context *gin.Context) {
		originURL := context.Query("origin")
		uuid, _ := uuid.NewUUID()
		uuidString := uuid.String()
		var arr []interface{}
		arr = append(arr, originURL)
		arr = append(arr, uuid)
		arr = append(arr, uuidString)
		context.JSON(200, arr)
		statement.Exec(uuidString, originURL, ptypes.TimestampNow())
		context.Done()
	})

	server.Run("0.0.0.0:8008")
}
