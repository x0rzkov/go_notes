package main

import(
	"fmt"
	"time"
	"crypto/sha1"
	"strings"
	"log"
)

var lp = log.Println

func pd(msg string) {
	if opts_intf["debug"].(bool) {
		lp(msg)
	}
}

func pl(params ...interface{}) {
	if opts_intf["verbose"].(bool) {
		fmt.Println(params...)
	}
}

func pf(msg string, params ...interface{}) {
	if opts_intf["verbose"].(bool) {
		fmt.Printf(msg, params...)
	}
}

func generate_sha1() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte("%$" + time.Now().String() + "e{")))
}

func short_sha(sha string) string{
	if len(sha) > 12 {
		return sha[:12]
	}
	return sha
}

func trim_whitespace(in_str string) string {
	return strings.Trim(in_str, " \n\r\t")
}
