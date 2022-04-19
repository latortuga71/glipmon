package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/progrium/macdriver/cocoa"
)

type ClipBoard struct {
	sync.Mutex
	ChangeCount int
	Contents    string
	Board       cocoa.NSPasteboard
}

func NewClipBoard() *ClipBoard {
	board := cocoa.NSPasteboard_GeneralPasteboard()
	changeCount := int(board.ChangeCount())
	contents := board.StringForType(cocoa.NSPasteboardTypeString)
	return &ClipBoard{
		ChangeCount: changeCount,
		Contents:    contents,
		Board:       board,
	}
}
func main() {
	sigChannel := make(chan os.Signal, 1)
	exit := make(chan interface{}, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	board := NewClipBoard()
	go func() {
		for {
			time.Sleep(time.Second * 1)
			board.Lock()
			changeCount := board.Board.ChangeCount()
			if board.ChangeCount != int(changeCount) {
				board.Contents = board.Board.StringForType(cocoa.NSPasteboardTypeString)
				log.Printf("%s\n", board.Contents)
				board.ChangeCount = int(changeCount)
			}
			board.Unlock()
		}
	}()
	go func() {
		test := <-sigChannel
		exit <- test
	}()
	<-exit
	fmt.Fprintf(os.Stderr, "Exiting...")
}
