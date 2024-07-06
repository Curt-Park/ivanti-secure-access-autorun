package applescripts

import (
	_ "embed"
)

//go:embed connect_btn_finder.applescript
var ConnectBtnFinderApple string

//go:embed otp_number_finder.applescript
var OtpNumberFinderApple string

//go:embed otp_input_finder.applescript
var OtpInputFinderApple string
