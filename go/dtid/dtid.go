// Package dtid provides a function to generate a time-based IDs.
package dtid

import "time"

// New generates a time-based ID in the format "20060102150405" representing the
// current time in UTC.
func New() string {
	now := time.Now().UTC()
	// 2006 - year, 01 - month, 02 - day, 15 - hour, 04 - minute, 05 - second
	return now.Format("20060102150405")
}
