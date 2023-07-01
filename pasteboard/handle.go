/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package pasteboard

type Handle interface {
	Handle(*Pasteboard, string) any
	OnCopy(*Pasteboard, string)
	AfterHandle(*Pasteboard, any)
}
