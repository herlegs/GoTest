package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"sync"
	"strconv"
	"time"
)

/***
Test Result:
use 'for update' in transaction will block
without 'for update' use transaction isolation level serializable will block
without transaction will not block

In transaction:
use 'for update' where unique_key = sth will just block resulting rows (unique key or primary key)
use non-unique column to select will block entire table
 */

var db *sql.DB
var workNum = 10

func main(){
	//?tx_isolation='SERIALIZABLE'
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer db.Close()
	wait := &sync.WaitGroup{}
	wait.Add(workNum)
	for i := 0; i < workNum; i++ {
		go testUpdateWhere(db, i, wait)
	}
	wait.Wait()
}

func insertWithTransaction(db *sql.DB, index int, wait *sync.WaitGroup){
	defer wait.Done()
	threadName := "thread" + strconv.Itoa(index)
	fmt.Println(threadName,"started")
	tx,err := db.Begin()
	if err != nil {
		fmt.Println("err creating transaction in",threadName)
		return
	}
	rows,err := tx.Query("select * from test where name = ? for update", "name")
	if err != nil {
		fmt.Println("err reading in thread", index)
		tx.Rollback()
		return
	}
	defer rows.Close()
	time.Sleep(time.Duration(index)*time.Second)
	if !rows.Next(){
		fmt.Println(threadName,"see no rows, planning to insert")
		result,err := tx.Exec("insert into test (id, name) values(?,?)", index, "name")
		fmt.Printf(threadName + "insert result: %v | %v\n",result,err)
	}else{
		fmt.Println(threadName,"detected exist")
	}
	tx.Commit()
}

func lockOnUniqueColumn(db *sql.DB, index int, wait *sync.WaitGroup){
	defer wait.Done()
	threadName := "thread" + strconv.Itoa(index)
	fmt.Println(threadName,"started")
	tx,err := db.Begin()
	if err != nil {
		fmt.Println("err creating transaction in",threadName)
		return
	}
	rows,err := tx.Query("select * from test for update")
	if err != nil {
		fmt.Println("err reading in thread", index)
		tx.Rollback()
		return
	}
	defer rows.Close()
	if !rows.Next(){
		fmt.Println(threadName,"see no rows, planning to insert")
		result,err := tx.Exec("insert into test (id, name) values(?,?)", index, threadName)
		fmt.Printf(threadName + "insert result: %v | %v\n",result,err)
	}else{
		fmt.Println(threadName,"detected exist")
	}
	tx.Commit()
}

//for update will block and wait when race condition
func testIncr(db *sql.DB, index int, wait *sync.WaitGroup){
	defer wait.Done()
	threadName := "thread" + strconv.Itoa(index)
	fmt.Println(threadName,"started")
	tx,err := db.Begin()
	if err != nil {
		fmt.Println("err creating transaction in",threadName)
		return
	}
	rows,err := tx.Query("select uni from test where id = ? for update", 0)
	if err != nil {
		fmt.Println("err reading in thread", index)
		tx.Rollback()
		return
	}
	if rows.Next(){
		uni := 0
		rows.Scan(&uni)
		rows.Close()
		fmt.Printf(threadName + "planning to update from %v to %v\n",uni, uni + 1)
		result,err := tx.Exec("update test set uni = ? where id = ?", uni + 1, 0)
		fmt.Printf(threadName + "update result: %v | %v\n",result,err)
		res,_ := tx.Query("select uni from test where id = 0")
		if res.Next(){
			res.Scan(&uni)
		}
		res.Close()
		fmt.Println(threadName+" the new uni:",uni)
	}
	tx.Commit()

}

//update set increment will failed and return error when race condition
//only when there is modification it will lock transaction
func testUpdateIncr(db *sql.DB, index int, wait *sync.WaitGroup){
	defer wait.Done()
	threadName := "thread" + strconv.Itoa(index)
	fmt.Println(threadName,"started")
	tx,_ := db.Begin()
	uni := 0
	res,_ := tx.Query("select uni from test where id = 0")
	if res.Next(){
		res.Scan(&uni)
	}
	res.Close()
	fmt.Println(threadName+" the old uni:",uni)

	fmt.Println(threadName+" will update")
	_,err := tx.Exec("update test set uni = uni + 1 where id = 0")
	fmt.Println(threadName+" has updated")
	if err != nil {
		fmt.Println(threadName+"err updating:"+err.Error())
	}
	//time.Sleep(time.Millisecond*time.Duration(100))

	res,_ = tx.Query("select uni from test where id = 0")
	if res.Next(){
		res.Scan(&uni)
	}
	res.Close()
	fmt.Println(threadName+" the new uni:",uni)
	tx.Commit()
}

func testUpdateWhere(db *sql.DB, index int, wait *sync.WaitGroup){
	defer wait.Done()
	threadName := "thread" + strconv.Itoa(index)
	fmt.Println(threadName,"started")
	tx,_ := db.Begin()
	_,err := tx.Exec("update test set uni = uni + 1 where id = 0 and uni < 5")
	fmt.Println(threadName+" has updated")
	if err != nil {
		fmt.Println(threadName+"err updating:"+err.Error())
	}
	//time.Sleep(time.Millisecond*time.Duration(100))
	uni := 0
	res,_ := tx.Query("select uni from test where id = 0")
	if res.Next(){
		res.Scan(&uni)
	}
	res.Close()
	fmt.Println(threadName+" the new uni:",uni)
	tx.Commit()
}