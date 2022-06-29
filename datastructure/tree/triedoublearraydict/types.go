package triedoublearraydict

import (
	tda "github.com/caicaispace/gohelper/datastructure/tree/triedoublearray"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *tda.Cedar
	Values [][]string
}
