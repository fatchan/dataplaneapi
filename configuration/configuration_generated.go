// Code generated by go generate; DO NOT EDIT.
// Copyright 2021 HAProxy Technologies
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

package configuration

import (
	"time"

	"github.com/go-openapi/runtime/flagext"
	"github.com/jessevdk/go-flags"

	"github.com/haproxytech/client-native/v2/models"

	"github.com/haproxytech/dataplaneapi/misc"
)

type configTypeDataplaneapi struct {
	EnabledListeners *[]string              `yaml:"scheme,omitempty" hcl:"scheme,omitempty"`
	CleanupTimeout   *time.Duration         `yaml:"cleanup-timeout,omitempty" hcl:"cleanup-timeout,omitempty"`
	GracefulTimeout  *time.Duration         `yaml:"graceful-timeout,omitempty" hcl:"graceful-timeout,omitempty"`
	MaxHeaderSize    *flagext.ByteSize      `yaml:"max-header-size,omitempty" hcl:"max-header-size,omitempty"`
	SocketPath       *flags.Filename        `yaml:"socket-path,omitempty" hcl:"socket-path,omitempty"`
	Host             *string                `yaml:"host,omitempty" hcl:"host,omitempty"`
	Port             *int                   `yaml:"port,omitempty" hcl:"port,omitempty"`
	ListenLimit      *int                   `yaml:"listen-limit,omitempty" hcl:"listen-limit,omitempty"`
	KeepAlive        *time.Duration         `yaml:"keep-alive,omitempty" hcl:"keep-alive,omitempty"`
	ReadTimeout      *time.Duration         `yaml:"read-timeout,omitempty" hcl:"read-timeout,omitempty"`
	WriteTimeout     *time.Duration         `yaml:"write-timeout,omitempty" hcl:"write-timeout,omitempty"`
	ShowSystemInfo   *bool                  `yaml:"show-system-info,omitempty" hcl:"show-system-info,omitempty"`
	DisableInotify   *bool                  `yaml:"disable-inotify,omitempty" hcl:"disable-inotify,omitempty"`
	Tls              *configTypeTls         `yaml:"tls,omitempty" hcl:"tls,omitempty"`
	Userlist         *configTypeUserlist    `yaml:"userlist,omitempty" hcl:"userlist,omitempty"`
	Transaction      *configTypeTransaction `yaml:"transaction,omitempty" hcl:"transaction,omitempty"`
	Resources        *configTypeResources   `yaml:"resources,omitempty" hcl:"resources,omitempty"`
	Advertised       *configTypeAdvertised  `yaml:"advertised,omitempty" hcl:"advertised,omitempty"`
}

type configTypeTls struct {
	TLSHost           *string         `yaml:"tls-host,omitempty" hcl:"tls-host,omitempty"`
	TLSPort           *int            `yaml:"tls-port,omitempty" hcl:"tls-port,omitempty"`
	TLSCertificate    *flags.Filename `yaml:"tls-certificate,omitempty" hcl:"tls-certificate,omitempty"`
	TLSCertificateKey *flags.Filename `yaml:"tls-key,omitempty" hcl:"tls-key,omitempty"`
	TLSCACertificate  *flags.Filename `yaml:"tls-ca,omitempty" hcl:"tls-ca,omitempty"`
	TLSListenLimit    *int            `yaml:"tls-listen-limit,omitempty" hcl:"tls-listen-limit,omitempty"`
	TLSKeepAlive      *time.Duration  `yaml:"tls-keep-alive,omitempty" hcl:"tls-keep-alive,omitempty"`
	TLSReadTimeout    *time.Duration  `yaml:"tls-read-timeout,omitempty" hcl:"tls-read-timeout,omitempty"`
	TLSWriteTimeout   *time.Duration  `yaml:"tls-write-timeout,omitempty" hcl:"tls-write-timeout,omitempty"`
}

type configTypeHaproxy struct {
	ConfigFile       *string           `yaml:"config-file,omitempty" hcl:"config-file,omitempty"`
	HAProxy          *string           `yaml:"haproxy-bin,omitempty" hcl:"haproxy-bin,omitempty"`
	MasterRuntime    *string           `yaml:"master-runtime,omitempty" hcl:"master-runtime,omitempty"`
	NodeIDFile       *string           `yaml:"fid,omitempty" hcl:"fid,omitempty"`
	MasterWorkerMode *bool             `yaml:"master-worker-mode,omitempty" hcl:"master-worker-mode,omitempty"`
	Reload           *configTypeReload `yaml:"reload,omitempty" hcl:"reload,omitempty"`
}

type configTypeUserlist struct {
	Userlist     *string `yaml:"userlist,omitempty" hcl:"userlist,omitempty"`
	UserListFile *string `yaml:"userlist-file,omitempty" hcl:"userlist-file,omitempty"`
}

type configTypeReload struct {
	ReloadDelay     *int    `yaml:"reload-delay,omitempty" hcl:"reload-delay,omitempty"`
	ReloadCmd       *string `yaml:"reload-cmd,omitempty" hcl:"reload-cmd,omitempty"`
	RestartCmd      *string `yaml:"restart-cmd,omitempty" hcl:"restart-cmd,omitempty"`
	ReloadRetention *int    `yaml:"reload-retention,omitempty" hcl:"reload-retention,omitempty"`
	ValidateCmd     *string `yaml:"validate-cmd,omitempty" hcl:"validate-cmd,omitempty"`
}

type configTypeTransaction struct {
	TransactionDir      *string `yaml:"transaction-dir,omitempty" hcl:"transaction-dir,omitempty"`
	BackupsNumber       *int    `yaml:"backups-number,omitempty" hcl:"backups-number,omitempty"`
	BackupsDir          *string `yaml:"backups-dir,omitempty" hcl:"backups-dir,omitempty"`
	MaxOpenTransactions *int64  `yaml:"max-open-transactions,omitempty" hcl:"max-open-transactions,omitempty"`
}

type configTypeResources struct {
	MapsDir              *string `yaml:"maps-dir,omitempty" hcl:"maps-dir,omitempty"`
	SSLCertsDir          *string `yaml:"ssl-certs-dir,omitempty" hcl:"ssl-certs-dir,omitempty"`
	UpdateMapFiles       *bool   `yaml:"update-map-files,omitempty" hcl:"update-map-files,omitempty"`
	UpdateMapFilesPeriod *int64  `yaml:"update-map-files-period,omitempty" hcl:"update-map-files-period,omitempty"`
	SpoeDir              *string `yaml:"spoe-dir,omitempty" hcl:"spoe-dir,omitempty"`
	SpoeTransactionDir   *string `yaml:"spoe-transaction-dir,omitempty" hcl:"spoe-transaction-dir,omitempty"`
}

type configTypeCluster struct {
	ClusterTLSCertDir  *string `yaml:"cluster-tls-dir,omitempty" hcl:"cluster-tls-dir,omitempty"`
	ID                 *string `yaml:"id,omitempty" hcl:"id,omitempty"`
	BootstrapKey       *string `yaml:"bootstrap_key,omitempty" hcl:"bootstrap_key,omitempty"`
	ActiveBootstrapKey *string `yaml:"active_bootstrap_key,omitempty" hcl:"active_bootstrap_key,omitempty"`
	Token              *string `yaml:"token,omitempty" hcl:"token,omitempty"`
	URL                *string `yaml:"url,omitempty" hcl:"url,omitempty"`
	Port               *string `yaml:"port,omitempty" hcl:"port,omitempty"`
	APIBasePath        *string `yaml:"api_base_path,omitempty" hcl:"api_base_path,omitempty"`
	APINodesPath       *string `yaml:"api_nodes_path,omitempty" hcl:"api_nodes_path,omitempty"`
	APIRegisterPath    *string `yaml:"api_register_path,omitempty" hcl:"api_register_path,omitempty"`
	CertificateDir     *string `yaml:"cert-path,omitempty" hcl:"cert-path,omitempty"`
	CertificateFetched *bool   `yaml:"cert-fetched,omitempty" hcl:"cert-fetched,omitempty"`
	Name               *string `yaml:"name,omitempty" hcl:"name,omitempty"`
	Description        *string `yaml:"description,omitempty" hcl:"description,omitempty"`
}

type configTypeAdvertised struct {
	APIAddress *string `yaml:"api-address,omitempty" hcl:"api-address,omitempty"`
	APIPort    *int64  `yaml:"api-port,omitempty" hcl:"api-port,omitempty"`
}

type configTypeServiceDiscovery struct {
	Consuls *[]*models.Consul `yaml:"consuls,omitempty" hcl:"consuls,omitempty"`
}

type configTypeSyslog struct {
	SyslogSrv      *string `yaml:"syslog-server,omitempty" hcl:"syslog-server,omitempty"`
	SyslogPort     *uint   `yaml:"syslog-port,omitempty" hcl:"syslog-port,omitempty"`
	SyslogProto    *string `yaml:"syslog-protocol,omitempty" hcl:"syslog-protocol,omitempty"`
	SyslogTag      *string `yaml:"syslog-tag,omitempty" hcl:"syslog-tag,omitempty"`
	SyslogPriority *string `yaml:"syslog-priority,omitempty" hcl:"syslog-priority,omitempty"`
	SyslogFacility *string `yaml:"syslog-facility,omitempty" hcl:"syslog-facility,omitempty"`
}

type configTypeLog struct {
	LogTo     *string           `yaml:"log-to,omitempty" hcl:"log-to,omitempty"`
	LogFile   *string           `yaml:"log-file,omitempty" hcl:"log-file,omitempty"`
	LogLevel  *string           `yaml:"log-level,omitempty" hcl:"log-level,omitempty"`
	LogFormat *string           `yaml:"log-format,omitempty" hcl:"log-format,omitempty"`
	ACLFormat *string           `yaml:"apache-common-log-format,omitempty" hcl:"apache-common-log-format,omitempty"`
	Syslog    *configTypeSyslog `yaml:"syslog,omitempty" hcl:"syslog,omitempty"`
}

type StorageDataplaneAPIConfiguration struct {
	Version                *int                        `yaml:"config_version,omitempty" hcl:"config_version,omitempty"`
	Name                   *string                     `yaml:"name,omitempty" hcl:"name,omitempty"`
	Mode                   *string                     `yaml:"mode,omitempty" hcl:"mode,omitempty"`
	DeprecatedBootstrapKey *string                     `yaml:"bootstrap_key,omitempty" hcl:"bootstrap_key,omitempty"`
	Status                 *string                     `yaml:"status,omitempty" hcl:"status,omitempty"`
	Dataplaneapi           *configTypeDataplaneapi     `yaml:"dataplaneapi,omitempty" hcl:"dataplaneapi,omitempty"`
	Haproxy                *configTypeHaproxy          `yaml:"haproxy,omitempty" hcl:"haproxy,omitempty"`
	Cluster                *configTypeCluster          `yaml:"cluster,omitempty" hcl:"cluster,omitempty"`
	ServiceDiscovery       *configTypeServiceDiscovery `yaml:"service_discovery,omitempty" hcl:"service_discovery,omitempty"`
	Log                    *configTypeLog              `yaml:"log,omitempty" hcl:"log,omitempty"`
}

func copyToConfiguration(cfg *Configuration) {
	cfgStorage := cfg.storage.Get()
	if cfgStorage.Name != nil {
		cfg.Name.Store(*cfgStorage.Name)
	}
	if cfgStorage.Mode != nil {
		cfg.Mode.Store(*cfgStorage.Mode)
	}
	if cfgStorage.DeprecatedBootstrapKey != nil {
		cfg.DeprecatedBootstrapKey.Store(*cfgStorage.DeprecatedBootstrapKey)
	}
	if cfgStorage.Status != nil {
		cfg.Status.Store(*cfgStorage.Status)
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.ShowSystemInfo != nil && !misc.HasOSArg("i", "show-system-info", "") {
		cfg.HAProxy.ShowSystemInfo = *cfgStorage.Dataplaneapi.ShowSystemInfo
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.DisableInotify != nil && !misc.HasOSArg("", "disable-inotify", "") {
		cfg.HAProxy.DisableInotify = *cfgStorage.Dataplaneapi.DisableInotify
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.ConfigFile != nil && !misc.HasOSArg("c", "config-file", "") {
		cfg.HAProxy.ConfigFile = *cfgStorage.Haproxy.ConfigFile
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.HAProxy != nil && !misc.HasOSArg("b", "haproxy-bin", "") {
		cfg.HAProxy.HAProxy = *cfgStorage.Haproxy.HAProxy
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.MasterRuntime != nil && !misc.HasOSArg("m", "master-runtime", "") {
		cfg.HAProxy.MasterRuntime = *cfgStorage.Haproxy.MasterRuntime
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.NodeIDFile != nil && !misc.HasOSArg("", "fid", "") {
		cfg.HAProxy.NodeIDFile = *cfgStorage.Haproxy.NodeIDFile
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.MasterWorkerMode != nil && !misc.HasOSArg("", "master-worker-mode", "") {
		cfg.HAProxy.MasterWorkerMode = *cfgStorage.Haproxy.MasterWorkerMode
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Userlist != nil && cfgStorage.Dataplaneapi.Userlist.Userlist != nil && !misc.HasOSArg("u", "userlist", "") {
		cfg.HAProxy.Userlist = *cfgStorage.Dataplaneapi.Userlist.Userlist
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Userlist != nil && cfgStorage.Dataplaneapi.Userlist.UserListFile != nil && !misc.HasOSArg("", "userlist-file", "") {
		cfg.HAProxy.UserListFile = *cfgStorage.Dataplaneapi.Userlist.UserListFile
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.Reload != nil && cfgStorage.Haproxy.Reload.ReloadDelay != nil && !misc.HasOSArg("d", "reload-delay", "") {
		cfg.HAProxy.ReloadDelay = *cfgStorage.Haproxy.Reload.ReloadDelay
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.Reload != nil && cfgStorage.Haproxy.Reload.ReloadCmd != nil && !misc.HasOSArg("r", "reload-cmd", "") {
		cfg.HAProxy.ReloadCmd = *cfgStorage.Haproxy.Reload.ReloadCmd
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.Reload != nil && cfgStorage.Haproxy.Reload.RestartCmd != nil && !misc.HasOSArg("s", "restart-cmd", "") {
		cfg.HAProxy.RestartCmd = *cfgStorage.Haproxy.Reload.RestartCmd
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.Reload != nil && cfgStorage.Haproxy.Reload.ReloadRetention != nil && !misc.HasOSArg("", "reload-retention", "") {
		cfg.HAProxy.ReloadRetention = *cfgStorage.Haproxy.Reload.ReloadRetention
	}
	if cfgStorage.Haproxy != nil && cfgStorage.Haproxy.Reload != nil && cfgStorage.Haproxy.Reload.ValidateCmd != nil && !misc.HasOSArg("", "validate-cmd", "") {
		cfg.HAProxy.ValidateCmd = *cfgStorage.Haproxy.Reload.ValidateCmd
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Transaction != nil && cfgStorage.Dataplaneapi.Transaction.TransactionDir != nil && !misc.HasOSArg("t", "transaction-dir", "") {
		cfg.HAProxy.TransactionDir = *cfgStorage.Dataplaneapi.Transaction.TransactionDir
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Transaction != nil && cfgStorage.Dataplaneapi.Transaction.BackupsNumber != nil && !misc.HasOSArg("n", "backups-number", "") {
		cfg.HAProxy.BackupsNumber = *cfgStorage.Dataplaneapi.Transaction.BackupsNumber
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Transaction != nil && cfgStorage.Dataplaneapi.Transaction.BackupsDir != nil && !misc.HasOSArg("", "backups-dir", "") {
		cfg.HAProxy.BackupsDir = *cfgStorage.Dataplaneapi.Transaction.BackupsDir
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Transaction != nil && cfgStorage.Dataplaneapi.Transaction.MaxOpenTransactions != nil && !misc.HasOSArg("", "max-open-transactions", "") {
		cfg.HAProxy.MaxOpenTransactions = *cfgStorage.Dataplaneapi.Transaction.MaxOpenTransactions
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Resources != nil && cfgStorage.Dataplaneapi.Resources.MapsDir != nil && !misc.HasOSArg("p", "maps-dir", "") {
		cfg.HAProxy.MapsDir = *cfgStorage.Dataplaneapi.Resources.MapsDir
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Resources != nil && cfgStorage.Dataplaneapi.Resources.SSLCertsDir != nil && !misc.HasOSArg("", "ssl-certs-dir", "") {
		cfg.HAProxy.SSLCertsDir = *cfgStorage.Dataplaneapi.Resources.SSLCertsDir
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Resources != nil && cfgStorage.Dataplaneapi.Resources.UpdateMapFiles != nil && !misc.HasOSArg("", "update-map-files", "") {
		cfg.HAProxy.UpdateMapFiles = *cfgStorage.Dataplaneapi.Resources.UpdateMapFiles
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Resources != nil && cfgStorage.Dataplaneapi.Resources.UpdateMapFilesPeriod != nil && !misc.HasOSArg("", "update-map-files-period", "") {
		cfg.HAProxy.UpdateMapFilesPeriod = *cfgStorage.Dataplaneapi.Resources.UpdateMapFilesPeriod
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Resources != nil && cfgStorage.Dataplaneapi.Resources.SpoeDir != nil && !misc.HasOSArg("", "spoe-dir", "") {
		cfg.HAProxy.SpoeDir = *cfgStorage.Dataplaneapi.Resources.SpoeDir
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Resources != nil && cfgStorage.Dataplaneapi.Resources.SpoeTransactionDir != nil && !misc.HasOSArg("", "spoe-transaction-dir", "") {
		cfg.HAProxy.SpoeTransactionDir = *cfgStorage.Dataplaneapi.Resources.SpoeTransactionDir
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.ClusterTLSCertDir != nil && !misc.HasOSArg("", "cluster-tls-dir", "") {
		cfg.HAProxy.ClusterTLSCertDir = *cfgStorage.Cluster.ClusterTLSCertDir
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.ID != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.ID.Store(*cfgStorage.Cluster.ID)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.BootstrapKey != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.BootstrapKey.Store(*cfgStorage.Cluster.BootstrapKey)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.ActiveBootstrapKey != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.ActiveBootstrapKey.Store(*cfgStorage.Cluster.ActiveBootstrapKey)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.Token != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.Token.Store(*cfgStorage.Cluster.Token)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.URL != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.URL.Store(*cfgStorage.Cluster.URL)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.Port != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.Port.Store(*cfgStorage.Cluster.Port)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.APIBasePath != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.APIBasePath.Store(*cfgStorage.Cluster.APIBasePath)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.APINodesPath != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.APINodesPath.Store(*cfgStorage.Cluster.APINodesPath)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.APIRegisterPath != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.APIRegisterPath.Store(*cfgStorage.Cluster.APIRegisterPath)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.CertificateDir != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.CertificateDir.Store(*cfgStorage.Cluster.CertificateDir)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.CertificateFetched != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.CertificateFetched.Store(*cfgStorage.Cluster.CertificateFetched)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.Name != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.Name.Store(*cfgStorage.Cluster.Name)
	}
	if cfgStorage.Cluster != nil && cfgStorage.Cluster.Description != nil && !misc.HasOSArg("", "", "") {
		cfg.Cluster.Description.Store(*cfgStorage.Cluster.Description)
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Advertised != nil && cfgStorage.Dataplaneapi.Advertised.APIAddress != nil && !misc.HasOSArg("", "api-address", "") {
		cfg.APIOptions.APIAddress = *cfgStorage.Dataplaneapi.Advertised.APIAddress
	}
	if cfgStorage.Dataplaneapi != nil && cfgStorage.Dataplaneapi.Advertised != nil && cfgStorage.Dataplaneapi.Advertised.APIPort != nil && !misc.HasOSArg("", "api-port", "") {
		cfg.APIOptions.APIPort = *cfgStorage.Dataplaneapi.Advertised.APIPort
	}
	if cfgStorage.ServiceDiscovery != nil && cfgStorage.ServiceDiscovery.Consuls != nil && !misc.HasOSArg("", "", "") {
		cfg.ServiceDiscovery.Consuls = *cfgStorage.ServiceDiscovery.Consuls
	}
	if cfgStorage.Log != nil && cfgStorage.Log.Syslog != nil && cfgStorage.Log.Syslog.SyslogSrv != nil && !misc.HasOSArg("", "syslog-server", "") {
		cfg.Syslog.SyslogSrv = *cfgStorage.Log.Syslog.SyslogSrv
	}
	if cfgStorage.Log != nil && cfgStorage.Log.Syslog != nil && cfgStorage.Log.Syslog.SyslogPort != nil && !misc.HasOSArg("", "syslog-port", "") {
		cfg.Syslog.SyslogPort = *cfgStorage.Log.Syslog.SyslogPort
	}
	if cfgStorage.Log != nil && cfgStorage.Log.Syslog != nil && cfgStorage.Log.Syslog.SyslogProto != nil && !misc.HasOSArg("", "syslog-protocol", "") {
		cfg.Syslog.SyslogProto = *cfgStorage.Log.Syslog.SyslogProto
	}
	if cfgStorage.Log != nil && cfgStorage.Log.Syslog != nil && cfgStorage.Log.Syslog.SyslogTag != nil && !misc.HasOSArg("", "syslog-tag", "") {
		cfg.Syslog.SyslogTag = *cfgStorage.Log.Syslog.SyslogTag
	}
	if cfgStorage.Log != nil && cfgStorage.Log.Syslog != nil && cfgStorage.Log.Syslog.SyslogPriority != nil && !misc.HasOSArg("", "syslog-priority", "") {
		cfg.Syslog.SyslogPriority = *cfgStorage.Log.Syslog.SyslogPriority
	}
	if cfgStorage.Log != nil && cfgStorage.Log.Syslog != nil && cfgStorage.Log.Syslog.SyslogFacility != nil && !misc.HasOSArg("", "syslog-facility", "") {
		cfg.Syslog.SyslogFacility = *cfgStorage.Log.Syslog.SyslogFacility
	}
	if cfgStorage.Log != nil && cfgStorage.Log.LogTo != nil && !misc.HasOSArg("", "log-to", "") {
		cfg.Logging.LogTo = *cfgStorage.Log.LogTo
	}
	if cfgStorage.Log != nil && cfgStorage.Log.LogFile != nil && !misc.HasOSArg("", "log-file", "") {
		cfg.Logging.LogFile = *cfgStorage.Log.LogFile
	}
	if cfgStorage.Log != nil && cfgStorage.Log.LogLevel != nil && !misc.HasOSArg("", "log-level", "") {
		cfg.Logging.LogLevel = *cfgStorage.Log.LogLevel
	}
	if cfgStorage.Log != nil && cfgStorage.Log.LogFormat != nil && !misc.HasOSArg("", "log-format", "") {
		cfg.Logging.LogFormat = *cfgStorage.Log.LogFormat
	}
	if cfgStorage.Log != nil && cfgStorage.Log.ACLFormat != nil && !misc.HasOSArg("", "apache-common-log-format", "") {
		cfg.Logging.ACLFormat = *cfgStorage.Log.ACLFormat
	}
}

func copyConfigurationToStorage(cfg *Configuration) {
	cfgStorage := cfg.storage.Get()

	version := 2
	cfgStorage.Version = &version

	valueName := cfg.Name.Load()
	cfgStorage.Name = &valueName

	valueMode := cfg.Mode.Load()
	cfgStorage.Mode = &valueMode

	cfgStorage.DeprecatedBootstrapKey = nil

	valueStatus := cfg.Status.Load()
	cfgStorage.Status = &valueStatus

	valueClusterID := cfg.Cluster.ID.Load()
	if valueClusterID != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.ID = &valueClusterID
	}

	valueClusterBootstrapKey := cfg.Cluster.BootstrapKey.Load()
	if valueClusterBootstrapKey != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.BootstrapKey = &valueClusterBootstrapKey
	}

	valueClusterActiveBootstrapKey := cfg.Cluster.ActiveBootstrapKey.Load()
	if valueClusterActiveBootstrapKey != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.ActiveBootstrapKey = &valueClusterActiveBootstrapKey
	}

	valueClusterToken := cfg.Cluster.Token.Load()
	if valueClusterToken != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.Token = &valueClusterToken
	}

	valueClusterURL := cfg.Cluster.URL.Load()
	if valueClusterURL != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.URL = &valueClusterURL
	}

	valueClusterPort := cfg.Cluster.Port.Load()
	if valueClusterPort != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.Port = &valueClusterPort
	}

	valueClusterAPIBasePath := cfg.Cluster.APIBasePath.Load()
	if valueClusterAPIBasePath != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.APIBasePath = &valueClusterAPIBasePath
	}

	valueClusterAPINodesPath := cfg.Cluster.APINodesPath.Load()
	if valueClusterAPINodesPath != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.APINodesPath = &valueClusterAPINodesPath
	}

	valueClusterAPIRegisterPath := cfg.Cluster.APIRegisterPath.Load()
	if valueClusterAPIRegisterPath != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.APIRegisterPath = &valueClusterAPIRegisterPath
	}

	valueClusterCertificateDir := cfg.Cluster.CertificateDir.Load()
	if valueClusterCertificateDir != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.CertificateDir = &valueClusterCertificateDir
	}

	valueClusterCertificateFetched := cfg.Cluster.CertificateFetched.Load()
	if valueClusterCertificateFetched && cfgStorage.Cluster == nil {
		cfgStorage.Cluster = &configTypeCluster{}
	}
	if cfgStorage.Cluster != nil {
		cfgStorage.Cluster.CertificateFetched = &valueClusterCertificateFetched
	}

	valueClusterName := cfg.Cluster.Name.Load()
	if valueClusterName != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.Name = &valueClusterName
	}

	valueClusterDescription := cfg.Cluster.Description.Load()
	if valueClusterDescription != "" {
		if cfgStorage.Cluster == nil {
			cfgStorage.Cluster = &configTypeCluster{}
		}
		cfgStorage.Cluster.Description = &valueClusterDescription
	}

	if cfgStorage.ServiceDiscovery == nil {
		cfgStorage.ServiceDiscovery = &configTypeServiceDiscovery{}
	}
	cfgStorage.ServiceDiscovery.Consuls = &cfg.ServiceDiscovery.Consuls
}
