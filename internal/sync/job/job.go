package job

import (
    "log"
)

type Job struct {
    Name        string
    Prog        Prog
    Cmd         Cmd
    Src         string
    Dst         string
    Log         string
}

func NewJob(name string, src string, dst string) *Job {
    job := Job{
        Name: name,
        Src: src,
        Dst: dst,
    }

    return &job
}

func (j *Job) Run() {
}
