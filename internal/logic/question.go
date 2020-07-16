// Package logic provides ...
package logic

import "github.com/glepnir/jarvis/pkg/cli"

func PluginManage() string {
	message := "What is plugin manage do you use?"
	options := []string{"dein", "vim-plug"}
	return cli.SingleSelectTemplate(message, options)
}

func LeaderKey() string {
	message := "What is your Leader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

func LocalLeaderKey() string {
	message := "What is your LocalLeader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

func LspPlugin() string {
	message := "What is your  Lsp plugin?"
	options := []string{"coc.nvim", "nvim-lsp"}
	return cli.SingleSelectTemplate(message, options)
}

func ExplorerPlugin() string {
	message := "What is your explorer plugin?"
	options := []string{"defx.nvim", "nerdtree", "coc-explorer"}
	return cli.SingleSelectTemplate(message, options)
}