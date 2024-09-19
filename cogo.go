package cogo

import (
    "github.com/rewgs/cogo/rclone"
    "github.com/rewgs/cogo/robocopy"
    "github.com/rewgs/cogo/rsync"
)

func _test() {
    rclone := new(rclone.Rclone)
    robocopy := new(robocopy.Robocopy)
    rsync := new(rsync.Rsync)
}
