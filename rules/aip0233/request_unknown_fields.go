// Copyright 2020 Google LLC
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

package aip0233

import (
	"fmt"

	"bitbucket.org/creachadair/stringset"

	"github.com/commure/api-linter/lint"
	"github.com/jhump/protoreflect/desc"
)

// Batch Create methods should not have unrecognized fields.
var requestUnknownFields = &lint.FieldRule{
	Name: lint.NewRuleName(233, "request-unknown-fields"),
	OnlyIf: func(f *desc.FieldDescriptor) bool {
		return isBatchCreateRequestMessage(f.GetOwner())
	},
	LintField: func(field *desc.FieldDescriptor) []lint.Problem {
		allowedFields := stringset.New(
			"parent",        // AIP-233
			"request_id",    // AIP-155
			"requests",      // AIP-233
			"validate_only", // AIP-163
		)
		if !allowedFields.Contains(field.GetName()) {
			return []lint.Problem{{
				Message: fmt.Sprintf(
					"Unexpected field: Batch Create RPCs must only contain fields explicitly described in https://aip.dev/233, not %q.",
					string(field.GetName()),
				),
				Descriptor: field,
			}}
		}

		return nil
	},
}
