// Copyright 2020, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !interface_string_no_reflect

package value

import "reflect"

// interfaceName returns the name of an interface type.
func interfaceName(b []byte, t reflect.Type, qualified bool) []byte {
	b = append(b, "interface{ "...)
	for i := 0; i < t.NumMethod(); i++ {
		if i > 0 {
			b = append(b, "; "...)
		}
		m := t.Method(i)
		if qualified && m.PkgPath != "" {
			b = append(b, '"')
			b = append(b, m.PkgPath...)
			b = append(b, '"')
			b = append(b, '.')
		}
		b = append(b, m.Name...)
		b = appendTypeName(b, m.Type, qualified, true)
	}
	if b[len(b)-1] == ' ' {
		b = b[:len(b)-1]
	} else {
		b = append(b, ' ')
	}
	b = append(b, '}')
	return b
}
