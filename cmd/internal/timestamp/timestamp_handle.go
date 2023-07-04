/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */

package timestamp

import (
	"github.com/jimyag/mactools/pkg/clipboard"
	"github.com/jimyag/mactools/pkg/log"
	"github.com/jimyag/mactools/pkg/notification"
	"strconv"
	"time"
)

func init() {
	handle := &TimeStampHandle{
		timeFormat: time.DateTime,
	}
	clipboard.GetClipboard().Register(func(data clipboard.Data) {
		if data.Type != clipboard.ClipboardItemTypeString {
			handle.AfterHandle(handle.Handle(data.Content.(string)))
		}
	})
}

type TimeStampHandle struct {
	time       time.Time
	timeFormat string
}

func (h *TimeStampHandle) AfterHandle(res any) {
	if res == nil {
		return
	}
	h.show(res.(string))
}

func (h *TimeStampHandle) Handle(content string) any {
	log.Debug("handle: ", content)
	if content == "" {
		return nil
	}

	if len(content) < 10 {
		return nil
	}

	timeStamp, err := strconv.Atoi(content)
	if err != nil {
		return nil
	}

	switch len(content) {
	// 秒
	case 10:
		h.time = time.Unix(int64(timeStamp), 0)
	// 毫秒
	case 13:
		h.time = time.UnixMilli(int64(timeStamp))
	// 微秒
	case 16:
		h.time = time.UnixMicro(int64(timeStamp))
	// 百纳秒
	case 17:
		h.time = time.Unix(0, int64(timeStamp)*100)
	// 纳秒
	case 19:
		h.time = time.Unix(0, int64(timeStamp))
	default:
		return nil
	}
	return content
}

func (h *TimeStampHandle) show(content string) {
	notification.New().
		SetTitle(content).
		SetInformativeText(h.time.Format(h.timeFormat)).
		Show()
}
