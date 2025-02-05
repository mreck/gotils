package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMatchesAny(t *testing.T) {
	assert.True(t, StringMatchesAny("foo", []string{"foo", "bar", "baz"}))
	assert.True(t, StringMatchesAny("bar", []string{"foo", "bar", "baz"}))
	assert.True(t, StringMatchesAny("baz", []string{"foo", "bar", "baz"}))

	assert.False(t, StringMatchesAny("fu", []string{"foo", "bar", "baz"}))
	assert.False(t, StringMatchesAny("buz", []string{"foo", "bar", "baz"}))
}

func TestStringContainsAny(t *testing.T) {
	assert.True(t, StringContainsAny("myfoo", []string{"foo", "bar", "baz"}))
	assert.True(t, StringContainsAny("bar2", []string{"foo", "bar", "baz"}))
	assert.True(t, StringContainsAny("1bazz", []string{"foo", "bar", "baz"}))

	assert.False(t, StringContainsAny("fo", []string{"foo", "bar", "baz"}))
	assert.False(t, StringContainsAny("1buzz", []string{"foo", "bar", "baz"}))
}

func TestStringContainsAll(t *testing.T) {
	assert.True(t, StringContainsAll("foobarbaz", []string{"foo", "bar", "baz"}))
	assert.True(t, StringContainsAll("foo bar baz", []string{"foo", "bar", "baz"}))
	assert.True(t, StringContainsAll("barbazfoo", []string{"foo", "bar", "baz"}))

	assert.False(t, StringContainsAll("foobar", []string{"foo", "bar", "baz"}))
	assert.False(t, StringContainsAll("bazbaz", []string{"foo", "bar", "baz"}))
}

func TestStringContainsNone(t *testing.T) {
	assert.True(t, StringContainsNone("foburbuz", []string{"foo", "bar", "baz"}))

	assert.False(t, StringContainsNone("foburbaz", []string{"foo", "bar", "baz"}))
}
