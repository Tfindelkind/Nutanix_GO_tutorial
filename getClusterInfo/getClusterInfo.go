package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type clustersGet struct {
	Metadata struct {
		GrandTotalEntities int    `json:"grandTotalEntities"`
		TotalEntities      int    `json:"totalEntities"`
		FilterCriteria     string `json:"filterCriteria"`
		SortCriteria       string `json:"sortCriteria"`
		Page               int    `json:"page"`
		Count              int    `json:"count"`
		StartIndex         int    `json:"startIndex"`
		EndIndex           int    `json:"endIndex"`
	} `json:"metadata"`
	Entities []struct {
		ID                                   string        `json:"id"`
		UUID                                 string        `json:"uuid"`
		ClusterIncarnationID                 int64         `json:"clusterIncarnationId"`
		ClusterUUID                          string        `json:"clusterUuid"`
		Name                                 string        `json:"name"`
		ClusterExternalIPAddress             string        `json:"clusterExternalIPAddress"`
		ClusterExternalDataServicesIPAddress string        `json:"clusterExternalDataServicesIPAddress"`
		Timezone                             string        `json:"timezone"`
		SupportVerbosityType                 string        `json:"supportVerbosityType"`
		NumNodes                             int           `json:"numNodes"`
		BlockSerials                         []string      `json:"blockSerials"`
		Version                              string        `json:"version"`
		FullVersion                          string        `json:"fullVersion"`
		ExternalSubnet                       string        `json:"externalSubnet"`
		InternalSubnet                       string        `json:"internalSubnet"`
		NccVersion                           string        `json:"nccVersion"`
		EnableLockDown                       bool          `json:"enableLockDown"`
		EnablePasswordRemoteLoginToCluster   bool          `json:"enablePasswordRemoteLoginToCluster"`
		FingerprintContentCachePercentage    int           `json:"fingerprintContentCachePercentage"`
		SsdPinningPercentageLimit            int           `json:"ssdPinningPercentageLimit"`
		EnableShadowClones                   bool          `json:"enableShadowClones"`
		GlobalNfsWhiteList                   []string      `json:"globalNfsWhiteList"`
		NameServers                          []string      `json:"nameServers"`
		NtpServers                           []string      `json:"ntpServers"`
		ServiceCenters                       []interface{} `json:"serviceCenters"`
		HTTPProxies                          []interface{} `json:"httpProxies"`
		RackableUnits                        []struct {
			ID               int         `json:"id"`
			RackableUnitUUID string      `json:"rackableUnitUuid"`
			Model            string      `json:"model"`
			ModelName        string      `json:"modelName"`
			Location         interface{} `json:"location"`
			Serial           string      `json:"serial"`
			Positions        []string    `json:"positions"`
			Nodes            []int       `json:"nodes"`
			NodeUuids        []string    `json:"nodeUuids"`
		} `json:"rackableUnits"`
		PublicKeys []struct {
			Name string `json:"name"`
			Key  string `json:"key"`
		} `json:"publicKeys"`
		SMTPServer             interface{} `json:"smtpServer"`
		HypervisorTypes        []string    `json:"hypervisorTypes"`
		ClusterRedundancyState struct {
			CurrentRedundancyFactor int `json:"currentRedundancyFactor"`
			DesiredRedundancyFactor int `json:"desiredRedundancyFactor"`
			RedundancyStatus        struct {
				KCassandraPrepareDone bool `json:"kCassandraPrepareDone"`
				KZookeeperPrepareDone bool `json:"kZookeeperPrepareDone"`
			} `json:"redundancyStatus"`
		} `json:"clusterRedundancyState"`
		Multicluster             bool `json:"multicluster"`
		Cloudcluster             bool `json:"cloudcluster"`
		HasSelfEncryptingDrive   bool `json:"hasSelfEncryptingDrive"`
		IsUpgradeInProgress      bool `json:"isUpgradeInProgress"`
		SecurityComplianceConfig struct {
			Schedule                   string `json:"schedule"`
			EnableAide                 bool   `json:"enableAide"`
			EnableCore                 bool   `json:"enableCore"`
			EnableHighStrengthPassword bool   `json:"enableHighStrengthPassword"`
			EnableBanner               bool   `json:"enableBanner"`
			EnableSNMPv3Only           bool   `json:"enableSNMPv3Only"`
		} `json:"securityComplianceConfig"`
		HypervisorSecurityComplianceConfig struct {
			Schedule                   string `json:"schedule"`
			EnableAide                 bool   `json:"enableAide"`
			EnableCore                 bool   `json:"enableCore"`
			EnableHighStrengthPassword bool   `json:"enableHighStrengthPassword"`
			EnableBanner               bool   `json:"enableBanner"`
		} `json:"hypervisorSecurityComplianceConfig"`
		Domain                          interface{} `json:"domain"`
		NosClusterAndHostsDomainJoined  bool        `json:"nosClusterAndHostsDomainJoined"`
		AllHypervNodesInFailoverCluster bool        `json:"allHypervNodesInFailoverCluster"`
		Credential                      interface{} `json:"credential"`
		Stats                           struct {
			HypervisorAvgIoLatencyUsecs          string `json:"hypervisor_avg_io_latency_usecs"`
			NumReadIops                          string `json:"num_read_iops"`
			HypervisorWriteIoBandwidthKBps       string `json:"hypervisor_write_io_bandwidth_kBps"`
			TimespanUsecs                        string `json:"timespan_usecs"`
			ControllerNumReadIops                string `json:"controller_num_read_iops"`
			ReadIoPpm                            string `json:"read_io_ppm"`
			ControllerNumIops                    string `json:"controller_num_iops"`
			TotalReadIoTimeUsecs                 string `json:"total_read_io_time_usecs"`
			ControllerTotalReadIoTimeUsecs       string `json:"controller_total_read_io_time_usecs"`
			ReplicationTransmittedBandwidthKBps  string `json:"replication_transmitted_bandwidth_kBps"`
			HypervisorNumIo                      string `json:"hypervisor_num_io"`
			ControllerTotalTransformedUsageBytes string `json:"controller_total_transformed_usage_bytes"`
			HypervisorCPUUsagePpm                string `json:"hypervisor_cpu_usage_ppm"`
			ControllerNumWriteIo                 string `json:"controller_num_write_io"`
			AvgReadIoLatencyUsecs                string `json:"avg_read_io_latency_usecs"`
			ContentCacheLogicalSsdUsageBytes     string `json:"content_cache_logical_ssd_usage_bytes"`
			ControllerTotalIoTimeUsecs           string `json:"controller_total_io_time_usecs"`
			ControllerTotalReadIoSizeKbytes      string `json:"controller_total_read_io_size_kbytes"`
			ControllerNumSeqIo                   string `json:"controller_num_seq_io"`
			ControllerReadIoPpm                  string `json:"controller_read_io_ppm"`
			ContentCacheNumLookups               string `json:"content_cache_num_lookups"`
			ControllerTotalIoSizeKbytes          string `json:"controller_total_io_size_kbytes"`
			ContentCacheHitPpm                   string `json:"content_cache_hit_ppm"`
			ControllerNumIo                      string `json:"controller_num_io"`
			HypervisorAvgReadIoLatencyUsecs      string `json:"hypervisor_avg_read_io_latency_usecs"`
			ContentCacheNumDedupRefCountPph      string `json:"content_cache_num_dedup_ref_count_pph"`
			NumWriteIops                         string `json:"num_write_iops"`
			ControllerNumRandomIo                string `json:"controller_num_random_io"`
			NumIops                              string `json:"num_iops"`
			ReplicationReceivedBandwidthKBps     string `json:"replication_received_bandwidth_kBps"`
			HypervisorNumReadIo                  string `json:"hypervisor_num_read_io"`
			HypervisorTotalReadIoTimeUsecs       string `json:"hypervisor_total_read_io_time_usecs"`
			ControllerAvgIoLatencyUsecs          string `json:"controller_avg_io_latency_usecs"`
			HypervisorHypervCPUUsagePpm          string `json:"hypervisor_hyperv_cpu_usage_ppm"`
			NumIo                                string `json:"num_io"`
			ControllerNumReadIo                  string `json:"controller_num_read_io"`
			HypervisorNumWriteIo                 string `json:"hypervisor_num_write_io"`
			ControllerSeqIoPpm                   string `json:"controller_seq_io_ppm"`
			ControllerReadIoBandwidthKBps        string `json:"controller_read_io_bandwidth_kBps"`
			ControllerIoBandwidthKBps            string `json:"controller_io_bandwidth_kBps"`
			HypervisorHypervMemoryUsagePpm       string `json:"hypervisor_hyperv_memory_usage_ppm"`
			HypervisorTimespanUsecs              string `json:"hypervisor_timespan_usecs"`
			HypervisorNumWriteIops               string `json:"hypervisor_num_write_iops"`
			ReplicationNumTransmittedBytes       string `json:"replication_num_transmitted_bytes"`
			TotalReadIoSizeKbytes                string `json:"total_read_io_size_kbytes"`
			HypervisorTotalIoSizeKbytes          string `json:"hypervisor_total_io_size_kbytes"`
			AvgIoLatencyUsecs                    string `json:"avg_io_latency_usecs"`
			HypervisorNumReadIops                string `json:"hypervisor_num_read_iops"`
			ContentCacheSavedSsdUsageBytes       string `json:"content_cache_saved_ssd_usage_bytes"`
			ControllerWriteIoBandwidthKBps       string `json:"controller_write_io_bandwidth_kBps"`
			ControllerWriteIoPpm                 string `json:"controller_write_io_ppm"`
			HypervisorAvgWriteIoLatencyUsecs     string `json:"hypervisor_avg_write_io_latency_usecs"`
			HypervisorTotalReadIoSizeKbytes      string `json:"hypervisor_total_read_io_size_kbytes"`
			ReadIoBandwidthKBps                  string `json:"read_io_bandwidth_kBps"`
			HypervisorEsxMemoryUsagePpm          string `json:"hypervisor_esx_memory_usage_ppm"`
			HypervisorMemoryUsagePpm             string `json:"hypervisor_memory_usage_ppm"`
			HypervisorNumIops                    string `json:"hypervisor_num_iops"`
			HypervisorIoBandwidthKBps            string `json:"hypervisor_io_bandwidth_kBps"`
			ControllerNumWriteIops               string `json:"controller_num_write_iops"`
			TotalIoTimeUsecs                     string `json:"total_io_time_usecs"`
			HypervisorKvmCPUUsagePpm             string `json:"hypervisor_kvm_cpu_usage_ppm"`
			ContentCachePhysicalSsdUsageBytes    string `json:"content_cache_physical_ssd_usage_bytes"`
			ControllerRandomIoPpm                string `json:"controller_random_io_ppm"`
			ControllerAvgReadIoSizeKbytes        string `json:"controller_avg_read_io_size_kbytes"`
			TotalTransformedUsageBytes           string `json:"total_transformed_usage_bytes"`
			AvgWriteIoLatencyUsecs               string `json:"avg_write_io_latency_usecs"`
			NumReadIo                            string `json:"num_read_io"`
			WriteIoBandwidthKBps                 string `json:"write_io_bandwidth_kBps"`
			HypervisorReadIoBandwidthKBps        string `json:"hypervisor_read_io_bandwidth_kBps"`
			RandomIoPpm                          string `json:"random_io_ppm"`
			ContentCacheNumHits                  string `json:"content_cache_num_hits"`
			TotalUntransformedUsageBytes         string `json:"total_untransformed_usage_bytes"`
			HypervisorTotalIoTimeUsecs           string `json:"hypervisor_total_io_time_usecs"`
			NumRandomIo                          string `json:"num_random_io"`
			HypervisorKvmMemoryUsagePpm          string `json:"hypervisor_kvm_memory_usage_ppm"`
			ControllerAvgWriteIoSizeKbytes       string `json:"controller_avg_write_io_size_kbytes"`
			ControllerAvgReadIoLatencyUsecs      string `json:"controller_avg_read_io_latency_usecs"`
			NumWriteIo                           string `json:"num_write_io"`
			HypervisorEsxCPUUsagePpm             string `json:"hypervisor_esx_cpu_usage_ppm"`
			TotalIoSizeKbytes                    string `json:"total_io_size_kbytes"`
			IoBandwidthKBps                      string `json:"io_bandwidth_kBps"`
			ContentCachePhysicalMemoryUsageBytes string `json:"content_cache_physical_memory_usage_bytes"`
			ReplicationNumReceivedBytes          string `json:"replication_num_received_bytes"`
			ControllerTimespanUsecs              string `json:"controller_timespan_usecs"`
			NumSeqIo                             string `json:"num_seq_io"`
			ContentCacheSavedMemoryUsageBytes    string `json:"content_cache_saved_memory_usage_bytes"`
			SeqIoPpm                             string `json:"seq_io_ppm"`
			WriteIoPpm                           string `json:"write_io_ppm"`
			ControllerAvgWriteIoLatencyUsecs     string `json:"controller_avg_write_io_latency_usecs"`
			ContentCacheLogicalMemoryUsageBytes  string `json:"content_cache_logical_memory_usage_bytes"`
		} `json:"stats"`
		UsageStats struct {
			StorageReservedFreeBytes                     string `json:"storage.reserved_free_bytes"`
			StorageTierDasSataUsageBytes                 string `json:"storage_tier.das-sata.usage_bytes"`
			DataReductionCompressionSavedBytes           string `json:"data_reduction.compression.saved_bytes"`
			DataReductionSavingRatioPpm                  string `json:"data_reduction.saving_ratio_ppm"`
			DataReductionErasureCodingPostReductionBytes string `json:"data_reduction.erasure_coding.post_reduction_bytes"`
			StorageTierSsdPinnedUsageBytes               string `json:"storage_tier.ssd.pinned_usage_bytes"`
			StorageReservedUsageBytes                    string `json:"storage.reserved_usage_bytes"`
			DataReductionErasureCodingSavingRatioPpm     string `json:"data_reduction.erasure_coding.saving_ratio_ppm"`
			StorageTierDasSataCapacityBytes              string `json:"storage_tier.das-sata.capacity_bytes"`
			StorageTierDasSataFreeBytes                  string `json:"storage_tier.das-sata.free_bytes"`
			StorageUsageBytes                            string `json:"storage.usage_bytes"`
			DataReductionErasureCodingSavedBytes         string `json:"data_reduction.erasure_coding.saved_bytes"`
			DataReductionCompressionPreReductionBytes    string `json:"data_reduction.compression.pre_reduction_bytes"`
			StorageTierDasSataPinnedUsageBytes           string `json:"storage_tier.das-sata.pinned_usage_bytes"`
			DataReductionPreReductionBytes               string `json:"data_reduction.pre_reduction_bytes"`
			StorageTierSsdCapacityBytes                  string `json:"storage_tier.ssd.capacity_bytes"`
			StorageTierSsdFreeBytes                      string `json:"storage_tier.ssd.free_bytes"`
			DataReductionDedupPreReductionBytes          string `json:"data_reduction.dedup.pre_reduction_bytes"`
			DataReductionErasureCodingPreReductionBytes  string `json:"data_reduction.erasure_coding.pre_reduction_bytes"`
			StorageCapacityBytes                         string `json:"storage.capacity_bytes"`
			DataReductionDedupPostReductionBytes         string `json:"data_reduction.dedup.post_reduction_bytes"`
			StorageLogicalUsageBytes                     string `json:"storage.logical_usage_bytes"`
			DataReductionSavedBytes                      string `json:"data_reduction.saved_bytes"`
			StorageFreeBytes                             string `json:"storage.free_bytes"`
			StorageTierSsdUsageBytes                     string `json:"storage_tier.ssd.usage_bytes"`
			DataReductionCompressionPostReductionBytes   string `json:"data_reduction.compression.post_reduction_bytes"`
			DataReductionPostReductionBytes              string `json:"data_reduction.post_reduction_bytes"`
			DataReductionDedupSavedBytes                 string `json:"data_reduction.dedup.saved_bytes"`
			DataReductionCompressionSavingRatioPpm       string `json:"data_reduction.compression.saving_ratio_ppm"`
			DataReductionDedupSavingRatioPpm             string `json:"data_reduction.dedup.saving_ratio_ppm"`
			StorageTierSsdPinnedBytes                    string `json:"storage_tier.ssd.pinned_bytes"`
			StorageReservedCapacityBytes                 string `json:"storage.reserved_capacity_bytes"`
		} `json:"usageStats"`
		EnforceRackableUnitAwarePlacement bool `json:"enforceRackableUnitAwarePlacement"`
		DisableDegradedNodeMonitoring     bool `json:"disableDegradedNodeMonitoring"`
	} `json:"entities"`
}

// EncodeCredentials this func is encoding the Username and Password with base64 encoding which is
// required for Nutanix
func EncodeCredentials(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

// v0_8 returns the main entry point for the v0.8 Nutanix API
func v0_8(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/api/nutanix/v0.8/"

}

// v1_0 returns the main entry point for the v1.0 Nutanix API
func v1_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v1/"

}

func main() {

	// PRISM user
	var username = "admin"
	// PRISM user password
	var password = "nutanix/4u"
	// Nutanix Cluster IP/DNSName CVM IP/DNSName
	var NutanixHost = "192.168.178.130"

	// Ignores certificates which can not be validated
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// create a HTTP client
	var httpClient = http.Client{Transport: tr}

	// create a http Request pointer
	var req *http.Request
	var err error

	// Defines the HTTP Request
	// send a GET to the NUTANIX API and receives the cluster info
	// https://192.168.178.130:9440/PrismGateway/services/rest/v1/clusters
	req, err = http.NewRequest("GET", v1_0(NutanixHost)+"/clusters", nil)

	// before the request is send set the HTTP Header key "Authorization" with
	// the value of base64 encoded Username and Password
	req.Header.Set("Authorization", "Basic "+EncodeCredentials(username, password))

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read the data from the resp.body into bodyText
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// create the struct
	var clustersGetResp clustersGet

	// Parse/Unmarshal JSON into the struct
	json.Unmarshal(bodyText, &clustersGetResp)

	// Print the parsed data to stdout
	println("ID: " + clustersGetResp.Entities[0].ID)
	println("Name: " + clustersGetResp.Entities[0].Name)

}
