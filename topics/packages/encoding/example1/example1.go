// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to unmarshal a JSON document into
// a user defined struct type.
package main

import (
	"encoding/json"
	"fmt"
)

// document contains a JSON document.
var document = `{
"credentials": {
    "token": "06142010_1:75bf6a413327dd71ebe8f3f30c5a4210a9b11e93c028d6e11abfca7ff"
},
"valid": true,
"locale": "en_US",
"tnc_version": 2,
"preference_info": {
    "currency_code": "USD",
    "time_zone": "PST",
    "number_format": {
        "decimal_separator": ".",
        "grouping_separator": ",",
        "group_pattern": "###,##0.##"
    }
 }
}`

// Fields to be encoded/decoded must be exported else the
// json encoding functions can't see the fields.

type (
	// preferenceInfo represents preference information.
	preferenceInfo struct {
		CurrencyCode string `json:"currency_code"`
		TimeZone     string `json:"time_zone"`
		NumberFormat struct {
			DecimalSaparator  string `json:"decimal_separator"`
			GroupingSeparator string `json:"grouping_separator"`
			GroupPattern      string `json:"group_pattern"`
		} `json:"number_format"`
	}

	// userContext contains information for the user.
	userContext struct {
		Credentials struct {
			Token string `json:"token"`
		} `json:"credentials"`
		Valid          bool           `json:"valid"`
		Locale         string         `json:"locale"`
		TncVersion     int            `json:"tnc_version"`
		PreferenceInfo preferenceInfo `json:"preference_info"`
	}
)

func main() {

	// Declare a variable of type UserContext.
	var uc userContext

	// Unmarshal the JSON document into the variable.
	if err := json.Unmarshal([]byte(document), &uc); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n\n", uc)

	// Declare a pointer of type UserContext.
	var ucp *userContext

	// Unmarshal the JSON document into the nil pointer.
	if err := json.Unmarshal([]byte(document), &ucp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", *ucp)
}
