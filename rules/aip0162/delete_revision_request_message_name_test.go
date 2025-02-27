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

package aip0162

import (
	"testing"

	"github.com/commure/api-linter/rules/internal/testutils"
)

func TestDeleteRevisionRequestMessageName(t *testing.T) {
	// Set up the testing permutations.
	tests := []struct {
		testName       string
		MethodName     string
		ReqMessageName string
		problems       testutils.Problems
	}{
		{"Valid", "DeleteBookRevision", "DeleteBookRevisionRequest", nil},
		{"Invalid", "DeleteBookRevision", "DeleteBookRequest", testutils.Problems{{Suggestion: "DeleteBookRevisionRequest"}}},
		{"Irrelevant", "AcquireBook", "GetBookRequest", nil},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			f := testutils.ParseProto3Tmpl(t, `
				import "google/protobuf/empty.proto";
				service Library {
					rpc {{.MethodName}}({{.ReqMessageName}}) returns (google.protobuf.Empty) {}
				}
				message {{.ReqMessageName}} {}
			`, test)
			m := f.GetServices()[0].GetMethods()[0]
			if diff := test.problems.SetDescriptor(m).Diff(deleteRevisionRequestMessageName.Lint(f)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
