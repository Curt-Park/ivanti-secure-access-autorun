package main

import (
	_ "embed"
	"log"
	"runtime"
)

type Commands struct {
	vpnExecutor  []string
	btnFinder    []string
	vpnPwFinder  []string
	otpFinder    []string
}

//go:embed btn_finder.applescript
var btnFinderApple string

//go:embed otp_finder.applescript
var otpFinderApple string

func InitCommands() *Commands {
	var vpnExecutor, btnFinder, vpnPwFinder, otpFinder []string

	// MacOS
	if runtime.GOOS == "darwin" {
		vpnExecutor = []string{"bash", "-c", "open -n /Applications/Ivanti\\ Secure\\ Access.app"}
		btnFinder = []string{"osascript", "-e", btnFinderApple}
		vpnPwFinder = []string{"bash", "-c", "security find-generic-password -s 'auto.vpn' -w"}
		otpFinder = []string{"osascript", "-e", otpFinderApple}
	} else {
		log.Fatal("Not supported OS:", runtime.GOOS)
	}
	return &Commands{vpnExecutor: vpnExecutor, btnFinder: btnFinder, vpnPwFinder: vpnPwFinder, otpFinder: otpFinder}
}
