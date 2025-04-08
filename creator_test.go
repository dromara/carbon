package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateFromStdTime(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromStdTime(time.Now(), "").HasError())
		assert.Empty(t, CreateFromStdTime(time.Now(), "").ToString())
		assert.True(t, CreateFromStdTime(time.Now(), "xxx").HasError())
		assert.Empty(t, CreateFromStdTime(time.Now(), "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		now := time.Now()
		assert.Equal(t, now.Unix(), CreateFromStdTime(now).Timestamp())
	})

	t.Run("with timezone", func(t *testing.T) {
		now := time.Now().In(time.Local)
		assert.Equal(t, now.Unix(), CreateFromStdTime(now).Timestamp())
		assert.Equal(t, now.Unix(), CreateFromStdTime(now, UTC).Timestamp())
	})
}

func TestCreateFromTimestamp(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimestamp(0, "").HasError())
		assert.Empty(t, CreateFromTimestamp(0, "").ToString())
		assert.True(t, CreateFromTimestamp(0, "xxx").HasError())
		assert.Empty(t, CreateFromTimestamp(0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "1969-12-31 23:59:59 +0000 UTC", CreateFromTimestamp(-1).ToString())
		assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestamp(0).ToString())
		assert.Equal(t, "1970-01-01 00:00:01 +0000 UTC", CreateFromTimestamp(1).ToString())
		assert.Equal(t, "2022-04-12 03:55:55 +0000 UTC", CreateFromTimestamp(1649735755).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "1970-01-01 07:59:59 +0800 CST", CreateFromTimestamp(-1, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestamp(0, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:01 +0800 CST", CreateFromTimestamp(1, PRC).ToString())
		assert.Equal(t, "2022-04-12 11:55:55 +0800 CST", CreateFromTimestamp(1649735755, PRC).ToString())
	})
}

func TestCreateFromTimestampMilli(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimestampMilli(0, "").HasError())
		assert.True(t, CreateFromTimestampMilli(0, "").HasError())
		assert.Empty(t, CreateFromTimestampMilli(0, "xxx").ToString())
		assert.Empty(t, CreateFromTimestampMilli(0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "1969-12-31 23:59:59.999 +0000 UTC", CreateFromTimestampMilli(-1).ToString())
		assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampMilli(0).ToString())
		assert.Equal(t, "1970-01-01 00:00:00.001 +0000 UTC", CreateFromTimestampMilli(1).ToString())
		assert.Equal(t, "2022-04-12 03:55:55.981 +0000 UTC", CreateFromTimestampMilli(1649735755981).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "1970-01-01 07:59:59.999 +0800 CST", CreateFromTimestampMilli(-1, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestampMilli(0, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00.001 +0800 CST", CreateFromTimestampMilli(1, PRC).ToString())
		assert.Equal(t, "2022-04-12 11:55:55.981 +0800 CST", CreateFromTimestampMilli(1649735755981, PRC).ToString())
	})
}

func TestCreateFromTimestampMicro(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimestampMicro(0, "").HasError())
		assert.True(t, CreateFromTimestampMicro(0, "").HasError())
		assert.Empty(t, CreateFromTimestampMicro(0, "xxx").ToString())
		assert.Empty(t, CreateFromTimestampMicro(0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "1969-12-31 23:59:59.999999 +0000 UTC", CreateFromTimestampMicro(-1).ToString())
		assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampMicro(0).ToString())
		assert.Equal(t, "1970-01-01 00:00:00.000001 +0000 UTC", CreateFromTimestampMicro(1).ToString())
		assert.Equal(t, "2022-04-12 03:55:55.981566 +0000 UTC", CreateFromTimestampMicro(1649735755981566).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "1970-01-01 07:59:59.999999 +0800 CST", CreateFromTimestampMicro(-1, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestampMicro(0, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00.000001 +0800 CST", CreateFromTimestampMicro(1, PRC).ToString())
		assert.Equal(t, "2022-04-12 11:55:55.981566 +0800 CST", CreateFromTimestampMicro(1649735755981566, PRC).ToString())
	})
}

func TestCreateFromTimestampNano(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimestampNano(0, "").HasError())
		assert.True(t, CreateFromTimestampNano(0, "").HasError())
		assert.Empty(t, CreateFromTimestampNano(0, "xxx").ToString())
		assert.Empty(t, CreateFromTimestampNano(0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "1969-12-31 23:59:59.999999999 +0000 UTC", CreateFromTimestampNano(-1).ToString())
		assert.Equal(t, "1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampNano(0).ToString())
		assert.Equal(t, "1970-01-01 00:00:00.000000001 +0000 UTC", CreateFromTimestampNano(1).ToString())
		assert.Equal(t, "2022-04-12 03:55:55.981566888 +0000 UTC", CreateFromTimestampNano(1649735755981566888).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "1970-01-01 07:59:59.999999999 +0800 CST", CreateFromTimestampNano(-1, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00 +0800 CST", CreateFromTimestampNano(0, PRC).ToString())
		assert.Equal(t, "1970-01-01 08:00:00.000000001 +0800 CST", CreateFromTimestampNano(1, PRC).ToString())
		assert.Equal(t, "2022-04-12 11:55:55.981566888 +0800 CST", CreateFromTimestampNano(1649735755981566888, PRC).ToString())
	})
}

func TestCreateFromDateTime(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateTime(0, 0, 0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateTime(0, 0, 0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateTime(0, 0, 0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateTime(0, 0, 0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTime(0, 0, 0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTime(0, 0, 0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", CreateFromDateTime(2020, 8, 5, 13, 14, 15, PRC).ToString())
	})
}

func TestCreateFromDateTimeMilli(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 13:14:15.999 +0000 UTC", CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15.999 +0800 CST", CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999, PRC).ToString())
	})
}

func TestCreateFromDateTimeMicro(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0000 UTC", CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15.999999 +0800 CST", CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999, PRC).ToString())
	})
}

func TestCreateFromDateTimeNano(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0000 UTC", CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15.999999999 +0800 CST", CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999, PRC).ToString())
	})
}

func TestCreateFromDate(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDate(0, 0, 0, "").HasError())
		assert.True(t, CreateFromDate(0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDate(0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDate(0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDate(0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", CreateFromDate(2020, 8, 5).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDate(0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", CreateFromDate(2020, 8, 5, PRC).ToString())
	})
}

func TestCreateFromDateMilli(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateMilli(0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateMilli(0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateMilli(0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateMilli(0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateMilli(0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 00:00:00.999 +0000 UTC", CreateFromDateMilli(2020, 8, 5, 999).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateMilli(0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00.999 +0800 CST", CreateFromDateMilli(2020, 8, 5, 999, PRC).ToString())
	})
}

func TestCreateFromDateMicro(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateMicro(0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateMicro(0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateMicro(0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateMicro(0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateMicro(0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0000 UTC", CreateFromDateMicro(2020, 8, 5, 999999).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateMicro(0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00.999999 +0800 CST", CreateFromDateMicro(2020, 8, 5, 999999, PRC).ToString())
	})
}

func TestCreateFromDateNano(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromDateNano(0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromDateNano(0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromDateNano(0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromDateNano(0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0000 UTC", CreateFromDateNano(0, 0, 0, 0).ToString())
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0000 UTC", CreateFromDateNano(2020, 8, 5, 999999999).ToString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "-0001-11-30 00:00:00 +0805 LMT", CreateFromDateNano(0, 0, 0, 0, PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00.999999999 +0800 CST", CreateFromDateNano(2020, 8, 5, 999999999, PRC).ToString())
	})
}

func TestCreateFromTime(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTime(0, 0, 0, "").HasError())
		assert.True(t, CreateFromTime(0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromTime(0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromTime(0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTime(0, 0, 0).ToTimeString())
		assert.Equal(t, "13:14:15", CreateFromTime(13, 14, 15).ToTimeString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTime(0, 0, 0, PRC).ToTimeString())
		assert.Equal(t, "13:14:15", CreateFromTime(13, 14, 15, PRC).ToTimeString())
	})
}

func TestCreateFromTimeMilli(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimeMilli(0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromTimeMilli(0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromTimeMilli(0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromTimeMilli(0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTimeMilli(0, 0, 0, 0).ToTimeMilliString())
		assert.Equal(t, "13:14:15.999", CreateFromTimeMilli(13, 14, 15, 999).ToTimeMilliString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTimeMilli(0, 0, 0, 0, PRC).ToTimeMilliString())
		assert.Equal(t, "13:14:15.999", CreateFromTimeMilli(13, 14, 15, 999, PRC).ToTimeMilliString())
	})
}

func TestCreateFromTimeMicro(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimeMicro(0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromTimeMicro(0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromTimeMicro(0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromTimeMicro(0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTimeMicro(0, 0, 0, 0).ToTimeMicroString())
		assert.Equal(t, "13:14:15.999999", CreateFromTimeMicro(13, 14, 15, 999999).ToTimeMicroString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTimeMicro(0, 0, 0, 0, PRC).ToTimeMicroString())
		assert.Equal(t, "13:14:15.999999", CreateFromTimeMicro(13, 14, 15, 999999, PRC).ToTimeMicroString())
	})
}

func TestCreateFromTimeNano(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, CreateFromTimeNano(0, 0, 0, 0, "").HasError())
		assert.True(t, CreateFromTimeNano(0, 0, 0, 0, "").HasError())
		assert.Empty(t, CreateFromTimeNano(0, 0, 0, 0, "xxx").ToString())
		assert.Empty(t, CreateFromTimeNano(0, 0, 0, 0, "xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTimeNano(0, 0, 0, 0).ToTimeNanoString())
		assert.Equal(t, "13:14:15.999999999", CreateFromTimeNano(13, 14, 15, 999999999).ToTimeNanoString())
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, "00:00:00", CreateFromTimeNano(0, 0, 0, 0, PRC).ToTimeNanoString())
		assert.Equal(t, "13:14:15.999999999", CreateFromTimeNano(13, 14, 15, 999999999, PRC).ToTimeNanoString())
	})
}
