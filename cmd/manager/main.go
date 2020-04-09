// Copyright 2019 The Kubernetes Authors.
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

package main

import (
	"flag"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/spf13/pflag"

	"github.com/open-cluster-management/multicloud-operators-channel/pkg/log/zap"

	"github.com/open-cluster-management/multicloud-operators-channel/cmd/manager/exec"
)

func main() {
	defer klog.Flush()
	exec.ProcessFlags()

	klog.InitFlags(nil)
	pflag.CommandLine.AddFlagSet(zap.FlagSet())

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	exec.HidKlogFlag(pflag.CommandLine)
	pflag.Parse()

	exec.RunManager(signals.SetupSignalHandler())
}
