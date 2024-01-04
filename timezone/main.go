package main

import (
	"fmt"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// load location by name
	tzArr := []string{"Asia/Shanghai", "America/New_York", "America/Lo", "Europe/London", "UTC"}
	for _, tz := range tzArr {
		t, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Println("load location tz failed, err: ", err)
			continue
		}
		fmt.Println(tz, time.Now().In(t))
	}

	// set mysql timezone session variable
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&time_zone=UTC")
	if err != nil {
		fmt.Println("open mysql failed, err: ", err)
		return
	}

	for _, tz := range tzArr {
		_, execErr := db.Exec("set time_zone = ?", tz)
		if execErr != nil {
			fmt.Println("set time_zone failed, err: ", execErr)
			return
		}
		var tz string
		db.QueryRow("select @@time_zone").Scan(&tz)
		fmt.Println("mysql time_zone: ", tz)
	}

	db.Close()
}
