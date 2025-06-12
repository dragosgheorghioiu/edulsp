package lsp

type GotoDefinitionRequest struct {
	Request
	Params GotoDefinitionParams `json:"params"`
}

type GotoDefinitionParams struct {
	TextDocumentPositionParams
}

type GotoDefinitionResponse struct {
	Response
	Result Location `json:"result"`
}

