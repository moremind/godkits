/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Copyright 2018 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import "testing"

func TestEqualEquals(t *testing.T) {
	if got, want := equal(2, "a", "a"), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestEqualFails(t *testing.T) {
	if got, want := equal(2, "a", "b"), "\tassert_test.go:16: got a want b \n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanics(t *testing.T) {
	if got, want := doesPanic(2, func() { panic("foo") }, ""), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanicsAndMatches(t *testing.T) {
	if got, want := doesPanic(2, func() { panic("foo") }, "foo"), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanicsAndDoesNotMatch(t *testing.T) {
	if got, want := doesPanic(2, func() { panic("foo") }, "bar"), "\tassert.go:62: got foo which does not match bar\n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanicsAndDoesNotPanic(t *testing.T) {
	if got, want := doesPanic(2, func() {}, "bar"), "\tassert.go:65: did not panic\n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestMatchesMatches(t *testing.T) {
	if got, want := matches(2, "aaa", "a"), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestMatchesDoesNotMatch(t *testing.T) {
	if got, want := matches(2, "aaa", "b"), "\tassert_test.go:52: got aaa which does not match b\n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
