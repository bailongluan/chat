package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
func connnectSql()  {
    database, err := sql.Open("mysql", "root:389121@tcp(127.0.0.1:3306)/chat")
    if err != nil {
        fmt.Println(err)
    }
    db=  database
    db.Ping()
}

func insert(username,password string)  {
    sqlStr := "insert into userinfo (username,password) values (?,?)"
    r, err := db.Exec(sqlStr, "liuxiao", "123455")
    if err != nil {
        fmt.Println("insert cmd error", err)
        return
    }

    id, execErr := r.LastInsertId()
    if execErr != nil {
        fmt.Println("execute insert cmd  error", execErr)
        return
    }
    fmt.Println("insert success", id)
}
//  查询单条数据
func querysingle(id int){
    fmt.Println("querysingle")
       rows, err := db.Query("select username,password from userinfo where id = ?", id)
       if nil != err {
         fmt.Println("query db error: ", err.Error())
         return
       }
       //fmt.Println(rows)
       for rows.Next() {
         var username, password string
         //将值存入变量name中
         err = rows.Scan(&username, &password)
         if err != nil {
             fmt.Println(err)
             // panic(err.Error())
         }
         fmt.Println(username, password)
       }
}

func queryAllUsers(){
    fmt.Println("queryAllUsers")
    rows, err := db.Query("select  * from userinfo ")
    if nil != err {
        fmt.Println("query db error: ", err.Error())
        return
    }
    //fmt.Println(rows)
    for rows.Next() {
        var username, password string
        var id int
        err = rows.Scan(&id,&username, &password)
        if err != nil {
            fmt.Println(err)
            // panic(err.Error())
        }
        fmt.Println(id,username, password)
    }
}

func deleteUser(username string)  {
    res, err := db.Exec("delete from userinfo where username=?", username)
       if err != nil {
           fmt.Println("exec failed, ", err)
           return
       }

       row,err := res.RowsAffected()
       if err != nil {
           fmt.Println("rows failed, ",err)
       }
       fmt.Println("delete succ: ",row)
}

func main()  {
    fmt.Println("mysql start")
    fmt.Println("connect mysql")

    connnectSql()
    insert("xiaoqiang","123456")
    querysingle(3)
    queryAllUsers()
    deleteUser("liuxiao")

    db.Close()
}