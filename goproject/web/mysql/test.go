package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid      int    `json:"uid" sql:"uid"`
	Username string `json:"username" sql:"username"`
	Password string `json:"password" sql:"password"`
	Email    string `json:"email" sql:"email"`
}

/*
连接数据库，输入数据库名字，返回指向和error
*/
func LinkMYSQL(databasename string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "saber:suxiaobai@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(("use " + databasename))
	if err != nil {
		fmt.Println(err.Error())
	}
	return db, err
	//well done
}

/*
插入元组，输入数据库指向，表名，以及User数据结构，返回error
*/
func Insertrows(db *sql.DB, tablename string, user *User) error {

	stmt, err := db.Prepare("insert into " + tablename + " set uid=?,username=?,password=?,email=?")
	if err != nil {
		fmt.Println(err)
	}
	searchuidmax(db, tablename, &user.Uid)
	user.Uid++
	_, err = stmt.Exec(user.Uid, user.Username, user.Password, user.Email)
	if err != nil {
		fmt.Println(err)
	}
	return err
	//well done
}

/*
查找当前数据库中最大的uid，因为uid是primary key，不可重复
*/
func searchuidmax(db *sql.DB, tablename string, user *int) {
	db.QueryRow("select max(uid) from " + tablename + ";").Scan(user)
	//well done
}

/*
查找姓名或者邮箱是否已注册
*/
func uniquesearch(db *sql.DB, tablename, searchmode, content string) bool {
	// var flag int
	err := db.QueryRow("select email from " + tablename + " where " + searchmode + " = '" + content + "';").Scan(&content)
	// fmt.Println(err)
	return err == nil
	//well done
}

/*
修改密码
*/
func ChangePsw(db *sql.DB, username, newpassword string) {
	db.QueryRow("update user set password = '" + newpassword + "' where username= '" + username + "';")
	//well done
}

/*
查找是否存在用户
*/
func IsExistUser(db *sql.DB, username string) bool {
	var fuck string
	err := db.QueryRow("select username from user where username= '" + username + "';").Scan(&fuck)
	fmt.Println(err)
	return err == nil
	//well done
}
func main() {
	var user User
	user.Uid = 2
	user.Username = "archer"
	user.Password = "suxiaobai"
	user.Email = ""
	// test()
	db, err := LinkMYSQL("userbase")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var flag int
	searchuidmax(db, "user", &flag)
	fmt.Println(flag)
	ChangePsw(db, "archer", "123")
}
