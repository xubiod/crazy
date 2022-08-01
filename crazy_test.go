package crazy

import (
	"testing"
)

func TestCrazyRune(t *testing.T) {
	checkAgainst := [][]rune{
		{'0', '0', '1'},
		{'0', '1', '0'},
		{'0', '2', '0'},
		{'1', '0', '1'},
		{'1', '1', '0'},
		{'1', '2', '2'},
		{'2', '0', '2'},
		{'2', '1', '2'},
		{'2', '2', '1'},
	}

	var r rune
	var err error

	for _, item := range checkAgainst {
		r, err = CrazyRune(item[0], item[1])
		if err != nil {
			t.Errorf("crazytest: errored\n%s", err.Error())
			t.Fail()
		}

		if r != item[2] {
			t.Errorf("crazytest: crz %d %d was invalid", item[0], item[1])
			t.FailNow()
		}
	}
}

func TestCrazyString(t *testing.T) {
	answer := "1001022211"

	question, err := CrazyString("0001112220", "0120120120")
	if err != nil {
		t.Errorf("crazytest: errored\n%s", err.Error())
		t.Fail()
	}

	if question != answer {
		t.Error("crazytest: question incorrect")
		t.Fail()
	}
}
