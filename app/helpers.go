package app

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func GetInput() string {
	fmt.Println("find: ")
	line := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
		line = Sanitize_inputs(line)
	}
	return line
}

func Sanitize_inputs(str string) string {
	newStr := ""
	for _, char := range str {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			newStr = newStr + "\\" + string(char)
		} else {
			newStr = newStr + string(char)
		}
	}
	return newStr
}

func ValidateFileExtension(input string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	if re.MatchString(input) {
		return true
	} else {
		return false
	}
}

func ValidateConfigPath(file string) bool {
	if _, err := os.Stat(file); err == nil {
		fmt.Printf("File exists\n")
		return true
	} else {
		fmt.Printf("File does not exist\n")
		return false
	}
}

func RandomString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789~!@#$%^&*()_+`><?")
	hash := make([]rune, n)
	for i := range hash {
		hash[i] = chars[rand.Intn(len(chars))]
	}
	return string(hash)
}

func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func parseResults(resultSlice []string) (fileNames []string, lineNums []string, results []string) {
	for _, value := range resultSlice {
		tmpValue := value
		fileRegex := regexp.MustCompile("^./[^:]+")
		fileMatch := fileRegex.FindStringSubmatch(value)
		if len(fileMatch) > 0 {
			tmpValue = strings.Replace(tmpValue, fileMatch[0], "", 1)
			fileNames = append(fileNames, fileMatch[0])
		} else {
			continue
		}

		lineRegex := regexp.MustCompile("^:[0-9]+:")
		lineMatch := lineRegex.FindStringSubmatch(tmpValue)
		if len(lineMatch) > 0 {
			tmpValue = strings.Replace(tmpValue, lineMatch[0], "", 1)
			lineMatch[0] = strings.Replace(lineMatch[0], ":", "", -1)
			lineNums = append(lineNums, lineMatch[0])
		}
		if len(tmpValue) > 0 {
			results = append(results, tmpValue)
		}
	}
	return
}
