/*
 * Copyright (c) 2023 by jimyag, All Rights Reserved.
 * Licensed under the MIT License. See LICENSE file in the project root for license information.
 */
package clipboard

type ClipboardItemType string

const (
	ClipboardItemTypeString ClipboardItemType = "String"
)

type Data struct {
	Type    ClipboardItemType
	Content any
}

type OnChangedHandler func(data Data)

type Clipboard interface {
	Register(handler OnChangedHandler)
	Write(data Data) error
	Listen() error
}
