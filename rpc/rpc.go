package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	//  NOTE: Receives a message, and serializes it to JSON
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	//  NOTE: %d for number | %s for string
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", 0, errors.New("Did not find separator")
	}

	// Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	//  NOTE: String convert ascii to integer
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", 0, err
	}

	_ = content

	var baseMessage BaseMessage
	//  NOTE: unmarshal takes a reference/pointer to baseMessage, to then populate it
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", 0, err
	}

	return baseMessage.Method, contentLength, nil
}
