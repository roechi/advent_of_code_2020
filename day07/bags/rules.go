package bags

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRule(rule string, g *graph) *graph {

	split := strings.Split(rule, "contain ")
	mainBagRule := strings.Split(split[0], " ")
	mainBagColor := strings.Join(mainBagRule[:len(mainBagRule)-2], " ")
	containedBags := strings.Split(split[1], ", ")

	colors := make([]string, len(containedBags))
	counts := make([]int, len(containedBags))

	for i, bagString := range containedBags {
		splitBagString := strings.Split(bagString, " ")
		count, _ := strconv.Atoi(splitBagString[0])
		color := strings.Join(splitBagString[1:len(splitBagString)-1], " ")

		colors[i] = color
		counts[i] = count
	}
	g.AddNode(mainBagColor)
	for i, _ := range containedBags {
		g.AddNode(colors[i])
		g.AddChild(mainBagColor, colors[i], counts[i])
	}

	return g
}

func ParseRules(rules []string) {
	graph := NewGraph()
	rulesParsed := 0
	for _, rule := range rules {
		ParseRule(rule, graph)
		rulesParsed++
	}

	sgCount := 0
	for _, node := range graph.GetElemMap() {
		if node.ChildrenContain("shiny gold") {
			sgCount++
		}
	}

	fmt.Println("Bags containing shiny gold bag: ", sgCount)

	n := graph.keysAndNodes["shiny gold"]
	fmt.Println("Bags contained in shiny gold bag: ", n.Size()-1)
}
