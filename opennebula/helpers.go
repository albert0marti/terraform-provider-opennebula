package opennebula

import (
	"fmt"
	"strings"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/shared"
)

func inArray(val string, array []string) (index int) {
	var ok bool
	for i := range array {
		if ok = array[i] == val; ok {
			return i
		}
	}
	return -1
}

// appendTemplate add attribute and value to an existing string
func appendTemplate(template, attribute, value string) string {
	return fmt.Sprintf("%s\n%s = \"%s\"", template, attribute, value)
}

// ArrayToString convert an array into a string
func ArrayToString(list []interface{}, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(list)), delim), "[]")
}

// StringToLockLevel convers a string into a "lock" object
func StringToLockLevel(str string, lock *shared.LockLevel) error {
	if str == "USE" {
		*lock = shared.LockUse
		return nil
	}
	if str == "MANAGE" {
		*lock = shared.LockManage
		return nil
	}
	if str == "ADMIN" {
		*lock = shared.LockAdmin
		return nil
	}
	if str == "ALL" {
		*lock = shared.LockAll
		return nil
	}
	return fmt.Errorf("Unexpected Lock level %s", str)
}

// LockLevelToString retuns a string when we provide a "lock" int
func LockLevelToString(lock int) string {
	if lock == 1 {
		return "USE"
	}
	if lock == 2 {
		return "MANAGE"
	}
	if lock == 3 {
		return "ADMIN"
	}
	if lock == 4 {
		return "ALL"
	}
	return ""
}
