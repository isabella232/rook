/*
Copyright 2021 The Rook Authors. All rights reserved.

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

package mirror

import (
	"fmt"

	"github.com/rook/rook/pkg/daemon/ceph/client"
	"github.com/rook/rook/pkg/operator/ceph/config"
	"github.com/rook/rook/pkg/operator/ceph/config/keyring"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	keyringTemplate = `
[client.fs-mirror.%s]
	key = %s
	caps mon = "allow r"
	caps mgr = "allow r"
	caps mds = "allow r"
	caps osd = "'allow rw tag cephfs metadata=*, allow r tag cephfs data=*'"

`
)

// daemonConfig for a single rbd-mirror
type daemonConfig struct {
	ResourceName string              // the name rook gives to mirror resources in k8s metadata
	DaemonID     string              // the ID of the Ceph daemon ("a", "b", ...)
	DataPathMap  *config.DataPathMap // location to store data in container
	ownerRef     metav1.OwnerReference
}

// PeerToken is the content of the peer token
type PeerToken struct {
	ClusterFSID string `json:"fsid"`
	ClientID    string `json:"client_id"`
	Key         string `json:"key"`
	MonHost     string `json:"mon_host"`
}

func (r *ReconcileFilesystemMirror) generateKeyring(clusterInfo *client.ClusterInfo, daemonConfig *daemonConfig) (string, error) {
	user := fullDaemonName(daemonConfig.DaemonID)
	access := []string{
		"mon", "allow r",
		"mgr", "allow r",
		"mds", "allow r",
		"osd", "allow rw tag cephfs metadata=*, allow r tag cephfs data=*",
	}
	s := keyring.GetSecretStore(r.context, clusterInfo, &daemonConfig.ownerRef)

	key, err := s.GenerateKey(user, access)
	if err != nil {
		return "", err
	}

	keyring := fmt.Sprintf(keyringTemplate, daemonConfig.DaemonID, key)
	return keyring, s.CreateOrUpdate(daemonConfig.ResourceName, keyring)
}

func fullDaemonName(daemonID string) string {
	return fmt.Sprintf("client.fs-mirror.%s", daemonID)
}
