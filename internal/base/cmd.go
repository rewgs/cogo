package base

import (
    "os/exec"
)


type SyncCmd interface {
    NewCmd()
    GetOptions() []string // Returns the options for this particular command, i.e. `rsync --archive --verbose` would return a slice `[--archive, --verbose]`
    SetOptions() // 
}

type SyncCmdImpl struct {
    Name        string
    Method      SyncMethod
    Cmd         exec.Cmd
    Options     []string
}
