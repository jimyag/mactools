/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package clipboard

import (
	"errors"
	"github.com/jimyag/mactools/pkg/log"
	"time"

	"github.com/progrium/macdriver/cocoa"
)

var (
	impl Clipboard
)

func GetPasteboard() Clipboard {
	if impl != nil {
		return impl
	}
	gp := cocoa.NSPasteboard_GeneralPasteboard()
	impl = &MacImpl{
		PB:             gp,
		PreChangeCount: int(gp.ChangeCount()),
	}
	return impl
}

type MacImpl struct {
	PB             cocoa.NSPasteboard
	PreChangeCount int
	handles        []OnChangedHandler
}

func (pb *MacImpl) IsChanged() bool {
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

func (pb *MacImpl) Register(handle OnChangedHandler) {
	pb.handles = append(pb.handles, handle)
}

func (pb *MacImpl) Write(data Data) error {
	if data.Type == ClipboardItemTypeString {
		pb.PB.ClearContents()
		pb.PB.SetStringForType(data.Content, cocoa.NSPasteboardTypeString)
		return nil
	}
	return errors.New("not support type")
}

func (pb *MacImpl) Listen() error {
	for {
		if !pb.IsChanged() {
			log.Debug("sleep 1000ms")
			time.Sleep(time.Millisecond * 1000)
			continue
		}

		content := pb.PB.StringForType(cocoa.NSPasteboardTypeString)
		log.Debug("read from pasteboard: %v", content)

		for _, handle := range pb.handles {
			handle(content)
		}

	}
}
