// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package cosdk

import "encoding/json"

func Pretty(obj interface{}) string {
	s, _ := json.MarshalIndent(obj, "", "  ")
	return string(s)
}
