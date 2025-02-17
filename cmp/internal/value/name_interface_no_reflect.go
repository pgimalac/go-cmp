// Copyright 2020, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build interface_string_no_reflect

package value

import "reflect"

// interfaceName returns the name of an interface type.
//
// When the build tag interface_string_no_reflect is set, just use reflect's string representation.
// This avoids disabling deadcode elimination due to using `reflect.Type.Method`.
// One caveat is that it doesn't follow the `qualified` parameter.
func interfaceName(b []byte, t reflect.Type, _ bool) []byte {
	return append(b, t.String()...)
}
