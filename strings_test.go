package gotils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mreck/gotils"
)

func TestStringMatchesAny(t *testing.T) {
	assert.True(t, gotils.StringMatchesAny("foo", []string{"foo", "bar", "baz"}))
	assert.True(t, gotils.StringMatchesAny("bar", []string{"foo", "bar", "baz"}))
	assert.True(t, gotils.StringMatchesAny("baz", []string{"foo", "bar", "baz"}))

	assert.False(t, gotils.StringMatchesAny("fu", []string{"foo", "bar", "baz"}))
	assert.False(t, gotils.StringMatchesAny("buz", []string{"foo", "bar", "baz"}))
}

func TestStringContainsAny(t *testing.T) {
	assert.True(t, gotils.StringContainsAny("myfoo", []string{"foo", "bar", "baz"}))
	assert.True(t, gotils.StringContainsAny("bar2", []string{"foo", "bar", "baz"}))
	assert.True(t, gotils.StringContainsAny("1bazz", []string{"foo", "bar", "baz"}))

	assert.False(t, gotils.StringContainsAny("fo", []string{"foo", "bar", "baz"}))
	assert.False(t, gotils.StringContainsAny("1buzz", []string{"foo", "bar", "baz"}))
}

func TestStringContainsAll(t *testing.T) {
	assert.True(t, gotils.StringContainsAll("foobarbaz", []string{"foo", "bar", "baz"}))
	assert.True(t, gotils.StringContainsAll("foo bar baz", []string{"foo", "bar", "baz"}))
	assert.True(t, gotils.StringContainsAll("barbazfoo", []string{"foo", "bar", "baz"}))

	assert.False(t, gotils.StringContainsAll("foobar", []string{"foo", "bar", "baz"}))
	assert.False(t, gotils.StringContainsAll("bazbaz", []string{"foo", "bar", "baz"}))
}

func TestStringContainsNone(t *testing.T) {
	assert.True(t, gotils.StringContainsNone("foburbuz", []string{"foo", "bar", "baz"}))

	assert.False(t, gotils.StringContainsNone("foburbaz", []string{"foo", "bar", "baz"}))
}
