package notification

import "os/exec"

type LinuxImpl struct {
	title *string
	text  *string
}

func New() Notification {
	return &LinuxImpl{}
}

func (n *LinuxImpl) SetTitle(title string) Notification {
	n.title = &title
	return n
}

func (n *LinuxImpl) SetInformativeText(text string) Notification {
	n.text = &text
	return n
}

func (n *LinuxImpl) Show() error {
	var args []string
	if n.title != nil {
		args = append(args, "--app-name", *n.title)
	}
	if n.text != nil {
		args = append(args, *n.text)
	}
	return exec.Command("notify-send", args...).Run()
}
