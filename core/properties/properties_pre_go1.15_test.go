//go:build !go1.15
// +build !go1.15

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

package properties

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

// TestMustGetParsedDuration works with go before go1.15 where the panic
// message was changed slightly. We keep this test (!) here to demonstrate the
// backwards compatibility and to keep the author happy as long as it does not
// affect any real users. Thank you! Frank :)
//
// See https://github.com/magiconair/properties/pull/63
func TestMustGetParsedDuration(t *testing.T) {
	input := "key = 123ms\nkey2 = ghi"
	p := mustParse(t, input)
	assert.Equal(t, p.MustGetParsedDuration("key"), 123*time.Millisecond)
	assert.Panic(t, func() { p.MustGetParsedDuration("key2") }, `time: invalid duration ghi`)
	assert.Panic(t, func() { p.MustGetParsedDuration("invalid") }, "unknown property: invalid")
}
