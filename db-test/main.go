package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"git.ibbd.net/dsp/go-config"
)

var (
	db   *sql.DB
	stmt = map[string]*sql.Stmt{
		"one":  nil,
		"many": nil,
	}
)

func init() {
	db, err := sql.Open("mysql", goConfig.MysqlAddress)
	if err != nil {
		println("open error")
	}

	stmt["one"], _ = db.Prepare("SELECT `id`, `name` FROM `ad_plan` WHERE `id` = ?")
	stmt["many"], _ = db.Prepare("SELECT `id`, `name` FROM `ad_plan`")
}

func main() {
	rows, err := stmt["one"].Query(3)
	if err != nil {
		println("rows err")
	}
	defer rows.Close()

	var (
		id   uint32
		name string
	)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			println("scan err")
		}
		fmt.Println(id)
		fmt.Println(name)
	}
	if err = rows.Err(); err != nil {
		println("rows err")
	}

	rows2, err := stmt["many"].Query()
	if err != nil {
		println("rows err")
	}
	defer rows2.Close()

	for rows2.Next() {
		err = rows2.Scan(&id, &name)
		if err != nil {
			println("scan err")
		}
		fmt.Println(id)
		fmt.Println(name)
	}
	if err = rows2.Err(); err != nil {
		println("rows err")
	}

	fmt.Println("Over")
}
