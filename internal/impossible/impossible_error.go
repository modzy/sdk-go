package impossible

import "fmt"

// HandleError - This package has no coverage, and this method provides a way to remove
// "impossible" paths from coverage.
func HandleError(err error) {
	if err != nil {
		panic(fmt.Sprintf("impossible error: %v", err))
	}
}
