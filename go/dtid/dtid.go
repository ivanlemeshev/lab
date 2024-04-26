package dtid

import "time"

func New() string {
	now := time.Now()
	// 2006 - year, 01 - month, 02 - day, 15 - hour, 04 - minute, 05 - second
	return now.Format("20060102150405")
}
