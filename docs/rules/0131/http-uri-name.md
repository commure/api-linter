---
rule:
  aip: 131
  name: [core, '0131', http-uri-name]
  summary: Get methods must map the resource_name field to the URI.
permalink: /131/http-uri-name
redirect_from:
  - /0131/http-uri-name
---

# Get methods: HTTP URI resource_name field

This rule enforces that all `Get` RPCs map the `rsource_name` field to the HTTP URI, as
mandated in [AIP-131][].

## Details

This rule looks at any message matching beginning with `Get`, and complains if
the `resource_name` variable is not included in the URI. It _does_ check additional
bindings if they are present.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc GetBook(GetBookRequest) returns (Book) {
  option (google.api.http) = {
    get: "/v1/publishers/*/books/*"  // The `resource_name` field should be extracted.
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc GetBook(GetBookRequest) returns (Book) {
  option (google.api.http) = {
    get: "/v1/{resource_name=publishers/*/books/*}"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aip.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0131::http-uri-name=disabled
//     aip.dev/not-precedent: We need to do this because reasons. --)
rpc GetBook(GetBookRequest) returns (Book) {
  option (google.api.http) = {
    get: "/v1/publishers/*/books/*"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aip-131]: https://aip.dev/131
[aip.dev/not-precedent]: https://aip.dev/not-precedent
