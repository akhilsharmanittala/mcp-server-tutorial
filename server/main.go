package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

/* ---------- Greet tool ---------- */

type GreetInput struct {
	Name string `json:"name" jsonschema:"name of the person"`
}

type GreetOutput struct {
	Greeting string `json:"greeting"`
}

func Greet(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input GreetInput,
) (*mcp.CallToolResult, GreetOutput, error) {
	return nil, GreetOutput{
		Greeting: "Hi " + input.Name,
	}, nil
}

/* ---------- Add tool ---------- */

type AddInput struct {
	A int `json:"a" jsonschema:"first number"`
	B int `json:"b" jsonschema:"second number"`
}

type AddOutput struct {
	Sum int `json:"sum"`
}

func Add(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input AddInput,
) (*mcp.CallToolResult, AddOutput, error) {
	return nil, AddOutput{
		Sum: input.A + input.B,
	}, nil
}

/* ---------- main ---------- */

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "demo-server",
			Version: "v1.0.0",
		},
		nil,
	)

	// Register multiple tools on ONE server
	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "Say hi to someone",
	}, Greet)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "add",
		Description: "Add two numbers",
	}, Add)

	log.Println("MCP server running on stdio")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
