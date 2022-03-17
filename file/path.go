package fileUtils

import (
	"bytes"
	"os/exec"
	"path"
	"path/filepath"

	sysUtils "github.com/lipence/utils/sys"
)

func TargetPath(base string, target string) string {
	if filepath.IsAbs(target) {
		return filepath.Clean(target)
	}
	return filepath.Clean(filepath.Join(base, target))
}

func CleanPath(p string) (string, error) {
	if sysUtils.IsUsingCygwinGit() {
		if out, err := exec.Command("cygpath", p).Output(); err != nil {
			return "", err
		} else {
			p = string(bytes.TrimSpace(out))
		}
		return path.Clean(p), nil
	}
	return filepath.Clean(p), nil
}
