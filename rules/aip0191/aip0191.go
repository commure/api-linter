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

// Package aip0191 contains rules defined in https://aip.dev/191.
package aip0191

import (
	"regexp"

	"github.com/commure/api-linter/lint"
	"github.com/jhump/protoreflect/desc"
)

// AddRules adds all of the AIP-191 rules to the provided registry.
func AddRules(r lint.RuleRegistry) error {
	return r.Register(
		191,
		csharpNamespace,
		filename,
		fileLayout,
		fileOptionConsistency,
		//javaMultipleFiles,
		//javaOuterClassname,
		//javaPackage,
		phpNamespace,
		//protoPkg, //TODO: Shawn please make this not look at the path from Root.
		rubyPackage,
		syntax,
	)
}

func hasPackage(f *desc.FileDescriptor) bool {
	return f.GetPackage() != ""
}

var versionRegexp = regexp.MustCompile("^v[0-9]+(p[0-9]+)?((alpha|beta)[0-9]*)?$")
