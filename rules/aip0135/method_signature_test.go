// Copyright 2019 Google LLC
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

package aip0135

import (
	"testing"

	"github.com/commure/api-linter/rules/internal/testutils"
)

func TestMethodSignature(t *testing.T) {
	for _, test := range []struct {
		name       string
		MethodName string
		Signature  string
		problems   testutils.Problems
	}{
		{"Valid", "DeleteBook", `option (google.api.method_signature) = "name";`, testutils.Problems{}},
		{"Missing", "DeleteBook", "", testutils.Problems{{Message: `(google.api.method_signature) = "name"`}}},
		{
			"Wrong",
			"DeleteBook",
			`option (google.api.method_signature) = "book";`,
			testutils.Problems{{Suggestion: `option (google.api.method_signature) = "name";`}},
		},
		{"Irrelevant", "BurnBook", "", testutils.Problems{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			f := testutils.ParseProto3Tmpl(t, `
				import "google/api/client.proto";
				import "google/protobuf/empty.proto";
				service Library {
					rpc {{.MethodName}}({{.MethodName}}Request) returns (google.protobuf.Empty) {
						{{.Signature}}
					}
				}
				message {{.MethodName}}Request {}
			`, test)
			m := f.GetServices()[0].GetMethods()[0]
			if diff := test.problems.SetDescriptor(m).Diff(methodSignature.Lint(f)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
