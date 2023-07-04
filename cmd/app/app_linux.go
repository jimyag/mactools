package app

import "github.com/jimyag/mactools/pkg/clipboard"

func Run() {
	err := clipboard.GetClipboard().Listen()
	if err != nil {
		panic(err)
	}
}
