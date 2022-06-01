package gamecolor

import (
	"regexp"
	"testing"
)

// TODO: These should be derived from []smap so that extending it
// automatically extends the tests properly
var OKColors = []GColor{Black, White, Red, Green}
var OKColorNames = []string{"Black", "White", "Red", "Green"}

var BADColors = []GColor{Black - 1, Green + 1, 99999}

func TestValidColors(t *testing.T) {
	for _, c := range OKColors {
		if !c.IsValid() {
			t.Errorf("Color %v (%d) incorrectly considered INVALID", c, c)
		}
	}
	for _, c := range BADColors {
		if c.IsValid() {
			t.Errorf("Color %v (%d) incorrectly considered VALID", c, c)
		}
	}
}

func TestValidStrings(t *testing.T) {
	for key, color := range smap {
		cstring := color.String()
		if cstring != key {
			t.Errorf("Incorrect name for color %v (%d): %s", color, color, cstring)
		}
	}
	for _, c := range BADColors {
		cstring := c.String()
		re := regexp.MustCompile("INVALID\\([-]{0,1}[0-9]+\\)")
		if !re.MatchString(cstring) {
			t.Errorf("Incorrect name for invalid color %d: %s", c, cstring)
		}
	}
}

func TestInitFromString(t *testing.T) {
	for idx, c := range OKColorNames {
		gc := FromString(c)
		if gc != OKColors[idx] {
			t.Errorf("Invalid conversion from string \"%s\" to color %d (got %d)",
				c, OKColors[idx], gc)
		}
	}
	for _, c := range [...]string{"Blueplatz", "", "INVALID"} {
		gc := FromString(c)
		if gc.IsValid() {
			t.Errorf("Invalid conversion from invalid string \"%s\" to valid color %v (%d)",
				c, gc, gc)
		}
	}
}
