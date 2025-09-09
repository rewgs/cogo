package syncer

import ()

type Job struct {
	Name string
	Src  string
	Dst  string
	Log  string
}

func New(name string, src string, dst string) *Job {
	job := Job{
		Name: name,
		Src:  src,
		Dst:  dst,
	}

	return &job
}

func (j *Job) Run() {
}
