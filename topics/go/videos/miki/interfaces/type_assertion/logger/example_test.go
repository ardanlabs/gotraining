package logger

import "os"

func ExampleLogger_Info() {
	log := New(os.Stdout)
	log.Info("CPU at %.2f", 12.34)

	// Output:
	// INFO: CPU at 12.34
}
