package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func formatjson(part string) string {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(part), &data)
	if err != nil {
		return part
	}

	msg, ok := data["msg"]
	if !ok {
		message, ok := data["message"]
		if ok {
			msg = message
		} else {
			return part
		}
	}

	level, ok := data["level"]
	isError := ok && level == "error"
	isWarning := ok && level == "warn"

	if isError {
		return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
	} else if isWarning {
		return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
	} else {
		return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Split(line, "\x1B")
		if len(parts) > 2 {
			lastPart := parts[len(parts)-2]
			firstBracket := strings.Index(lastPart, "{")
			if firstBracket != -1 {
				lastPart = lastPart[firstBracket:]
			}
			lastBracket := strings.LastIndex(lastPart, "}")
			if lastBracket != -1 {
				lastPart = lastPart[:lastBracket+1]
			}
			formatted := formatjson(lastPart)
			fmt.Println(formatted)
		} else {
			fmt.Println(line)
		}
	}
}
