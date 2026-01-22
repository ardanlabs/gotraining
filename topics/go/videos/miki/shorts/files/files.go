package files

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Do a real health check
	fmt.Fprintln(w, "OK")
}
