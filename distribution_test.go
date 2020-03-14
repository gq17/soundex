package soundex_test

import (
	"fmt"
	"testing"

	"github.com/gq17/soundex"
)

func TestGetIntNumber(t *testing.T) {
	num := soundex.GetIntNumber()
	fmt.Println("num: ", num)
}
