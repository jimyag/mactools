/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package app

import (
	"github.com/jimyag/mactools/log"
	"github.com/jimyag/mactools/pasteboard"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

var (
	App cocoa.NSApplication
)

func init() {
	log.SetLevel(log.InfoLevel)

	App = cocoa.NSApp_WithDidLaunch(func(_ objc.Object) {
		go pasteboard.PB.Run()
	})

	// https://github.com/progrium/macdriver/blob/main/examples/notification/main.go
	nsbundle := cocoa.NSBundle_Main().Class()
	nsbundle.AddMethod("__bundleIdentifier", func(_ objc.Object) objc.Object {
		return core.String("com.example.fake")
	})
	nsbundle.Swizzle("bundleIdentifier", "__bundleIdentifier")

	App.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
}
