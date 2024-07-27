// Package hamming provides function to calculate Hamming distance
package hamming

import "errors"

// Distance calculates Hamming Distance and returns it only if two given strands of DNA have equal length otherwise it returns an error
func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return 0, errors.New("strands of DNA do not have the same length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
