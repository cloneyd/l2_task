package search

import (
	"errors"
	"fmt"
	"strings"
)

const (
	lineNumFormat = "%d:%s\n" // Строка с номером
	lineFormat    = "%s\n"    // Строка
)

var (
	noEntry    = errors.New("no entry found")
	negAfter   = errors.New("after value should be positive")
	negBefore  = errors.New("before value should be positive")
	negContext = errors.New("context value should be positive")
	negCount   = errors.New("count value should be positive")
)

func OutputResult(lines []string, target string, after int, before int, context int, count bool, ignoreCase bool, invert bool, fixed bool, lineNum bool) error {
	tempLines := lines
	tempTarget := target

	var targetIdxs []int

	switch {
	case ignoreCase:
		tempLines = linesToLower(lines)
		tempTarget = strings.ToLower(target)
		fallthrough
	//
	//case after < 0:
	//	return negAfter
	//case after > 0:
	//	printAfter(lines, idxs, after, lineNum)
	//	return nil
	//
	//case before < 0:
	//	return negBefore
	//case before > 0:
	//	printBefore(lines, idxs, before)
	//	return nil
	//
	//case context < 0:
	//	return negContext
	//case context > 0:
	//	res = lines[entryIdx-context : entryIdx+context]

	default:
		targetIdxs = findEntries(tempLines, tempTarget, invert, fixed)
	}

	if len(targetIdxs) == 0 {
		return nil
	}

	if count {
		fmt.Println(len(targetIdxs))
	} else {
		printLines(lines, targetIdxs, lineNum)
	}

	return nil
}

// Вывод строк в STDOUT
func printLines(lines []string, idxs []int, lineNum bool) {
	for _, idx := range idxs {
		if lineNum {
			fmt.Printf(lineNumFormat, idx, lines[idx])
		} else {
			fmt.Printf(lineFormat, lines[idx])
		}
	}
}

// Поиск строк на основе переданных флагов
func findEntries(lines []string, target string, invert bool, fixed bool) []int {
	var indecies []int

	for idx, line := range lines {
		var condition bool

		if fixed {
			condition = line == target
		} else {
			condition = strings.Contains(line, target)
		}

		if invert {
			condition = !condition
		}

		if condition {
			indecies = append(indecies, idx)
		}
	}

	return indecies
}

func linesToLower(lines []string) []string {
	lowerLines := make([]string, len(lines))

	for idx, line := range lines {
		lowerLines[idx] = strings.ToLower(line)
	}

	return lowerLines
}

//func afterIndecies(lines []string, idxs []int, after int) []int {
//	var res []int
//
//	sizeLines := len(lines)
//	sizeIdxs := len(idxs)
//
//	for idx, foundIdx := range idxs {
//		cur := foundIdx
//		next :=
//	}
//}

//func printAfter(lines []string, idxs []int, after int, lineNum bool) {
//	sizeLines := len(lines)
//	sizeIdxs := len(idxs)
//
//	for idx, foundIdx := range idxs {
//		curFoundIdx := foundIdx
//		var nextFoundIdx int
//		if idx+1 < sizeIdxs {
//			nextFoundIdx = idxs[idx+1]
//		}
//
//		if foundIdx+after+1 > sizeLines {
//			printLines(lines[foundIdx:sizeLines], lineNum)
//			return
//		}
//
//		printLines(lines[foundIdx:foundIdx+after+1], lineNum)
//
//		if nextFoundIdx != 0 && curFoundIdx+after < nextFoundIdx-1 {
//			fmt.Println("-")
//		}
//	}
//}
//
//func printBefore(lines []string, idxs []int, before int, lineNum bool) {
//	for idx, foundIdx := range idxs {
//		curFoundIdx := foundIdx
//		var prevFoundIdx int
//		if idx-1 >= 0 {
//			prevFoundIdx = idxs[idx-1]
//		}
//
//		if foundIdx-before < 0 {
//			printLines(lines[0:foundIdx+1], lineNum)
//			return
//		}
//
//		if curFoundIdx-before > prevFoundIdx && idx != 0 {
//			fmt.Println("-")
//		}
//
//		printLines(lines[foundIdx-before:foundIdx+1], lineNum)
//	}
//}

// TODO: finish function
//func printContext(lines []string, idxs []int, context int) {
//	sizeLines := len(lines)
//	sizeIdxs := len(idxs)
//
//	for idx, foundIdx := range idxs {
//
//	}
//}
