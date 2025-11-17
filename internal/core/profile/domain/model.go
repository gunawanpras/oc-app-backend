package domain

import "time"

type Profile struct {
	ID        int64
	Name      string
	Detail    []byte
	CreatedAt time.Time
}

type Profiles []Profile

type GetProfileOptions struct {
	Proc bool
}
