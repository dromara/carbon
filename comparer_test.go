package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_HasError(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().HasError())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").HasError())
		assert.True(t, Parse("0").HasError())
		assert.True(t, Parse("xxx").HasError())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Now().HasError())
	})
}

func TestCarbon_IsEmpty(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsEmpty())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("0").IsEmpty())
		assert.False(t, Parse("xxx").IsEmpty())
		assert.True(t, Parse("").IsEmpty())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05").IsEmpty())
	})
}

func TestCarbon_IsZero(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		stdTime1 := time.Date(0001, 1, 1, 00, 00, 00, 00, time.UTC)
		carbon1 := CreateFromDateTimeNano(0001, 1, 1, 00, 00, 00, 00, UTC)
		assert.True(t, carbon1.IsZero())
		assert.Equal(t, stdTime1.IsZero(), carbon1.IsZero())

		stdTime2 := time.Time{}
		carbon2 := NewCarbon()
		assert.True(t, carbon2.IsZero())
		assert.Equal(t, stdTime2.IsZero(), carbon2.IsZero())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsZero())
		assert.False(t, Parse("0").IsZero())
		assert.False(t, Parse("xxx").IsZero())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("0000-00-00 00:00:00").IsZero())
		assert.False(t, Parse("2020-08-05").IsZero())
		assert.False(t, Parse("0000-00-00").IsZero())
	})
}

func TestCarbon_IsEpoch(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		carbon1 := CreateFromDateTimeNano(0001, 1, 1, 00, 00, 00, 00, UTC)
		assert.False(t, carbon1.IsEpoch())

		carbon2 := NewCarbon()
		assert.False(t, carbon2.IsEpoch())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsEpoch())
		assert.False(t, Parse("0").IsEpoch())
		assert.False(t, Parse("xxx").IsEpoch())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, CreateFromDateTimeNano(1970, 1, 1, 0, 0, 0, 0, UTC).IsEpoch())
		assert.True(t, CreateFromTimestamp(0).IsEpoch())
		assert.False(t, Parse("2020-08-05").IsEpoch())
		assert.False(t, Parse("0000-00-00").IsEpoch())
	})
}

func TestCarbon_IsValid(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsValid())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsValid())
		assert.False(t, Parse("0").IsValid())
		assert.False(t, Parse("xxx").IsValid())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-08").IsValid())
	})
}

func TestCarbon_IsInvalid(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsInvalid())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.True(t, Parse("").IsInvalid())
		assert.True(t, Parse("0").IsInvalid())
		assert.True(t, Parse("xxx").IsInvalid())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-08").IsInvalid())
	})
}

func TestCarbon_IsDST(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsDST())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsDST())
		assert.False(t, Parse("0").IsDST())
		assert.False(t, Parse("xxx").IsDST())
	})

	t.Run("valid time", func(t *testing.T) {
		tzWithDST, tzWithoutDST := "Australia/Sydney", "Australia/Brisbane"
		assert.True(t, Parse("2009-01-01", tzWithDST).IsDST())
		assert.False(t, Parse("2009-01-01", tzWithoutDST).IsDST())
	})
}

func TestCarbon_IsAM(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsAM())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsAM())
		assert.False(t, Parse("0").IsAM())
		assert.False(t, Parse("xxx").IsAM())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-08 00:00:00").IsAM())
		assert.False(t, Parse("2020-08-08 12:00:00").IsAM())
		assert.False(t, Parse("2020-08-08 23:59:59").IsAM())
	})
}

func TestCarbon_IsPM(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsPM())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsPM())
		assert.False(t, Parse("0").IsPM())
		assert.False(t, Parse("xxx").IsPM())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-08 00:00:00").IsPM())
		assert.True(t, Parse("2020-08-08 12:00:00").IsPM())
		assert.True(t, Parse("2020-08-08 23:59:59").IsPM())
	})
}

func TestCarbon_IsLeapYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsLeapYear())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsLeapYear())
		assert.False(t, Parse("0").IsLeapYear())
		assert.False(t, Parse("xxx").IsLeapYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2015-01-01").IsLeapYear())
		assert.True(t, Parse("2016-01-01").IsLeapYear())
	})
}

func TestCarbon_IsLongYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsLongYear())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsLongYear())
		assert.False(t, Parse("0").IsLongYear())
		assert.False(t, Parse("xxx").IsLongYear())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2015-01-01").IsLongYear())
		assert.False(t, Parse("2016-01-01").IsLongYear())
	})
}

func TestCarbon_IsJanuary(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsJanuary())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsJanuary())
		assert.False(t, Parse("0").IsJanuary())
		assert.False(t, Parse("xxx").IsJanuary())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-01-01").IsJanuary())
		assert.False(t, Parse("2020-02-01").IsJanuary())
	})
}

func TestCarbon_IsFebruary(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsFebruary())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsFebruary())
		assert.False(t, Parse("0").IsFebruary())
		assert.False(t, Parse("xxx").IsFebruary())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-02-01").IsFebruary())
		assert.False(t, Parse("2020-03-01").IsFebruary())
	})
}

func TestCarbon_IsMarch(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsMarch())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsMarch())
		assert.False(t, Parse("0").IsMarch())
		assert.False(t, Parse("xxx").IsMarch())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-03-01").IsMarch())
		assert.False(t, Parse("2020-04-01").IsMarch())
	})
}

func TestCarbon_IsApril(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsApril())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsApril())
		assert.False(t, Parse("0").IsApril())
		assert.False(t, Parse("xxx").IsApril())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-04-01").IsApril())
		assert.False(t, Parse("2020-05-01").IsApril())
	})
}

func TestCarbon_IsMay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsMay())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsMay())
		assert.False(t, Parse("0").IsMay())
		assert.False(t, Parse("xxx").IsMay())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-05-01").IsMay())
		assert.False(t, Parse("2020-06-01").IsMay())
	})
}

func TestCarbon_IsJune(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsJune())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsJune())
		assert.False(t, Parse("0").IsJune())
		assert.False(t, Parse("xxx").IsJune())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-06-01").IsJune())
		assert.False(t, Parse("2020-07-01").IsJune())
	})
}

func TestCarbon_IsJuly(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsJuly())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsJuly())
		assert.False(t, Parse("0").IsJuly())
		assert.False(t, Parse("xxx").IsJuly())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-07-01").IsJuly())
		assert.False(t, Parse("2020-08-01").IsJuly())
	})
}

func TestCarbon_IsAugust(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsAugust())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsAugust())
		assert.False(t, Parse("0").IsAugust())
		assert.False(t, Parse("xxx").IsAugust())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-01").IsAugust())
		assert.False(t, Parse("2020-09-01").IsAugust())
	})
}

func TestCarbon_IsSeptember(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsSeptember())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsSeptember())
		assert.False(t, Parse("0").IsSeptember())
		assert.False(t, Parse("xxx").IsSeptember())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-09-01").IsSeptember())
		assert.False(t, Parse("2020-10-01").IsSeptember())
	})
}

func TestCarbon_IsOctober(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsOctober())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsOctober())
		assert.False(t, Parse("0").IsOctober())
		assert.False(t, Parse("xxx").IsOctober())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-10-01").IsOctober())
		assert.False(t, Parse("2020-11-01").IsOctober())
	})
}

func TestCarbon_IsNovember(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsNovember())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsNovember())
		assert.False(t, Parse("0").IsNovember())
		assert.False(t, Parse("xxx").IsNovember())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-11-01").IsNovember())
		assert.False(t, Parse("2020-12-01").IsNovember())
	})
}

func TestCarbon_IsDecember(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsDecember())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsDecember())
		assert.False(t, Parse("0").IsDecember())
		assert.False(t, Parse("xxx").IsDecember())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-12-01").IsDecember())
		assert.False(t, Parse("2020-01-01").IsDecember())
	})
}

func TestCarbon_IsMonday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsMonday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsMonday())
		assert.False(t, Parse("0").IsMonday())
		assert.False(t, Parse("xxx").IsMonday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-03").IsMonday())
		assert.False(t, Parse("2025-03-04").IsMonday())
	})
}

func TestCarbon_IsTuesday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsTuesday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsTuesday())
		assert.False(t, Parse("0").IsTuesday())
		assert.False(t, Parse("xxx").IsTuesday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-04").IsTuesday())
		assert.False(t, Parse("2025-03-05").IsTuesday())
	})
}

func TestCarbon_IsWednesday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsWednesday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsWednesday())
		assert.False(t, Parse("0").IsWednesday())
		assert.False(t, Parse("xxx").IsWednesday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-05").IsWednesday())
		assert.False(t, Parse("2025-03-06").IsWednesday())
	})
}

func TestCarbon_IsThursday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsThursday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsThursday())
		assert.False(t, Parse("0").IsThursday())
		assert.False(t, Parse("xxx").IsThursday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-06").IsThursday())
		assert.False(t, Parse("2025-03-07").IsThursday())
	})
}

func TestCarbon_IsFriday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsFriday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsFriday())
		assert.False(t, Parse("0").IsFriday())
		assert.False(t, Parse("xxx").IsFriday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-07").IsFriday())
		assert.False(t, Parse("2025-03-08").IsFriday())
	})
}

func TestCarbon_IsSaturday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsSaturday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsSaturday())
		assert.False(t, Parse("0").IsSaturday())
		assert.False(t, Parse("xxx").IsSaturday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-08").IsSaturday())
		assert.False(t, Parse("2025-03-09").IsSaturday())
	})
}

func TestCarbon_IsSunday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsSunday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsSunday())
		assert.False(t, Parse("0").IsSunday())
		assert.False(t, Parse("xxx").IsSunday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-09").IsSunday())
		assert.False(t, Parse("2025-03-10").IsSunday())
	})
}

func TestCarbon_IsWeekday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsWeekday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsWeekday())
		assert.False(t, Parse("0").IsWeekday())
		assert.False(t, Parse("xxx").IsWeekday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2025-03-01").IsWeekday())
		assert.False(t, Parse("2025-03-02").IsWeekday())
		assert.True(t, Parse("2025-03-03").IsWeekday())
	})
}

func TestCarbon_IsWeekend(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsWeekend())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsWeekend())
		assert.False(t, Parse("0").IsWeekend())
		assert.False(t, Parse("xxx").IsWeekend())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2025-03-01").IsWeekend())
		assert.True(t, Parse("2025-03-02").IsWeekend())
		assert.False(t, Parse("2025-03-03").IsWeekend())
	})
}

func TestCarbon_IsNow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsNow())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsNow())
		assert.False(t, Parse("0").IsNow())
		assert.False(t, Parse("xxx").IsNow())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Yesterday().IsNow())
		assert.True(t, Now().IsNow())
		assert.False(t, Tomorrow().IsNow())
		assert.False(t, Parse("2025-03-01").IsNow())
	})
}

func TestCarbon_IsFuture(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsFuture())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsFuture())
		assert.False(t, Parse("0").IsFuture())
		assert.False(t, Parse("xxx").IsFuture())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Yesterday().IsFuture())
		assert.False(t, Now().IsFuture())
		assert.True(t, Tomorrow().IsFuture())
		assert.False(t, Parse("2025-03-01").IsFuture())
	})
}

func TestCarbon_IsPast(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsPast())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsPast())
		assert.False(t, Parse("0").IsPast())
		assert.False(t, Parse("xxx").IsPast())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Yesterday().IsPast())
		assert.False(t, Now().IsPast())
		assert.False(t, Tomorrow().IsPast())
		assert.True(t, Parse("2025-03-01").IsPast())
	})
}

func TestCarbon_IsYesterday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsYesterday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsYesterday())
		assert.False(t, Parse("0").IsYesterday())
		assert.False(t, Parse("xxx").IsYesterday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Yesterday().IsYesterday())
		assert.False(t, Now().IsYesterday())
		assert.False(t, Tomorrow().IsYesterday())
		assert.False(t, Parse("2025-03-01").IsYesterday())
	})
}

func TestCarbon_IsToday(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsToday())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsToday())
		assert.False(t, Parse("0").IsToday())
		assert.False(t, Parse("xxx").IsToday())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Yesterday().IsToday())
		assert.True(t, Now().IsToday())
		assert.False(t, Tomorrow().IsToday())
		assert.False(t, Parse("2025-03-01").IsToday())
	})
}

func TestCarbon_IsTomorrow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().IsTomorrow())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("").IsTomorrow())
		assert.False(t, Parse("0").IsTomorrow())
		assert.False(t, Parse("xxx").IsTomorrow())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Yesterday().IsTomorrow())
		assert.False(t, Now().IsTomorrow())
		assert.True(t, Tomorrow().IsTomorrow())
		assert.False(t, Parse("2025-03-01").IsTomorrow())
	})
}

func TestCarbon_IsSameCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameCentury(NewCarbon()))
		assert.False(t, Now().IsSameCentury(NewCarbon()))
		assert.False(t, NewCarbon().IsSameCentury(Now()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameCentury(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameCentury(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameCentury(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05").IsSameCentury(Parse("2010-01-01")))
		assert.True(t, Parse("2020-08-05").IsSameCentury(Parse("2030-12-31")))
		assert.False(t, Parse("2020-08-05").IsSameCentury(Parse("2100-08-05")))
	})
}

func TestCarbon_IsSameDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameDecade(NewCarbon()))
		assert.False(t, Now().IsSameDecade(NewCarbon()))
		assert.False(t, NewCarbon().IsSameDecade(Now()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameDecade(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameDecade(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameDecade(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05").IsSameDecade(Parse("2020-01-01")))
		assert.True(t, Parse("2020-08-05").IsSameDecade(Parse("2020-12-31")))
		assert.False(t, Parse("2020-08-05").IsSameDecade(Parse("2010-08-05")))
	})
}

func TestCarbon_IsSameYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameYear(NewCarbon()))
		assert.False(t, Now().IsSameYear(NewCarbon()))
		assert.False(t, NewCarbon().IsSameYear(Now()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameYear(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameYear(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameYear(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05").IsSameYear(Parse("2010-08-05")))
		assert.True(t, Parse("2020-08-05").IsSameYear(Parse("2020-01-01")))
		assert.True(t, Parse("2020-08-05").IsSameYear(Parse("2020-12-31")))
	})
}

func TestCarbon_IsSameQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameQuarter(NewCarbon()))
		assert.False(t, Parse("2020-08-05").IsSameQuarter(NewCarbon()))
		assert.False(t, NewCarbon().IsSameQuarter(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameQuarter(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameQuarter(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameQuarter(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05").IsSameQuarter(Parse("2020-08-06")))
		assert.False(t, Parse("2020-08-05").IsSameQuarter(Parse("2010-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameQuarter(Parse("2010-01-05")))
	})
}

func TestCarbon_IsSameMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameMonth(NewCarbon()))
		assert.False(t, Parse("2020-08-05").IsSameMonth(NewCarbon()))
		assert.False(t, NewCarbon().IsSameMonth(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameMonth(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameMonth(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameMonth(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05").IsSameMonth(Parse("2020-08-01")))
		assert.False(t, Parse("2020-08-05").IsSameMonth(Parse("2021-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameMonth(Parse("2020-09-05")))
	})
}

func TestCarbon_IsSameDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameDay(NewCarbon()))
		assert.False(t, Parse("2020-08-05").IsSameDay(NewCarbon()))
		assert.False(t, NewCarbon().IsSameDay(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameDay(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameDay(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameDay(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 00:00:00").IsSameDay(Parse("2020-08-05 23:59:59")))
		assert.False(t, Parse("2020-08-05 00:00:00").IsSameDay(Parse("2021-08-05 00:00:00")))
		assert.False(t, Parse("2020-08-05 00:00:00").IsSameDay(Parse("2020-09-05 00:00:00")))
	})
}

func TestCarbon_IsSameHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameHour(NewCarbon()))
		assert.False(t, Parse("2020-08-05").IsSameHour(NewCarbon()))
		assert.False(t, NewCarbon().IsSameHour(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameHour(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameHour(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameHour(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-08-05 22:59:59")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameHour(Parse("2021-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-09-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-08-06 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_IsSameMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameMinute(NewCarbon()))
		assert.False(t, Parse("2020-08-05").IsSameMinute(NewCarbon()))
		assert.False(t, NewCarbon().IsSameMinute(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameMinute(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameMinute(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameMinute(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-05 22:00:59")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2021-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-06 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-05 23:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-05 22:01:00")))
	})
}

func TestCarbon_IsSameSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().IsSameSecond(NewCarbon()))
		assert.False(t, Parse("2020-08-05").IsSameSecond(NewCarbon()))
		assert.False(t, NewCarbon().IsSameSecond(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").IsSameSecond(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").IsSameSecond(Parse("xxx")))
		assert.False(t, Parse("xxx").IsSameSecond(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 22:00:00.999999")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2021-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-09-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-06 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 23:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 22:01:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 22:00:01")))
	})
}

func TestCarbon_Compare(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().Compare("=", NewCarbon()))
		assert.False(t, Parse("2020-08-05").Compare("=", NewCarbon()))
		assert.False(t, NewCarbon().Compare("=", Parse("2020-08-05")))
	})

	t.Run("invalid operator", func(t *testing.T) {
		assert.False(t, Now().Compare("", Yesterday()))
		assert.False(t, Now().Compare("%", Yesterday()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Compare("<", Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Compare(">", Parse("xxx")))
		assert.False(t, Parse("xxx").Compare("!=", Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").Compare("=", Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Compare(">=", Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Compare("<=", Parse("2020-08-05 22:00:00")))

		assert.True(t, Parse("2020-08-05 22:00:00").Compare(">", Parse("2020-07-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Compare(">=", Parse("2020-07-05 22:00:00")))

		assert.True(t, Parse("2020-08-05 22:00:00").Compare("<", Parse("2020-09-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Compare("<=", Parse("2020-09-05 22:00:00")))

		assert.True(t, Parse("2020-08-05 22:00:00").Compare("<>", Parse("2020-06-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Compare("!=", Parse("2020-06-05 22:00:00")))
	})
}

func TestCarbon_Gt(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().Gt(NewCarbon()))
		assert.True(t, Parse("2020-08-05").Gt(NewCarbon()))
		assert.False(t, NewCarbon().Gt(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Gt(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Gt(Parse("xxx")))
		assert.False(t, Parse("xxx").Gt(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").Gt(Parse("2020-08-05 21:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Gt(Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Gt(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_Lt(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().Lt(NewCarbon()))
		assert.False(t, Parse("2020-08-05").Lt(NewCarbon()))
		assert.True(t, NewCarbon().Lt(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Lt(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Lt(Parse("xxx")))
		assert.False(t, Parse("xxx").Lt(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05 22:00:00").Lt(Parse("2020-08-05 21:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Lt(Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Lt(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_Eq(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().Eq(NewCarbon()))
		assert.False(t, Parse("2020-08-05").Eq(NewCarbon()))
		assert.False(t, NewCarbon().Eq(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Eq(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Eq(Parse("xxx")))
		assert.False(t, Parse("xxx").Eq(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05 22:00:00").Eq(Parse("2020-08-05 21:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Eq(Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Eq(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_Ne(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().Ne(NewCarbon()))
		assert.True(t, Parse("2020-08-05").Ne(NewCarbon()))
		assert.True(t, NewCarbon().Ne(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Ne(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Ne(Parse("xxx")))
		assert.False(t, Parse("xxx").Ne(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").Ne(Parse("2020-08-05 21:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Ne(Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Ne(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_Gte(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().Gte(NewCarbon()))
		assert.True(t, Parse("2020-08-05").Gte(NewCarbon()))
		assert.False(t, NewCarbon().Gte(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Gte(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Gte(Parse("xxx")))
		assert.False(t, Parse("xxx").Gte(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").Gte(Parse("2020-08-05 21:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Gte(Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Gte(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_Lte(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().Lte(NewCarbon()))
		assert.False(t, Parse("2020-08-05").Lte(NewCarbon()))
		assert.True(t, NewCarbon().Lte(Parse("2020-08-05")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Lte(Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Lte(Parse("xxx")))
		assert.False(t, Parse("xxx").Lte(Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05 22:00:00").Lte(Parse("2020-08-05 21:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Lte(Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").Lte(Parse("2020-08-05 23:00:00")))
	})
}

func TestCarbon_Between(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().Between(NewCarbon(), NewCarbon()))
		assert.False(t, Parse("2020-08-05").Between(NewCarbon(), NewCarbon()))
		assert.False(t, NewCarbon().Between(NewCarbon(), Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").Between(NewCarbon(), Parse("2020-08-05")))
		assert.True(t, Parse("2020-08-05").Between(NewCarbon(), Parse("2020-08-06")))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").Between(Parse("xxx"), Parse("xxx")))

		assert.False(t, Parse("xxx").Between(Parse("xxx"), Parse("2020-08-05")))
		assert.False(t, Parse("xxx").Between(Parse("2020-08-05"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").Between(Parse("xxx"), Parse("xxx")))
		assert.False(t, Parse("2020-08-05").Between(Parse("xxx"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").Between(Parse("2020-08-05"), Parse("xxx")))
		assert.False(t, Parse("xxx").Between(Parse("2020-08-05"), Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").Between(Parse("2021-08-05 22:00:00"), Parse("2019-08-05 22:00:00")))
	})
}

func TestCarbon_BetweenIncludedStart(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().BetweenIncludedStart(NewCarbon(), NewCarbon()))
		assert.False(t, Parse("2020-08-05").BetweenIncludedStart(NewCarbon(), NewCarbon()))
		assert.True(t, NewCarbon().BetweenIncludedStart(NewCarbon(), Parse("2020-08-05")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedStart(NewCarbon(), Parse("2020-08-05")))
		assert.True(t, Parse("2020-08-05").BetweenIncludedStart(NewCarbon(), Parse("2020-08-06")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedStart(Parse("2020-08-05"), NewCarbon()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").BetweenIncludedStart(Parse("xxx"), Parse("xxx")))

		assert.False(t, Parse("xxx").BetweenIncludedStart(Parse("xxx"), Parse("2020-08-05")))
		assert.False(t, Parse("xxx").BetweenIncludedStart(Parse("2020-08-05"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").BetweenIncludedStart(Parse("xxx"), Parse("xxx")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedStart(Parse("xxx"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").BetweenIncludedStart(Parse("2020-08-05"), Parse("xxx")))
		assert.False(t, Parse("xxx").BetweenIncludedStart(Parse("2020-08-05"), Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2022-08-05 22:00:00"), Parse("2021-08-05 22:00:00")))
	})
}

func TestCarbon_BetweenIncludedEnd(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.False(t, NewCarbon().BetweenIncludedEnd(NewCarbon(), NewCarbon()))
		assert.False(t, Parse("2020-08-05").BetweenIncludedEnd(NewCarbon(), NewCarbon()))
		assert.False(t, NewCarbon().BetweenIncludedEnd(NewCarbon(), Parse("2020-08-05")))
		assert.True(t, Parse("2020-08-05").BetweenIncludedEnd(NewCarbon(), Parse("2020-08-05")))
		assert.True(t, Parse("2020-08-05").BetweenIncludedEnd(NewCarbon(), Parse("2020-08-06")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedEnd(Parse("2020-08-05"), NewCarbon()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").BetweenIncludedEnd(Parse("xxx"), Parse("xxx")))

		assert.False(t, Parse("xxx").BetweenIncludedEnd(Parse("xxx"), Parse("2020-08-05")))
		assert.False(t, Parse("xxx").BetweenIncludedEnd(Parse("2020-08-05"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").BetweenIncludedEnd(Parse("xxx"), Parse("xxx")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedEnd(Parse("xxx"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").BetweenIncludedEnd(Parse("2020-08-05"), Parse("xxx")))
		assert.False(t, Parse("xxx").BetweenIncludedEnd(Parse("2020-08-05"), Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2022-08-05 22:00:00"), Parse("2021-08-05 22:00:00")))
	})
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.True(t, NewCarbon().BetweenIncludedBoth(NewCarbon(), NewCarbon()))
		assert.False(t, Parse("2020-08-05").BetweenIncludedBoth(NewCarbon(), NewCarbon()))
		assert.True(t, NewCarbon().BetweenIncludedBoth(NewCarbon(), Parse("2020-08-05")))
		assert.True(t, Parse("2020-08-05").BetweenIncludedBoth(NewCarbon(), Parse("2020-08-05")))
		assert.True(t, Parse("2020-08-05").BetweenIncludedBoth(NewCarbon(), Parse("2020-08-06")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedBoth(Parse("2020-08-05"), NewCarbon()))
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.False(t, Parse("xxx").BetweenIncludedBoth(Parse("xxx"), Parse("xxx")))

		assert.False(t, Parse("xxx").BetweenIncludedBoth(Parse("xxx"), Parse("2020-08-05")))
		assert.False(t, Parse("xxx").BetweenIncludedBoth(Parse("2020-08-05"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").BetweenIncludedBoth(Parse("xxx"), Parse("xxx")))
		assert.False(t, Parse("2020-08-05").BetweenIncludedBoth(Parse("xxx"), Parse("2020-08-05")))

		assert.False(t, Parse("2020-08-05").BetweenIncludedBoth(Parse("2020-08-05"), Parse("xxx")))
		assert.False(t, Parse("xxx").BetweenIncludedBoth(Parse("2020-08-05"), Parse("xxx")))
	})

	t.Run("valid time", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		assert.True(t, Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		assert.False(t, Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2022-08-05 22:00:00"), Parse("2021-08-05 22:00:00")))
	})
}
