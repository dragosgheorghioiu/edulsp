package lsp

type CodeActionRequest struct {
	Request
	Params CodeActionParams `json:"params"`
}

type CodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"codeActionContext"`
}

type CodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeActionContext struct {}

type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkspaceEdit `json:"edit"`
	Command *Command       `json:"command,omitempty"`
}

type Command struct {
	Title     string  `json:"title"`
	Command   string  `json:"command"`
	Arguments []any   `json:"arguments,omitempty"`
}
