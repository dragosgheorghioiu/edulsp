package rpc_test

import (
	"testing"

	"github.com/dragosgheorghioiu/edulsp/src/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecoding(t *testing.T) {
	incoming := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	if err != nil {
		t.Fatal(err)
	}

	if contentLength := len(content); contentLength != 15 {
		t.Fatalf("Expected: 15, Actual: %d", contentLength)
	}

	if method != "hi" {
		t.Fatalf("Expected: \"hi\", Actual: %s", method)
	}
}
