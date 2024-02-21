package spliceio_test

import (
	"github.com/Schaudge/grailbase/file/fsnodefuse"
	"github.com/Schaudge/grailbase/ioctx/spliceio"
)

// Check this here to avoid circular package dependency with fsnodefuse.
var _ fsnodefuse.Writable = (*spliceio.OSFile)(nil)
