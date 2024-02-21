package main

import (
	"os"

	"github.com/Schaudge/grailbase/file"
	"github.com/Schaudge/grailbase/file/fsnodefuse"
	"github.com/Schaudge/grailbase/file/gfilefs"
	"github.com/Schaudge/grailbase/file/s3file"
	"github.com/Schaudge/grailbase/must"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func main() {
	mount := os.Args[1]
	file.RegisterImplementation("s3", func() file.Implementation {
		return s3file.NewImplementation(s3file.NewDefaultProvider(), s3file.Options{})
	})
	root := fsnodefuse.NewRoot(gfilefs.New("s3://", "s3"))
	mountOpts := fuse.MountOptions{FsName: "s3"}
	fsnodefuse.ConfigureDefaultMountOptions(&mountOpts)
	fsnodefuse.ConfigureRequiredMountOptions(&mountOpts)
	server, err := fs.Mount(mount, root, &fs.Options{MountOptions: mountOpts})
	must.Nil(err)
	server.Wait()
}
