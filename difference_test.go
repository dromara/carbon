package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DiffInYears(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInYears())
		assert.Zero(t, now.DiffInYears(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInYears(now))
		assert.Zero(t, now.DiffInYears(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInYears())
		assert.Equal(t, int64(0), Parse("2021-01-01 13:14:15").DiffInYears())
		assert.Equal(t, int64(-1), Parse("2021-08-28 13:14:59").DiffInYears())
		assert.Equal(t, int64(1), Parse("2018-08-28 13:14:59").DiffInYears())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-12-31 13:14:15").DiffInYears(Parse("2021-01-01 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2021-08-28 13:14:59")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffAbsInYears(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInYears())
		assert.Zero(t, now.DiffAbsInYears(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInYears())
		assert.Zero(t, now.DiffAbsInYears(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInYears())
		assert.Equal(t, int64(0), Parse("2021-01-01 13:14:15").DiffAbsInYears())
		assert.Equal(t, int64(1), Parse("2021-08-28 13:14:59").DiffAbsInYears())
		assert.Equal(t, int64(1), Parse("2018-08-28 13:14:59").DiffAbsInYears())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-12-31 13:14:15").DiffAbsInYears(Parse("2021-01-01 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2021-08-28 13:14:59")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffInMonths(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInMonths())
		assert.Zero(t, now.DiffInMonths(Parse("")))
		assert.Zero(t, Parse("").DiffInMonths(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInMonths())
		assert.Zero(t, now.DiffInMonths(Parse("xxx")))
		assert.Zero(t, Parse("xxx").DiffInMonths(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMonths())
		assert.Equal(t, int64(0), Parse("2020-07-28 13:14:00").DiffInMonths())
		assert.Equal(t, int64(-1), Parse("2020-09-06 13:14:59").DiffInMonths())
		assert.Equal(t, int64(23), Parse("2018-08-28 13:14:59").DiffInMonths())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-09-06 13:14:59")))
		assert.Equal(t, int64(-23), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffAbsInMonths(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInMonths())
		assert.Zero(t, now.DiffAbsInMonths(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInMonths())
		assert.Zero(t, now.DiffAbsInMonths(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths())
		assert.Equal(t, int64(0), Parse("2020-07-28 13:14:00").DiffAbsInMonths())
		assert.Equal(t, int64(1), Parse("2020-09-06 13:14:59").DiffAbsInMonths())
		assert.Equal(t, int64(23), Parse("2018-08-28 13:14:59").DiffAbsInMonths())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-09-06 13:14:59")))
		assert.Equal(t, int64(23), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2018-08-28 13:14:59")))
	})
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInWeeks())
		assert.Zero(t, now.DiffInWeeks(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInWeeks())
		assert.Zero(t, now.DiffInWeeks(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInWeeks())
		assert.Equal(t, int64(1), Parse("2020-07-28 13:14:00").DiffInWeeks())
		assert.Equal(t, int64(-1), Parse("2020-08-12 13:14:15").DiffInWeeks())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-08-12 13:14:15")))
	})
}

func TestCarbon_DiffAbsInWeeks(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInWeeks())
		assert.Zero(t, now.DiffAbsInWeeks(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInWeeks())
		assert.Zero(t, now.DiffAbsInWeeks(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInWeeks())
		assert.Equal(t, int64(1), Parse("2020-07-28 13:14:00").DiffAbsInWeeks())
		assert.Equal(t, int64(1), Parse("2020-08-12 13:14:15").DiffAbsInWeeks())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-07-28 13:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-08-12 13:14:15")))
	})
}

func TestCarbon_DiffInDays(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInDays())
		assert.Zero(t, now.DiffInDays(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInDays())
		assert.Zero(t, now.DiffInDays(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInDays())
		assert.Equal(t, int64(0), Parse("2020-08-04 13:14:59").DiffInDays())
		assert.Equal(t, int64(-1), Parse("2020-08-06 13:14:15").DiffInDays())
		assert.Equal(t, int64(1), Parse("2020-08-04 13:00:00").DiffInDays())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-04 13:14:59")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-06 13:14:15")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-04 13:00:00")))
	})
}

func TestCarbon_DiffAbsInDays(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInDays())
		assert.Zero(t, now.DiffAbsInDays(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInDays())
		assert.Zero(t, now.DiffAbsInDays(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays())
		assert.Equal(t, int64(0), Parse("2020-08-04 13:14:59").DiffAbsInDays())
		assert.Equal(t, int64(1), Parse("2020-08-06 13:14:15").DiffAbsInDays())
		assert.Equal(t, int64(1), Parse("2020-08-04 13:00:00").DiffAbsInDays())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-04 13:14:59")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-06 13:14:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-04 13:00:00")))
	})
}

func TestCarbon_DiffInHours(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInHours())
		assert.Zero(t, now.DiffInHours(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInHours())
		assert.Zero(t, now.DiffInHours(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInHours())
		assert.Equal(t, int64(0), Parse("2020-08-05 14:13:15").DiffInHours())
		assert.Equal(t, int64(1), Parse("2020-08-05 12:14:00").DiffInHours())
		assert.Equal(t, int64(-1), Parse("2020-08-05 14:14:15").DiffInHours())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 14:13:15")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 12:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 14:14:15")))
	})
}

func TestCarbon_DiffAbsInHours(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInHours())
		assert.Zero(t, now.DiffAbsInHours(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInHours())
		assert.Zero(t, now.DiffAbsInHours(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours())
		assert.Equal(t, int64(0), Parse("2020-08-05 14:13:15").DiffAbsInHours())
		assert.Equal(t, int64(1), Parse("2020-08-05 12:14:00").DiffAbsInHours())
		assert.Equal(t, int64(1), Parse("2020-08-05 14:14:15").DiffAbsInHours())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 14:13:15")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 12:14:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 14:14:15")))
	})
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInMinutes())
		assert.Zero(t, now.DiffInMinutes(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInMinutes())
		assert.Zero(t, now.DiffInMinutes(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes())
		assert.Equal(t, int64(0), Parse("2020-08-05 13:15:10").DiffInMinutes())
		assert.Equal(t, int64(1), Parse("2020-08-05 13:13:00").DiffInMinutes())
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:15:15").DiffInMinutes())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:15:10")))
		assert.Equal(t, int64(-1), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:13:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:15:15")))
	})
}

func TestCarbon_DiffAbsInMinutes(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInMinutes())
		assert.Zero(t, now.DiffAbsInMinutes(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInMinutes())
		assert.Zero(t, now.DiffAbsInMinutes(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes())
		assert.Equal(t, int64(0), Parse("2020-08-05 13:15:10").DiffAbsInMinutes())
		assert.Equal(t, int64(1), Parse("2020-08-05 13:13:00").DiffAbsInMinutes())
		assert.Equal(t, int64(1), Parse("2020-08-05 13:15:15").DiffAbsInMinutes())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:15:10")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:13:00")))
		assert.Equal(t, int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:15:15")))
	})
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffInSeconds())
		assert.Zero(t, now.DiffInSeconds(Parse("")))
		assert.Zero(t, Parse("xxx").DiffInSeconds())
		assert.Zero(t, now.DiffInSeconds(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds())
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15.999999").DiffInSeconds())
		assert.Equal(t, int64(-5), Parse("2020-08-05 13:14:20").DiffInSeconds())
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:10").DiffInSeconds())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:15.999999")))
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:20")))
		assert.Equal(t, int64(-5), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:10")))
	})
}

func TestCarbon_DiffAbsInSeconds(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, Parse("").DiffAbsInSeconds())
		assert.Zero(t, now.DiffAbsInSeconds(Parse("")))
		assert.Zero(t, Parse("xxx").DiffAbsInSeconds())
		assert.Zero(t, now.DiffAbsInSeconds(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds())
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15.999999").DiffAbsInSeconds())
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:20").DiffAbsInSeconds())
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:10").DiffAbsInSeconds())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:15.999999")))
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:20")))
		assert.Equal(t, int64(5), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:10")))
	})
}

func TestCarbon_DiffInString(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffInString())
		assert.Empty(t, now.DiffInString(Parse("")))
		assert.Empty(t, Parse("xxx").DiffInString())
		assert.Empty(t, now.DiffInString(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, "just now", now.DiffInString())
		assert.Equal(t, "-1 year", now.AddYearsNoOverflow(1).DiffInString())
		assert.Equal(t, "1 year", now.SubYearsNoOverflow(1).DiffInString())
		assert.Equal(t, "-1 month", now.AddMonthsNoOverflow(1).DiffInString())
		assert.Equal(t, "1 month", now.SubMonthsNoOverflow(1).DiffInString())
		assert.Equal(t, "-1 week", now.AddWeeks(1).DiffInString())
		assert.Equal(t, "1 week", now.SubWeeks(1).DiffInString())
		assert.Equal(t, "-1 day", now.AddDays(1).DiffInString())
		assert.Equal(t, "1 day", now.SubDays(1).DiffInString())
		assert.Equal(t, "-1 hour", now.AddHours(1).DiffInString())
		assert.Equal(t, "1 hour", now.SubHours(1).DiffInString())
		assert.Equal(t, "-1 minute", now.AddMinutes(1).DiffInString())
		assert.Equal(t, "1 minute", now.SubMinutes(1).DiffInString())
		assert.Equal(t, "-1 second", now.AddSeconds(1).DiffInString())
		assert.Equal(t, "1 second", now.SubSeconds(1).DiffInString())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, "just now", now.DiffInString(now))
		assert.Equal(t, "-1 year", now.AddYearsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "1 year", now.SubYearsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "-1 month", now.AddMonthsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "1 month", now.SubMonthsNoOverflow(1).DiffInString(now))
		assert.Equal(t, "-1 week", now.AddWeeks(1).DiffInString(now))
		assert.Equal(t, "1 week", now.SubWeeks(1).DiffInString(now))
		assert.Equal(t, "-1 day", now.AddDays(1).DiffInString(now))
		assert.Equal(t, "1 day", now.SubDays(1).DiffInString(now))
		assert.Equal(t, "-1 hour", now.AddHours(1).DiffInString(now))
		assert.Equal(t, "1 hour", now.SubHours(1).DiffInString(now))
		assert.Equal(t, "-1 minute", now.AddMinutes(1).DiffInString(now))
		assert.Equal(t, "1 minute", now.SubMinutes(1).DiffInString(now))
		assert.Equal(t, "-1 second", now.AddSeconds(1).DiffInString(now))
		assert.Equal(t, "1 second", now.SubSeconds(1).DiffInString(now))
	})
}

func TestCarbon_DiffAbsInString(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffAbsInString())
		assert.Empty(t, now.DiffAbsInString(Parse("")))
		assert.Empty(t, Parse("xxx").DiffAbsInString())
		assert.Empty(t, now.DiffAbsInString(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, "just now", now.DiffAbsInString())
		assert.Equal(t, "1 year", now.AddYearsNoOverflow(1).DiffAbsInString())
		assert.Equal(t, "1 year", now.SubYearsNoOverflow(1).DiffAbsInString())
		assert.Equal(t, "1 month", now.AddMonthsNoOverflow(1).DiffAbsInString())
		assert.Equal(t, "1 month", now.SubMonthsNoOverflow(1).DiffAbsInString())
		assert.Equal(t, "1 day", now.AddDays(1).DiffAbsInString())
		assert.Equal(t, "1 day", now.SubDays(1).DiffAbsInString())
		assert.Equal(t, "1 hour", now.AddHours(1).DiffAbsInString())
		assert.Equal(t, "1 hour", now.SubHours(1).DiffAbsInString())
		assert.Equal(t, "1 minute", now.AddMinutes(1).DiffAbsInString())
		assert.Equal(t, "1 minute", now.SubMinutes(1).DiffAbsInString())
		assert.Equal(t, "1 second", now.AddSeconds(1).DiffAbsInString())
		assert.Equal(t, "1 second", now.SubSeconds(1).DiffAbsInString())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, "just now", Parse("2020-08-05 13:14:15").DiffAbsInString(now))
		assert.Equal(t, "1 year", now.AddYearsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 year", now.SubYearsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 month", now.AddMonthsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 month", now.SubMonthsNoOverflow(1).DiffAbsInString(now))
		assert.Equal(t, "1 day", now.AddDays(1).DiffAbsInString(now))
		assert.Equal(t, "1 day", now.SubDays(1).DiffAbsInString(now))
		assert.Equal(t, "1 hour", now.AddHours(1).DiffAbsInString(now))
		assert.Equal(t, "1 hour", now.SubHours(1).DiffAbsInString(now))
		assert.Equal(t, "1 minute", now.AddMinutes(1).DiffAbsInString(now))
		assert.Equal(t, "1 minute", now.SubMinutes(1).DiffAbsInString(now))
		assert.Equal(t, "1 second", now.AddSeconds(1).DiffAbsInString(now))
		assert.Equal(t, "1 second", now.SubSeconds(1).DiffAbsInString(now))
	})
}

func TestCarbon_DiffInDuration(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffInDuration())
		assert.Empty(t, now.DiffInDuration(Parse("")))
		assert.Empty(t, Parse("xxx").DiffInDuration())
		assert.Empty(t, now.DiffInDuration(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, "0s", now.DiffInDuration().String())
		assert.Equal(t, "-8760h0m0s", now.AddYearsNoOverflow(1).DiffInDuration().String())
		assert.Equal(t, "8784h0m0s", now.SubYearsNoOverflow(1).DiffInDuration().String())
		assert.Equal(t, "-744h0m0s", now.AddMonthsNoOverflow(1).DiffInDuration().String())
		assert.Equal(t, "744h0m0s", now.SubMonthsNoOverflow(1).DiffInDuration().String())
		assert.Equal(t, "-24h0m0s", now.AddDays(1).DiffInDuration().String())
		assert.Equal(t, "24h0m0s", now.SubDays(1).DiffInDuration().String())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, "0s", now.DiffInDuration(now).String())
		assert.Equal(t, "-8760h0m0s", now.AddYearsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "8784h0m0s", now.SubYearsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "-744h0m0s", now.AddMonthsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "744h0m0s", now.SubMonthsNoOverflow(1).DiffInDuration(now).String())
		assert.Equal(t, "-24h0m0s", now.AddDays(1).DiffInDuration(now).String())
		assert.Equal(t, "24h0m0s", now.SubDays(1).DiffInDuration(now).String())
	})
}

func TestCarbon_DiffAbsInDuration(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffAbsInDuration())
		assert.Empty(t, now.DiffAbsInDuration(Parse("")))
		assert.Empty(t, Parse("xxx").DiffAbsInDuration())
		assert.Empty(t, now.DiffAbsInDuration(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, "0s", now.DiffAbsInDuration().String())
		assert.Equal(t, "8760h0m0s", now.AddYearsNoOverflow(1).DiffAbsInDuration().String())
		assert.Equal(t, "8784h0m0s", now.SubYearsNoOverflow(1).DiffAbsInDuration().String())
		assert.Equal(t, "744h0m0s", now.AddMonthsNoOverflow(1).DiffAbsInDuration().String())
		assert.Equal(t, "744h0m0s", now.SubMonthsNoOverflow(1).DiffAbsInDuration().String())
		assert.Equal(t, "24h0m0s", now.AddDays(1).DiffAbsInDuration().String())
		assert.Equal(t, "24h0m0s", now.SubDays(1).DiffAbsInDuration().String())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, "0s", now.DiffAbsInDuration(now).String())
		assert.Equal(t, "8760h0m0s", now.AddYearsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "8784h0m0s", now.SubYearsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "744h0m0s", now.AddMonthsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "744h0m0s", now.SubMonthsNoOverflow(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "24h0m0s", now.AddDays(1).DiffAbsInDuration(now).String())
		assert.Equal(t, "24h0m0s", now.SubDays(1).DiffAbsInDuration(now).String())
	})
}

func TestCarbon_DiffForHumans(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Empty(t, Parse("").DiffForHumans())
		assert.Empty(t, now.DiffForHumans(Parse("")))
		assert.Empty(t, Parse("xxx").DiffForHumans())
		assert.Empty(t, now.DiffForHumans(Parse("xxx")))
	})

	t.Run("without parameter", func(t *testing.T) {
		assert.Equal(t, "just now", Parse("2020-08-05 13:14:15").DiffForHumans())
		assert.Equal(t, "2 days ago", Parse("2020-08-03 13:14:15").DiffForHumans())
		assert.Equal(t, "2 days from now", Parse("2020-08-07 13:14:15").DiffForHumans())
		assert.Equal(t, "1 year from now", now.AddYearsNoOverflow(1).DiffForHumans())
		assert.Equal(t, "1 year ago", now.SubYearsNoOverflow(1).DiffForHumans())
		assert.Equal(t, "1 month from now", now.AddMonthsNoOverflow(1).DiffForHumans())
		assert.Equal(t, "1 month ago", now.SubMonthsNoOverflow(1).DiffForHumans())
		assert.Equal(t, "1 day from now", now.AddDays(1).DiffForHumans())
		assert.Equal(t, "1 day ago", now.SubDays(1).DiffForHumans())
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, "just now", Parse("2020-08-05 13:14:15").DiffForHumans(now))
		assert.Equal(t, "2 days before", Parse("2020-08-03 13:14:15").DiffForHumans(now))
		assert.Equal(t, "2 days after", Parse("2020-08-07 13:14:15").DiffForHumans(now))
		assert.Equal(t, "1 year after", now.AddYearsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 year before", now.SubYearsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 month after", now.AddMonthsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 month before", now.SubMonthsNoOverflow(1).DiffForHumans(now))
		assert.Equal(t, "1 day after", now.AddDays(1).DiffForHumans(now))
		assert.Equal(t, "1 day before", now.SubDays(1).DiffForHumans(now))
	})
}

func TestCarbon_getDiffInMonths(t *testing.T) {
	defer ClearTestNow()
	now := Parse("2020-08-05 13:14:15")
	SetTestNow(now)

	t.Run("invalid carbon", func(t *testing.T) {
		assert.Zero(t, getDiffInMonths(Parse(""), Parse("")))
		assert.Zero(t, getDiffInMonths(Parse(""), Parse("xxx")))
		assert.Zero(t, getDiffInMonths(Parse("xxx"), Parse("xxx")))
		assert.Zero(t, getDiffInMonths(Parse("xxx"), Parse("")))
	})

	t.Run("with parameter", func(t *testing.T) {
		assert.Equal(t, int64(0), getDiffInMonths(Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(1), getDiffInMonths(Parse("2020-07-05 13:14:15"), Parse("2020-08-05 13:14:15")))
		assert.Equal(t, int64(-1), getDiffInMonths(Parse("2020-09-05 13:14:15"), Parse("2020-08-05 13:14:15")))
	})
}

// https://github.com/dromara/carbon/issues/255
func TestCarbon_Issue255(t *testing.T) {
	assert.Equal(t, int64(0), Parse("2024-10-11").DiffInMonths(Parse("2024-11-10")))
	assert.Equal(t, int64(0), Parse("2024-11-10").DiffInMonths(Parse("2024-10-11")))
	assert.Equal(t, int64(1), Parse("2024-10-11").DiffInMonths(Parse("2024-11-11")))
	assert.Equal(t, int64(-1), Parse("2024-11-11").DiffInMonths(Parse("2024-10-11")))
	assert.Equal(t, int64(0), Parse("2024-10-11 23:59:00").DiffInMonths(Parse("2024-11-11 00:00:00")))
	assert.Equal(t, int64(-23), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2018-08-28 13:14:59")))
	assert.Equal(t, int64(23), Parse("2018-08-28 13:14:59").DiffInMonths(Parse("2020-08-05 13:14:15")))
	assert.Equal(t, int64(11999), Parse("1024-12-25 13:14:20").DiffInMonths(Parse("2024-12-25 13:14:19")))
}
