package getformat

import (
	"os"
	"strings"
	"unicode"
)

const (
	standard   = "./fonts/standard.txt"
	thinkertoy = "./fonts/thinkertoy.txt"
	shadow     = "./fonts/shadow.txt"
)

func FinalOutput(input string, outputFormat string) (string, bool) {
	data := strings.ReplaceAll(input, "\\n", "\n")
	inputStr := strings.Split(data, "\n")
	var pathToFile string
	switch outputFormat {
	case "standard":
		pathToFile = standard
	case "thinkertoy":
		pathToFile = thinkertoy
	case "shadow":
		pathToFile = shadow
	default:
		return "Please choose the font", false
	}
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		return "Bad Request", false
	}
	res := getFormat(file, inputStr)
	lenInput := len(input)
	lenInputStr := len(inputStr)
	if lenInput < lenInputStr {
		res = res[:len(res)-1]
	}
	resultString := ""
	for i := 0; i < len(res); i++ {
		resultString += output(res[i])
	}
	return resultString, true
}

func getFormat(file []byte, inputStr []string) [][][]string {
	arr := splitDataTxt(strings.ReplaceAll(string(file), "\r", ""))
	maps := createMap(arr)
	var finalres [][]string
	for i := 0; i < len(inputStr); i++ {
		finalres = append(finalres, getStringMap(maps, inputStr[i]))
	}
	var res [][][]string
	for i := 0; i < len(finalres); i++ {
		res = append(res, getMatrixOflines(finalres[i]))
	}
	return res
}

// split by counting new lines
func splitDataTxt(str string) []string {
	cnt := 0
	l := len(str)
	var arr []string
	j := 0
	for i := 1; i < l; i++ {
		if str[i] == '\n' {
			cnt++
		}
		if cnt == 9 {
			if j+1 < l {
				arr = append(arr, str[j+1:i])
			}
			j = i
			cnt = 0
		}
	}
	arr = append(arr, str[j+1:l])
	return arr
}

// create map
func createMap(str []string) map[byte]string {
	maps := make(map[byte]string)
	var i byte
	for i = ' '; i <= '~'; i++ {
		maps[i] = str[i-32]
	}
	// maps['\n'] = "\n"
	return maps
}

// check the existence of element in map
func getStringMap(data map[byte]string, input string) []string {
	res := []string{}
	if len(input) == 0 {
		res = append(res, "\n")
	}
	for i := range input {
		if _, ok := data[input[i]]; ok {
			str := data[input[i]]
			res = append(res, str)
		}
	}
	return res
}

// gets matrix of each letter with each lines
func getMatrixOflines(data []string) [][]string {
	res := make([][]string, len(data))
	i := 0
	c := 0
	for _, v := range data {
		i = 0
		for j := range v {
			if v[j] == '\n' {
				res[c] = append(res[c], v[i:j])
				i = j + 1
			}
		}
		c++
	}
	return res
}

// the final output
func output(res [][]string) string {
	var abc []string
	s := ""
	lr := len(res)    // length of row
	lc := len(res[0]) // length of column
	for j := 0; j < lc; j++ {
		for i := 0; i < lr; i++ {
			s = s + res[i][j]
		}
		abc = append(abc, s+"\n")
		s = ""
	}
	result := ""
	for _, v := range abc {
		result += v
	}
	return result
}
func CheckLang(str string) bool {
	for _, v := range str {
		if v > unicode.MaxASCII {
			return false
		}
	}
	return true
}
