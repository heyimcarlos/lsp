# My own LSP

A language server protocol (LSP) communicates with a protocol that uses two main parts:

## Editors transport mechanisms with the LSP is stdin stdout or TCP

- Header
  - Content-Length(number): the length of content in bytes. Required.
  - Content-Type(string): The mime type of the content. Defaults to
  application/vscode-jsonrpc; charset=utf-8.

- Content
