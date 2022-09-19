package search

import (
	"errors"
	"strings"
)

var (
	noEntry    = errors.New("no entry found")
	negAfter   = errors.New("after value should be positive value")
	negContext = errors.New("context value should be positive value")
	negCount   = errors.New("count value should be positive value")
)

func FindEntryInLine(lines []string, entry string) []int {
	var indecies []int

	lowerLines := make([]string, len(lines))
	for idx, line := range lines {
		lowerLines[idx] = strings.ToLower(line)
	}
	lowerEntry := strings.ToLower(entry)

	for idx, line := range lowerLines {
		if strings.Contains(line, lowerEntry) {
			indecies = append(indecies, idx)
		}
	}

	return indecies
}

//func OutputResult(lines []string, entry string, after int, before int, context int, count int, ignoreCase bool, invert bool, fixed bool, lineNum bool) error {
//	entryIdxs := findEntryInLine(lines, entry)
//	if len(entryIdxs) == 0 {
//		return noEntry
//	}
//
//	if after | context | before | count == 0 {
//		fmt.Println(lines[entryIdxs[0]])
//		return nil
//	}
//
//	if after < 0 {
//		return negAfter
//	} else {
//		fmt.Println(lines[entryIdxs[0] : entryIdx+after])
//		return nil
//	}
//
//	if after < 0 {
//		return negAfter
//	} else {
//		fmt.Println(lines[entryIdx : entryIdx+after])
//		return nil
//	}
//
//	if context < 0 {
//		return negContext
//	}
//	res = lines[entryIdx-context : entryIdx+context]
//
//	switch count
//}
