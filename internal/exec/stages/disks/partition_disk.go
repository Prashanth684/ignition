// Copyright 2020 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// +build !s390x

package disks

import "github.com/coreos/ignition/v2/config/v3_2_experimental/types"

// partitionDisk partitions disks using the sgdisk command. For s390x, we need to
// check whether DASDs are being used and use different commands.
func (s stage) partitionDisk(dev types.Disk, devAlias string) error {
	return s.partitionGPTDisk(dev, devAlias)
}
