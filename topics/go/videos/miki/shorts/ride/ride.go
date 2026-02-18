package ride

import "math"

// RideFare returns ride fare in ¢
func RideFare(distance float64, shared bool) int {
	if distance == 0 {
		return 0
	}

	fare := 250 // initial fare
	fare += int(math.Ceil(distance)) * 150

	if shared {
		fare = int(float64(fare) * 0.9)
	}

	return fare
}
