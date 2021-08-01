package main

import (
	Persistence2 "GeekTimeHomework/src/Week2/Persistence"
	"fmt"
)

func main(){
	emails,err := Persistence2.GetEmailAccountList()
	if err!=nil{
		for _,email := range emails{
			fmt.Println(email.ID,email.Email)
		}
	}
	fmt.Println("Done")
}