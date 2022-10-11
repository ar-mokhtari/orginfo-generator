package convertor

import (
	"errors"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"github.com/dlclark/regexp2"
	"reflect"
	"sync"
	"time"
)

//TODO::: split func to separate file(s)

// return time.Time from string
func StringToTime(myTime string) (time.Time, error) {
	layouts := DateLayouts()
	timeObj := time.Time{}
	err := errors.New("NO Layout match (StringToTime)")
	var wg sync.WaitGroup
	wg.Add(len(layouts))
	for _, v := range layouts {
		go func(s string) {
			defer wg.Done()
			tim1, err1 := time.Parse(s, myTime)
			if err1 == nil {
				timeObj = tim1
				err = nil
			}
		}(v)
	}
	wg.Wait()
	return timeObj, err
}

func DateLayouts() []string {
	return []string{
		"2006/01/02T15:04:05-07:00",
		"2006-01-02T15:04:05-07:00",
		"2006/01/02T15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02 15",
		"2006/01/02",
		"2006-01-02",
		"2006.01.02",
	}
}

// Returns true if input matches the passed pattern
func RegexRule(pattern string) func(interface{}) bool {
	regex := regexp2.MustCompile(pattern, regexp2.None)
	return func(input interface{}) bool {
		inputValue := reflect.ValueOf(input)
		switch inputValue.Kind() {
		case reflect.String:
			output, _ := regex.MatchString(input.(string))
			return output
		default:
			return false
		}
	}
}

func CheckIfComplex(input interface{}) bool {
	return RegexRule(env.RegULN8)(input)
}

func RemoveUnderlineFirstLast(input string) (res string) {
	lenInput := len(input)
	res = input
	if lenInput != 0 && input[:1] == "_" {
		res = input[1:]
	}
	if lenInput != 0 && res[len(res)-1:] == "_" {
		res = res[:len(res)-1]
	}
	return res
}
