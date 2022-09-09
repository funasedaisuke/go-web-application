package store

import (
	"time"
)

type Clocker interface{
	Now() time.Time

}

type RealClocker struct{}

func(r RealClocker) Now() time.Time{
	return time.Now()
}

type FixedClocker struct{}

func (fv FixedClocker) Now() time.Time{
	return time.Date(2022,5,5,5,5,5,5,time.UTC)
}
