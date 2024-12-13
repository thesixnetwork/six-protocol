package types

import (
	"strconv"
	"time"
)

func (m *Metadata) getBlockHeight() int64 {
	return m.GetBlockHeightFunction()
}

func (m *Metadata) getBlockTimestamp() time.Time {
	return m.GetBlockTimeFunction()
}

func (m *Metadata) GetBlockHeight() string {
	return strconv.FormatInt(m.getBlockHeight(), 10)
}

// 2022-12-15T07:00:57Z
func (m *Metadata) GetUTCBlockTimestamp(format string) string {
	return m.getBlockTimestamp().UTC().Format(format)
}

// TZ name is from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
func (m *Metadata) GetBlockTimestampByZone(zone string, format string) string {
	loc, _ := time.LoadLocation(zone)
	blockTIme := m.getBlockTimestamp().In(loc)
	return blockTIme.Format(format)
}

func (m *Metadata) BlockTimeUTCBefore(t string, format string) bool {
	_t, _ := time.Parse(format, t) // to UTC(time.RFC3339, time)
	return m.getBlockTimestamp().Before(_t)
}

func (m *Metadata) BlockTimeUTCAfter(t string, format string) bool {
	_t, _ := time.Parse(format, t) // to UTC(time.RFC3339, time)
	return m.getBlockTimestamp().After(_t)
}

func (m *Metadata) BlockTimeBeforeByZone(t string, zone string, format string) bool {
	loc, _ := time.LoadLocation(zone)
	_t, _ := time.ParseInLocation(format, t, loc)
	return m.getBlockTimestamp().Before(_t)
}

func (m *Metadata) BlockTimeAfterByZone(t string, zone string, format string) bool {
	loc, _ := time.LoadLocation(zone)
	_t, _ := time.ParseInLocation(format, t, loc)
	return m.getBlockTimestamp().After(_t)
}
