package fileUtils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Stat(path string) (string, os.FileInfo, error) {
	var err error
	path, err = filepath.Abs(strings.TrimSpace(path))
	if err != nil {
		return "", nil, err
	}
	var info os.FileInfo
	if info, err = os.Stat(path); err != nil {
		return path, nil, err
	}
	return path, info, nil
}

func Handle(path string, read, write, create bool, additionalOpts ...int) (file *os.File, err error) {
	var opt int
	var perm fs.FileMode
	if write {
		if read {
			opt = os.O_RDWR
		} else {
			opt = os.O_WRONLY
		}
	} else {
		opt = os.O_RDONLY
	}
	if create {
		opt |= os.O_CREATE
		perm = 0640
	}
	for _, o := range additionalOpts {
		opt |= o
	}
	return os.OpenFile(path, opt, perm)
}

func List(path string, pattern *regexp.Regexp, recursion bool) (list []string, err error) {
	var files []os.DirEntry
	if files, err = os.ReadDir(path); err != nil {
		return nil, err
	}
	list = make([]string, 0, len(files))
	for _, f := range files {
		var fileInfo fs.FileInfo
		var filePath = filepath.Join(path, f.Name())
		if fileInfo, err = f.Info(); err != nil {
			return nil, err
		}
		if fileInfo.Mode()&os.ModeSymlink > 0 {
			var realPath string
			if realPath, err = os.Readlink(filePath); err != nil {
				return nil, fmt.Errorf("%w (path: %s)", err, f.Name())
			}
			if _, fileInfo, err = Stat(realPath); err != nil {
				return nil, fmt.Errorf("%w (path: %s)", err, f.Name())
			}
		}
		if fileInfo.IsDir() && recursion {
			var sublist []string
			if sublist, err = List(filePath, pattern, recursion); err != nil {
				return nil, fmt.Errorf("%w (path: %s)", err, f.Name())
			}
			list = append(list, sublist...)
			continue
		}
		if pattern != nil && !pattern.MatchString(filePath) {
			continue
		}
		list = append(list, filePath)
	}
	return list, nil
}
