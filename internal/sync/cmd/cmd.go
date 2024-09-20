package cmd

import (
    "os/exec"

    "github.com/rewgs/cogo/internal/sync/prog"
)


type Cmd interface {
    NewCmd()
    GetOptions() []string // Returns the options for this particular command, i.e. `rsync --archive --verbose` would return a slice `[--archive, --verbose]`
    SetOptions() // 
}

type CmdImpl struct {
    Name        string
    Prog        prog.Prog
    Exec        exec.Cmd
    Options     []string
}
