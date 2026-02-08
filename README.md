# run client and server based testing

akhil@Akhil-Nittalas-MacBook-Pro mcp-go-demo % go mod init mcp-demo
go get github.com/modelcontextprotocol/go-sdk/mcp

# build server
go build -o mcp-server server/main.go 

# run client
go run my_client.go greet --name Akhil

go run my_client.go add --a 10 --b 100

# output
2026/02/08 13:38:42 {"greeting":"Hi Akhil"}
2026/02/08 13:38:42 {"sum":110}

# test via UI using mcp-inspector
mcp-inspector ./mcp-server 




# connect via claude desktop

> Install claude desktop
> go to claude icon on top > settings > developer > local mcp server > paste the below content in ~/.config/claude/claude_desktop_config.json


{
  "mcpServers": {
    "akhil-mcp": {
      "command": "/Users/akhil/mcp-server-tutorial/mcp-server"
    }
  },
  "preferences": {
    "coworkScheduledTasksEnabled": false,
    "sidebarMode": "chat"
  }
}

> reload claude desktop
> in chat "greet Akhil"

check the response and logs in settings> developer> view logs