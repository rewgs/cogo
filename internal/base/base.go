// TODO: Probably put all this in `internal/base`?
package base

import (
    "github.com/rewgs/cogo/internal/sync"
    "github.com/rewgs/cogo/rclone"
    "github.com/rewgs/cogo/robocopy"
    "github.com/rewgs/cogo/rsync"
)

func getInstalledSyncProgs() []sync.Prog {
    installed = []sync.Prog
    for i, m range(sync.Supported) {
        if m.Installed {
            // TODO:
            // add to installed
        }
    }
    return installed

}

func isValid(prog string) bool {
    for i, p range(sync.Supported) {
        if prog == p {
            return true
        }
    }
    return false
}

func isInstalled(prog string) bool {
    if ! (isValid(prog)) {
        return false
    }
}

func _test() {
    rclone := new(rclone.Rclone)
    robocopy := new(robocopy.Robocopy)
    rsync := new(rsync.Rsync)
}
