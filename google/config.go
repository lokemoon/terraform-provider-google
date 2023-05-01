// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package google

import (
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func MultiEnvSearch(ks []string) string {
	return transport_tpg.MultiEnvSearch(ks)
}
