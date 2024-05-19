package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {
	commands := InitCommands()

	// Open the VPN app window.
	err := exec.Command("bash", "-c", commands.vpnExecution).Start()
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
	out, err := exec.Command("bash", "-c", commands.btnFinder).CombinedOutput()
	if err != nil {
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
	vpnPW, err := exec.Command("bash", "-c", commands.vpnPwFinder).CombinedOutput()
	if err != nil {
		log.Fatal(err, ":Failed finding the vpn password")
	}
	robotgo.TypeStr(string(vpnPW))
	robotgo.KeyTap("enter")
	log.Println("Typed VPN password")

	// Type OTP
	log.Println("Wait for the otp number...")
	out, err = exec.Command("bash", "-c", commands.otpFinder).CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal(err, ":Failed fetching the otp number")
	}
	otpNum := string(out)
	log.Println("OTP:", otpNum)
}
