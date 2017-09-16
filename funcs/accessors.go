// Copyright © 2017 liclac
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package funcs

import (
	"reflect"
)

func init() {
	Registry["first"] = First
	Registry["last"] = Last
}

// Returns the first item in an array, or nil.
func First(items reflect.Value) interface{} {
	switch items.Kind() {
	case reflect.Array, reflect.Slice:
		if items.Len() > 0 {
			return items.Index(0).Interface()
		}
		return nil
	default:
		return items
	}
}

// Returns the last item in an array, or nil.
func Last(items reflect.Value) interface{} {
	switch items.Kind() {
	case reflect.Array, reflect.Slice:
		if l := items.Len(); l > 0 {
			return items.Index(l - 1).Interface()
		}
		return nil
	default:
		return items
	}
}
