// auth: kunlun
// date: 2019-01-15
// description:
package yihuo

import (
	"fmt"
	"sync"
	"time"
)

type Rate struct {
	SecLimit    LimitRate
	MinuteLimit LimitRate
	HourLimit   LimitRate
	DayLimit    LimitRate
	IsInit      bool
	Limit       string
}

type Count struct {
	SuccessCount LimitRate
	FailureCount LimitRate
	TotalCount   LimitRate
	CurrentCount LimitRate
}

type LimitRate struct {
	rate  int        //速率
	begin time.Time  //开始时间
	end   int        //结束
	count int        //数量
	lock  sync.Mutex //锁
}

func (l *LimitRate) SecLimit() bool {
	result := true
	l.lock.Lock()
	now := time.Now()
	if now.Second() != l.begin.Second() {
		l.begin = now
		l.count = 0
	}
	if l.rate != 0 {
		if l.count == l.rate {
			result = false
		} else {
			l.count++
		}
	}
	fmt.Println("Second count:")
	fmt.Println(l.count)
	l.lock.Unlock()
	return result
}

func (l *LimitRate) MinLimit() bool {
	result := true
	l.lock.Lock()
	now := time.Now()
	if now.Minute() != l.begin.Minute() {
		l.begin = now
		l.count = 0
	}

	if l.rate != 0 {
		if l.count == l.rate {
			result = false
		} else {
			l.count++
		}
	}
	fmt.Println("Minute count:")
	fmt.Println(l.count)
	l.lock.Unlock()
	return result
}

func (l *LimitRate) HourLimit() bool {
	result := true
	l.lock.Lock()
	now := time.Now()
	if l.begin.Hour() != now.Hour() {
		l.begin = now
		l.count = 0
	}
	if l.rate != 0 {
		if l.rate == l.count {
			result = false
		} else {
			l.count++
		}
	}
	fmt.Println("Hour count:")
	fmt.Println(l.count)
	l.lock.Unlock()
	return result
}

func (l *LimitRate) DayLimit() bool {
	result := false
	l.lock.Lock()
	now := time.Now()
	if now.Day() != l.begin.Day() {
		l.begin = now
		l.count = 0
	}
	if l.rate != 0 {
		t := now.Hour()
		bh := l.begin.Hour()
		if bh <= t && t < l.end || bh > l.end && (t < bh && t < l.end) {
			if l.count == l.rate {
				result = false
			} else {
				l.count++
			}
		}
	}
	fmt.Println("Day count:")
	fmt.Println(l.count)

	l.lock.Unlock()
	return result
}

func (l *LimitRate) SetRate(rate int, end int, rateType string) {
	l.rate = rate
	l.end = end
	now := time.Now()
	if rateType == "day" {
		if now.Day() != l.begin.Day() {
			l.begin = now
			l.count = 0
		}
	} else {
		l.begin = now
		l.count = 0
	}
}

func (l *LimitRate) IsNeedReset() bool {
	now := time.Now()
	return l.begin.Hour()+1 == now.Hour()
}

func (l *LimitRate) GetCount() int {
	return l.count
}

func (l *LimitRate) UpdateDayCount() {
	l.lock.Lock()
	now := time.Now()
	if now.Day() != l.begin.Day() {
		l.begin = now
		l.count = 0
	}
	l.count++
	l.lock.Unlock()
}

func (l *LimitRate) UpdateCurrentCount() {
	l.lock.Lock()
	now := time.Now()
	if now.Minute() != l.begin.Minute() {
		l.begin = now
		l.count = 0
	}
	l.count++
	l.lock.Unlock()
}
