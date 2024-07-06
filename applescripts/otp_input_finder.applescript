set VPN_APP to "Ivanti Secure Access"

tell application "System Events"
    tell process VPN_APP
        log "0"
        set windowList to windows
        repeat with w in windowList
            log "1"
            set elementList to entire contents of w
            repeat with e in elementList
                log "2"
                try
                    set elementName to name of e
                    set elementRole to role of e
                    log {elementName, elementRole, elementPosition, elementSize}
                end try
            end repeat
        end repeat
    end tell
end tell
