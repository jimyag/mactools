/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package pasteboard

import (
	"time"

	"github.com/jimyag/mactools/log"
	"github.com/progrium/macdriver/cocoa"
)

var (
	PB *Pasteboard
)

func init() {
	if PB == nil {
		PB = new()
	}

}

type Pasteboard struct {
	PB             cocoa.NSPasteboard
	PreChangeCount int
	handles        []Handle
}

func new() *Pasteboard {
	gp := cocoa.NSPasteboard_GeneralPasteboard()
	return &Pasteboard{
		PB:             gp,
		PreChangeCount: int(gp.ChangeCount()),
	}
}

func (pb *Pasteboard) IsChanged() bool {
	changeCount := int(pb.PB.ChangeCount())
	defer func() {
		if pb.PreChangeCount != changeCount {
			pb.PreChangeCount = changeCount
		}
	}()
	log.Debug("change count: %v", changeCount)
	log.Debug("pre change count: %v", pb.PreChangeCount)
	return pb.PreChangeCount != changeCount
}

func (pb *Pasteboard) Register(handle Handle) {
	pb.handles = append(pb.handles, handle)
}

func (pb *Pasteboard) Run() {
	for {
		if !pb.IsChanged() {
			log.Debug("sleep 1000ms")
			time.Sleep(time.Millisecond * 1000)
			continue
		}

		content := pb.PB.StringForType(cocoa.NSPasteboardTypeString)
		log.Debug("read from pasteboard: %v", content)

		for _, handle := range pb.handles {
			log.Debug("OnCopy: %v", content)
			handle.OnCopy(pb, content)
			res := handle.Handle(pb, content)
			log.Debug("Handle res: %v\n", res)
			handle.AfterHandle(pb, res)
			log.Debug("AfterHandle")
		}

	}
}
