package dsutils

import "bufio"
import "strings" 

func LinesToArray(lines string) ([]string, error) {
	var result = []string{}
	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if scanner.Err() != nil {
		return []string{}, scanner.Err()
	}
	return result, nil
}

func ReverseArray(lines []string) []string {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}

func Trunc(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num]
	}
	return bnoden
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
