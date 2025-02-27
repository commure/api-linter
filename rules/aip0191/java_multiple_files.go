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

	"github.com/commure/api-linter/lint"
	"github.com/commure/api-linter/locations"
	"github.com/jhump/protoreflect/desc"
)

var javaMultipleFiles = &lint.FileRule{
	Name: lint.NewRuleName(191, "java-multiple-files"),
	OnlyIf: func(f *desc.FileDescriptor) bool {
		return hasPackage(f) && !strings.HasSuffix(f.GetPackage(), ".master")
	},
	LintFile: func(f *desc.FileDescriptor) []lint.Problem {
		if !f.GetFileOptions().GetJavaMultipleFiles() {
			return []lint.Problem{{
				Descriptor: f,
				Location:   locations.FilePackage(f),
				Message:    "Proto files must set `option java_multiple_files = true;`",
			}}
		}
		return nil
	},
}
