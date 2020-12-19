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
	literal        string
}

var ruleNumRegex, literalRegex, orRuleRegex *regexp.Regexp

var recursionLimit int = 3

func main() {
	// file, err := os.Open("test19-2.txt")
	file, err := os.Open("input19-2.txt")
	// file, err := os.Open("../test19.txt")
	// file, err := os.Open("../input19.txt")

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
	exp := composeLiterals(0, 0, &ruleBook)
	// log.Println(exp)
	messageCheck := regexp.MustCompile(`\b` + exp + `\b`)
	// messageCheck := regexp.MustCompile(exp)
	log.Println(messageCheck)
	var res int
	for _, message := range messages {
		if messageCheck.MatchString(message) {
			log.Println(message)
			res++
		}
	}
	// log.Println(ruleBook)
	// log.Println(messages)

	fmt.Println(res)
}

func composeLiterals(ruleNum int, recursionCount int, ruleBook *map[int]rule) string {
	currentRule := (*ruleBook)[ruleNum]
	if currentRule.literal != "" {
		return currentRule.literal
	}

	var sb strings.Builder

	if ruleNum == 8 {
		sb.WriteString("((")
		sb.WriteString(composeLiterals(42, recursionCount+1, ruleBook))
		sb.WriteString(")+)")
		literal := sb.String()
		(*ruleBook)[ruleNum] = rule{number: ruleNum, block1: currentRule.block1, block2: currentRule.block2, literal: literal}
		return literal
	}

	sb.WriteString("((")
	for _, num := range currentRule.block1 {
		if num == ruleNum && recursionCount > recursionLimit {
			continue
		}
		sb.WriteString(composeLiterals(num, recursionCount+1, ruleBook))
	}
	sb.WriteString(")")

	if len(currentRule.block2) == 0 {
		sb.WriteString(")")
		literal := sb.String()
		// regexp.MustCompile(literal)
		(*ruleBook)[ruleNum] = rule{number: ruleNum, block1: currentRule.block1, literal: literal}
		return literal
	}

	sb.WriteString("|(")

	for _, num := range currentRule.block2 {
		if num == ruleNum && recursionCount > recursionLimit {
			continue
		}
		sb.WriteString(composeLiterals(num, recursionCount+1, ruleBook))
	}

	sb.WriteString("))")
	literal := sb.String()
	// regexp.MustCompile(literal)
	(*ruleBook)[ruleNum] = rule{number: ruleNum, block1: currentRule.block1, block2: currentRule.block2, literal: literal}
	return literal
}

func parseRule(in string) rule {
	part1submatches := ruleNumRegex.FindStringSubmatch(in)
	ruleNumber, _ := strconv.Atoi(part1submatches[1])
	ruleInfo := part1submatches[2]
	if strings.Contains(ruleInfo, "\"") {
		lit := parseLiteral(ruleInfo)
		return rule{number: ruleNumber, literal: lit}
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

func parseLiteral(in string) string {
	submatches := literalRegex.FindStringSubmatch(in)
	return submatches[1]
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
