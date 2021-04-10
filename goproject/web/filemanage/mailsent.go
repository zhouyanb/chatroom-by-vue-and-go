package web

import (
	"log"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

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

// func main() {
// 	code := SendMail("278457198@qq.com")
// 	println(code)
// }
