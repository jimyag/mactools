package clipboard

import (
	"github.com/jimyag/mactools/pkg/log"
	"io"
	"os/exec"
	"time"
)

var (
	impl Clipboard
)

func GetClipboard() Clipboard {
	if impl != nil {
		return impl
	} else {
		impl = &LinuxImpl{}
		return impl
	}
}

type LinuxImpl struct {
	lastContent *string
	handles     []OnChangedHandler
}

func (c *LinuxImpl) Register(handle OnChangedHandler) {
	log.Debug("register handle: %v", handle)
	c.handles = append(c.handles, handle)
}

func (c *LinuxImpl) Write(data Data) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer func(stdin io.WriteCloser) {
		err := stdin.Close()
		if err != nil {
			panic(err)
		}
	}(stdin)
	if err := cmd.Start(); err != nil {
		return err
	}
	if data.Type == ClipboardItemTypeString {
		if _, err := stdin.Write([]byte(data.Content.(string))); err != nil {
			return err
		}
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func readStringFromClipboard() (string, error) {
	cmd := exec.Command("xclip", "-selection", "clipboard", "-o")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return "", err
	}
	buf, err := io.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return string(buf), nil
}
func (c *LinuxImpl) Listen() error {
	for {
		if s, err := readStringFromClipboard(); err != nil {
			return err
		} else {
			if c.lastContent == nil || s == *c.lastContent {
				log.Trace("no change wait 1s")
				time.Sleep(time.Millisecond * 1000)
			} else {
				log.Debug("read from clipboard: %v", s)
				for _, handle := range c.handles {
					handle(Data{
						Type:    ClipboardItemTypeString,
						Content: s,
					})
				}
			}
			c.lastContent = &s
		}
	}
}
