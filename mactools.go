/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */

package main

import (
	"github.com/jimyag/mactools/app"
	"github.com/jimyag/mactools/log"
)

func main() {
	app := app.App
	log.Info("mactools running!")
	app.Run()
}
