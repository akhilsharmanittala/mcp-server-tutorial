package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: client.go <greet|add> [flags]")
	}

	tool := os.Args[1]

	ctx := context.Background()

	client := mcp.NewClient(
		&mcp.Implementation{
			Name:    "demo-client",
			Version: "v1.0.0",
		},
		nil,
	)

	transport := &mcp.CommandTransport{
		Command: exec.Command("./mcp-server"),
	}

	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	switch tool {

	case "greet":
		greetFlags := flag.NewFlagSet("greet", flag.ExitOnError)
		name := greetFlags.String("name", "world", "name to greet")
		greetFlags.Parse(os.Args[2:])

		res, err := session.CallTool(ctx, &mcp.CallToolParams{
			Name: "greet",
			Arguments: map[string]any{
				"name": *name,
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		printResult(res)

	case "add":
		addFlags := flag.NewFlagSet("add", flag.ExitOnError)
		a := addFlags.Int("a", 0, "first number")
		b := addFlags.Int("b", 0, "second number")
		addFlags.Parse(os.Args[2:])

		res, err := session.CallTool(ctx, &mcp.CallToolParams{
			Name: "add",
			Arguments: map[string]any{
				"a": *a,
				"b": *b,
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		printResult(res)

	default:
		log.Fatalf("unknown tool: %s", tool)
	}
}

func printResult(res *mcp.CallToolResult) {
	if res.IsError {
		log.Fatal("tool execution failed")
	}

	for _, c := range res.Content {
		if t, ok := c.(*mcp.TextContent); ok {
			log.Println(t.Text)
		}
	}
}
