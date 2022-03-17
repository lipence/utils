package sysUtils

import (
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	CMDCygPath   = "cygpath"
	CMDCygHelper = "cygwin-console-helper"
)

var usingCygwinGit bool

func init() {
	usingCygwinGit = isUsingCygwinGit()
}

func isUsingCygwinGit() bool {
	if runtime.GOOS != "windows" {
		return false
	}
	var err error
	var cygPathDir, cygHelperDir string
	if cygPathDir, err = exec.LookPath(CMDCygPath); err != nil {
		return false
	}
	if cygHelperDir, err = exec.LookPath(CMDCygHelper); err == nil {
		return false
	}
	return filepath.Dir(cygPathDir) == filepath.Dir(cygHelperDir)
}

func IsUsingCygwinGit() bool {
	return usingCygwinGit
}
