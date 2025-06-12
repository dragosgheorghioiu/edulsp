package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/dragosgheorghioiu/edulsp/src/analysis"
	"github.com/dragosgheorghioiu/edulsp/src/lsp"
	"github.com/dragosgheorghioiu/edulsp/src/rpc"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Ouch")
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Println(method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not parse this: %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg, logger)

		logger.Println("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: Could not parse this: %s", err)
		}

		logger.Printf("Opened: %s\n %s",
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text,
		)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: Could not parse this: %s", err)
		}

		logger.Printf("Changed: %s\n %s",
			request.Params.TextDocument.URI,
			request.Params.ContentChange,
		)
		for _, change := range request.Params.ContentChange {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: Could not parse this: %s", err)
		}
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response, logger)
	case "textDocument/definition":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition: Could not parse this: %s", err)
		}
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response, logger)
	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition: Could not parse this: %s", err)
		}
		response := state.CodeAction(request.ID, request.Params.TextDocument.URI)
		writeResponse(writer, response, logger)
	}
}

func writeResponse(writer io.Writer, msg any, logger *log.Logger) {
	reply := rpc.EncodeMessage(msg)
	logger.Println(reply)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("please provide a file")
	}

	return log.New(logfile, "[edulsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
