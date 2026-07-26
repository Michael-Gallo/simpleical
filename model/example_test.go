// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package model_test

import (
	"fmt"
	"time"

	"github.com/michael-gallo/simpleical/model"
)

func ExampleNewUTCDateTime() {
	dt := model.NewUTCDateTime(time.Date(1997, 7, 14, 17, 0, 0, 0, time.UTC))
	fmt.Println(dt.IsUTC(), dt.Time.Format(time.RFC3339))
	// Output: true 1997-07-14T17:00:00Z
}

func ExampleNewFloatingDateTime() {
	dt := model.NewFloatingDateTime(time.Date(1998, 1, 18, 23, 0, 0, 0, time.UTC))
	fmt.Println(dt.Form == model.DateTimeFormFloating)
	fmt.Println(dt.TZID == "")
	// Output:
	// true
	// true
}

func ExampleNewDate() {
	dt := model.NewDate(time.Date(1997, 11, 2, 0, 0, 0, 0, time.UTC))
	fmt.Println(dt.IsDate(), dt.Time.Day())
	// Output: true 2
}

func ExampleNewLocalTZDateTime() {
	dt := model.NewLocalTZDateTime(time.Date(1998, 1, 19, 2, 0, 0, 0, time.UTC), "America/New_York")
	fmt.Println(dt.Form == model.DateTimeFormLocalTZ, dt.TZID)
	// Output: true America/New_York
}

func ExampleDateTime_IsDate() {
	dt := model.NewDate(time.Date(1997, 11, 2, 0, 0, 0, 0, time.UTC))
	fmt.Println(dt.IsDate())
	// Output: true
}

func ExampleDateTime_IsUTC() {
	dt := model.NewUTCDateTime(time.Date(1997, 7, 14, 17, 0, 0, 0, time.UTC))
	fmt.Println(dt.IsUTC())
	// Output: true
}

func ExampleDateTime_IsZero() {
	var unset model.DateTime
	fmt.Println(unset.IsZero())
	fmt.Println(model.NewDate(time.Date(1997, 11, 2, 0, 0, 0, 0, time.UTC)).IsZero())
	// Output:
	// true
	// false
}
