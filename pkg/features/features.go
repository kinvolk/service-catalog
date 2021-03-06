/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package features

import (
	utilfeature "k8s.io/apiserver/pkg/util/feature"
)

const (
	// Every feature gate should add method here following this template:
	//
	// // owner: @username
	// // alpha: v1.4
	// MyFeature() bool

	// OriginatingIdentity controls whether the controller should include originating identity in the header of requests
	// sent to brokers
	//
	// owner: @pmorie
	// alpha: v1.7
	OriginatingIdentity utilfeature.Feature = "OriginatingIdentity"

	// AsyncBindingOperations controls whether the controller should
	// attempt asynchronous binding operations
	//
	// owner: @mkibbe
	// alpha: v1.7
	AsyncBindingOperations utilfeature.Feature = "AsyncBindingOperations"
)

func init() {
	utilfeature.DefaultFeatureGate.Add(defaultServiceCatalogFeatureGates)
}

// defaultServiceCatalogFeatureGates consists of all known
// feature keys that are specific to the service-catalog.
// To add a new feature, define a key for it above and add it here.
// The features will be available throughout service-catalog binaries.
var defaultServiceCatalogFeatureGates = map[utilfeature.Feature]utilfeature.FeatureSpec{
	OriginatingIdentity:    {Default: false, PreRelease: utilfeature.Alpha},
	AsyncBindingOperations: {Default: false, PreRelease: utilfeature.Alpha},
}
