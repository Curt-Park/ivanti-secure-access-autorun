# Ivanti Secure Access Auto-Run
It's only available for MacOS.
For more OS supports, you need to add commands in `commands.go`.

## Prerequisites
- Manually login at the beginning.
- [Set up iPhone to get SMS messages on Mac](https://support.apple.com/ko-kr/guide/messages/icht8a28bb9a/mac)
- `chmod +x *.applescript`
- `security add-generic-password -s 'auto.vpn' -a '' -w 'YourPW'`
    - How to remove: `security delete-generic-password -s 'auto.vpn'`
- `security add-generic-password -s 'auto.otp.sender' -a '' -w 'PhoneNumberWithoutBlank'`
    - How to remove: `security delete-generic-password -s 'auto.otp.sender'`

## Build
```bash
go build
sudo mv ivanti-secure-access-autorun /usr/bin
```

## Run
```bash
ivanti-secure-access-autorun 
```
