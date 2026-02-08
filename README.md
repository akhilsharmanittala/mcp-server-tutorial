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
akhil@Akhil-Nittalas-MacBook-Pro mcp-go-demo % 