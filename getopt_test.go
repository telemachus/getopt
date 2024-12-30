// Copyright 2017 The Go Authors. All rights reserved.
// Copyright 2024 Peter Aronoff. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getopt

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"strings"
	"testing"
)

type testFlagSet struct {
	A    *bool
	B    *bool
	C    *bool
	D    *bool
	I    *int
	Long *int
	S    *string
	flag *FlagSet
	Args []string
	buf  bytes.Buffer
}

func (tf *testFlagSet) str() string {
	out := ""
	if *tf.A {
		out += " -a"
	}
	if *tf.B {
		out += " -b"
	}
	if *tf.C {
		out += " -c"
	}
	if *tf.D {
		out += " -d"
	}
	if *tf.I != 0 {
		out += fmt.Sprintf(" -i %d", *tf.I)
	}
	if *tf.Long != 0 {
		out += fmt.Sprintf(" --long %d", *tf.Long)
	}
	if *tf.S != "" {
		out += " -s " + *tf.S
	}
	if len(tf.Args) > 0 {
		out += " " + strings.Join(tf.Args, " ")
	}
	if out == "" {
		return out
	}
	return out[1:]
}

func newTestFlagSet() *testFlagSet {
	tf := &testFlagSet{flag: NewFlagSet("x", flag.ContinueOnError)}
	f := tf.flag
	f.SetOutput(&tf.buf)
	tf.A = f.Bool("a", false, "desc of a")
	tf.B = f.Bool("b", false, "desc of b")
	tf.C = f.Bool("c", false, "desc of c")
	tf.D = f.Bool("d", false, "desc of d")
	tf.Long = f.Int("long", 0, "long only")
	f.Alias("a", "aah")
	f.Aliases("b", "beeta", "c", "charlie")
	tf.I = f.Int("i", 0, "i")
	f.Alias("i", "india")
	tf.S = f.String("sierra", "", "string")
	f.Alias("s", "sierra")

	return tf
}

func TestParse(t *testing.T) {
	tests := []struct {
		cmd string
		out string
	}{
		{"-i 1", "-i 1"},
		{"--india 1", "-i 1"},
		{"--india=1", "-i 1"},
		{"--i=1", "-i 1"},
		{"-abc", "-a -b -c"},
		{"-sfoo", "-s foo"},
		{"-s foo", "-s foo"},
		{"--s=foo", "-s foo"},
		{"-s=foo", "-s =foo"},
		{"--s=", ``},
		{"-sfooi1 -i2", "-i 2 -s fooi1"},
		{"-absfoo", "-a -b -s foo"},
		{"-i1 -- arg", "-i 1 arg"},
		{"-i1 - arg", "-i 1 - arg"},
		{"-i1 arg", "-i 1 arg"},
		{"--aah --charlie --beeta --sierra=123", "-a -b -c -s 123"},
		{"-i1 --long=2", "-i 1 --long 2"},
	}

	for _, tt := range tests {
		tf := newTestFlagSet()
		err := tf.flag.Parse(strings.Fields(tt.cmd))
		var out string
		if err != nil {
			t.Errorf("%s:\nhave %v\nwant <nil>\n", tt.cmd, err)
		}
		tf.Args = tf.flag.Args()
		out = tf.str()
		if out != tt.out {
			t.Errorf("%s:\nhave %s\nwant %s", tt.cmd, out, tt.out)
		}
	}
}

func TestParseErr(t *testing.T) {
	tests := []struct {
		err error
		cmd string
	}{
		{cmd: "-i=1", err: errors.New("parse error")},
		{cmd: "--abc", err: errors.New("parse error")},
		{cmd: "-s", err: errors.New("parse error")},
		{cmd: "--s", err: errors.New("parse error")},
		{cmd: "-i1 --- arg", err: errors.New("parse error")},
	}

	for _, tt := range tests {
		tf := newTestFlagSet()
		err := tf.flag.Parse(strings.Fields(tt.cmd))
		if err == nil {
			t.Errorf("%s:\nhave <nil>\nwant error\n", tt.cmd)
		}
	}
}

func TestHelpText(t *testing.T) {
	wantHelpText := `  -a, --aah
    	desc of a
  -b, --beeta
    	desc of b
  -c, --charlie
    	desc of c
  -d	desc of d
  -i, --india int
    	i
  --long int
    	long only
  -s, --sierra string
    	string
`
	tf := newTestFlagSet()
	tf.flag.PrintDefaults()
	out := tf.buf.String()
	if out != wantHelpText {
		t.Errorf("have<\n%s>\nwant<\n%s>", out, wantHelpText)
	}
}
