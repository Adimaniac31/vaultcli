// main.go
package main

import (
	"flag"
	"fmt"
	"vaultcli/cmd"
)

func main() {
	cmdName := flag.String("cmd", "", "Command to run: login | store | read")
	userID := flag.String("user", "", "User ID")
	key := flag.String("key", "", "Key")
	value := flag.String("value", "", "Value")
	token := flag.String("token", "", "JWT token")

	flag.Parse()

	switch *cmdName {
	case "login":
		if *userID == "" {
			fmt.Println("❌ user ID required for login")
			return
		}
		cmd.Login(*userID)
	case "store":
		if *token == "" || *key == "" || *value == "" {
			fmt.Println("❌ token, key, and value are required for storing")
			return
		}
		cmd.Store(*key, *value, *token)
	case "read":
		if *token == "" || *key == "" {
			fmt.Println("❌ token and key are required for reading")
			return
		}
		cmd.Read(*key, *token)
	default:
		fmt.Println("❌ Unknown command. Use -cmd=login|store|read")
	}
}
