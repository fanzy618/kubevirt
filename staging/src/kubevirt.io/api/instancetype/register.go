/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2021 Red Hat, Inc.
 *
 */

package instancetype

// GroupName is the group name used in this package
const (
	GroupName = "instancetype.kubevirt.io"

	SingularResourceName = "virtualmachineinstancetype"
	PluralResourceName   = SingularResourceName + "s"

	ClusterSingularResourceName = "virtualmachineclusterinstancetype"
	ClusterPluralResourceName   = ClusterSingularResourceName + "s"

	SingularPreferenceResourceName = "virtualmachinepreference"
	PluralPreferenceResourceName   = SingularPreferenceResourceName + "s"

	ClusterSingularPreferenceResourceName = "virtualmachineclusterpreference"
	ClusterPluralPreferenceResourceName   = ClusterSingularPreferenceResourceName + "s"
)

const (
	DefaultInstancetypeAnnotation     = "instancetype.kubevirt.io/default-instancetype"
	DefaultInstancetypeKindAnnotation = "instancetype.kubevirt.io/default-instancetype-kind"
	DefaultPreferenceAnnotation       = "instancetype.kubevirt.io/default-preference"
	DefaultPreferenceKindAnnotation   = "instancetype.kubevirt.io/default-preference-kind"
)
