package constants

type FoodScanStatus string

const (
	FoodScanPending FoodScanStatus = "PENDING"
	FoodScanSuccess FoodScanStatus = "SUCCESS"
	FoodScanFailed  FoodScanStatus = "FAILED"
)
