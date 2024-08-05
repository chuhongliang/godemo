package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	// userMap := make(map[int]any)
	// userMap[1] = User{Id: 1, Name: "张三"}

	res := []int{}
	for i, v := range slice {
		// slice = append(slice, v)
		res = append(res, v)
		println(i, v)
	}
	fmt.Printf("%v\n", res)
	// fmt.Printf("%v\n", userMap)
	fmt.Println(res)
	size := len(arr)
	fmt.Println(size)
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")

	userArr := []User{}
	userArr = append(userArr, User{Id: 1, Name: "张三"})
	userArr = append(userArr, User{Id: 2, Name: "李四"})
	userArr = append(userArr, User{Id: 3, Name: "王五"})

	jsonByte, _ := json.Marshal(userArr)
	fmt.Println("jsonByte", string(jsonByte))

	var userArr2 []User
	err := json.Unmarshal(jsonByte, &userArr2)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("userArr2", userArr2)

	var userMap map[int]User = make(map[int]User)
	userMap[1] = User{Id: 1, Name: "张三"}
	userMap[2] = User{Id: 2, Name: "李四"}
	userMap[3] = User{Id: 3, Name: "王五"}

	jsonByte2, _ := json.Marshal(userMap)
	fmt.Println("jsonByte2", string(jsonByte2))

	var userMap2 map[int]User = make(map[int]User)
	err = json.Unmarshal(jsonByte2, &userMap2)

	fmt.Println("userMap2", userMap2)

}
