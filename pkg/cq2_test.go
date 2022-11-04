package pkg

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestFindContiguousHistory(t *testing.T) {
	user0 := []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"}
	user1 := []string{"/start", "/pink", "/register", "/orange", "/red", "a"}
	user2 := []string{"a", "/one", "/two"}
	user3 := []string{"/pink", "/orange", "/yellow", "/plum", "/blue", "/tan", "/red", "/amber", "/HotRodPink", "/CornflowerBlue", "/LightGoldenRodYellow", "/BritishRacingGreen"}
	user4 := []string{"/pink", "/orange", "/amber", "/BritishRacingGreen", "/plum", "/blue", "/tan", "/red", "/lavender", "/HotRodPink", "/CornflowerBlue", "/LightGoldenRodYellow"}
	user5 := []string{"a"}
	user6 := []string{"/pink", "/orange", "/six", "/plum", "/seven", "/tan", "/red", "/amber"}

	ret := findContiguousHistory(user0, user1)
	assert.Equal(t, ret, []string{"/pink", "/register", "/orange"})
	ret = findContiguousHistory(user0, user2)
	assert.Equal(t, ret, []string{})
	ret = findContiguousHistory(user0, user0)
	assert.Equal(t, ret, []string{"/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"})
	ret = findContiguousHistory(user2, user1)
	assert.Equal(t, ret, []string{"a"})
	ret = findContiguousHistory(user5, user2)
	assert.Equal(t, ret, []string{"a"})
	ret = findContiguousHistory(user3, user4)
	assert.Equal(t, ret, []string{"/plum", "/blue", "/tan", "/red"})
	ret = findContiguousHistory(user4, user3)
	assert.Equal(t, ret, []string{"/plum", "/blue", "/tan", "/red"})
	ret = findContiguousHistory(user3, user6)
	assert.Equal(t, ret, []string{"/tan", "/red", "/amber"})
}
