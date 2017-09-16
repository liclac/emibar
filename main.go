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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/liclac/emibar/i3bar"
	yaml "gopkg.in/yaml.v2"
)

// The protocol header we present to i3bar.
var Header = i3bar.Header{
	Version:    1,
	StopSignal: syscall.SIGSTOP,
	ContSignal: syscall.SIGCONT,
}

// Reusable buffer used for rendering.
var Buffer bytes.Buffer

func Render(confs []BlockConfig) []i3bar.Block {
	var blocks []i3bar.Block
	for _, bconf := range confs {
		Buffer.Reset()
		if err := bconf.Template.Execute(&Buffer, nil); err != nil {
			blocks = append(blocks, ErrorBlock(err))
			continue
		}
		block := i3bar.Block{FullText: strings.TrimSpace(string(Buffer.Bytes()))}
		if block.FullText != "" {
			if bconf.Label != "" {
				block.FullText = bconf.Label + " " + block.FullText
			}
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func Update(w io.Writer, out *json.Encoder, blocks ...i3bar.Block) {
	if err := out.Encode(blocks); err != nil {
		Update(w, out, ErrorBlock(err))
		return
	}
	fmt.Fprintln(w, ",")
}

func ErrorBlock(err error) i3bar.Block {
	return i3bar.Block{
		FullText: err.Error(),
		Urgent:   true,
	}
}

func Execute(w io.Writer, out *json.Encoder, args []string) error {
	// Start the block stream.
	if err := out.Encode(Header); err != nil {
		return err
	}
	fmt.Println("[")

	// Read the config file.
	path := "$HOME/.config/emibar.yml"
	if len(args) > 1 {
		path = args[1]
	}
	data, err := ioutil.ReadFile(os.ExpandEnv(path))
	if err != nil {
		return err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	// Render the first frame.
	Update(w, out, Render(config.Blocks)...)

	// Listen for some signals
	sigs := make(chan os.Signal)
	signal.Notify(sigs, Header.ContSignal, Header.StopSignal, syscall.SIGTERM, os.Interrupt)
	defer signal.Stop(sigs)

	frozen := false
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case sig := <-sigs:
			switch sig {
			case Header.StopSignal:
				frozen = true
			case Header.ContSignal:
				frozen = false
			default:
				return nil
			}
		case <-timer.C:
			if frozen {
				break
			}
			Update(w, out, Render(config.Blocks)...)
		}
	}
}

func main() {
	w := os.Stdout
	out := json.NewEncoder(w)
	if err := Execute(w, out, os.Args); err != nil {
		Update(w, out, ErrorBlock(err))
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
		defer signal.Stop(quit)
		<-quit
	}
}
