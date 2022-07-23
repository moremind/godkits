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

// Copyright 2013-2014 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package properties

import (
	"fmt"
	"testing"
)

// Benchmarks the decoder by creating a property file with 1000 key/value pairs.
func BenchmarkLoad(b *testing.B) {
	input := ""
	for i := 0; i < 1000; i++ {
		input += fmt.Sprintf("key%d=value%d\n", i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Load([]byte(input), ISO_8859_1); err != nil {
			b.Fatal(err)
		}
	}
}
