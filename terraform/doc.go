/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

// Package terraform is the representation of a Terraform Module.
//
// It contains:
//
// • Header:        Module header found in shape of multi line '*.tf' comments or an entire file
//
// • Footer:        Module footer found in shape of multi line '*.tf' comments or an entire file
//
// • Inputs:        List of input 'variables' extracted from the Terraform module .tf files
//
// • ModuleCalls:   List of 'modules' extracted from the Terraform module .tf files
//
// • Outputs:       List of 'outputs' extracted from Terraform module .tf files
//
// • Providers:     List of 'providers' extracted from resources used in Terraform module
//
// • Requirements:  List of 'requirements' extracted from the Terraform module .tf files
//
// • Resources:     List of 'resources' extracted from the Terraform module .tf files
//
// Usage
//
//	options := &terraform.Options{
//	    Path:           "./examples",
//	    ShowHeader:     true,
//	    HeaderFromFile: "main.tf",
//	    ShowFooter:     true,
//	    FooterFromFile: "footer.md",
//	    SortBy: &terraform.SortBy{
//	        Name: true,
//	    },
//	    ReadComments: true,
//	}
//
//	tfmodule, err := terraform.LoadWithOptions(options)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	...
package terraform
