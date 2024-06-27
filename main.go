package main

import (
	"fmt"
	_ "github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"log"
	"time"
)

func main() {
	//secret "修改为自己的密钥"
	secret := "xxxxxx"
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		log.Fatalf("生成TOTP错误: %v", err)
	}
	fmt.Println(code)
}
