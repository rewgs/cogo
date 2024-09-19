package base

import (
)


type SyncJob struct {
    Name        string
    Method      SyncMethod
    Cmd         SyncCmd
    Src         string
    Dst         string
}

func NewSyncJob(name string, src string, dst string) *SyncJob {
    b := SyncJob{
        Name: name,
        Src: src,
        Dst: dst,
    }

    return &b
}

func (j *SyncJob) RunJob() {
}
