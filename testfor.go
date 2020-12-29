package main

import(
	"time"
	"fmt"
	"errors"
	"reflect"
)

func main(){
	// for {
	// 	fmt.Println("test start:")
	// 	for i :=0; i<=10; i++{	
	// 		if i == 10 { i=0 }
	// 		fmt.Println("test i:",i)
	// 		time.Sleep(2 * time.Second)
	// 		continue
	// 	}
	// }

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
	// i :=8
	// switch i{
	// case 1:
	// 	fmt.Println("i is 1")
	// case 2:
	// 	fmt.Println("i is 2")
	// case 8:
	// 	fmt.Println("i is 8")
	// default:
	// 	fmt.Println("i is ?")

	// }

	// select 8{
	// case i>3:
	// 	fmt.Println("i > 1")
	// case i<9:
	// 	fmt.Println("i > 2")
	// case i==8:
	// 	fmt.Println("i > 8")
	// default:
	// 	fmt.Println("i > ?")

	// }

	type testF func(int,int)error

	// n :=199
	// var testf1 =func(a int, b int)error{

	// 	if n ==199 {n =299;    return errors.New("n 是199")}
	// 	if n !=199 {return errors.New("n 不是199")}

	// 	if a ==0||b==0 {return errors.New("a与b都不能为0")}
	// 	fmt.Println("a:",a,"b:",b)
	// 	return nil
	// }(0,0)
	// fmt.Println("type:", reflect.TypeOf(testf1), "value:", testf1)


	m :=199
	var testf2 =func(a int, b int)error{

		if m ==199 {m =299;    return errors.New("m 是199")}
		if m !=199 {return errors.New(fmt.Sprintf("%s%d", "m 不是199,而是", m))}

		if a ==0||b==0 {return errors.New("a与b都不能为0")}
		fmt.Println("a:",a,"b:",b)

		return nil
	}

	fmt.Println("type:", reflect.TypeOf(testf2), "value:", testf2(1,2))
	time.Sleep(5*time.Second)
	fmt.Println("type:", reflect.TypeOf(testf2), "value:", testf2(1,2))
	fmt.Println("外层m:",m)

	// testf2 :=func(a int, b int){
	// 	fmt.Println("a:",a,"b:",b)
	// }(2,3)
	//_ =testf2
	//fmt.Println(testf2)

	select{}
}