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

package idutil

import (
	"github.com/openingo/godkits/gox/stringsx"
	"testing"
)

func TestRandomId(t *testing.T) {
	tests := []struct {
		length int
	}{
		{
			length: 60,
		},
	}
	for _, tt := range tests {
		t.Run(string(rune(tt.length)), func(t *testing.T) {
			result := RandomId(tt.length)
			println(result)
			if stringsx.Empty(result) {
				t.Errorf("RandomId() reulst is null")
				return
			}
		})
	}
}
