package helper

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"notifier/driver"
)

type Helper interface {
	RegexMatch(pattern, testString string) error
	UnMarshal(data []byte, dest interface{})
	Marshal(src interface{}) []byte
	ParseInt64(value string) (int64, error)
}

type helper struct {
	driver driver.Driver
}

func NewHelper(d driver.Driver) Helper {
	return &helper{
		driver: d,
	}
}

func (h *helper) RegexMatch(pattern, testString string) error {
	compiler, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("regexMatch: unable to compile regex: %s", err)
	}

    found := compiler.Match([]byte(testString))
	if !found {
		return fmt.Errorf("regexMatch: pattern does not match")
	}

	return nil
}

func (h *helper) UnMarshal(data []byte, dest interface{}) {
	_ = json.Unmarshal(data, dest)
}

func (h *helper) Marshal(src interface{}) []byte {
	res, _ := json.Marshal(src)
	return res
}

func (h *helper) ParseInt64(value string) (int64, error) {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("parseInt64: unable to convert value to int64: %s", err)
	}

	return val, nil
}
