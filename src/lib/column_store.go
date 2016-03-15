package sybil

// Before we save the new record list in a table, we tend to sort by time
type RecordList []*Record
type SortRecordsByTime struct {
	RecordList
}

func (a RecordList) Len() int      { return len(a) }
func (a RecordList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortRecordsByTime) Less(i, j int) bool {
	t1 := a.RecordList[i].Timestamp
	t2 := a.RecordList[j].Timestamp

	return t1 < t2
}

type SavedIntBucket struct {
	Value   int64
	Records []uint32
}

type SavedSetBucket struct {
	Value   int32
	Records []uint32
}

type SavedStrBucket struct {
	Value   int32
	Records []uint32
}

type SavedColumnInfo struct {
	NumRecords int32

	StrInfoMap SavedStrInfo
	IntInfoMap SavedIntInfo
}

type SavedIntColumn struct {
	Name            string
	DeltaEncodedIDs bool
	ValueEncoded    bool
	BucketEncoded   bool
	Bins            []SavedIntBucket
	Values          []int64
}

type SavedStrColumn struct {
	Name            string
	DeltaEncodedIDs bool
	BucketEncoded   bool
	Bins            []SavedStrBucket
	Values          []int32
	StringTable     []string
}

type SavedSetColumn struct {
	Name            string
	Bins            []SavedSetBucket
	Values          [][]int32
	StringTable     []string
	DeltaEncodedIDs bool
	BucketEncoded   bool
}
