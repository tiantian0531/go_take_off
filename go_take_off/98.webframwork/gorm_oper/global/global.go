package global

import (
	users "first/98.webframwork/gorm_oper/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
	username := "sjzy_dev_user"                   //账号
	password := "818f83e5"                        //密码
	host := "mysql-dev.default.svc.cluster.local" //数据库地址，可以是Ip或者域名
	port := 3306                                  //数据库端口
	Dbname := "test_b"                            //数据库名
	timeout := "10s"                              //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_",
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	DB = db
	// 连接成功
	fmt.Println(db)
}

func Migrate() {
	user := &users.UserModel{}
	err := DB.Debug().AutoMigrate(user)
	if err != nil {
		fmt.Println(err)
	}
}
