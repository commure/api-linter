# Commure API Linter

[![ci](https://github.com/commure/api-linter/actions/workflows/ci.yaml/badge.svg)](https://github.com/commure/api-linter/actions/workflows/ci.yaml)
![latest release](https://img.shields.io/github/v/release/commure/api-linter)
![go version](https://img.shields.io/github/go-mod/go-version/commure/api-linter)

Note: This is a fork of the
[Google API Linter](https://github.com/googleapis/api-linter), made to allow
some customizations of its framework that are not possible through
configuration. Most notably, the `name` field used for resource keys has been
changed to `resource_name`.

The API linter provides real-time checks for compliance with many of Google's
API standards, documented using [API Improvement Proposals][]. It operates on
API surfaces defined in [protocol buffers][].

It identifies common mistakes and inconsistencies in API surfaces:

```proto
// Incorrect.
message GetBookRequest {
  // This is wrong; it should be spelled `name`.
  string book = 1;
}
```

When able, it also offers a suggestion for the correct fix.

[_Read more ≫_](https://linter.aip.dev/)

## Versioning

The Commure API linter does **not** follow semantic versioning. Semantic
versioning is challenging for a tool like a linter because the addition or
correction of virtually any rule is "breaking" (in the sense that a file that
previously reported no problems may now do so).

Therefore, the version numbers refer to the linter's core interface. In
general:

- Releases with only documentation, chores, dependency upgrades, and/or
  bugfixes are patch releases.
- Releases with new rules (or potentially removed rules) are minor releases.
- Releases with core interface alterations are major releases. This could
  include changes to the internal Go interface or the CLI user interface.

**Note:** Releases that increment the Go version will be considered minor.

This is an attempt to follow the spirit of semantic versioning while still
being useful.

## License

This software is made available under the [Apache 2.0][] license.

[apache 2.0]: https://www.apache.org/licenses/LICENSE-2.0
[api improvement proposals]: https://aip.dev/
[protocol buffers]: https://developers.google.com/protocol-buffers
[rule documentation]: ./rules/index.md
