package main

import (
	. "./customsanswers"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCountUniqueAnswers(t *testing.T) {
	given := []string{"abc"} //\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	want := 3

	got := CountUniqueAnswers(given)

	assert.Equal(t, got, want)
}

func TestCountUniqueAnswers2(t *testing.T) {
	given := []string{"a", "b", "c"}
	want := 3

	got := CountUniqueAnswers(given)

	assert.Equal(t, got, want)
}

func TestCountUniqueAnswers3(t *testing.T) {
	given := []string{"ab", "ac"}
	want := 3

	got := CountUniqueAnswers(given)

	assert.Equal(t, got, want)
}

func TestCountUniqueAnswers4(t *testing.T) {
	given := []string{"a", "a", "a", "a"}
	want := 1

	got := CountUniqueAnswers(given)

	assert.Equal(t, got, want)
}

func TestCountUniqueAnswers5(t *testing.T) {
	given := []string{"b"}
	want := 1

	got := CountUniqueAnswers(given)

	assert.Equal(t, got, want)
}

func TestCountAllYes(t *testing.T) {
	given := []string{"abc"}
	want := 3

	got := CountAllYes(given)

	assert.Equal(t, got, want)
}

func TestCountAllYes2(t *testing.T) {
	given := []string{"a", "b", "c"}
	want := 0

	got := CountAllYes(given)

	assert.Equal(t, got, want)
}

func TestCountAllYes3(t *testing.T) {
	given := []string{"ab", "ac"}
	want := 1

	got := CountAllYes(given)

	assert.Equal(t, got, want)
}

func TestCountAllYes4(t *testing.T) {
	given := []string{"a", "a", "a", "a"}
	want := 1

	got := CountAllYes(given)

	assert.Equal(t, got, want)
}

func TestCountAllYes5(t *testing.T) {
	given := []string{"b"}
	want := 1

	got := CountAllYes(given)

	assert.Equal(t, got, want)
}
