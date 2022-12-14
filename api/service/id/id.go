// id summon
// 111111111111111111111111111111111111111111 111111111111
// 54                                      13 12         0
//                           timestamp(42bit)  step(12bit)

package id

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const Epoch int64 = 1669824000000 // 2022-12-01 00:00:00
const (
	TimeShift = 12
	MaxStep   = 1 << TimeShift
	IDBit     = 54
)

var GlobalIDMake *SummonID

func init() {
	GlobalIDMake = NewSummonID()
}

type SummonID struct {
	mu    sync.Mutex
	time  int64
	step  int64
	Epoch time.Time
}

func NewSummonID() *SummonID {
	return &SummonID{
		Epoch: time.Unix(Epoch/1e3, (Epoch%1e3)*1e6),
		mu:    sync.Mutex{},
		step:  0,
	}
}

func (s *SummonID) Generate() ID {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Since(s.Epoch).Milliseconds()
	if now == s.time {
		s.step = (s.step + 1) &^ MaxStep
		if s.step == 0 {
			for now <= s.time {
				now = time.Since(s.Epoch).Milliseconds()
			}
		}
	} else {
		s.step = 0
	}
	s.time = now

	return ID((now << TimeShift) | s.step)
}

// ID is a 54-bit ID.
type ID int64

func (f ID) Int64() int64   { return int64(f) }
func (f ID) String() string { return strconv.FormatInt(int64(f), 10) }
func (f ID) Base2() string  { return fmt.Sprintf("%0*v", IDBit, strconv.FormatInt(int64(f), 2)) }
func (f ID) Bytes() []byte  { return []byte(f.String()) }
func (f ID) Base64() string { return base64.StdEncoding.EncodeToString(f.Bytes()) }

func (f ID) Time() time.Time {
	t := (int64(f) >> TimeShift) + Epoch
	return time.Unix(t/1e3, (t%1e3)*1e6)
}

// string to ID, if not return -1
func StringToID(s string) ID {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return ID(i)
}
