// Copyright (C) 2017 ScyllaDB

package rclone

import (
	"fmt"
	"strconv"

	"github.com/scylladb/go-set/strset"
)

//go:generate ./options_gen.sh

const (
	// In order to reduce memory footprint, by default we allow at most two
	// concurrent requests.
	// upload_concurrency * chunk_size gives rough estimate how much upload
	// buffers will be allocated.
	defaultUploadConcurrency = 2

	// Default value of 5MB caused that we encountered problems with S3
	// returning 5xx. In order to reduce number of requests to S3, we are
	// increasing chunk size by ten times, which should decrease number of
	// requests by ten times.
	defaultChunkSize = "50M"
)

var s3Providers = strset.New(
	"AWS", "Minio", "Alibaba", "Ceph", "DigitalOcean",
	"IBMCOS", "Wasabi", "Dreamhost", "Netease", "Other",
)

// DefaultS3Options returns a S3Options initialized with default values.
func DefaultS3Options() S3Options {
	return S3Options{
		Provider:        "AWS",
		ChunkSize:       defaultChunkSize,
		DisableChecksum: "true",
		EnvAuth:         "true",
		// Because of access denied issues with Minio.
		// see https://github.com/rclone/rclone/issues/4633
		NoCheckBucket:     "true",
		UploadConcurrency: strconv.Itoa(defaultUploadConcurrency),
	}
}

// Validate returns error if option values are not set properly.
func (o *S3Options) Validate() error {
	if o.Endpoint != "" && o.Provider == "" {
		return fmt.Errorf("specify provider for the endpoint %s, available providers are: %s", o.Endpoint, s3Providers)
	}

	if o.Provider != "" && !s3Providers.Has(o.Provider) {
		return fmt.Errorf("unknown provider: %s", o.Provider)
	}

	return nil
}

// AutoFill sets region (if empty) from identity service, it only works when
// running in AWS.
func (o *S3Options) AutoFill() {
	if o.Region == "" && o.Endpoint == "" {
		o.Region = awsRegionFromMetadataAPI()
	}
}