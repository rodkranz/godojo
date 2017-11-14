// Copyright 2017 The OLX Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package result

import (
	"strings"
)

func palindrome(s string) bool {
	whitoutSpaces := strings.Replace(s, " ", "", -1)

	if len(whitoutSpaces) < 2 {
		return true
	} else {
		r := []rune(whitoutSpaces)
		return r[0] == r[len(r)-1] && palindrome(string(r[1:len(r)-1]))
	}
}
