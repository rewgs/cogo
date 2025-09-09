package sync

// Syncer defines an application that can sync files (e.g. rclone, robocopy, rsync, etc).
type Syncer interface {
	Copy(src string, dst string)

	runJob(job Job) error

	// Returns an absolute path of the executable
	getExec() string

	// Returns the possible options defined by the specified backup method (rclone, robocopy, or rsync)
	getOptions() []string

	isInstalled() bool
}
