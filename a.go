package main

import (
	"database/sql"
	"fmt"
	"log"
	//"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sql_table := `create table if not exists "userinfo"(
                    "uid" integer primary key autoincrement,
		    "username" varchar(64) null,
		    "departname" varchar(64) null,
		    "created" timestamp default(datetime('now','localtime'))
		    );
		    create table if not exists "userdetail"(
		    "uid" int(10) null,
		    "intro" text null,
		    "profile" text null,
		    primary key(uid)
		    );`
	db.Exec(sql_table)
	/*
		tx, err := db.Begin()
		if err != nil {
			fmt.Println(err)
		}
		stmt, err := tx.Prepare("delete from userinfo where uid=?")
		if err != nil {
			fmt.Println(err)
		}
		res, err := stmt.Exec(18)
		if err != nil {
			fmt.Println(err)
		}
		err = tx.Commit()
		//tx.Rollback()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.LastInsertId())*/

	stmt, err := db.Prepare("select username from userinfo where uid=?")
	if err != nil {
		fmt.Println(err)
	}
	row := stmt.QueryRow(10)
	if err != nil {
		fmt.Println(err)
	}
	var username string
	row.Scan(&username)
	fmt.Println(username)

	/*
		rows, err := db.Query("select * from userinfo where uid=?", 10)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		var uid int
		var username string
		var department string
		var created time.Time
		for rows.Next() {
			err = rows.Scan(&uid, &username, &department, &created)
			fmt.Println(err)
			fmt.Println(uid)
			fmt.Println(username)
			fmt.Println(department)
			fmt.Println(created)
		}
					stmt, err := db.Prepare("insert into userinfo(username,departname) values(?,?),(?,?),(?,?)")
					if err != nil {
						fmt.Println(err)
					}
					res, err := stmt.Exec("gin", "研发部", "echo", "运维部", "helo", "运维部")
					if err != nil {
						fmt.Println(err)
					}
					id, err := res.LastInsertId()
					if err != nil {
						fmt.Println(err)
					}
					rows, err := res.RowsAffected()
					fmt.Println(id)
					fmt.Println(rows)
				stmt, err := db.Prepare("update userinfo set username=? where uid=?")
				if err != nil {
					fmt.Println(err)
				}
				res, err := stmt.Exec("athlon", 3)
				if err != nil {
					fmt.Println(err)
				}
				affect, err := res.RowsAffected()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(affect)
			stmt, err := db.Prepare("delete from userinfo where uid=?")
			if err != nil {
				fmt.Println(err)
			}
			res, err := stmt.Exec(4)
			if err != nil {
				fmt.Println(err)
			}
			affect, err := res.RowsAffected()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(affect)
	*/

}
