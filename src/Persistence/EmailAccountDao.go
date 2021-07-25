package Persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type EmailAccount struct {
	ID int
	Email string
	EmailPwd string
	EmailHost string
	EmailPort string
}

func GetEmailAccountList() ([]EmailAccount,error){
	result := make([]EmailAccount,0)
	db,_:=sql.Open("mysql","root:****@(127.0.0.1:3306)/TicketsMem")
	err:=db.Ping()
	if err!=nil{
		fmt.Println("Error occured during connecting db.")
	}
	defer db.Close()
	rows,_:= db.Query("SELECT * FROM EmailAccount")
	var ID int
	var Email string
	for rows.Next(){
		rows.Scan(&ID,&Email)
		result = append(result,EmailAccount{ID: ID,Email: Email})
	}
	return result,rows.Err()
}