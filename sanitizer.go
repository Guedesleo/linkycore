package linkycore

import (
	"github.com/microcosm-cc/bluemonday"
)

var SanitizerDefault *bluemonday.Policy

func InitSanitizer(opts *LinkyCoreOptions) {
	// Default police:
	SanitizerDefault = bluemonday.UGCPolicy()
	SanitizerDefault.AllowDataURIImages()
}

func GetSanitizer() *bluemonday.Policy {
	return SanitizerDefault
}
