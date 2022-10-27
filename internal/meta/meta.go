package meta

import (
	"fmt"
	"os"
	"time"
)

var (
	buildVersion string
	buildDate    string
)

var (
	BuildVersion string
	BuildDate    time.Time
)

func init() {
	if buildVersion != "" {
		BuildVersion = buildVersion
	} else {
		BuildVersion = "unknown"
	}
	if date, err := time.Parse("2006-01-02 15:04:05", buildDate); err == nil {
		BuildDate = date
	} else {
		fmt.Fprintf(os.Stderr, "invalid build date: %s (buildDate=%q\n)", err, buildDate)
	}
}
