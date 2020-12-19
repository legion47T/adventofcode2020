package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	number         int
	block1, block2 []int
	literals       []string
}

var ruleNumRegex, literalRegex, orRuleRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test19.txt")
	file, err := os.Open("../input19.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	ruleNumRegex = regexp.MustCompile(`(\d*): (.*)`)
	literalRegex = regexp.MustCompile(`"(\w*)"`)
	orRuleRegex = regexp.MustCompile(`(.*) \| (.*)`)

	ruleBook := make(map[int]rule)
	messages := make([]string, 0)
	mode := 0
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			mode++
			continue
		}
		switch mode {
		case 0:
			currentRule := parseRule(line)
			ruleBook[currentRule.number] = currentRule
		case 1:
			messages = append(messages, line)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	literals := unique(composeLiterals(0, &ruleBook))
	var res int
	for _, message := range messages {
		for _, lit := range literals {
			if message == lit {
				res++
				break
			}
		}
	}
	// log.Println(ruleBook)
	// log.Println(messages)

	fmt.Println(res)
}

func composeLiterals(ruleNum int, ruleBook *map[int]rule) []string {
	currentRule := (*ruleBook)[ruleNum]
	if len(currentRule.literals) > 0 {
		return currentRule.literals
	}
	literals := make([]string, 0)
	subBlocks := make([][]string, 0)
	for _, num := range currentRule.block1 {
		subBlocks = append(subBlocks, composeLiterals(num, ruleBook))
	}
	literals = append(literals, composeBlock(subBlocks)...)
	if len(currentRule.block2) == 0 {
		(*ruleBook)[ruleNum] = rule{number: currentRule.number, block1: currentRule.block1, block2: currentRule.block2, literals: literals}
		return literals
	}
	subBlocks = make([][]string, 0)
	for _, num := range currentRule.block2 {
		subBlocks = append(subBlocks, composeLiterals(num, ruleBook))
	}
	literals = append(literals, composeBlock(subBlocks)...)
	(*ruleBook)[ruleNum] = rule{number: currentRule.number, block1: currentRule.block1, block2: currentRule.block2, literals: literals}
	return literals
}

func composeBlock(subBlocks [][]string) []string {
	if len(subBlocks) == 1 {
		return subBlocks[0]
	}
	var blocks []string
	for _, block := range subBlocks[0] {
		for _, subBlock := range composeBlock(subBlocks[1:]) {
			var sb strings.Builder
			sb.WriteString(block)
			sb.WriteString(subBlock)
			blocks = append(blocks, sb.String())
		}
	}
	return blocks
}

func parseRule(in string) rule {
	part1submatches := ruleNumRegex.FindStringSubmatch(in)
	ruleNumber, _ := strconv.Atoi(part1submatches[1])
	ruleInfo := part1submatches[2]
	if strings.Contains(ruleInfo, "\"") {
		lit := parseLiteral(ruleInfo)
		return rule{number: ruleNumber, literals: lit}
	} else if strings.Contains(ruleInfo, "|") {
		block1, block2 := parseOrRule(ruleInfo)
		return rule{number: ruleNumber, block1: block1, block2: block2}
	} else {
		block := parsePlainRule(ruleInfo)
		return rule{number: ruleNumber, block1: block}
	}
}

func parsePlainRule(in string) []int {
	nums := strings.Split(in, " ")
	block := make([]int, 0)
	for _, num := range nums {
		numVal, _ := strconv.Atoi(num)
		block = append(block, numVal)
	}
	return block
}

func parseOrRule(in string) ([]int, []int) {
	submatches := orRuleRegex.FindStringSubmatch(in)
	return parsePlainRule(submatches[1]), parsePlainRule(submatches[2])
}

func parseLiteral(in string) []string {
	submatches := literalRegex.FindStringSubmatch(in)
	return []string{submatches[1]}
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
