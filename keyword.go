package keyword

import (
	"strings"
)

const (
	orSeparator  = ","
	andSeparator = "&&"
)

// Checker .
type Checker struct {
	sensitive bool

	keyword [][]string
	exclude [][]string
}

// New .
func New(keyword, exclude string, sensitive bool) *Checker {
	// lower case all if no sensitive
	if !sensitive {
		keyword = strings.ToLower(keyword)
		exclude = strings.ToLower(exclude)
	}

	return &Checker{
		sensitive: sensitive,

		keyword: transform(keyword),
		exclude: transform(exclude),
	}
}

// CheckKeyword .
func (k *Checker) CheckKeyword(str string) bool {
	if !k.sensitive {
		str = strings.ToLower(str)
	}

	// if keyword is empty -> always true
	if len(k.keyword) == 0 {
		return true
	}

	return test(k.keyword, str)
}

// CheckExclude .
func (k *Checker) CheckExclude(str string) bool {
	if !k.sensitive {
		str = strings.ToLower(str)
	}

	// if exclude is empty -> always false
	if len(k.exclude) == 0 {
		return false
	}

	return test(k.exclude, str)
}

// Check .
func (k *Checker) Check(str string) bool {
	return k.CheckKeyword(str) && !k.CheckExclude(str)
}

func test(strArr [][]string, str string) bool {
	for i := 0; i < len(strArr); i++ {
		if !testArr(strArr[i], str) {
			return false
		}
	}

	return true
}

func testArr(strArr []string, str string) bool {
	for i := 0; i < len(strArr); i++ {
		if strings.Contains(str, strArr[i]) {
			return true
		}
	}

	return false
}

func trimArr(arr []string) (trimmedArr []string) {
	for i := 0; i < len(arr); i++ {
		str := strings.TrimSpace(arr[i])

		// remove empty string
		if str == "" {
			continue
		}

		trimmedArr = append(trimmedArr, str)
	}
	return
}

func transform(str string) (res [][]string) {
	// split and
	andArr := strings.Split(str, andSeparator)
	andArr = trimArr(andArr)

	// split or
	for i := 0; i < len(andArr); i++ {
		orArr := strings.Split(andArr[i], orSeparator)
		orArr = trimArr(orArr)

		// skip empty
		if len(orArr) == 0 {
			continue
		}

		res = append(res, orArr)
	}

	return
}
