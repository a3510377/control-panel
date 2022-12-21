package system

import (
	"fmt"
	"math"
)

type ByteSize float64

const maxByte = 1 << 10

const (
	Byte ByteSize = 1 << (10 * iota)
	KByte
	MByte
	GByte
	TByte
	PByte
	EByte
)

func stringBytesSize(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}

	e := math.Floor(math.Log(float64(s)) / math.Log(base))
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	format := "%.0f %s"

	if val < 10 {
		format = "%.1f %s"
	}

	return fmt.Sprintf(format, val, sizes[int(e)])
}

func BytesString(s uint64) string {
	return stringBytesSize(s, maxByte, []string{
		"B", "kB", "MB",
		"GB", "TB", "PB", "EB",
	})
}
