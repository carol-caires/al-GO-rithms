package search

import "github.com/labstack/gommon/log"

func main() {

}

// BinarySearch searchs for a value in a list
func BinarySearch(list []string, item string) int {
	l := 0
	u := len(list) - 1

	log.Infof("Searching %s in a list with %d items", item, u)
	for n := range list {
		m := int((l + u) / 2)

		if list[m] == item {
			log.Infof("Got it after %d tries! Position of the item %d", n, m)
			return m
		} else if list[m] > item {
			log.Infof("%d is a superestimated guess", m)
			u = m - 1
		} else if list[m] < item {
			log.Infof("%d is a underestimated guess", m)
			l = m + 1
		}
	}
	log.Info("Cannot find item :(")
	return -1
}
