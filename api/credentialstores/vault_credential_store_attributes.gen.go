// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentialstores

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type VaultCredentialStoreAttributes struct {
	Address                  string `json:"address,omitempty"`
	Namespace                string `json:"namespace,omitempty"`
	CaCert                   string `json:"ca_cert,omitempty"`
	TlsServerName            string `json:"tls_server_name,omitempty"`
	TlsSkipVerify            bool   `json:"tls_skip_verify,omitempty"`
	Token                    string `json:"token,omitempty"`
	TokenHmac                string `json:"token_hmac,omitempty"`
	ClientCertificate        string `json:"client_certificate,omitempty"`
	ClientCertificateKey     string `json:"client_certificate_key,omitempty"`
	ClientCertificateKeyHmac string `json:"client_certificate_key_hmac,omitempty"`
	WorkerFilter             string `json:"worker_filter,omitempty"`
	TokenStatus              string `json:"token_status,omitempty"`
}

func AttributesMapToVaultCredentialStoreAttributes(in map[string]interface{}) (*VaultCredentialStoreAttributes, error) {
	if in == nil {
		return nil, fmt.Errorf("nil input map")
	}
	var out VaultCredentialStoreAttributes
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &out,
		TagName: "json",
	})
	if err != nil {
		return nil, fmt.Errorf("error creating mapstructure decoder: %w", err)
	}
	if err := dec.Decode(in); err != nil {
		return nil, fmt.Errorf("error decoding: %w", err)
	}
	return &out, nil
}

func (pt *CredentialStore) GetVaultCredentialStoreAttributes() (*VaultCredentialStoreAttributes, error) {
	if pt.Type != "vault" {
		return nil, fmt.Errorf("asked to fetch %s-type attributes but credential-store is of type %s", "vault", pt.Type)
	}
	return AttributesMapToVaultCredentialStoreAttributes(pt.Attributes)
}
