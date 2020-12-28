package main

import(
	"time"
	"fmt"
)

func main(){
	for {
		fmt.Println("test start:")
		for i :=0; i<=10; i++{	
			if i == 10 { i=0 }
			fmt.Println("test i:",i)
			time.Sleep(2 * time.Second)
			continue
		}
	}

	// j :=10
	// go func(){
	// 	for i := 0; i<=j; i++{
	// 		if i == j { i=0 }
	// 		fmt.Println("test i:",i)
	// 		time.Sleep(1*time.Second)		
	// 	}
	// }()

	// time.Sleep(15*time.Second)
	// j=20
	select{}
}