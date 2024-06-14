// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type ConfigSpec -type StatusSpec -type TunnelSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package siderolink

// DeepCopy generates a deep copy of ConfigSpec.
func (o ConfigSpec) DeepCopy() ConfigSpec {
	var cp ConfigSpec = o
	return cp
}

// DeepCopy generates a deep copy of StatusSpec.
func (o StatusSpec) DeepCopy() StatusSpec {
	var cp StatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of TunnelSpec.
func (o TunnelSpec) DeepCopy() TunnelSpec {
	var cp TunnelSpec = o
	return cp
}
