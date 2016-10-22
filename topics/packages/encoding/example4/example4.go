// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how write a custom Unmarshal and Marshal functions.
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// NullableTime represents a nullable time for our sample.
type NullableTime struct {
	time.Time
}

// MarshalJSON implements the Marshaler interface.
func (t NullableTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}

	return []byte(`"` + t.Format(time.RFC3339) + `"`), nil
}

// UnmarshalJSON implements the Unmarshaler interface.
func (t *NullableTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		t.Time = time.Time{}
		return nil
	}

	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	t.Time = tt
	return nil
}

func main() {
	var n NullableTime
	if err := json.Unmarshal([]byte("null"), &n); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", n)
}
