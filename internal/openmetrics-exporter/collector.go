package collectors

import (
	"context"
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/collectors"
)

func Collector(ctx context.Context, metrics string, registry *prometheus.Registry, fbclient *client.FBClient) bool {
	filesystems := fbclient.GetFileSystems()
	buckets := fbclient.GetBuckets()
	arrayCollector := NewArraysCollector(fbclient)
	registry.MustRegister(
		// collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		// collectors.NewGoCollector(),
		arrayCollector,
	)
	if metrics == "all" || metrics == "array" {
		perfCollector := NewPerfCollector(fbclient)
		s3perfCollector := NewS3PerfCollector(fbclient)
		httpPerfCollector := NewHttpPerfCollector(fbclient)
		nfsPerfCollector := NewNfsPerfCollector(fbclient)
		perfReplCollector := NewPerfReplicationCollector(fbclient)
		arraySpaceCollector := NewArraySpaceCollector(fbclient)
		alertsCollector := NewAlertsCollector(fbclient)
		hardwareCollector := NewHardwareCollector(fbclient)
		hwPerfConnectorsCollector := NewHwConnectorsPerfCollector(fbclient)
		objstoreacctsCollector := NewObjectStoreAccountsCollector(fbclient)
		registry.MustRegister(
			perfCollector,
			s3perfCollector,
			httpPerfCollector,
			nfsPerfCollector,
			perfReplCollector,
			arraySpaceCollector,
			alertsCollector,
			hardwareCollector,
			hwPerfConnectorsCollector,
			objstoreacctsCollector,
		)
	}
	if metrics == "all" || metrics == "filesystems" {
		filesystemsPerfCollector := NewFileSystemsPerfCollector(fbclient, filesystems)
		filesystemsSpaceCollector := NewFileSystemsSpaceCollector(filesystems)
		registry.MustRegister(
			filesystemsSpaceCollector,
			filesystemsPerfCollector,
		)
	}
	if metrics == "all" || metrics == "buckets" {
		bucketsPerfCollector := NewBucketsPerfCollector(fbclient, buckets)
		buckestS3PerfCollector := NewBucketsS3PerfCollector(fbclient, buckets)
		bucketsSpaceCollector := NewBucketsSpaceCollector(buckets)
		registry.MustRegister(
			bucketsPerfCollector,
			buckestS3PerfCollector,
			bucketsSpaceCollector,
		)
	}
	if metrics == "all" || metrics == "clients" {
		clientsPerfCollector := NewClientsPerfCollector(fbclient)
		registry.MustRegister(clientsPerfCollector)
	}
	if metrics == "all" || metrics == "usage" {
		usageCollector := NewUsageCollector(fbclient, filesystems)
		registry.MustRegister(usageCollector)
	}
	if metrics == "all" || metrics == "policies" {
		policiesCollector := NewNfsPoliciesCollector(fbclient)
		registry.MustRegister(policiesCollector)
	}
	return true
}
