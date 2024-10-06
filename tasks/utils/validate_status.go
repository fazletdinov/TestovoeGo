package utils

import "strings"

var statusTasks = []string{"Не выполнено", "Выполнено"}

func ValidStatus(status string) bool {
	for _, value := range statusTasks {
		if strings.EqualFold(value, status) {
			return true
		}
	}
	return false
}
