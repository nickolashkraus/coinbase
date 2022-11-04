package pkg

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetClicksByDomain(t *testing.T) {
	arr := []string{
		"900,google.com",
		"60,mail.yahoo.com",
		"10,mobile.sports.yahoo.com",
		"40,sports.yahoo.com",
		"300,yahoo.com",
		"10,stackoverflow.com",
		"20,overflow.com",
		"5,com.com",
		"2,en.wikipedia.org",
		"1,m.wikipedia.org",
		"1,mobile.sports",
		"1,google.co.uk",
	}
	exp := map[string]int{
		"co.uk":                   1,
		"com":                     1345,
		"com.com":                 5,
		"en.wikipedia.org":        2,
		"google.co.uk":            1,
		"google.com":              900,
		"m.wikipedia.org":         1,
		"mail.yahoo.com":          60,
		"mobile.sports":           1,
		"mobile.sports.yahoo.com": 10,
		"org":                     3,
		"overflow.com":            20,
		"sports":                  1,
		"sports.yahoo.com":        50,
		"stackoverflow.com":       10,
		"uk":                      1,
		"wikipedia.org":           3,
		"yahoo.com":               410,
	}
	ret := getClicksByDomain(arr)
	assert.Equal(t, exp, ret)
}
