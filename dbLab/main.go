package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USER   = "root"
	PWD    = "123456"
	DBIP   = "127.0.0.1"
	DBPORT = "3306"
	DBNAME = "gorm"
)

type Admin struct {
	ID       int64
	User     string
	Password string
}

type Finish struct {
	ID         int
	Callid     string `gorm:"size:50; not null"`
	Subid      string `gorm:"size:100"`
	Aid        int64  `gorm:"index"`
	CreateTime time.Time
}

type Account struct {
	//gorm.Model 是一个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的结构体
	gorm.Model
	Appkey  string `gorm:"type:varchar(15);index:idx_appkey;not null"`
	Company string `gorm:"column:company_name;size:30"`
	Status  int8   `gorm:"default:1"`
}

func (Admin) TableName() string {
	return "vn_finish"
}

func (Account) TableName() string {
	return "vn_account"
}

// 生成随机的字符串
func gainRandomString(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for v := range b {
		b[v] = s[rand.Intn(len(s))]
	}
	return string(b)
}

// 生成md5字符串
func gainRandomMd5String(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func gainData(d chan Admin) {
	for i := 0; i < 20; i++ {
		name := gainRandomString(9)
		pwd := gainRandomMd5String(name)
		data := Admin{User: name, Password: pwd}
		d <- data
	}
	close(d)
}

func main() {
	info := USER + ":" + PWD + "@tcp(" + DBIP + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	fmt.Println(info)

	db, err := gorm.Open("mysql", info)
	defer db.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	} else {
		fmt.Println("mysql connect success")
	}

	// 对连接池的配置
	// 设置最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 设置最大打开连接数
	db.DB().SetMaxOpenConns(100)
	db.Create(&Admin{})
	if !db.HasTable(Admin{}.TableName()) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Admin{})
	}
	if !db.HasTable(&Finish{}) {
		// 创建表
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Finish{})
	}
	if !db.HasTable(&Account{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Account{})
	}

	// rand.Seed(time.Now().UnixNano())
	// use1 := "Test" + strconv.Itoa(rand.Intn(99))
	// pwd := func(s string) string {
	// 	str := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	// 	return str
	// }(use1)
	// data := Admin{User: use1, Password: pwd}
	// db.NewRecord(data)
	// db.Create(&data)
	// 模拟业务逻辑
	// 定义一个channel 缓冲为10
	// datas := make(chan Admin, 20)
	// // 生产数据
	// go gainData(datas)
	// // 插入数据
	// for v := range datas {
	// 	// NewRecord check if value's primary key is blank
	// 	db.NewRecord(v)
	// 	// Create insert the value into database
	// 	db.Create(&v)
	// }
	//
	// a1 := Admin{}
	// db.Select([]string{"id", "user", "password"}).
	// 	Where("id = ? AND user = ?", 1, "Test51")
	// fmt.Println(a1)
	//
	// a2 := Admin{}
	// db.Last(&a2)
	// fmt.Println(a2)
	// // 获取所有记录
	// var admins []Admin
	// db.Where("id > 20").Find(&admins)
	// fmt.Println(admins)
	// // 更新操作
	// a3 := Admin{}
	// a3.User = "Test-demo"
	// a3.ID = 20
	// a3.Password = a1.Password
	// // save 更新或者保存
	// db.Save(&a3)
	// // 更新部分字段
	// a4 := Admin{}
	// a4.ID = 1
	// //更新字段user
	// db.Model(&a4).Update("user", "demo99")
	// // 删除操作
	// a5 := Admin{}
	// a5.ID = 2
	// errs := db.Delete(&a5).Error
	// if errs == nil{
	// 	fmt.Println("delete success")
	// }

	var users []Admin
	fields := "id,user,password"
	db.Select(fields).Where("id > ?", 10).Find(&users)
	// 遍历切片
	for k, v := range users {
		fmt.Println(k, v)
	}
}
