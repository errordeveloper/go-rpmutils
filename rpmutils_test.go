/*
 * Copyright (c) SAS Institute, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpmutils

import (
	"io/ioutil"
	"os"
	"testing"
	"testing/iotest"
)

func TestReadHeader(t *testing.T) {
	f, err := os.Open("./testdata/simple-1.0.1-1.i386.rpm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	hdr, err := ReadHeader(iotest.HalfReader(f))
	if err != nil {
		t.Fatal(err)
	}

	if hdr == nil {
		t.Fatal("no header found")
	}

	nevra, err := hdr.GetNEVRA()
	if err != nil {
		t.Fatal(err)
	}

	expectedNevra := NEVRA{
		Epoch:   "0",
		Name:    "simple",
		Version: "1.0.1",
		Release: "1",
		Arch:    "i386",
	}
	if expectedNevra != *nevra {
		t.Fatalf("incorrect nevra: %s (expected %s)",
			nevra.String(), expectedNevra.String())
	}

	files, err := hdr.GetFiles()
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 3 {
		t.Fatalf("incorrect number of files %d", len(files))
	}

	hdrRange := hdr.GetRange()
	expectedRange := HeaderRange{
		Start: 280,
		End:   1764,
	}
	if hdrRange != expectedRange {
		t.Errorf("incorrect header range %+v (expected %+v)",
			hdrRange, expectedRange)
	}
}

func TestEpoch(t *testing.T) {
	items := [][]string{
		{"testdata/zero-epoch-0.1-1.x86_64.rpm", "0"},
		{"testdata/one-epoch-0.1-1.x86_64.rpm", "1"},
	}
	for _, item := range items {
		filename, expected := item[0], item[1]
		t.Run(expected, func(t *testing.T) {
			f, err := os.Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			hdr, err := ReadHeader(f)
			if err != nil {
				t.Fatal(err)
			}
			nevra, err := hdr.GetNEVRA()
			if err != nil {
				t.Fatal(err)
			}
			if nevra.Epoch != expected {
				t.Errorf("%s: expected %q got %q", filename, expected, nevra.Epoch)
			}
		})
	}
}

func TestPayloadReader(t *testing.T) {
	f, err := os.Open("./testdata/simple-1.0.1-1.i386.rpm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	rpm, err := ReadRpm(iotest.HalfReader(f))
	if err != nil {
		t.Fatal(err)
	}

	pldr, err := rpm.PayloadReader()
	if err != nil {
		t.Fatal(err)
	}

	hdr, err := pldr.Next()
	if err != nil {
		t.Fatal(err)
	}

	if hdr.Size != 7 {
		t.Fatalf("wrong file size %d", hdr.Size)
	}

	if hdr.Name != "./config" {
		t.Fatalf("wrong file name %s", hdr.Name)
	}
}

func TestExpandPayload(t *testing.T) {
	f, err := os.Open("./testdata/simple-1.0.1-1.i386.rpm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	rpm, err := ReadRpm(iotest.HalfReader(f))
	if err != nil {
		t.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("", "rpmutil")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	if err := rpm.ExpandPayload(tmpdir); err != nil {
		t.Fatal(err)
	}
}
