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

package aip0134

import (
	"github.com/commure/api-linter/lint"
	"github.com/jhump/protoreflect/desc"
)

var requestMaskRequired = &lint.MessageRule{
	Name:   lint.NewRuleName(134, "request-mask-required"),
	OnlyIf: isUpdateRequestMessage,
	LintMessage: func(m *desc.MessageDescriptor) []lint.Problem {
		updateMask := m.FindFieldByName("update_mask")
		if updateMask == nil {
			return []lint.Problem{{
				Message:    "Update methods should have an `update_mask` field.",
				Descriptor: m,
			}}
		}
		return nil
	},
}
