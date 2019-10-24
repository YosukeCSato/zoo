package main

import ( 
	"fmt"

	"zoo/animals"
)

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出し */
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
