// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package target

type Option func(kt *KustTarget)

func WithDisableDeprecationWarningMessage() Option {
	return func(kt *KustTarget) {
		kt.disableDeprecationWarningMessage = true
	}
}
