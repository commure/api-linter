// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0191

import (
	"strings"
	"testing"

	"github.com/commure/api-linter/rules/internal/testutils"
)

func TestJavaPackage(t *testing.T) {
	for _, test := range []struct {
		name       string
		statements []string
		problems   testutils.Problems
	}{
		{"Valid", []string{"package foo.v1;", `option java_package = "com.foo.v1";`}, testutils.Problems{}},
		{"InvalidEmpty", []string{"package foo.v1;", ""}, testutils.Problems{{Message: "java_package"}}},
		{"InvalidWrong", []string{"package foo.v1;", `option java_package = "something.else";`}, testutils.Problems{{
			Suggestion: `option java_package = "com.foo.v1";`,
		}}},
		{"Ignored", []string{"", ""}, testutils.Problems{}},
		{"IgnoredMaster", []string{"package foo.master;", ""}, testutils.Problems{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			f := testutils.ParseProto3String(t, strings.Join(test.statements, "\n"))
			if diff := test.problems.SetDescriptor(f).Diff(javaPackage.Lint(f)); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
