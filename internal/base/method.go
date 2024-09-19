/*
This is the main package for `backups`. 
*/

package base

import (
)

type SupportedMethod string
const (
    Rclone      string = "rclone"
    Robocopy    string = "robocopy"
    Rsync       string = "rsync"
)

// rclone, robocopy, or rsync
type SyncMethod interface {
    Install() 
    IsInstalled() bool
    GetExec() string // Returns an absolute path of the executable
    GetOptions() []string // Returns the possible options defined by the specified backup method (rclone, robocopy, or rsync)
}

type SyncMethodImpl struct {
    Name        SupportedMethod
    Installed   bool
    Exec        string // Absolute path to the executable
    Options     []string
}