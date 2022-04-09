package weather

import (
	"fmt"
	"time"
)

func (d Day) Day() string {
	return fmt.Sprintf("%02d", d.unix().Day())
}

func (d Day) Month() string {
	return fmt.Sprintf("%02d", d.unix().Month())
}

func (d Day) unix() time.Time {
	return time.Unix(d.Time, 0)
}
