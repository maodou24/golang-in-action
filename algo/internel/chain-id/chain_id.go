package chainid

import "strings"

// ChainId 1.1， 1.1.1
type ChainID string

func (c ChainID) Parent() ChainID {
	strings.Split(string(c), ".")
	return ""
}
