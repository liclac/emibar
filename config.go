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

package main

import (
	"text/template"

	"github.com/liclac/emibar/funcs"
)

type Template struct {
	*template.Template
}

func (t *Template) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var src string
	if err := unmarshal(&src); err != nil {
		return err
	}
	tpl, err := template.New("").Funcs(funcs.Registry).Parse(src)
	if err != nil {
		return err
	}
	t.Template = tpl
	return err
}

// Block definition.
type BlockConfig struct {
	Label    string   `yaml:"label"`
	Template Template `yaml:"tpl"`
}

// Struct for the configuration file.
type Config struct {
	Blocks []BlockConfig `yaml:"blocks"`
}
