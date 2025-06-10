package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument  VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChange []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
