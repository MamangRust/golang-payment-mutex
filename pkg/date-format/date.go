package dateformat

import "time"

func DateFormat(date time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		return date
	}

	return date.In(loc)
}
