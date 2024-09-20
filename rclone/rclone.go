package rclone

import (
    "github.com/rewgs/cogo/internal/sync"
)

type Rclone struct {
    sync.Prog
}

func (r *Rclone) Init() {
    var name string = "rclone"
}
