package notification

type Notification interface {
	SetTitle(string) Notification
	SetInformativeText(string) Notification
	Show() error
}
