package main

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {
	commands := InitCommands()

	executeCmd(commands.vpnExecutor, "Failed opening the process")
	activateVPNWindow()
	connectBtnLocation := findElementLocation(commands.connectBtnFinder, 1)
	clickMouseOn(connectBtnLocation)
	fetchAndTypePassword(commands.vpnPwFinder)
	otp := fetchOTPfromSMS(commands.otpNumberFinder)
	otpInputLocation := findElementLocation(commands.otpInputFinder, 0)
	clickMouseOn(otpInputLocation)
	typeOTPtoInput(otp)
}

func executeCmd(args []string, errMsg string) string {
	out, err := exec.Command(args[0], args[1], args[2]).CombinedOutput()
	executionMsg := string(out)
	if err != nil {
		log.Println(executionMsg)
		log.Fatalf("%s: %v\n", errMsg, err)
	}
	return executionMsg
}

func clickMouseOn(location []int) {
	robotgo.MouseSleep = 100 // millisecond
	robotgo.Move(location[0]+location[2]/2, location[1]+location[3]/2)
	robotgo.Click()
	robotgo.Sleep(1)
}

func findElementLocation(elementFinder []string, elementIdx int) []int {
	out := executeCmd(elementFinder, "Failed finding the element")
	elementInfos := strings.Split(strings.ReplaceAll(out, "\r\n", "\n"), "\n")
	targetElementInfo := strings.Split(elementInfos[elementIdx], ", ")
	if len(targetElementInfo) < 6 {
		log.Fatalf("Couldn't fetch the element information. Is it valid situation?\n")
	}
	var elementLocation []int
	for i := 2; i < len(targetElementInfo); i++ {
		n, err := strconv.Atoi(targetElementInfo[i])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v\n", err)
		}
		elementLocation = append(elementLocation, n)
	}
	log.Println("x y w h:", elementLocation)
	return elementLocation
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

func fetchAndTypePassword(vpnPwFinder []string) {
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

func typeOTPtoInput(otp string) {
	robotgo.TypeStr(otp)
	robotgo.KeyTap("enter")
	log.Println("Entered the OTP to the input")
}
