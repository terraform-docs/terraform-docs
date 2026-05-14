module full-example{
    source = "github.com/terraform-docs/terraform-docs//examples"

    input_with_underscores = "input_any"
    list-2 = [
        "value_1",
        "value_2"
    ]
    map-2 = {
        "1" = "value"
        "2" = "value_2"
    }
    number-2 = 2
    string-2 = "2"
    string_no_default = "string_value"
    unquoted = test
}