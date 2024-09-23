package cogo

import (
    // "github.com/rewgs/cogo/internal/sync"
    "github.com/rewgs/cogo/rclone"
    // "github.com/rewgs/cogo/robocopy"
    // "github.com/rewgs/cogo/rsync"
)

// func getInstalledSyncProgs() []sync.Prog {
//     var installed []sync.Prog
// 	for _, p := range sync.Supported  {
//         if p.Installed {
//             // TODO:
//             // add to installed
//         }
//     }
//     return installed
// }

// func isValid(prog string) bool {
// 	for i, p := range sync.Supported  {
//         if prog == p {
//             return true
//         }
//     }
//     return false
// }

// func isInstalled(prog string) bool {
//     if ! (IsValid(prog)) {
//         return false
//     }
// 	return true
// }

func Test() {
    rclone := new(rclone.Rclone)
    // robocopy := new(robocopy.Robocopy)
    // rsync := new(rsync.Rsync)

	rclone.Init()
	// robocopy.Init()
	// rsync.Init()
}
