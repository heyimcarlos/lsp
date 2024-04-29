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

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find separator")
	}

	// Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	//  NOTE: String convert ascii to integer
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	_ = content

	var baseMessage BaseMessage
	//  NOTE: unmarshal takes a reference/pointer to baseMessage, to then populate it
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		//  NOTE: not ready yet.
		return 0, nil, nil
	}

	// Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	//  NOTE: String convert ascii to integer
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		//  NOTE: no number, return error.
		return 0, nil, err
	}

	if len(content) < contentLength {
		//  NOTE: we haven't read enough bytes, let's wait.
		return 0, nil, nil
	}

	//  NOTE: Use 4 because we're using a separator of '\r', '\n', '\r', '\n'
	totalLength := len(header) + 4 + contentLength

	return totalLength, data[:totalLength], nil
}
