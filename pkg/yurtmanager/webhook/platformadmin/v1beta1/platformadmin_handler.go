/*
Copyright 2024 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the License);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an AS IS BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	yurtClient "github.com/openyurtio/openyurt/cmd/yurt-manager/app/client"
	"github.com/openyurtio/openyurt/cmd/yurt-manager/names"
	"github.com/openyurtio/openyurt/pkg/apis/iot/v1beta1"
	"github.com/openyurtio/openyurt/pkg/yurtmanager/controller/platformadmin/config"
	webhookutil "github.com/openyurtio/openyurt/pkg/yurtmanager/webhook/util"
)

// SetupWebhookWithManager sets up Cluster webhooks.
func (webhook *PlatformAdminHandler) SetupWebhookWithManager(mgr ctrl.Manager) (string, string, error) {
	// init
	webhook.Client = yurtClient.GetClientByControllerNameOrDie(mgr, names.PlatformAdminController)

	if err := webhook.initManifest(); err != nil {
		return "", "", err
	}
	return webhookutil.RegisterWebhook(mgr, &v1beta1.PlatformAdmin{}, webhook)
}

func (webhook *PlatformAdminHandler) initManifest() error {
	webhook.Manifests = &config.Manifest{}

	manifestContent, err := config.EdgeXFS.ReadFile(config.ManifestPath)
	if err != nil {
		klog.Error(err, "File to open the embed EdgeX manifest file")
		return err
	}

	if err := yaml.Unmarshal(manifestContent, webhook.Manifests); err != nil {
		klog.Error(err, "Error manifest EdgeX configuration file")
		return err
	}

	return nil
}

// Cluster implements a validating and defaulting webhook for Cluster.
type PlatformAdminHandler struct {
	Client    client.Client
	Manifests *config.Manifest
}

var _ webhook.CustomDefaulter = &PlatformAdminHandler{}
var _ webhook.CustomValidator = &PlatformAdminHandler{}
