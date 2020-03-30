
# Module `testdata/overrides`

Provider Requirements:
* **null:** (any version)

## Input Variables
* `A` (required): The A variable OVERRIDDEN
* `B` (required): The B variable
* `C` (required): An entirely new variable C

## Output Values
* `A`: I am an overridden output!
* `B`: I am B

## Managed Resources
* `null_resource.A` from `null`
* `null_resource.B` from `null`

## Child Modules
* `foo` from `foo/bar/baz` (`1.0.2_override`)

