package rpc_test

import (
	"lsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

type DecodingExample struct {
	Testing int
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	//  NOTE: Instantiate a struct
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, contentLength, err := rpc.DecodeMessage([]byte(incomingMessage))

	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected: 15, Actual: %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected: 'hi', Actual: %s", method)
	}
}
