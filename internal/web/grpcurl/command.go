// Package grpcurl formats commands for display and clipboard use. It does not
// execute grpcurl or make network requests.
package grpcurl

import (
	"encoding/json"
	"strings"
)

// Command formats a unary gRPC request as a shell-safe grpcurl command.
func Command(address, method string, payload any) string {
	encoded, err := json.Marshal(payload)
	if err != nil {
		return ""
	}

	return strings.Join([]string{
		"grpcurl",
		"  -insecure",
		"  -d " + shellQuote(string(encoded)),
		"  " + address,
		"  " + method,
	}, " \\\n")
}

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}
