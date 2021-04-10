package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type Param struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"Password" binding:"required"`
	Email    string `json:"email"`
	Yesorno  string `json:"flag"`
}

func getRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func SendMail(postmail string) string {
	secretcode := getRandomString(6)
	msg := "验证码为" + secretcode
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "278457198@qq.com", "验证码") // 发件人邮箱，发件人名称
	m.SetHeader("To",                                     // 收件人
		m.FormatAddress(postmail, "Receiver"),
	)
	// println(postmail)
	m.SetHeader("Subject", "验证码") // 主题
	m.SetBody("text/plain", msg)  // 正文

	d := gomail.NewDialer("smtp.qq.com", 465, "278457198@qq.com", "leznqrclbmppbgfj") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail err:", err)
	}
	return secretcode
}

var (
	f   *os.File
	err error
)

/*
	检查文件是否存在在磁盘上
*/
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*
	这个函数用来判断用户是否存在且秘密是否正确
*/
func readfile(filename, name, psw string) bool {
	defer f.Close()
	f, err = os.Open(filename)

	if err != nil {
		println("can not open the file ,err is %+v", err)
	}
	//
	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			println("can not read ,the err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		// println(row[0])
		if row[0] == name && row[1] == psw {
			return true

		}
	}
	return false
}

/*
	新增用户
*/

func addfile(filename, name, psw, mail string) bool {
	defer f.Close()
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		println("can not add ,err is %+v", err)
	}
	writerCsv := csv.NewWriter(f)
	word := []string{name, psw, mail}
	err1 := writerCsv.Write(word)
	if err1 != nil {
		println("can not add ,err is %+v", err1)
	}
	writerCsv.Flush()
	return true
}

/*
	集成到这个函数立面实现
*/
func FileOrder(filename, order, name, psw, mail string) bool {
	if !checkFileIsExist(filename) {
		os.Create(filename)

	}
	if order == "search" {
		flag := readfile(filename, name, psw)
		return flag
	} else if order == "add" {
		flag := addfile(filename, name, psw, mail)
		return flag
	}
	return false
}

func main1() {
	var netcode string
	var person, onperson Param
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                            //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"} //允许请求的方法

	server.Use(cors.New(config))
	// server.Use(Cors.CORS(Cors.Options{Origin: "*"}))
	// server.Use(gin.Recovery())
	// 样本
	server.POST("/login", func(ctx *gin.Context) {
		ctx.BindJSON(&person)
		filepath := "/var/www/localdata/id.csv"
		flag := FileOrder(filepath, "search", person.Username, person.Password, person.Email)
		ctx.JSON(200, gin.H{
			"flag": flag,
		})

	})

	//login
	server.POST("/register", func(c *gin.Context) {
		c.BindJSON(&person)
		// fmt.Println(person.Username)
		onperson = person
		fmt.Println(onperson)
		netcode = SendMail(person.Email)
		fmt.Println(netcode)
	})

	server.POST("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": netcode,
		})
	})

	server.POST("/code", func(c *gin.Context) {
		var ppp Param
		c.ShouldBind(&ppp)
		flag := ppp.Yesorno
		fmt.Println(ppp.Yesorno)
		fmt.Println(ppp)
		if flag == netcode {
			filepath := "/var/www/localdata/id.csv"
			FileOrder(filepath, "add", person.Username, person.Password, person.Email)
			c.JSON(200, gin.H{
				"flag": "yes",
			})
		} else {
			c.JSON(200, gin.H{
				"flag": "no",
			})
		}
	})

	server.Run(":60")
}
