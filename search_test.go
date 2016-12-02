package stringutil

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"WTS Rabadons Cap", true},
		{"wtS a cap", true},
		{"wtS a staff", true},
		{"fungi staff", false},
	}
	for _, c := range cases {
		got := CaseInsenstiveContains(c.in, "WTS")
		if got != c.want {
			t.Errorf("CaseInsesitiveContains(%q) == %q, want %q", c.in, got, c.want)
		}
	}

	casesB := []struct {
		in string
		want int
	}{
		{"WTS Rabadons Cap WTB a cow", 17},
		{"wtS a cap", -1},
		{"wtS a staff", -1},
		{"fungi staff WTB two purple daggers", 12},
	}
	for _, c := range casesB {
		got := CaseInsensitiveIndexOf(c.in, "wTB")
		if got != c.want {
			t.Errorf("CaseInesnsitiveIndexOf(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}