package gentime

import (
	"fmt"
	"time"
)

// GenTime is function to generate time
func GenTime() string {
	cTime := time.Now()
	t := fmt.Sprint(cTime)

	return t
}
