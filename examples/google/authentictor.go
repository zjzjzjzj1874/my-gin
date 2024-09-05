package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"log"
)

func main() {
	// Generate a new key for the user
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Cloud-Host",
		AccountName: "user-zhangSan",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Display the key URL
	fmt.Println("Secret Key:", key.Secret())
	fmt.Println("URL for Google Authenticator:", key.URL())

	// Secret Key: ZYTW3HLDFVV47QQN5UBDTL7PKGV6VPNR
	// Authenticator: otpauth://totp/Cloud-Host:user-zhangSan?algorithm=SHA1&digits=6&issuer=Cloud-Host&period=30&secret=ZYTW3HLDFVV47QQN5UBDTL7PKGV6VPNR

	secret := key.Secret() // Replace with the user's secret key
	userOTP := "628870"    // Replace with the OTP input from the user
	if validateOTP(secret, userOTP) {
		fmt.Println("OTP is valid!")
	} else {
		fmt.Println("Invalid OTP.")
	}
}

// 创建待扫描的二维码
func genQR(issuer, account string) {
	// Generate a new key for the user
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Display the key URL
	fmt.Println("Secret Key:", key.Secret())
	fmt.Println("URL for Google Authenticator:", key.URL())
}

// 验证授权码
func checkAuthenticator(secret, otp string) bool {
	return validateOTP(secret, otp)
}

func validateOTP(secret, otp string) bool {
	return totp.Validate(otp, secret)
}
