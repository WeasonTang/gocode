package main

import (
	"fmt"
	"gocode/familyaccount/utils"
)

func main() {
	utils.NewFamilyAccount().MainMenu()
	fmt.Println("已退出家庭记账软件的额使用...")
}
