package helpers

import "time"

func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// kalau ulang tahun belum lewat tahun ini → kurangi 1
	if now.Month() < dob.Month() ||
		(now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}