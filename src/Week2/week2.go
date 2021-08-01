package Week2

import (
	"GeekTimeHomework/src/Week2/Persistence"
	"fmt"
)

func Do(){
	emails,err := Persistence.GetEmailAccountList()
	if err==nil{
		for _,email := range emails{
			fmt.Println(email.ID,email.Email)
		}
	}
	fmt.Println("Done")
}