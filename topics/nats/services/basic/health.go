// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
//
// Code provided by Kelsey Hightower: https://github.com/kelseyhightower/app
package main

import (
	"net/http"
	"sync/atomic"
)

var (
	healthzStatus   = int32(http.StatusOK)
	readinessStatus = int32(http.StatusOK)
)

// HealthzStatus returns the current state of the services health.
func HealthzStatus() int {
	return int(atomic.LoadInt32(&healthzStatus))
}

// ReadinessStatus returns the current state of the services readiness.
func ReadinessStatus() int {
	return int(atomic.LoadInt32(&readinessStatus))
}

// SetHealthzStatus changes the state of the services health.
func SetHealthzStatus(status int) {
	atomic.StoreInt32(&healthzStatus, int32(status))
}

// SetReadinessStatus changes the state of the services readiness.
func SetReadinessStatus(status int) {
	atomic.StoreInt32(&readinessStatus, int32(status))
}

// HealthzHandler responds to health check requests.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(HealthzStatus())
}

// ReadinessHandler responds to readiness check requests.
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(ReadinessStatus())
}

// ReadinessStatusHandler allows the health status to be changed.
func ReadinessStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch ReadinessStatus() {
	case http.StatusOK:
		SetReadinessStatus(http.StatusServiceUnavailable)

	case http.StatusServiceUnavailable:
		SetReadinessStatus(http.StatusOK)
	}

	w.WriteHeader(http.StatusOK)
}

// HealthzStatusHandler allows the readiness status to be changed.
func HealthzStatusHandler(w http.ResponseWriter, r *http.Request) {
	switch HealthzStatus() {
	case http.StatusOK:
		SetHealthzStatus(http.StatusServiceUnavailable)

	case http.StatusServiceUnavailable:
		SetHealthzStatus(http.StatusOK)
	}

	w.WriteHeader(http.StatusOK)
}
