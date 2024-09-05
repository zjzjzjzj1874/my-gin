package main

import "testing"

func Test_checkAuthenticator(t *testing.T) {
	type args struct {
		secret string
		otp    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "checkAuthenticator", args: args{secret: "ZYTW3HLDFVV47QQN5UBDTL7PKGV6VPNR", otp: "628870"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAuthenticator(tt.args.secret, tt.args.otp); got != tt.want {
				t.Errorf("checkAuthenticator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genQR(t *testing.T) {
	type args struct {
		issuer  string
		account string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "genQR", args: args{issuer: "Cloud-Host", account: "user-zhangSan"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genQR(tt.args.issuer, tt.args.account)
			// Secret Key: ZYTW3HLDFVV47QQN5UBDTL7PKGV6VPNR
			// Authenticator: otpauth://totp/Cloud-Host:user-zhangSan?algorithm=SHA1&digits=6&issuer=Cloud-Host&period=30&secret=ZYTW3HLDFVV47QQN5UBDTL7PKGV6VPNR
		})
	}
}
