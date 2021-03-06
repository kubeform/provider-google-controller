/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"encoding/json"
	"fmt"
	"os"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/constants/google"
)

func (i *Invocation) GoogleProviderRef(name string) (*core.Secret, error) {
	sa := google.ServiceAccountFromEnv()
	if sa == "" {
		return nil, fmt.Errorf("GOOGLE_SERVICE_ACCOUNT_JSON_KEY and GOOGLE_APPLICATION_CREDENTIALS are empty")
	}

	val := map[string]string{
		"credentials": sa,
		"region":      "us-central1",
		"project":     os.Getenv(google.GOOGLE_PROJECT_ID),
	}

	data, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		StringData: map[string]string{
			"provider": string(data),
		},
	}, nil
}
