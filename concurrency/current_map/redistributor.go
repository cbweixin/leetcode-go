package current_map

// BucketStatus 代表散列桶状态的类型。
type BucketStatus uint8

const (
	// BUCKET_STATUS_NORMAL 代表散列桶正常。
	BUCKET_STATUS_NORMAL BucketStatus = 0
	// BUCKET_STATUS_UNDERWEIGHT 代表散列桶过轻。
	BUCKET_STATUS_UNDERWEIGHT BucketStatus = 1
	// BUCKET_STATUS_OVERWEIGHT 代表散列桶过重。
	BUCKET_STATUS_OVERWEIGHT BucketStatus = 2
)

// PairRedistributor 代表针对键-元素对的再分布器。
// 用于当散列段内的键-元素对分布不均时进行重新分布。

type PairRedistrubited interface {
	//UpdateThreshold 会根据键-元素对总数和散列桶总数计算并更新阈值。
	UpdateThreshold(pairTotal uint64, bucketNumber int) //
	//CheckBucketStatus 用于检查散列桶的状态
	CheckBucketStatus(pairTotal uint64, bucketSize uint64) (bucketStatus BucketStatus)
	//Redistribe 用于实施键-元素对的再分布。
	Redistribe(bucketStatus BucketStatus, buckets []Bucket) (newBukets []Bucket, changed bool)
}

type myPairRedistributor struct {
	// loadFactor 代表装载因子。
	loadFactor float64
	// upperThreshold 代表散列桶重量的上阈限。
	// 当某个散列桶的尺寸增至此值时会触发再散列。
	upperThreshold uint64
	// overweightBucketCount 代表过重的散列桶的计数
	overweightBucketCount uint64
	// emptyBucketCount 代表空的散列桶的计数。
	emptyBucketCount uint64
}
