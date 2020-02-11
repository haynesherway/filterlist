package filterlist

import (
	"fmt"
	"strings"
)

// FilterList holds a list of values to be filtered on
type FilterList []string

// New returns an empty FilterList
func New() FilterList {
	return FilterList{}
}

// String returns a FilterList in a string format
func (f *FilterList) String() string {
	return fmt.Sprintf("%v", *f)
}

// Set sets the FilterList value
func (f *FilterList) Set(value string) error {
	*f = strings.Split(value, ",")
	return nil
}

// Add appends a value to the filter list
func (f *FilterList) Add(value string) {
	if !f.Has(value) {
		*f = append(*f, value)
	}
	return
}

// QueryString returns the FilterList as a sql query
func (f *FilterList) QueryString() string {
	if len(*f) == 0 {
		return ""
	} else if len(*f) == 1 {
		for _, val := range *f {
			return fmt.Sprintf("='%s'", val)
		}
	}

	return fmt.Sprintf("IN ('%s')", strings.Join(*f, "','"))
}

// Len returns the length of the FilterList
func (f *FilterList) Len() int {
	return len(*f)
}

// Has returns true if a FilterList has a value in it
func (f *FilterList) Has(find string) bool {
	for _, val := range *f {
		if val == find {
			return true
		}
	}
	return false
}
