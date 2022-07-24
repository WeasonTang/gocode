package main

import (
	"fmt"
	"sort"
	"math/rand"
)

//1.声明Hero结构体
type Hero struct {
	Name string
	Age int
}

//2.声明Hero结构体切片类型
type HeroSlice []Hero

//3.实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//下面这行代码 <=> 上面三行代码
	hs[i], hs[j] = hs[j], hs[i]
}


func main() {
	//实现对hero结构体的排序：sort.Sort(data Interface)
	
	//对intslice切片进行排序
	var intslice = []int{22, -1, 10, 7, 99}
	//1.冒泡排序...
	//2.系统提供的方法 func Ints(a []int)
	sort.Ints(intslice)
	fmt.Println(intslice)

	//对结构体切片进行排序
	//1.冒泡排序...
	//2.也可以使用系统提供的方法 func Sort(data Interface)
	
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", i),
			Age : rand.Intn(100),
		}
		//将 hero append到heroes
		heroes = append(heroes, hero)
	}
	for _ , v := range heroes {
		fmt.Println(v)
	}

	//调用sort
	sort.Sort(heroes)
	fmt.Println("排序后。。。")
	for _ , v := range heroes {
		fmt.Println(v)
	}
}