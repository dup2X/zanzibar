// Copyright (c) 2017 Uber Technologies, Inc.
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

package exampleReader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRequestSuccessfulRequestOKResponse(t *testing.T) {
	schema := (&exampleReaderMiddleware{}).JSONSchema()

	// TODO(sindelar): Read expected from test file
	expected := "{\n" +
		"    \"$schema\": \"http://json-schema.org/schema#\",\n" +
		"    \"type\": \"object\",\n" +
		"    \"properties\": {\n" +
		"        \"Foo\": {\n" +
		"            \"type\": \"string\"\n" +
		"        }\n" +
		"    },\n" +
		"    \"required\": [\n" +
		"        \"Foo\"\n" +
		"    ]\n" +
		"}"

	json, err := schema.Marshal()
	assert.Nil(t, err)
	assert.Equal(t, string(json), expected)
}
