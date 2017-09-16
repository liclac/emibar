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

package i3bar

const (
	AlignLeft   = "left"
	AlignCenter = "center"
	AlignRight  = "right"

	MarkupPango = "pango"
)

// A block fed to i3bar.
type Block struct {
	Name     string `json:"name,omitempty"`
	Instance string `json:"instance,omitempty"`

	FullText  string `json:"full_text,omitempty"`
	ShortText string `json:"short_text,omitempty"`
	Markup    string `json:"markup,omitempty"` // Blank or "pango".

	Color      string `json:"color,omitempty"`      // #RRGGBB colours.
	Background string `json:"background,omitempty"` // #RRGGBB colours.
	Border     string `json:"border,omitempty"`     // #RRGGBB colours.
	Urgent     bool   `json:"urgent,omitempty"`

	MinWidth int    `json:"min_width,omitempty"`
	Align    string `json:"align,omitempty"` // center, right, left

	Separator           string `json:"separator,omitempty"`
	SeparatorBlockWidth int    `json:"separator_block_width,omitempty"`
}
