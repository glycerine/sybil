//+build tdigest

package sybil

import "gopkg.in/caio/go-tdigest.v1"

var ENABLE_TDIGEST = true

func newTDigestHist(table *Table, info *IntInfo) Histogram {
	return table.NewTDigestHist(info)
}

// {{{ TDigest HIST
type TDigestHist struct {
	*tdigest.TDigest

	table *Table
	Info  *IntInfo

	Count int64
}

func (hc *TDigestHist) Min() int64 {

	return int64(hc.TDigest.Quantile(0))
}

func (hc *TDigestHist) Max() int64 {
	return int64(hc.TDigest.Quantile(1.0))
}

func (hc *TDigestHist) NewHist() Histogram {
	return hc.table.NewHist(hc.Info)
}

func (h *TDigestHist) Mean() float64 {
	return h.TDigest.Quantile(0.5)
}

func (h *TDigestHist) RecordValues(value int64, n int64) error {
	h.Count += n
	return h.TDigest.Add(float64(value), uint32(n))
}

func (t *Table) NewTDigestHist(info *IntInfo) *TDigestHist {
	hdr_hist := tdigest.New(1)
	outer_hist := TDigestHist{hdr_hist, t, info, 0}

	return &outer_hist

}

func (th *TDigestHist) Combine(oh interface{}) {
	hist := oh.(*TDigestHist)
	th.TDigest.Merge(hist.TDigest)
}

func (th *TDigestHist) GetVariance() float64 {
	std := th.StdDev()
	return std * std
}

func (th *TDigestHist) StdDev() float64 {
	return 0
}

func (th *TDigestHist) GetPercentiles() []int64 {

	ret := make([]int64, 100)
	for i := 0; i < 100; i++ {
		ret[i] = int64(th.TDigest.Quantile(float64(i) / 100.0))
	}

	return ret
}

func (th *TDigestHist) GetBuckets() map[string]int64 {
	ret := make(map[string]int64)
	// TODO: implement this!

	return ret

}

func (th *TDigestHist) TotalCount() int64 {
	return th.Count
}

// }}} TDigest HIST
