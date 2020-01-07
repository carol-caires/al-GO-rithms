package search

import "fmt"

func main() {

}

// BinarySearch searchs for a value in a list
func BinarySearch(list []string, item string) int {
	l := 0
	u := len(list) - 1

	fmt.Println("Searching", item, "in a list with", u, "items")
	for n := range list {
		m := int((l + u) / 2)

		if list[m] == item {
			fmt.Println("Got it after", n, "tries! Position of the item:", m)
			return m
		}
		if list[m] > item {
			fmt.Println(m, "superestimated guess")
			u = m - 1
		}
		if list[m] < item {
			fmt.Println(m, "underestimated guess")
			l = m + 1
		}
	}
	fmt.Println("Cannot find item :(")
	return -1
}
