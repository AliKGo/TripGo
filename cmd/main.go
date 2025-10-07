package main

import (
	"flag"

	"ride-hail/internal/logger"
)

func main() {
	svc := flag.String("service", "ride", "service to run: ride | driver | admin")
	flag.Parse()

	switch *svc {
	case "ride":
		// runRideService()
	case "driver":
		// runDriverService()
	case "admin":
		// runAdminService()
	default:
		logger.Fatal("unknown service: %s", *svc)
	}
}
