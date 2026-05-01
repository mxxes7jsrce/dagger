//go:build !unix
// +build !unix

package daggercmd

import "os/exec"

func ensureChildProcessesAreKilled(cmd *exec.Cmd) {
}
