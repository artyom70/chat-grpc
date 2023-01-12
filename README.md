# Help Commands:
- /jg ${group_name} - joins to existing group
- /lg ${group_name} - lefts from group
- /cg ${group_name} - creates group and joins
- /smg ${group_name} - send message in group
- /sm ${username} ${message} - sends message to provided user
- /exit - disconnect from chat

# How to run chat
1. You need to run server `make run-server`
2. You need to run client which connects to server with provided username `PORT=:8080 go run cmd/client/main.go --username=artyom`

