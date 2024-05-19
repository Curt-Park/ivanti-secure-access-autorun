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

	// Open the VPN app window.
	args := commands.vpnExecution
	err := exec.Command(args[0], args[1], args[2]).Start()
	if err != nil {
		log.Fatal(err, ":Failed opening the process")
	}

	// Activate the VPN app window.
	fpid, err := robotgo.FindIds("Ivanti Secure Access")
	if err != nil || len(fpid) == 0 {
		log.Fatal(err, ":Failed finding the process")
	}
	log.Println("The process ID:", fpid[0])
	err = robotgo.ActivePid(fpid[0])
	if err != nil {
		log.Fatal(err, ":Failed activating the window")
	}

	// Find the connection button location.
	args = commands.btnFinder
	out, err := exec.Command(args[0], args[1], args[2]).CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err, ":Failed finding the connection button")
	}
	btn_infos := strings.Split(strings.ReplaceAll(string(out), "\r\n", "\n"), "\n")
	target_btn_info := strings.Split(btn_infos[1], ", ")
	if len(target_btn_info) < 6 {
		log.Fatal(err, ":The connection is already established")
	}
	var btn_location []int
	for i := 2; i < len(target_btn_info); i++ {
		n, err := strconv.Atoi(target_btn_info[i])
		if err != nil {
			log.Fatal(err, ":Failed to convert string to int")
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
	vpnPW, err := exec.Command(args[0], args[1], args[2]).CombinedOutput()
	if err != nil {
		log.Fatal(err, ":Failed finding the vpn password")
	}
	robotgo.TypeStr(string(vpnPW))
	robotgo.KeyTap("enter")
	log.Println("Typed VPN password")

	// Get OTP from SMS.
	log.Println("Wait for the otp number...")
	args = commands.otpFinder
	out, err = exec.Command(args[0], args[1], args[2]).CombinedOutput()
	if err != nil {
		log.Fatal(err, ":Failed fetching the otp number")
	}
	re := regexp.MustCompile(`\[OTP:\s(\d+)\]`)
	match := re.FindStringSubmatch(string(out))
	if len(match) <= 1 {
		log.Fatal("No OTP number found")
	}
	otp := match[1]
	log.Println("Extracted OTP number: ", otp)
}
