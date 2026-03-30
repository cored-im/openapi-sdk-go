// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package cocore

type Marshaller = func(v interface{}) ([]byte, error)
type Unmarshaller = func(data []byte, v interface{}) error
