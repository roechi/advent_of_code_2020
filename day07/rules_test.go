package main

import (
	. "./bags"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestParseRule(t *testing.T) {
	ruleText := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	graph := NewGraph()

	graph = ParseRule(ruleText, graph)

	assert.Equal(t, graph.ElemCount(), 3)
}
