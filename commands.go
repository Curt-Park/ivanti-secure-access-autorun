package main

import (
	"log"
	"runtime"

	applescripts "github.com/Curt-Park/ivanti-secure-access-autorun/applescripts"
)

type Commands struct {
	vpnExecutor      []string
	connectBtnFinder []string
	vpnPwFinder      []string
	otpNumberFinder  []string
	otpInputFinder   []string
}

func InitCommands() *Commands {
	var vpnExecutor, connectBtnFinder, vpnPwFinder, otpNumberFinder, otpInputFinder []string

	// MacOS
	if runtime.GOOS == "darwin" {
		vpnExecutor = []string{"bash", "-c", "open -n /Applications/Ivanti\\ Secure\\ Access.app"}
		connectBtnFinder = []string{"osascript", "-e", applescripts.ConnectBtnFinderApple}
		vpnPwFinder = []string{"bash", "-c", "security find-generic-password -s 'auto.vpn' -w"}
		otpNumberFinder = []string{"osascript", "-e", applescripts.OtpNumberFinderApple}
		otpInputFinder = []string{"osascript", "-e", applescripts.OtpInputFinderApple}
	} else {
		log.Fatal("Not supported OS:", runtime.GOOS)
	}
	return &Commands{
		vpnExecutor:      vpnExecutor,
		connectBtnFinder: connectBtnFinder,
		vpnPwFinder:      vpnPwFinder,
		otpNumberFinder:  otpNumberFinder,
		otpInputFinder:   otpInputFinder,
	}
}
