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

package aip0131

import (
	"github.com/commure/api-linter/lint"
	"github.com/commure/api-linter/rules/internal/utils"
)

// Get methods should have a proper HTTP pattern.
var httpNameField = &lint.MethodRule{
	Name:       lint.NewRuleName(131, "http-uri-name"),
	OnlyIf:     isGetMethod,
	LintMethod: utils.LintHTTPURIHasNameVariable,
}
