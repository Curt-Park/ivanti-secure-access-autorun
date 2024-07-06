package main

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
)

func executeCmd(args []string, errMsg string) string {
	out, err := exec.Command(args[0], args[1], args[2]).CombinedOutput()
	executionMsg := string(out)
	if err != nil {
		log.Println(executionMsg)
		log.Fatalf("%s: %v\n", errMsg, err)
	}
	return executionMsg
}

func activateVPNWindow() {
	fpid, err := robotgo.FindIds("Ivanti Secure Access")
	if err != nil || len(fpid) == 0 {
		log.Fatalf("Failed finding the process: %v\n", err)
	}
	log.Println("The process ID:", fpid[0])
	err = robotgo.ActivePid(fpid[0])
	if err != nil {
		log.Fatalf("Failed activating the window: %v\n", err)
	}
}

func findBtnLocation(connectBtnFinder []string) []int {
	out := executeCmd(connectBtnFinder, "Failed finding the connection button")
	btnInfos := strings.Split(strings.ReplaceAll(out, "\r\n", "\n"), "\n")
	targetBtnInfo := strings.Split(btnInfos[1], ", ")
	if len(targetBtnInfo) < 6 {
		log.Fatalf("The connection is already established\n")
	}
	var btnLocation []int
	for i := 2; i < len(targetBtnInfo); i++ {
		n, err := strconv.Atoi(targetBtnInfo[i])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v\n", err)
		}
		btnLocation = append(btnLocation, n)
	}
	log.Println("x y w h:", btnLocation)
	return btnLocation
}

func moveMouseToButton(btnLocation []int) {
	robotgo.MouseSleep = 100 // millisecond
	robotgo.Move(btnLocation[0]+btnLocation[2]/2, btnLocation[1]+btnLocation[3]/2)
	robotgo.Click()
	robotgo.Sleep(1)
}

func typePassword(vpnPwFinder []string) {
	out := executeCmd(vpnPwFinder, "Failed finding the vpn password")
	robotgo.TypeStr(out)
	robotgo.KeyTap("enter")
	log.Println("Typed VPN password")
}

func fetchOTPfromSMS(otpNumberFinder []string) string {
	log.Println("Wait for the otp number...")
	out := executeCmd(otpNumberFinder, "Failed fetching the otp number")
	re := regexp.MustCompile(`\[OTP:\s(\d+)\]`)
	match := re.FindStringSubmatch(string(out))
	if len(match) <= 1 {
		log.Fatal("No OTP number found")
	}
	otp := match[1]
	return otp
}

func copyOTPtoClipboard(otp string) {
	err := clipboard.WriteAll(otp)
	if err != nil {
		log.Fatalf("Failed to copy text to clipboard: %v\n", err)
	}
	log.Println("Copied the OTP to the clipboard. Paste it!")
}

func main() {
	commands := InitCommands()

	executeCmd(commands.vpnExecutor, "Failed opening the process")
	activateVPNWindow()
	btnLocation := findBtnLocation(commands.connectBtnFinder)
	moveMouseToButton(btnLocation)
	typePassword(commands.vpnPwFinder)
	otp := fetchOTPfromSMS(commands.otpNumberFinder)
	copyOTPtoClipboard(otp)
}
