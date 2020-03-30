
# Module `testdata/legacy-block-labels`

Core Version Constraints:
* `>= 0.11.0`

Provider Requirements:
* **aws:** `1.0.0`
* **external:** (any version)
* **notnull:** (any version)
* **noversion:** (any version)

## Input Variables
* `foo` (default `"foo default"`): foo description

## Output Values
* `foo`: foo description

## Managed Resources
* `null_resource.foo` from `notnull`

## Data Resources
* `data.external.foo` from `external`

## Child Modules
* `foo` from `foo/bar/baz` (`1.2.3`)

