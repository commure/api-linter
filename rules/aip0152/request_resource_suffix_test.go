// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0152

import (
	"testing"

	"github.com/commure/api-linter/rules/internal/testutils"
)

func TestRequestResourceSuffix(t *testing.T) {
	for _, test := range []struct {
		name      string
		RPC       string
		Field     string
		FieldOpts string
		problems  testutils.Problems
	}{
		{"ValidPresent", "RunWriteBookJob", "resource_name", ` [(google.api.resource_reference) = { type: "WriteBookJob" }]`, nil},
		{"ValidMissing", "RunWriteBookJob", "resource_name", "", nil},
		{"IncorrectSuffix", "RunWriteBookJob", "resource_name", ` [(google.api.resource_reference) = { type: "Book" }]`,
			testutils.Problems{{Message: "google.api.resource_reference", Suggestion: "WriteBookJob"}}},
		{"IrrelevantMessage", "PurgeBooks", "resource_name", "", nil},
		{"IrrelevantField", "RunWriteBookJob", "something_else", "", nil},
	} {
		t.Run(test.name, func(t *testing.T) {
			f := testutils.ParseProto3Tmpl(t, `
				import "google/api/resource.proto";
				message {{.RPC}}Request {
					string {{.Field}} = 1{{.FieldOpts}};
				}
			`, test)
			field := f.GetMessageTypes()[0].GetFields()[0]
			if diff := test.problems.SetDescriptor(field).Diff(requestResourceSuffix.Lint(f)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
