package triedoublearraydict

import (
	"github.com/caicaispace/gohelper/tree/triedoublearray"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *triedoublearray.Cedar
	Values [][]string
}
