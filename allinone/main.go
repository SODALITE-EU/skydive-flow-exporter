/*
 * Copyright (C) 2019 IBM, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	awsflowlogs "github.com/skydive-project/skydive-flow-exporter/awsflowlogs/pkg"
	"github.com/skydive-project/skydive-flow-exporter/core"
	secadvisor "github.com/skydive-project/skydive-flow-exporter/secadvisor/pkg"
	prom "github.com/skydive-project/skydive-flow-exporter/prom_sky_con/pkg"
)

func main() {
	core.Main("/etc/skydive/allinone.yml")
}

func init() {
	core.ClassifierHandlers.Register("subnet_autodiscovery", secadvisor.NewClassifySubnetWithAutoDiscovery, false)
	core.ManglerHandlers.Register("logstatus", secadvisor.NewMangleLogStatus, false)
	core.EncoderHandlers.Register("secadvisor", secadvisor.NewEncode, false)
	core.TransformerHandlers.Register("awsflowlogs", awsflowlogs.NewTransform, false)
	core.TransformerHandlers.Register("secadvisor", secadvisor.NewTransform, false)
	core.StorerHandlers.Register("prom_sky_con", prom.NewStorePrometheusWrapper, false)
}
