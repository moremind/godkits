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

package arrayx

import "testing"

func Test(t *testing.T) {
	lst := List{}
	lst.Add("hello", "world")

	tests := []struct {
		name string
		lst  List
		want int
	}{
		{
			lst:  lst,
			want: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lst := &List{}
			lst.Add("hello", "world")
			if got := lst.Size(); got != test.want {
				t.Errorf("Size() = %v, want %v", got, test.want)
			}
		})
	}
}
