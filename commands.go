package main

import (
	_ "embed"
	"log"
	"runtime"
)

type Commands struct {
	vpnExecution string
	btnFinder    string
	vpnPwFinder  string
	otpFinder    string
}

//go:embed btn_finder.applescript
var btnFinderApple string

//go:embed otp_finder.applescript
var otpFinderApple string

func InitCommands() *Commands {
	var vpnExecution, btnFinder, vpnPwFinder, otpFinder string

	// MacOS
	if runtime.GOOS == "darwin" {
		vpnExecution = "open -n /Applications/Ivanti\\ Secure\\ Access.app"
		btnFinder = "osascript -e " + "'" + btnFinderApple + "'"
		vpnPwFinder = "security find-generic-password -s 'auto.vpn' -w"
		otpFinder = "osascript -e " + "'" + otpFinderApple + "'"
	} else {
		log.Fatal("Not supported OS:", runtime.GOOS)
	}
	return &Commands{vpnExecution: vpnExecution, btnFinder: btnFinder, vpnPwFinder: vpnPwFinder, otpFinder: otpFinder}
}
