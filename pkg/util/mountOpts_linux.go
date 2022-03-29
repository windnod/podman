package util

import (
	"os"

	"golang.org/x/sys/unix"
)

func getDefaultMountOptions(path string) (defaultMountOptions, error) {
	opts := defaultMountOptions{false, true, true}
	if path == "" {
		return opts, nil
	}
	_, err := isFileExist(path)
	if err != nil {
		if e := os.MkdirAll(path, 0755); e != nil {
			return opts, &os.PathError{Op: "statfs", Path: path, Err: e}
		}
	}

	var statfs unix.Statfs_t
	if e := unix.Statfs(path, &statfs); e != nil {
		return opts, &os.PathError{Op: "statfs", Path: path, Err: e}
	}
	opts.nodev = (statfs.Flags&unix.MS_NODEV == unix.MS_NODEV)
	opts.noexec = (statfs.Flags&unix.MS_NOEXEC == unix.MS_NOEXEC)
	opts.nosuid = (statfs.Flags&unix.MS_NOSUID == unix.MS_NOSUID)

	return opts, nil
}

func isFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	return false, err
}
