package main

import (
	file "web/filemanage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name  string `json:"username"`
	Psw   string `json:"Password"`
	Email string `json:"email"`
	Flag  string `json:"flag"`
}

func main() {
	var person user
	var ppp user
	var netcode string
	filepath := "1.csv"
	server := gin.Default()

	//allow all request
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}

	server.Use(cors.New(config))

	//login主页
	server.POST("/login", func(ctx *gin.Context) {
		// format json to person struct
		ctx.BindJSON(&person)

		flag, _ := file.FileOrder(filepath, "search", person.Name, person.Psw, person.Email)
		ctx.JSON(200, gin.H{
			"flag": flag,
		})
	})

	//url with register endding
	server.POST("/register", func(c *gin.Context) {
		c.BindJSON(&person)
		// check if username,usermail is used
		flagname, flagmail := file.FileOrder(filepath, "check", person.Name, person.Psw, person.Email)
		if flagmail {
			c.JSON(200, gin.H{
				"flagmail": "yes",
			})
		} else if flagname {
			c.JSON(200, gin.H{
				"flagname": "yes",
			})
		} else {
			c.JSON(200, gin.H{
				"flagmail": "no",
				"flagname": "no",
			})
			netcode = file.SendMail(person.Email)
		}
		// onperson = person
		// netcode = file.SendMail(person.Email)
	})

	// server.POST("/check", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"code": netcode,
	// 	})
	// })

	server.POST("/checkmail", func(c *gin.Context) {
		
		c.ShouldBind(&ppp)
		flag := ppp.Flag
		if flag == netcode {
			exist,_:=file.FileOrder(filepath,"check",person.Name,person.Psw,person.Email)
			if !exist{
				file.FileOrder(filepath, "add", person.Name, person.Psw, person.Email)
			}
			c.JSON(200, gin.H{
				"flag": "yes",
			})
		} else {
			c.JSON(200, gin.H{
				"flag": "no",
			})
		}
	})

	server.POST("/findpassword",func(c *gin.Context){
		c.ShouldBind(&ppp)
		// flag:=ppp.Email
		_,emailexist:=file.FileOrder(filepath,"check","","",ppp.Email)
		if !emailexist{
			c.JSON(200,gin.H{
				"flag":"no",
			})
		}else {
			netcode=file.SendMail(ppp.Email)
		}
	})

	server.POST("/inputnewpasswd",func(c *gin.Context){
		c.ShouldBind(&ppp)
		if ppp.Flag==netcode{
			
		}
		// c.JSON()
	})

	// go Run
	server.Run(":60")

}
