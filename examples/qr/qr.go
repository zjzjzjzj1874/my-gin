package main

import (
	"github.com/boombuler/barcode"
	"image/png"
	"os"

	"github.com/boombuler/barcode/qr"
)

func main() {
	// OTP Auth URL
	otpAuthURL := "otpauth://totp/Cloud-Host:superadmin?algorithm=SHA1&digits=6&issuer=Cloud-Host&period=30&secret=UJI3LKDBQRDFPXWXYYT4PGPKTJENNFTR"
	//otpAuthURL := "otpauth://totp/Cloud-Host:superadmin?algorithm=SHA1&digits=6&issuer=Cloud-Host&period=60&secret=QHHBHE7PS4TC55AQARL4B6JWCDSHEXPL"
	//otpAuthURL := "otpauth://totp/Cloud-Host:user-zhangSan?algorithm=SHA1&digits=6&issuer=Cloud-Host&period=30&secret=ZYTW3HLDFVV47QQN5UBDTL7PKGV6VPNR"

	// Create the barcode
	qrCode, _ := qr.Encode(otpAuthURL, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
