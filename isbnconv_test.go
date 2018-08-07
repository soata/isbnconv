package isbnconv

import (
	"testing"
)

func TestISBN13to10(t *testing.T) {
	list := []string{"9784873118222", "4873118220"}

	for i := 0; i < len(list); i += 2 {
		isbn10, err := ISBN13to10(list[i])
		if err != nil {
			t.Error(err.Error(), list[i], list[i+1], isbn10)
		}
		if isbn10 != list[i+1] {
			t.Error("not equal", list[i], list[i+1], isbn10)
		}
	}

}
