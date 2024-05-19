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

func main() {
	commands := InitCommands()

	// Open the VPN app window.
	args := commands.vpnExecution
	err := exec.Command(args[0], args[1], args[2]).Start()
	if err != nil {
		log.Fatalf("Failed opening the process: %v\n", err)
	}

	// Activate the VPN app window.
	fpid, err := robotgo.FindIds("Ivanti Secure Access")
	if err != nil || len(fpid) == 0 {
		log.Fatalf("Failed finding the process: %v\n", err)
	}
	log.Println("The process ID:", fpid[0])
	err = robotgo.ActivePid(fpid[0])
	if err != nil {
		log.Fatalf("Failed activating the window: %v\n", err)
	}

	// Find the connection button location.
	args = commands.btnFinder
	out, err := exec.Command(args[0], args[1], args[2]).CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatalf("Failed finding the connection button: %v\n", err)
	}
	btn_infos := strings.Split(strings.ReplaceAll(string(out), "\r\n", "\n"), "\n")
	target_btn_info := strings.Split(btn_infos[1], ", ")
	if len(target_btn_info) < 6 {
		log.Fatalf("The connection is already established\n")
	}
	var btn_location []int
	for i := 2; i < len(target_btn_info); i++ {
		n, err := strconv.Atoi(target_btn_info[i])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v\n", err)
		}
		btn_location = append(btn_location, n)
	}
	log.Println("x y w h:", btn_location)

	// Move the mouse pointer on the button.
	robotgo.MouseSleep = 100 // millisecond
	robotgo.Move(btn_location[0]+btn_location[2]/2, btn_location[1]+btn_location[3]/2)
	robotgo.Click()
	robotgo.Sleep(1)

	// Type the vpn password.
	args = commands.vpnPwFinder
	out, err = exec.Command(args[0], args[1], args[2]).CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatalf("Failed finding the vpn password: %v\n", err)
	}
	robotgo.TypeStr(string(out))
	robotgo.KeyTap("enter")
	log.Println("Typed VPN password")

	// Get OTP from SMS.
	log.Println("Wait for the otp number...")
	args = commands.otpFinder
	out, err = exec.Command(args[0], args[1], args[2]).CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatalf("Failed fetching the otp number: %v\n", err)
	}
	re := regexp.MustCompile(`\[OTP:\s(\d+)\]`)
	match := re.FindStringSubmatch(string(out))
	if len(match) <= 1 {
		log.Fatal("No OTP number found")
	}
	otp := match[1]

	// Copy the OTP to ClipBoard.
	err = clipboard.WriteAll(otp)
	if err != nil {
		log.Fatalf("Failed to copy text to clipboard: %v\n", err)
	}
	log.Println("Copied the OTP to the clipboard. Paste it!")
}
