-- Path to the Messages database
set messagesDB to POSIX path of (path to library folder from user domain) & "Messages/chat.db"

-- Function to get the latest message
on getLatestMessage(senderHandle)
	global messagesDB
    set query to "SELECT text, datetime(date / 1000000000 + strftime('%s', '2001-01-01'), 'unixepoch', 'localtime') as message_date FROM message WHERE handle_id = (SELECT ROWID FROM handle WHERE id = " & quoted form of senderHandle & ") AND service = 'SMS' ORDER BY date DESC LIMIT 1;"
    try
        set latestMessage to do shell script "sqlite3 " & quoted form of messagesDB & " " & quoted form of query
        return latestMessage
    on error errMsg
        return "Error fetching latest message: " & errMsg
    end try
end getLatestMessage

-- Function to get the message count
on getMessageCount(senderHandle)
	global messagesDB
    set query to "SELECT COUNT(*) FROM message WHERE handle_id = (SELECT ROWID FROM handle WHERE id = " & quoted form of senderHandle & ") AND service = 'SMS';"
    try
        set messageCount to do shell script "sqlite3 " & quoted form of messagesDB & " " & quoted form of query
        return messageCount as integer
    on error errMsg
        return "Error fetching message count: " & errMsg
    end try
end getMessageCount

-- Get the senders number
set senderHandle to do shell script "security find-generic-password -s 'auto.otp.sender' -w"

-- Get the initial message count
set initialCount to getMessageCount(senderHandle)

-- Execute the query and handle the result
repeat
    set currentCount to getMessageCount(senderHandle)
    if currentCount > initialCount then
        set latestMessage to getLatestMessage(senderHandle)
        if latestMessage is not "" and latestMessage does not start with "Error" then
            log "New message from " & senderHandle & ": " & latestMessage
            exit repeat
        end if
    end if
    delay 1 -- Wait for 1 second before checking again
end repeat
