package search

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	ignoreCaseRegex = `(?i)`
)

var (
	negAfter   = errors.New("after value should be positive")
	negBefore  = errors.New("before value should be positive")
	negContext = errors.New("context value should be positive")
)

func OutputResult(lines []string, target string, after int, before int, context int, count bool, ignoreCase bool, invert bool, fixed bool, lineNum bool) error {
	targetIdxs := findEntries(lines, target, invert, fixed, ignoreCase)

	if len(targetIdxs) == 0 {
		return nil
	}

	switch {
	case count:
		fmt.Println(len(targetIdxs))

	case after < 0:
		return negAfter

	case before < 0:
		return negBefore

	case context < 0:
		return negContext
	case context > 0:
		after = context
		before = context
		fallthrough

	default:
		printLines(lines, targetIdxs, after, before, lineNum)
	}

	return nil
}

// Поиск строк на основе переданных флагов
func findEntries(lines []string, target string, invert bool, fixed bool, ignoreCase bool) []int {
	var indecies []int
	var sb strings.Builder

	// Чувствительность к регистру символов
	if ignoreCase {
		sb.WriteString(ignoreCaseRegex)
	}

	// Строгое совпадение
	if fixed {
		sb.WriteRune('^')
		sb.WriteString(target)
		sb.WriteRune('$')
	} else {
		sb.WriteString(target)
	}

	pattern := sb.String()
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for idx, line := range lines {
		match := re.MatchString(line)

		if invert {
			match = !match
		}

		if match {
			indecies = append(indecies, idx)
		}
	}

	return indecies
}

// Вывод строк в STDOUT
func printLines(lines []string, idxs []int, after int, before int, lineNum bool) {
	for _, idx := range idxs {
		var context []string

		if lineNum {
			context = getIndexedLineContext(lines, idx, after, before)
		} else {
			context = getLineContext(lines, idx, after, before)
		}

		for _, line := range context {
			fmt.Println(line)
		}
	}
}

// Функция получения "контекста" строки (before символов до, after символов после)
func getLineContext(lines []string, idx int, after int, before int) []string {
	size := len(lines)

	var context []string

	if maxInt(idx-before, 0) > 0 && before != 0 {
		context = append(context, "-")
	}

	for i := maxInt(idx-before, 0); i < idx; i++ {
		context = append(context, lines[i])
	}

	context = append(context, lines[idx])

	for i := idx + 1; i < minInt(idx+after+1, size); i++ {
		context = append(context, lines[i])
	}

	if idx+after+1 < size && after != 0 {
		context = append(context, "-")
	}

	return context
}

// Функция получения "контекста" строки с её номером (before символов до, after символов после)
func getIndexedLineContext(lines []string, idx int, after int, before int) []string {
	size := len(lines)

	var context []string

	if maxInt(idx-before, 0) > 0 && before != 0 {
		context = append(context, "-")
	}

	for i := maxInt(idx-before, 0); i < idx; i++ {
		context = append(context, buildIndexedLine(i, lines[i]))
	}

	context = append(context, buildIndexedLine(idx, lines[idx]))

	for i := idx + 1; i < minInt(idx+after+1, size); i++ {
		context = append(context, buildIndexedLine(i, lines[i]))
	}

	if idx+after+1 < size && after != 0 {
		context = append(context, "-")
	}

	return context
}

func buildIndexedLine(idx int, line string) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(idx))
	sb.WriteRune(':')
	sb.WriteString(line)

	return sb.String()
}

func minInt(first int, second int) int {
	if first < second {
		return first
	}

	return second
}

func maxInt(first int, second int) int {
	if first > second {
		return first
	}

	return second
}
