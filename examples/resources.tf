// Description from comment
resource "azurerm_resource_group" "resource_with_static_count" {
  count = 1
}

// Description from comment
resource "azurerm_resource_group" "resource_with_dynamic_count" {
  count = "${var.count}"
}

resource "azurerm_resource_group" "resource_without_count" {
}