package datetimeutils

import (
	"github.com/golang-module/carbon"
)

func GetNow() carbon.Carbon {
	return carbon.Now()
}

func GetNowString() string {
	return GetNow().ToRfc3339String()
}
