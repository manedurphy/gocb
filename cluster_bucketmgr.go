package gocb

import (
	"context"
	"time"
)

// BucketType specifies the kind of bucket.
type BucketType string

const (
	// CouchbaseBucketType indicates a Couchbase bucket type.
	CouchbaseBucketType BucketType = "membase"

	// MemcachedBucketType indicates a Memcached bucket type.
	MemcachedBucketType BucketType = "memcached"

	// EphemeralBucketType indicates an Ephemeral bucket type.
	EphemeralBucketType BucketType = "ephemeral"
)

// ConflictResolutionType specifies the kind of conflict resolution to use for a bucket.
type ConflictResolutionType string

const (
	// ConflictResolutionTypeTimestamp specifies to use timestamp conflict resolution on the bucket.
	ConflictResolutionTypeTimestamp ConflictResolutionType = "lww"

	// ConflictResolutionTypeSequenceNumber specifies to use sequence number conflict resolution on the bucket.
	ConflictResolutionTypeSequenceNumber ConflictResolutionType = "seqno"

	// ConflictResolutionTypeCustom specifies to use a custom bucket conflict resolution.
	// In Couchbase Server 7.1, this feature is only available in "developer-preview" mode. See the UI XDCR settings
	// for the custom conflict resolution properties.
	// VOLATILE: This API is subject to change at any time.
	ConflictResolutionTypeCustom ConflictResolutionType = "custom"
)

// EvictionPolicyType specifies the kind of eviction policy to use for a bucket.
type EvictionPolicyType string

const (
	// EvictionPolicyTypeFull specifies to use full eviction for a couchbase bucket.
	EvictionPolicyTypeFull EvictionPolicyType = "fullEviction"

	// EvictionPolicyTypeValueOnly specifies to use value only eviction for a couchbase bucket.
	EvictionPolicyTypeValueOnly EvictionPolicyType = "valueOnly"

	// EvictionPolicyTypeNotRecentlyUsed specifies to use not recently used (nru) eviction for an ephemeral bucket.
	EvictionPolicyTypeNotRecentlyUsed EvictionPolicyType = "nruEviction"

	// EvictionPolicyTypeNRU specifies to use no eviction for an ephemeral bucket.
	EvictionPolicyTypeNoEviction EvictionPolicyType = "noEviction"
)

// CompressionMode specifies the kind of compression to use for a bucket.
type CompressionMode string

const (
	// CompressionModeOff specifies to use no compression for a bucket.
	CompressionModeOff CompressionMode = "off"

	// CompressionModePassive specifies to use passive compression for a bucket.
	CompressionModePassive CompressionMode = "passive"

	// CompressionModeActive specifies to use active compression for a bucket.
	CompressionModeActive CompressionMode = "active"
)

// StorageBackend specifies the storage type to use for the bucket.
type StorageBackend string

const (
	// StorageBackendCouchstore specifies to use the couchstore storage type.
	StorageBackendCouchstore StorageBackend = "couchstore"

	// StorageBackendMagma specifies to use the magma storage type. EE only.
	StorageBackendMagma StorageBackend = "magma"
)

// BucketSettings holds information about the settings for a bucket.
type BucketSettings struct {
	Name                 string
	FlushEnabled         bool
	ReplicaIndexDisabled bool // inverted so that zero value matches server default.
	RAMQuotaMB           uint64
	NumReplicas          uint32     // NOTE: If not set this will set 0 replicas.
	BucketType           BucketType // Defaults to CouchbaseBucketType.
	EvictionPolicy       EvictionPolicyType
	// Deprecated: Use MaxExpiry instead.
	MaxTTL                 time.Duration
	MaxExpiry              time.Duration
	CompressionMode        CompressionMode
	MinimumDurabilityLevel DurabilityLevel
	StorageBackend         StorageBackend

	// # UNCOMMITTED
	//
	// Specifies whether history retention should be enabled or disabled by default on collections in the bucket.
	// This API is UNCOMMITTED and may change in the future.
	HistoryRetentionCollectionDefault *bool

	// # UNCOMMITTED
	//
	// This API is UNCOMMITTED and may change in the future.
	HistoryRetentionBytes uint64

	// # UNCOMMITTED
	//
	// This API is UNCOMMITTED and may change in the future.
	HistoryRetentionDuration time.Duration
}

// BucketManager provides methods for performing bucket management operations.
// See BucketManager for methods that allow creating and removing buckets themselves.
type BucketManager struct {
	getProvider func() (bucketManagementProvider, error)
}

// GetBucketOptions is the set of options available to the bucket manager GetBucket operation.
type GetBucketOptions struct {
	Timeout       time.Duration
	RetryStrategy RetryStrategy
	ParentSpan    RequestSpan

	// Using a deadlined Context alongside a Timeout will cause the shorter of the two to cause cancellation, this
	// also applies to global level timeouts.
	// UNCOMMITTED: This API may change in the future.
	Context context.Context
}

// GetBucket returns settings for a bucket on the cluster.
func (bm *BucketManager) GetBucket(bucketName string, opts *GetBucketOptions) (*BucketSettings, error) {
	if opts == nil {
		opts = &GetBucketOptions{}
	}

	provider, err := bm.getProvider()
	if err != nil {
		return nil, err
	}

	return provider.GetBucket(bucketName, opts)
}

// GetAllBucketsOptions is the set of options available to the bucket manager GetAll operation.
type GetAllBucketsOptions struct {
	Timeout       time.Duration
	RetryStrategy RetryStrategy
	ParentSpan    RequestSpan

	// Using a deadlined Context alongside a Timeout will cause the shorter of the two to cause cancellation, this
	// also applies to global level timeouts.
	// UNCOMMITTED: This API may change in the future.
	Context context.Context
}

// GetAllBuckets returns a list of all active buckets on the cluster.
func (bm *BucketManager) GetAllBuckets(opts *GetAllBucketsOptions) (map[string]BucketSettings, error) {
	if opts == nil {
		opts = &GetAllBucketsOptions{}
	}

	provider, err := bm.getProvider()
	if err != nil {
		return nil, err
	}

	return provider.GetAllBuckets(opts)
}

// CreateBucketSettings are the settings available when creating a bucket.
type CreateBucketSettings struct {
	BucketSettings
	ConflictResolutionType ConflictResolutionType
}

// CreateBucketOptions is the set of options available to the bucket manager CreateBucket operation.
type CreateBucketOptions struct {
	Timeout       time.Duration
	RetryStrategy RetryStrategy
	ParentSpan    RequestSpan

	// Using a deadlined Context alongside a Timeout will cause the shorter of the two to cause cancellation, this
	// also applies to global level timeouts.
	// UNCOMMITTED: This API may change in the future.
	Context context.Context
}

// CreateBucket creates a bucket on the cluster.
func (bm *BucketManager) CreateBucket(settings CreateBucketSettings, opts *CreateBucketOptions) error {
	if opts == nil {
		opts = &CreateBucketOptions{}
	}

	provider, err := bm.getProvider()
	if err != nil {
		return err
	}

	return provider.CreateBucket(settings, opts)
}

// UpdateBucketOptions is the set of options available to the bucket manager UpdateBucket operation.
type UpdateBucketOptions struct {
	Timeout       time.Duration
	RetryStrategy RetryStrategy
	ParentSpan    RequestSpan

	// Using a deadlined Context alongside a Timeout will cause the shorter of the two to cause cancellation, this
	// also applies to global level timeouts.
	// UNCOMMITTED: This API may change in the future.
	Context context.Context
}

// UpdateBucket updates a bucket on the cluster.
func (bm *BucketManager) UpdateBucket(settings BucketSettings, opts *UpdateBucketOptions) error {
	if opts == nil {
		opts = &UpdateBucketOptions{}
	}

	provider, err := bm.getProvider()
	if err != nil {
		return err
	}

	return provider.UpdateBucket(settings, opts)
}

// DropBucketOptions is the set of options available to the bucket manager DropBucket operation.
type DropBucketOptions struct {
	Timeout       time.Duration
	RetryStrategy RetryStrategy
	ParentSpan    RequestSpan

	// Using a deadlined Context alongside a Timeout will cause the shorter of the two to cause cancellation, this
	// also applies to global level timeouts.
	// UNCOMMITTED: This API may change in the future.
	Context context.Context
}

// DropBucket will delete a bucket from the cluster by name.
func (bm *BucketManager) DropBucket(name string, opts *DropBucketOptions) error {
	if opts == nil {
		opts = &DropBucketOptions{}
	}

	provider, err := bm.getProvider()
	if err != nil {
		return err
	}

	return provider.DropBucket(name, opts)
}

// FlushBucketOptions is the set of options available to the bucket manager FlushBucket operation.
type FlushBucketOptions struct {
	Timeout       time.Duration
	RetryStrategy RetryStrategy
	ParentSpan    RequestSpan

	// Using a deadlined Context alongside a Timeout will cause the shorter of the two to cause cancellation, this
	// also applies to global level timeouts.
	// UNCOMMITTED: This API may change in the future.
	Context context.Context
}

// FlushBucket will delete all the of the data from a bucket.
// Keep in mind that you must have flushing enabled in the buckets configuration.
func (bm *BucketManager) FlushBucket(name string, opts *FlushBucketOptions) error {
	if opts == nil {
		opts = &FlushBucketOptions{}
	}

	provider, err := bm.getProvider()
	if err != nil {
		return err
	}

	return provider.FlushBucket(name, opts)
}
