# Ivanti Secure Access Auto-Run
It's only available for MacOS.
For more OS, you need to add commands in `commands.go`.

## Prerequisites
- Manually login at the beginning.
- `chmod +x btn_finder.applescript`
- `security add-generic-password -s 'auto.nsa.vpn' -a '' -w 'YourPW'`
    - How to remove: `security delete-generic-password s 'auto.nsa.vpn'`

## Build
```bash
go build
```

## Run
```bash
./ivanti-secure-access-autorun 
```
