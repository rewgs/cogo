/*
Handles the base implementation for the various sync programs.
*/

package sync

import (
)

// Perhaps look through the directories in the root of this repo that aren't `internal` instead?
// type Supported string
// const (
//     Rclone      string = "rclone"
//     Robocopy    string = "robocopy"
//     Rsync       string = "rsync"
// )

var Supported = [3]string{
	"rclone",
	"robocopy",
	"rsync",
}

// rclone, robocopy, or rsync
type Prog interface {
    Install() 
    getExec() string // Returns an absolute path of the executable
    getOptions() []string // Returns the possible options defined by the specified backup method (rclone, robocopy, or rsync)
    isInstalled() bool
}

type ProgImpl struct {
    Name        string
    Installed   bool
    Exec        string // Absolute path to the executable
    Options     []string
}

// TODO:
func (p *ProgImpl) Install() {}

// TODO:
// func (p *ProgImpl) getExec() string {}

// TODO:
// func (p *ProgImpl) getOptions() []string {}

// TODO:
// Searches $PATH for SyncMethodImpl.Name
// func (p *ProgImpl) isInstalled() bool {}
