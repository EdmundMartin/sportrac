# sportrac

Provides a Golang library for working with data which is exported from various fitness
tracking devices.

There is currently support for gpx and tcx file formats.

## GPX Examples
```
package main

import (
	"fmt"
	"github.com/EdmundMartin/sportrac/pkg/gpx"
)

func main() {
	
	gpxFile, err := gpx.LoadActivity("example.gpx")
	if err != nil {
		panic(err)
	}
	
	totalKm := gpxFile.TotalDistance()
	fmt.Println(totalKm)
	
	summary := gpxFile.GenerateSummary()
	fmt.Println(summary)
}

```