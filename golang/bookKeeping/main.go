package main

import (
	"bookKeeping/OO"
	"bookKeeping/db"
	"fmt"
)

func main() {
	// 接受用户输入
	key := ""
	//控制是否退出循环
	loop := true

	for loop {
		fmt.Println("----------- tam-info -----------")
		fmt.Println("            1. TAM列表")
		fmt.Println("            2. 全部客户")
		fmt.Println("            3. TAM对应客户")
		fmt.Println("            4. 退出")
		fmt.Print("请输入 (1-4): ")

		fmt.Scanln(&key)

		db.InitDB()

		switch key {
		case "1":
			fmt.Println("\n----------- TAM列表 -----------")
			OO.GetAllTam()
			fmt.Println("\n")
		case "2":
			fmt.Println("\n----------- 全部客户 -----------")
		case "3":
			fmt.Println("\n----------- TAM对应客户 -----------")
		case "4":
			fmt.Println("\n退出中...")
			loop = false
		default:
			fmt.Println("请输入正确选项。")
		}
	}

}
