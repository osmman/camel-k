/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package minishift contains utilities for Minishift deployments
package minishift

import (
	"strconv"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	registryNamespace = "kube-system"
)

// FindRegistry returns the Minishift registry location if any
func FindRegistry() (*string, error) {
	svcs := v1.ServiceList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.SchemeGroupVersion.String(),
			Kind:       "Service",
		},
	}
	options := metav1.ListOptions{
		LabelSelector: "kubernetes.io/minikube-addons=registry",
	}
	if err := sdk.List(registryNamespace, &svcs, sdk.WithListOptions(&options)); err != nil {
		return nil, err
	}
	if len(svcs.Items) == 0 {
		return nil, nil
	}
	svc := svcs.Items[0]
	ip := svc.Spec.ClusterIP
	portStr := ""
	if len(svc.Spec.Ports) > 0 {
		port := svc.Spec.Ports[0].Port
		if port > 0 && port != 80 {
			portStr = ":" + strconv.FormatInt(int64(port), 10)
		}
	}
	registry := ip + portStr
	return &registry, nil
}
