package main

import (
	"bufio"
	"encoding/json"
	// "fmt"
	"log"
	"lsp/lsp"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("/Users/carlos/Documents/code/personal/lsp/log.txt")
	logger.Println("Hey, I'm running!")
	//  Keep reading from stdin, until there's a message.

	//  NOTE: Scanner reads from it's param, until there's a new message; In this case reading from Stdin.
	scanner := bufio.NewScanner(os.Stdin)
	//  NOTE: scanner.Split takes a SplitFunc which overwrites the default Split function
	// called when a new message is received
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, this couldn't be parsed: %s", err)
		}

		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		// reply!
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, this couldn't be parsed: %s", err)
		}
		logger.Printf("Connected to: %s %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text, request.Params.TextDocument.LanguageID)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("File couldn't be opened")
	}

	return log.New(logfile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
