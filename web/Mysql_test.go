package web

import (
	"testing"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"strconv"
	"math/rand"
)

var MysqlUrl = "root:root@tcp(172.17.6.3:3306)/go_db?charset=utf8"

func TestMysql(t *testing.T) {
	ManyThreadTxWithGo()
}

func ManyThreadTxWithGo() {
	fmt.Println(time.Now().Unix())

	db, err := sql.Open("mysql", MysqlUrl)
	db.SetMaxOpenConns(1001)
	checkErr(db, err)
	del, err := db.Exec("delete from user")
	fmt.Println(del.RowsAffected())
	t1 := time.Now().Unix()

	ch := make(chan int, 100)
	intertnumber := 0
	go func() {
		for i := range ch {
			intertnumber += i
			//fmt.Println(intertnumber)
		}
	}()
	for i := 0; i < 1000; i++ {
		value := i
		go func(i int) {
			stmt, _ := db.Prepare("INSERT INTO user (username, password) VALUES (?,?)")
			ii := 0
			//fmt.Println(ii)
			tx, _ := db.Begin()
			for k := value; k < 100000; k = k + 1000 {
				rand.Seed(time.Now().Unix())
				x := rand.Int63()
				res, err := tx.Stmt(stmt).Exec(strconv.FormatInt(x, 36)+string(i*i*i), strconv.FormatInt(x, 36)+string(i+59+i))
				checkErr(res, err)
				affect, err := res.RowsAffected()

				ii += int(affect)
				ch <- int(affect)
				checkErr(affect, err)
			}
			tx.Commit()
		}(i)
	}
	for intertnumber < 100000 {
		time.Sleep(1 * time.Second)
	}
	close(ch)
	t2 := time.Now().Unix()
	fmt.Println(t2 - t1)
	fmt.Println(intertnumber)
}

func ManyThreadNoTxWithGo() {
	fmt.Println(time.Now().Unix())

	intertnumber := int64(0)
	db, err := sql.Open("mysql", MysqlUrl)
	db.SetMaxIdleConns(10)
	checkErr(db, err)
	del, _ := db.Exec("delete from user")
	fmt.Println(del.RowsAffected())
	t1 := time.Now().Unix()
	for i := 0; i < 10; i++ {
		value := i
		go func() {
			//db, _ := sql.Open("mysql", MysqlUrl)
			stmt, _ := db.Prepare("INSERT INTO user (username, password) VALUES (?,?)")
			for k := value; k < 1000; k = k + 10 {
				rand.Seed(time.Now().Unix())
				x := rand.Int63()
				res, err := stmt.Exec(strconv.FormatInt(x, 36)+string(i*i*i), strconv.FormatInt(x, 36)+string(i+59+i))
				checkErr(res, err)
				affect, err := res.RowsAffected()
				intertnumber += affect
				checkErr(affect, err)
			}
			intertnumber = intertnumber + 100
		}()
	}
	for intertnumber < 1000 {
		time.Sleep(1 * time.Second)
	}
	t2 := time.Now().Unix()
	fmt.Println(t2 - t1)
	fmt.Println(intertnumber)
}

func oneThreadWithTran() {
	db, err := sql.Open("mysql", MysqlUrl)
	checkErr(db, err)
	del, _ := db.Exec("delete from user")
	fmt.Println(del.RowsAffected())

	t1 := time.Now()
	fmt.Println(t1.Unix())
	stmt, err := db.Prepare("INSERT INTO user (username, password) VALUES (?,?)")
	tx, err := db.Begin()
	fmt.Println(tx, err)
	checkErr(stmt, err)
	ress := int64(0)
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().Unix())
		x := rand.Int63()
		res, err := tx.Stmt(stmt).Exec(strconv.FormatInt(x, 36)+string(i*i*i), strconv.FormatInt(x, 36)+string(i+59+i))
		checkErr(res, err)
		affect, err := res.RowsAffected()
		ress += affect
		checkErr(affect, err)
	}
	tx.Commit()
	t2 := time.Now()
	fmt.Println(t2.Unix())
	fmt.Println(t2.Unix() - t1.Unix())
	fmt.Println(ress)
}

func oneThread() {
	db, err := sql.Open("mysql", MysqlUrl)
	checkErr(db, err)
	del, _ := db.Exec("delete from user")
	fmt.Println(del.RowsAffected())

	t1 := time.Now()
	fmt.Println(t1.Unix())
	stmt, err := db.Prepare("INSERT INTO user (username, password) VALUES (?,?)")
	checkErr(stmt, err)
	ress := int64(0)
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().Unix())
		x := rand.Int63()
		res, err := stmt.Exec(strconv.FormatInt(x, 36)+string(i*i*i), strconv.FormatInt(x, 36)+string(i+59+i))
		checkErr(res, err)
		affect, err := res.RowsAffected()
		ress += affect
		checkErr(affect, err)
	}

	t2 := time.Now()
	fmt.Println(t2.Unix())
	fmt.Println(t2.Unix() - t1.Unix())
	fmt.Println(ress)
}

func checkErr(in interface{}, err error) {
	if err != nil {
		fmt.Println(in)
		panic(err)
	}
}
