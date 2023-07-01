/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package notification

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

// https://github.com/progrium/macdriver/blob/main/examples/notification/main.go

type Notification struct {
	notification       objc.Object
	notificationCenter objc.Object
}

func New() *Notification {
	return &Notification{
		notification:       objc.Get("NSUserNotification").Alloc().Init(),
		notificationCenter: objc.Get("NSUserNotificationCenter").Send("defaultUserNotificationCenter"),
	}
}

func (n *Notification) Show() {
	n.notificationCenter.Send("deliverNotification:", n.notification)
	n.notification.Release()
}

func (n *Notification) SetTitle(title string) *Notification {
	n.notification.Set("title:", core.String(title))
	return n
}

func (n *Notification) SetInformativeText(text string) *Notification {
	n.notification.Set("informativeText:", core.String(text))
	return n
}
