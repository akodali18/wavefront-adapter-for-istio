// Copyright 2018 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config_test

import (
	"testing"

	"github.com/vmware/wavefront-adapter-for-istio/wavefront/config"
)

func TestMetricInfo(t *testing.T) {
	table := []struct {
		metric *config.Params_MetricInfo
		name   string
	}{
		{&config.Params_MetricInfo{InstanceName: "instance"}, "instance"},
		{&config.Params_MetricInfo{Name: "name", InstanceName: "instance"}, "name"},
	}

	for _, entry := range table {
		if name := config.MetricName(entry.metric); name != entry.name {
			t.Errorf("Validation failed, got: %v, want: %v.", name, entry.name)
		}
	}
}
