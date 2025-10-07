package models

import "context"

type Ride struct {
	ID            string  `json:"id"`
	RideNumber    string  `json:"ride_number"`
	PassengerID   string  `json:"passenger_id"`
	DriverID      *string `json:"driver_id,omitempty"`
	Status        string  `json:"status"`
	EstimatedFare float64 `json:"estimated_fare"`
}

type RideRepo interface {
	CreateRide(ctx context.Context, r *Ride) error
	UpdateRideStatus(ctx context.Context, rideID, status string, driverID *string) error
	GetRide(ctx context.Context, rideID string) (*Ride, error)
}
