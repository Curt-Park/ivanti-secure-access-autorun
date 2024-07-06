set VPN_APP to "Ivanti Secure Access"

tell application "System Events"
    tell process VPN_APP
        set windowList to windows
        repeat with w in windowList
            set elementList to entire contents of w
            repeat with e in elementList
                try
                    set elementName to name of e
                    set elementRole to role of e
                    if elementRole is "AXTextField" then
                        set elementPosition to position of e
                        set elementSize to size of e
                        log {elementName, elementRole, elementPosition, elementSize}
                    end if
                end try
            end repeat
        end repeat
    end tell
end tell
