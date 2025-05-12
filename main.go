package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var jwt string

	// Accept token from argument or stdin
	if len(os.Args) == 2 {
		jwt = os.Args[1]
	} else {
		fmt.Println("Usage: JWT <JWT_TOKEN>")
		os.Exit(1)
	}

	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT format: should have 3 parts")
		os.Exit(1)
	}

	// Base64URL decode the token (part 1)
	token, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error decoding JWT payload:", err)
		os.Exit(1)
	}

	// Pretty print decoded token
	var payloadJSON map[string]interface{}
	if err := json.Unmarshal(token, &payloadJSON); err != nil {
		fmt.Println("Error parsing JSON payload:", err)
		os.Exit(1)
	}

	payloadPretty, _ := json.MarshalIndent(payloadJSON, "", "  ")
	fmt.Println("Decoded Payload:")
	fmt.Println(string(payloadPretty))

	// Check expiration
	if expVal, ok := payloadJSON["exp"]; ok {
		switch exp := expVal.(type) {
		case float64:
			expTime := time.Unix(int64(exp), 0)
			fmt.Printf("\nToken expires at: %s\n", expTime.Local())
		default:
			fmt.Println("\n'exp' field is not a number")
		}
	} else {
		fmt.Println("\nNo 'exp' field in token")
	}
}