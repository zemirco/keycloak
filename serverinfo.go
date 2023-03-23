package keycloak

import (
	"context"
	"net/http"
)

type ServerInfo struct {
	SystemInfo struct {
		Version        string    `json:"version"`
		SystemTime     string    `json:"serverTime"` // TODO: parse as time.Time
		UptimeMillis   int64     `json:"uptimeMillis"`
		JavaVersion    string    `json:"javaVersion"`
		JavaVendor     string    `json:"javaVendor"`
		JavaVM         string    `json:"javaVM"`
		JavaVMVersion  string    `json:"javaVMVersion"`
		JavaRuntime    string    `json:"javaRuntime"`
		JavaHome       string    `json:"javaHome"`
		OSName         string    `json:"osName"`
		OSArchitecture string    `json:"osArchitecture"`
		OSVersion      string    `json:"osVersion"`
		FileEncoding   string    `json:"fileEncoding"`
		UserName       string    `json:"userName"`
		UserDir        string    `json:"userDir"`
		UserTimezone   string    `json:"userTimezone"`
		UserLocale     string    `json:"userLocale"`
	} `json:"systemInfo"`
	MemoryInfo struct {
		Total          int64  `json:"total"`
		TotalFormatted string `json:"totalFormatted"`
		Used           int64  `json:"used"`
		UsedFormatted  int64  `json:"usedFormatted"`
		Free           int64  `json:"free"`
		FreeFormatted  int64  `json:"freeFormatted"`
		FreePercentage int    `json:"freePercentage"`
	} `json:"memoryInfo"`
	ProfileInfo struct {
		Name                 string   `json:"name"`
		DisabledFeatures     []string `json:"disabledFeatures"`
		PreviewFeatures      []string `json:"previewFeatures"`
		ExperimentalFeatures []string `json:"experimentalFeatures"`
	} `json:"profileInfo"`
	CryptoInfo struct {
		CryptoProvider         string   `json:"cryptoProvider"`
		SupportedKeystoreTypes []string `json:"supportedKeystoreTypes"`
	} `json:"cryptoInfo"`

	// Additional Theme, Locale, and Provider info omitted for now
}

func (k *Keycloak) GetServerInfo() (*ServerInfo, error) {
	if req, err := k.NewRequest(http.MethodGet, "/admin/serverinfo", nil); err == nil {
		serverInfo := &ServerInfo{}
		if _, e := k.Do(context.Background(), req, serverInfo); e != nil {
			return nil, e
		} else {
			return serverInfo, nil
		}
	} else {
		return nil, err
	}
}
