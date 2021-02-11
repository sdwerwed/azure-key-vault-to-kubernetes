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

package v2beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureKeyVaultSecret is a specification for a AzureKeyVaultSecret resource
type AzureKeyVaultSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureKeyVaultSecretSpec   `json:"spec"`
	Status AzureKeyVaultSecretStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureKeyVaultSecretList is a list of AzureKeyVaultSecret resources
type AzureKeyVaultSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AzureKeyVaultSecret `json:"items"`
}

// AzureKeyVaultSecretSpec is the spec for a AzureKeyVaultSecret resource
type AzureKeyVaultSecretSpec struct {
	Vault  AzureKeyVault       `json:"vault"`
	Output AzureKeyVaultOutput `json:"output,omitempty"`
}

// AzureKeyVault contains information needed to get the
// Azure Key Vault secret from Azure Key Vault
type AzureKeyVault struct {
	Name          string              `json:"name"`
	Object        AzureKeyVaultObject `json:"object"`
	AzureIdentity string              `json:"azureIdentity"`
}

// AzureKeyVaultObject has information about the Azure Key Vault
// object to get from Azure Key Vault
type AzureKeyVaultObject struct {
	Name        string                         `json:"name"`
	Type        AzureKeyVaultObjectType        `json:"type"`
	Version     string                         `json:"version"`
	Poll        bool                           `json:"bool"`
	ContentType AzureKeyVaultObjectContentType `json:"contentType"`
}

// AzureKeyVaultObjectType defines which Object type to get from Azure Key Vault
type AzureKeyVaultObjectType string

// AzureKeyVaultObjectContentType defines what content type a secret contains
type AzureKeyVaultObjectContentType string

const (
	// AzureKeyVaultObjectTypeSecret - get Secret object type from Azure Key Vault
	AzureKeyVaultObjectTypeSecret AzureKeyVaultObjectType = "secret"

	// AzureKeyVaultObjectTypeMultiKeyValueSecret - get Secret object type from Azure Key Vault containing multiple key/values
	AzureKeyVaultObjectTypeMultiKeyValueSecret = "multi-key-value-secret"

	// AzureKeyVaultObjectTypeCertificate - get Certificate object type from Azure Key Vault
	AzureKeyVaultObjectTypeCertificate = "certificate"

	// AzureKeyVaultObjectTypeKey - get Key object type from Azure Key Vault
	AzureKeyVaultObjectTypeKey = "key"

	// AzureKeyVaultObjectContentTypeJSON - object content is of type application/x-json
	AzureKeyVaultObjectContentTypeJSON AzureKeyVaultObjectContentType = "application/x-json"

	// AzureKeyVaultObjectContentTypeYaml - object content is of type application/x-yaml
	AzureKeyVaultObjectContentTypeYaml = "application/x-yaml"
)

// AzureKeyVaultOutput defines output sources, currently only support Secret
type AzureKeyVaultOutput struct {
	Secret    AzureKeyVaultOutputSecret    `json:"secret"`
	ConfigMap AzureKeyVaultOutputConfigMap `json:"configMap"`
	// +optional
	Transform []string `json:"transform,omitempty"`
}

// AzureKeyVaultOutputSecret has information needed to output
// a secret from Azure Key Vault to Kubertnetes as a Secret resource
type AzureKeyVaultOutputSecret struct {
	Name string `json:"name"`
	// +optional
	Type corev1.SecretType `json:"type,omitempty"`
	// +optional
	DataKey string `json:"dataKey,omitempty"`
	// +optional
	ChainOrder string `json:"chainOrder,omitempty"`
}

// AzureKeyVaultOutputConfigMap has information needed to output
// a secret from Azure Key Vault to Kubertnetes as a ConfigMap resource
type AzureKeyVaultOutputConfigMap struct {
	Name    string `json:"name"`
	DataKey string `json:"dataKey"`
}

// AzureKeyVaultSecretStatus is the status for a AzureKeyVaultSecret resource
type AzureKeyVaultSecretStatus struct {
	SecretHash      string      `json:"secretHash,omitempty"`
	SecretName      string      `json:"secretName,omitempty"`
	ConfigMapHash   string      `json:"configMapHash,omitempty"`
	ConfigMapName   string      `json:"configMapName,omitempty"`
	LastAzureUpdate metav1.Time `json:"lastAzureUpdate,omitempty"`
}