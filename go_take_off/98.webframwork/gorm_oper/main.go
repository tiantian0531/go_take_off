package main

import (
	"encoding/json"
	"first/98.webframwork/gorm_oper/global"
	users "first/98.webframwork/gorm_oper/models"
	"fmt"
	"strconv"
)

func main() {
	global.Connect()
	QueryUser()
	//name := "sa"
	//addr := "湖北省武汉市"
	//user := users.UserModel{
	//	Name:    name,
	//	Age:     18,
	//	Address: &addr,
	//	Sex:     1,
	//}
	//global.DB.Create(&user)

}

func BatchAdd() {
	var list []users.UserModel
	addr := "湖北省武汉市"
	for i := 0; i < 100; i++ {
		s := i%2 == 0
		var sex byte = 1
		if !s {
			sex = 0
		}
		list = append(list, users.UserModel{
			Name:    fmt.Sprintf("机器人 %s 号", strconv.Itoa(i+1)),
			Age:     i + 1,
			Sex:     sex,
			Address: &addr,
		})
	}
	global.DB.Debug().Create(&list)
}

func QueryUser() {

	var u users.UserModel
	global.DB.Debug().Take(&u)
	fmt.Println(u)

	global.DB.Debug().First(&u)
	fmt.Println(u)

	u.Id = 0
	global.DB.Debug().Last(&u)
	fmt.Println(u)

	var us []users.UserModel
	global.DB.Debug().Find(&us, []int{104, 105, 106})
	for _, u := range us {
		data, _ := json.Marshal(u)
		fmt.Println(string(data))
	}

	var us1 []users.UserModel
	global.DB.Debug().Select("Id", "Name").Find(&us1, "name like ?", "机器人%")
	for _, u1 := range us1 {
		data, _ := json.Marshal(u1)
		fmt.Println(string(data))
	}
}
