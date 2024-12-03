package mostpopular

import (
	"fmt"
	"testing"
)

func TestContentPopularityManager(t *testing.T) {
	pm := NewContentPopularityManager()
	pm.IncreasePopularity(7)
	pm.IncreasePopularity(7)
	pm.IncreasePopularity(8)
	fmt.Printf("most p=%v\n", pm.MostPopular())
	pm.IncreasePopularity(8)
	pm.IncreasePopularity(8)
	fmt.Printf("most p=%v\n", pm.MostPopular())
	pm.DecreasePopularity(8)
	pm.DecreasePopularity(8)
	fmt.Printf("most p=%v\n", pm.MostPopular())
	pm.DecreasePopularity(7)
	pm.DecreasePopularity(7)
	pm.DecreasePopularity(8)
	fmt.Printf("most p=%v\n", pm.MostPopular())
}
