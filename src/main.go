package main

import (
	"GeekTimeHomework/src/Persistence"
	"fmt"
)

func main(){
	emails,err := Persistence.GetEmailAccountList()
	if err!=nil{
		for _,email := range emails{
			fmt.Println(email.ID,email.Email)
		}
	}
	fmt.Println("Done")
}