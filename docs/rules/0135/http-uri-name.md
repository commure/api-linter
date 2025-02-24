---
rule:
  aip: 135
  name: [core, '0135', http-uri-name]
  summary: Delete methods must map the name field to the URI.
permalink: /135/http-uri-name
redirect_from:
  - /0135/http-uri-name
---

# Delete methods: HTTP URI name field

This rule enforces that all `Delete` RPCs map the `resource_name` field to the HTTP URI,
as mandated in [AIP-135][].

## Details

This rule looks at any message matching beginning with `Delete`, and complains
if the `resource_name` variable is not included in the URI. It _does_ check additional
bindings if they are present.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  option (google.api.http) = {
    delete: "/v1/publishers/*/books/*"  // The `resource_name` field should be extracted.
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  option (google.api.http) = {
    delete: "/v1/{resource_name=publishers/*/books/*}"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aip.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0135::http-uri-name=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)
rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty) {
  option (google.api.http) = {
    delete: "/v1/publishers/*/books/*"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aip-135]: https://aip.dev/135
[aip.dev/not-precedent]: https://aip.dev/not-precedent
