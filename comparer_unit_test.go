package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ComparerSuite struct {
	suite.Suite
}

func TestComparerSuite(t *testing.T) {
	suite.Run(t, new(ComparerSuite))
}

func (s *ComparerSuite) TestCarbon_HasError() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.HasError())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().HasError())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").HasError())
	})

	s.Run("error carbon", func() {
		s.True(Parse("xxx").HasError())
	})

	s.Run("valid carbon", func() {
		s.False(Now().HasError())
	})
}

func (s *ComparerSuite) TestCarbon_IsNil() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.True(c.IsNil())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsNil())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsNil())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsNil())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05").IsNil())
	})
}

func (s *ComparerSuite) TestCarbon_IsEmpty() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsEmpty())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsEmpty())
	})

	s.Run("empty carbon", func() {
		s.True(Parse("").IsEmpty())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsEmpty())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05").IsEmpty())
	})
}

func (s *ComparerSuite) TestCarbon_IsZero() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsZero())
	})

	s.Run("zero carbon", func() {
		stdTime1 := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
		carbon1 := CreateFromDateTimeNano(1, 1, 1, 0, 0, 0, 0, UTC)
		s.True(carbon1.IsZero())
		s.Equal(stdTime1.IsZero(), carbon1.IsZero())

		stdTime2 := time.Time{}
		carbon2 := NewCarbon()
		s.True(carbon2.IsZero())
		s.Equal(stdTime2.IsZero(), carbon2.IsZero())

		s.True(ZeroValue().IsZero())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsZero())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsZero())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("0000-00-00 00:00:00").IsZero())
		s.False(Parse("2020-08-05").IsZero())
		s.False(Parse("0000-00-00").IsZero())
	})
}

func (s *ComparerSuite) TestCarbon_IsEpoch() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsEpoch())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsEpoch())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsEpoch())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsEpoch())
	})

	s.Run("valid carbon", func() {
		s.True(CreateFromDateTimeNano(1970, 1, 1, 0, 0, 0, 0, UTC).IsEpoch())
		s.True(CreateFromTimestamp(0).IsEpoch())
		s.False(Parse("2020-08-05").IsEpoch())
		s.False(Parse("0000-00-00").IsEpoch())
	})
}

func (s *ComparerSuite) TestCarbon_IsValid() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsValid())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsValid())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsValid())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsValid())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-08").IsValid())
	})
}

func (s *ComparerSuite) TestCarbon_IsInvalid() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.True(c.IsInvalid())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsInvalid())
	})

	s.Run("empty carbon", func() {
		s.True(Parse("").IsInvalid())
	})

	s.Run("error carbon", func() {
		s.True(Parse("xxx").IsInvalid())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-08").IsInvalid())
	})
}

func (s *ComparerSuite) TestCarbon_IsDST() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsDST())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsDST())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsDST())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsDST())
	})

	s.Run("valid carbon", func() {
		tzWithDST, tzWithoutDST := "Australia/Sydney", "Australia/Brisbane"
		s.True(Parse("2009-01-01", tzWithDST).IsDST())
		s.False(Parse("2009-01-01", tzWithoutDST).IsDST())
	})
}

func (s *ComparerSuite) TestCarbon_IsAM() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsAM())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsAM())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsAM())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsAM())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-08 00:00:00").IsAM())
		s.False(Parse("2020-08-08 12:00:00").IsAM())
		s.False(Parse("2020-08-08 23:59:59").IsAM())
	})
}

func (s *ComparerSuite) TestCarbon_IsPM() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsPM())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsPM())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsPM())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsPM())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-08 00:00:00").IsPM())
		s.True(Parse("2020-08-08 12:00:00").IsPM())
		s.True(Parse("2020-08-08 23:59:59").IsPM())
	})
}

func (s *ComparerSuite) TestCarbon_IsLeapYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsLeapYear())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsLeapYear())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsLeapYear())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsLeapYear())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2015-01-01").IsLeapYear())
		s.True(Parse("2016-01-01").IsLeapYear())
	})
}

func (s *ComparerSuite) TestCarbon_IsLongYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsLongYear())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsLongYear())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsLongYear())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsLongYear())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2015-01-01").IsLongYear())
		s.False(Parse("2016-01-01").IsLongYear())
	})
}

func (s *ComparerSuite) TestCarbon_IsJanuary() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsJanuary())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsJanuary())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsJanuary())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsJanuary())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-01-01").IsJanuary())
		s.False(Parse("2020-02-01").IsJanuary())
	})
}

func (s *ComparerSuite) TestCarbon_IsFebruary() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsFebruary())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsFebruary())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsFebruary())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsFebruary())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-02-01").IsFebruary())
		s.False(Parse("2020-03-01").IsFebruary())
	})
}

func (s *ComparerSuite) TestCarbon_IsMarch() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsMarch())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsMarch())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsMarch())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsMarch())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-03-01").IsMarch())
		s.False(Parse("2020-04-01").IsMarch())
	})
}

func (s *ComparerSuite) TestCarbon_IsApril() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsApril())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsApril())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsApril())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsApril())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-04-01").IsApril())
		s.False(Parse("2020-05-01").IsApril())
	})
}

func (s *ComparerSuite) TestCarbon_IsMay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsMay())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsMay())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsMay())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsMay())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-05-01").IsMay())
		s.False(Parse("2020-06-01").IsMay())
	})
}

func (s *ComparerSuite) TestCarbon_IsJune() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsJune())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsJune())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsJune())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsJune())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-06-01").IsJune())
		s.False(Parse("2020-07-01").IsJune())
	})
}

func (s *ComparerSuite) TestCarbon_IsJuly() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsJuly())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsJuly())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsJuly())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsJuly())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-07-01").IsJuly())
		s.False(Parse("2020-08-01").IsJuly())
	})
}

func (s *ComparerSuite) TestCarbon_IsAugust() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsAugust())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsAugust())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsAugust())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsAugust())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-01").IsAugust())
		s.False(Parse("2020-09-01").IsAugust())
	})
}

func (s *ComparerSuite) TestCarbon_IsSeptember() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSeptember())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsSeptember())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsSeptember())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsSeptember())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-09-01").IsSeptember())
		s.False(Parse("2020-10-01").IsSeptember())
	})
}

func (s *ComparerSuite) TestCarbon_IsOctober() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsOctober())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsOctober())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsOctober())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsOctober())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-10-01").IsOctober())
		s.False(Parse("2020-11-01").IsOctober())
	})
}

func (s *ComparerSuite) TestCarbon_IsNovember() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsNovember())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsNovember())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsNovember())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsNovember())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-11-01").IsNovember())
		s.False(Parse("2020-12-01").IsNovember())
	})
}

func (s *ComparerSuite) TestCarbon_IsDecember() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsDecember())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsDecember())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsDecember())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsDecember())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-12-01").IsDecember())
		s.False(Parse("2020-01-01").IsDecember())
	})
}

func (s *ComparerSuite) TestCarbon_IsMonday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsMonday())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsMonday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsMonday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsMonday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-03").IsMonday())
		s.False(Parse("2025-03-04").IsMonday())
	})
}

func (s *ComparerSuite) TestCarbon_IsTuesday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsTuesday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsTuesday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsTuesday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsTuesday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-04").IsTuesday())
		s.False(Parse("2025-03-05").IsTuesday())
	})
}

func (s *ComparerSuite) TestCarbon_IsWednesday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsWednesday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsWednesday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsWednesday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsWednesday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-05").IsWednesday())
		s.False(Parse("2025-03-06").IsWednesday())
	})
}

func (s *ComparerSuite) TestCarbon_IsThursday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsThursday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsThursday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsThursday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsThursday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-06").IsThursday())
		s.False(Parse("2025-03-07").IsThursday())
	})
}

func (s *ComparerSuite) TestCarbon_IsFriday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsFriday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsFriday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsFriday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsFriday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-07").IsFriday())
		s.False(Parse("2025-03-08").IsFriday())
	})
}

func (s *ComparerSuite) TestCarbon_IsSaturday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSaturday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsSaturday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsSaturday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsSaturday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-08").IsSaturday())
		s.False(Parse("2025-03-09").IsSaturday())
	})
}

func (s *ComparerSuite) TestCarbon_IsSunday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSunday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsSunday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsSunday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsSunday())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-09").IsSunday())
		s.False(Parse("2025-03-10").IsSunday())
	})
}

func (s *ComparerSuite) TestCarbon_IsWeekday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsWeekday())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsWeekday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsWeekday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsWeekday())
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2025-03-01").IsWeekday())
		s.False(Parse("2025-03-02").IsWeekday())
		s.True(Parse("2025-03-03").IsWeekday())
	})
}

func (s *ComparerSuite) TestCarbon_IsWeekend() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsWeekend())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsWeekend())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsWeekend())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsWeekend())
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2025-03-01").IsWeekend())
		s.True(Parse("2025-03-02").IsWeekend())
		s.False(Parse("2025-03-03").IsWeekend())
	})
}

func (s *ComparerSuite) TestCarbon_IsNow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsNow())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsNow())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsNow())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsNow())
	})

	s.Run("valid carbon", func() {
		s.False(Yesterday().IsNow())
		s.True(Now().IsNow())
		s.False(Tomorrow().IsNow())
		s.False(Parse("2025-03-01").IsNow())
	})
}

func (s *ComparerSuite) TestCarbon_IsFuture() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsFuture())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsFuture())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsFuture())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsFuture())
	})

	s.Run("valid carbon", func() {
		s.False(Yesterday().IsFuture())
		s.False(Now().IsFuture())
		s.True(Tomorrow().IsFuture())
		s.False(Parse("2025-03-01").IsFuture())
	})
}

func (s *ComparerSuite) TestCarbon_IsPast() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsPast())
	})

	s.Run("zero carbon", func() {
		s.True(NewCarbon().IsPast())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsPast())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsPast())
	})

	s.Run("valid carbon", func() {
		s.True(Yesterday().IsPast())
		s.False(Now().IsPast())
		s.False(Tomorrow().IsPast())
		s.True(Parse("2025-03-01").IsPast())
	})
}

func (s *ComparerSuite) TestCarbon_IsYesterday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsYesterday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsYesterday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsYesterday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsYesterday())
	})

	s.Run("valid carbon", func() {
		s.True(Yesterday().IsYesterday())
		s.False(Now().IsYesterday())
		s.False(Tomorrow().IsYesterday())
		s.False(Parse("2025-03-01").IsYesterday())
	})
}

func (s *ComparerSuite) TestCarbon_IsToday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsToday())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsToday())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsToday())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsToday())
	})

	s.Run("valid carbon", func() {
		s.False(Yesterday().IsToday())
		s.True(Now().IsToday())
		s.False(Tomorrow().IsToday())
		s.False(Parse("2025-03-01").IsToday())
	})
}

func (s *ComparerSuite) TestCarbon_IsTomorrow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsTomorrow())
	})

	s.Run("zero carbon", func() {
		s.False(NewCarbon().IsTomorrow())
	})

	s.Run("empty carbon", func() {
		s.False(Parse("").IsTomorrow())
	})

	s.Run("error carbon", func() {
		s.False(Parse("xxx").IsTomorrow())
	})

	s.Run("valid carbon", func() {
		s.False(Yesterday().IsTomorrow())
		s.False(Now().IsTomorrow())
		s.True(Tomorrow().IsTomorrow())
		s.False(Parse("2025-03-01").IsTomorrow())
	})
}

func (s *ComparerSuite) TestCarbon_IsSameCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameCentury(NewCarbon()))
		s.False(Now().IsSameCentury(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameCentury(Now()))
		s.False(Now().IsSameCentury(c))
		s.True(c.IsSameCentury(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameCentury(Now()))
		s.False(Now().IsSameCentury(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameCentury(Now()))
		s.False(Now().IsSameCentury(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05").IsSameCentury(Parse("2010-01-01")))
		s.True(Parse("2020-08-05").IsSameCentury(Parse("2030-12-31")))
		s.False(Parse("2020-08-05").IsSameCentury(Parse("2100-08-05")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameDecade(NewCarbon()))
		s.False(Now().IsSameDecade(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameDecade(Now()))
		s.False(Now().IsSameDecade(c))
		s.True(c.IsSameDecade(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameDecade(Now()))
		s.False(Now().IsSameDecade(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameDecade(Now()))
		s.False(Now().IsSameDecade(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05").IsSameDecade(Parse("2020-01-01")))
		s.True(Parse("2020-08-05").IsSameDecade(Parse("2020-12-31")))
		s.False(Parse("2020-08-05").IsSameDecade(Parse("2010-08-05")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameYear(NewCarbon()))
		s.False(Now().IsSameYear(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameYear(Now()))
		s.False(Now().IsSameYear(c))
		s.True(c.IsSameYear(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameYear(Now()))
		s.False(Now().IsSameYear(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameYear(Now()))
		s.False(Now().IsSameYear(c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05").IsSameYear(Parse("2010-08-05")))
		s.True(Parse("2020-08-05").IsSameYear(Parse("2020-01-01")))
		s.True(Parse("2020-08-05").IsSameYear(Parse("2020-12-31")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameQuarter(NewCarbon()))
		s.False(Now().IsSameQuarter(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameQuarter(Now()))
		s.False(Now().IsSameQuarter(c))
		s.True(c.IsSameQuarter(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameQuarter(Now()))
		s.False(Now().IsSameQuarter(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameQuarter(Now()))
		s.False(Now().IsSameQuarter(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05").IsSameQuarter(Parse("2020-08-06")))
		s.False(Parse("2020-08-05").IsSameQuarter(Parse("2010-08-05")))
		s.False(Parse("2020-08-05").IsSameQuarter(Parse("2010-01-05")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameMonth(NewCarbon()))
		s.False(Now().IsSameMonth(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameMonth(Now()))
		s.False(Now().IsSameMonth(c))
		s.True(c.IsSameMonth(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameMonth(Now()))
		s.False(Now().IsSameMonth(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameMonth(Now()))
		s.False(Now().IsSameMonth(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05").IsSameMonth(Parse("2020-08-01")))
		s.False(Parse("2020-08-05").IsSameMonth(Parse("2021-08-05")))
		s.False(Parse("2020-08-05").IsSameMonth(Parse("2020-09-05")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameDay(NewCarbon()))
		s.False(Now().IsSameDay(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameDay(Now()))
		s.False(Now().IsSameDay(c))
		s.True(c.IsSameDay(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameDay(Now()))
		s.False(Now().IsSameDay(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameDay(Now()))
		s.False(Now().IsSameDay(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 00:00:00").IsSameDay(Parse("2020-08-05 23:59:59")))
		s.False(Parse("2020-08-05 00:00:00").IsSameDay(Parse("2021-08-05 00:00:00")))
		s.False(Parse("2020-08-05 00:00:00").IsSameDay(Parse("2020-09-05 00:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameHour(NewCarbon()))
		s.False(Now().IsSameHour(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameHour(Now()))
		s.False(Now().IsSameHour(c))
		s.True(c.IsSameHour(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameHour(Now()))
		s.False(Now().IsSameHour(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameHour(Now()))
		s.False(Now().IsSameHour(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-08-05 22:59:59")))
		s.False(Parse("2020-08-05 22:00:00").IsSameHour(Parse("2021-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-09-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-08-06 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameHour(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameMinute(NewCarbon()))
		s.False(Now().IsSameMinute(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameMinute(Now()))
		s.False(Now().IsSameMinute(c))
		s.True(c.IsSameMinute(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameMinute(Now()))
		s.False(Now().IsSameMinute(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameMinute(Now()))
		s.False(Now().IsSameMinute(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-05 22:00:59")))
		s.False(Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2021-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-06 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-05 23:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameMinute(Parse("2020-08-05 22:01:00")))
	})
}

func (s *ComparerSuite) TestCarbon_IsSameSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsSameSecond(NewCarbon()))
		s.False(Now().IsSameSecond(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsSameSecond(Now()))
		s.False(Now().IsSameSecond(c))
		s.True(c.IsSameSecond(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsSameSecond(Now()))
		s.False(Now().IsSameSecond(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsSameSecond(Now()))
		s.False(Now().IsSameSecond(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 22:00:00.999999")))
		s.False(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2021-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-09-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-06 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 23:00:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 22:01:00")))
		s.False(Parse("2020-08-05 22:00:00").IsSameSecond(Parse("2020-08-05 22:00:01")))
	})
}

func (s *ComparerSuite) TestCarbon_IsBirthday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.IsBirthday())
		s.False(c.IsBirthday(NewCarbon()))
		s.False(Now().IsBirthday(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.IsBirthday(Parse("2020-08-05")))
		s.False(Parse("2020-08-05").IsBirthday(c))
		s.True(c.IsBirthday(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.IsBirthday())
		s.False(c.IsBirthday(Now()))
		s.False(Now().IsBirthday(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.IsBirthday())
		s.False(c.IsBirthday(Now()))
		s.False(Now().IsBirthday(c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2014-09-26").IsBirthday(Parse("1987-04-23")))
		s.True(Parse("2014-04-23").IsBirthday(Parse("1987-04-23")))
		s.True(Parse("1987-04-23 00:00:00").IsBirthday(Parse("2014-04-23 23:59:59")))
		s.False(Parse("1987-04-23").IsBirthday(Parse("1987-05-23")))
		s.False(Parse("1987-04-23").IsBirthday(Parse("1987-04-24")))
	})

	s.Run("without date", func() {
		SetTestNow(Parse("2020-08-05 13:14:15"))
		defer ClearTestNow()

		s.True(Parse("1990-08-05").IsBirthday())
		s.False(Parse("1990-08-04").IsBirthday())
		s.False(Parse("1990-08-06").IsBirthday())
	})

	s.Run("february 29", func() {
		s.True(Parse("2000-02-29").IsBirthday(Parse("2024-02-29")))
		s.False(Parse("2000-02-29").IsBirthday(Parse("2023-02-28")))
		s.False(Parse("2000-02-29").IsBirthday(Parse("2023-03-01")))
	})

	s.Run("without date in another timezone", func() {
		SetTestNow(Parse("2020-08-05 23:30:00", UTC))
		defer ClearTestNow()

		s.True(Parse("1990-08-06", Tokyo).IsBirthday())
		s.False(Parse("1990-08-05", Tokyo).IsBirthday())
	})

	s.Run("dates in different timezones", func() {
		s.False(Parse("1990-08-05", UTC).IsBirthday(Parse("2020-08-06 01:00:00", Tokyo)))
		s.True(Parse("1990-08-06", UTC).IsBirthday(Parse("2020-08-06 01:00:00", Tokyo)))
	})
}

func (s *ComparerSuite) TestCarbon_Compare() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Compare("<", Now()))
		s.False(Now().Compare("<", c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.Compare("=", Now()))
		s.False(Now().Compare("=", c))
		s.True(c.Compare("=", c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Compare("<", Now()))
		s.False(Now().Compare(">", c))
		s.False(c.Compare("!=", c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Compare("<", Now()))
		s.False(Now().Compare(">", c))
		s.False(c.Compare("!=", c))
	})

	s.Run("invalid operator", func() {
		s.False(Now().Compare("", Yesterday()))
		s.False(Now().Compare("%", Yesterday()))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").Compare("=", Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Compare(">=", Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Compare("<=", Parse("2020-08-05 22:00:00")))

		s.True(Parse("2020-08-05 22:00:00").Compare(">", Parse("2020-07-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Compare(">=", Parse("2020-07-05 22:00:00")))

		s.True(Parse("2020-08-05 22:00:00").Compare("<", Parse("2020-09-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Compare("<=", Parse("2020-09-05 22:00:00")))

		s.True(Parse("2020-08-05 22:00:00").Compare("<>", Parse("2020-06-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Compare("!=", Parse("2020-06-05 22:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Gt() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Gt(Now()))
		s.False(Now().Gt(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.Gt(Now()))
		s.True(Now().Gt(c))
		s.False(c.Gt(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Gt(Now()))
		s.False(Now().Gt(c))
		s.False(c.Gt(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Gt(Now()))
		s.False(Now().Gt(c))
		s.False(c.Gt(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").Gt(Parse("2020-08-05 21:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Gt(Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Gt(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Lt() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Lt(Now()))
		s.False(Now().Lt(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.True(c.Lt(Now()))
		s.False(Now().Lt(c))
		s.False(c.Lt(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Lt(Now()))
		s.False(Now().Lt(c))
		s.False(c.Lt(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Lt(Now()))
		s.False(Now().Lt(c))
		s.False(c.Lt(c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05 22:00:00").Lt(Parse("2020-08-05 21:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Lt(Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Lt(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Eq() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Eq(Now()))
		s.False(Now().Eq(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.Eq(Now()))
		s.False(Now().Eq(c))
		s.True(c.Eq(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Eq(Now()))
		s.False(Now().Eq(c))
		s.False(c.Eq(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Eq(Now()))
		s.False(Now().Eq(c))
		s.False(c.Eq(c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05 22:00:00").Eq(Parse("2020-08-05 21:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Eq(Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Eq(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Ne() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Ne(Now()))
		s.False(Now().Ne(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.True(c.Ne(Now()))
		s.True(Now().Ne(c))
		s.False(c.Ne(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Ne(Now()))
		s.False(Now().Ne(c))
		s.False(c.Ne(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Ne(Now()))
		s.False(Now().Ne(c))
		s.False(c.Ne(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").Ne(Parse("2020-08-05 21:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Ne(Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Ne(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Gte() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Gte(Now()))
		s.False(Now().Gte(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.Gte(Now()))
		s.True(Now().Gte(c))
		s.True(c.Gte(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Gte(Now()))
		s.False(Now().Gte(c))
		s.False(c.Gte(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Gte(Now()))
		s.False(Now().Gte(c))
		s.False(c.Gte(c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").Gte(Parse("2020-08-05 21:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Gte(Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Gte(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Lte() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Lte(Now()))
		s.False(Now().Lte(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.True(c.Lte(Now()))
		s.False(Now().Lte(c))
		s.True(c.Lte(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Lte(Now()))
		s.False(Now().Lte(c))
		s.False(c.Lte(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Lte(Now()))
		s.False(Now().Lte(c))
		s.False(c.Lte(c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05 22:00:00").Lte(Parse("2020-08-05 21:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Lte(Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").Lte(Parse("2020-08-05 23:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_Between() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.Between(Now(), Tomorrow()))
		s.False(Now().Between(c, Tomorrow()))
		s.False(Yesterday().Between(Now(), c))
		s.False(Yesterday().Between(c, c))
		s.False(c.Between(c, Tomorrow()))
		s.False(c.Between(c, c))

	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.Between(Now(), Tomorrow()))
		s.True(Now().Between(c, Tomorrow()))
		s.False(Yesterday().Between(Now(), c))
		s.False(Yesterday().Between(c, c))
		s.False(c.Between(c, Tomorrow()))
		s.False(c.Between(c, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.Between(Now(), Tomorrow()))
		s.False(Now().Between(c, Tomorrow()))
		s.False(Yesterday().Between(Now(), c))
		s.False(Yesterday().Between(c, c))
		s.False(c.Between(c, Tomorrow()))
		s.False(c.Between(c, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.Between(Now(), Tomorrow()))
		s.False(Now().Between(c, Tomorrow()))
		s.False(Yesterday().Between(Now(), c))
		s.False(Yesterday().Between(c, c))
		s.False(c.Between(c, Tomorrow()))
		s.False(c.Between(c, c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Between(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").Between(Parse("2021-08-05 22:00:00"), Parse("2019-08-05 22:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_BetweenIncludedStart() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.BetweenIncludedStart(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedStart(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedStart(Now(), c))
		s.False(Yesterday().BetweenIncludedStart(c, c))
		s.False(c.BetweenIncludedStart(c, Tomorrow()))
		s.False(c.BetweenIncludedStart(c, c))

	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.BetweenIncludedStart(Now(), Tomorrow()))
		s.True(Now().BetweenIncludedStart(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedStart(Now(), c))
		s.False(Yesterday().BetweenIncludedStart(c, c))
		s.True(c.BetweenIncludedStart(c, Tomorrow()))
		s.True(c.BetweenIncludedStart(c, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.BetweenIncludedStart(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedStart(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedStart(Now(), c))
		s.False(Yesterday().BetweenIncludedStart(c, c))
		s.False(c.BetweenIncludedStart(c, Tomorrow()))
		s.False(c.BetweenIncludedStart(c, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.BetweenIncludedStart(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedStart(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedStart(Now(), c))
		s.False(Yesterday().BetweenIncludedStart(c, c))
		s.False(c.BetweenIncludedStart(c, Tomorrow()))
		s.False(c.BetweenIncludedStart(c, c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedStart(Parse("2022-08-05 22:00:00"), Parse("2021-08-05 22:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_BetweenIncludedEnd() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.BetweenIncludedEnd(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedEnd(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedEnd(Now(), c))
		s.False(Yesterday().BetweenIncludedEnd(c, c))
		s.False(c.BetweenIncludedEnd(c, Tomorrow()))
		s.False(c.BetweenIncludedEnd(c, c))

	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.BetweenIncludedEnd(Now(), Tomorrow()))
		s.True(Now().BetweenIncludedEnd(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedEnd(Now(), c))
		s.False(Yesterday().BetweenIncludedEnd(c, c))
		s.False(c.BetweenIncludedEnd(c, Tomorrow()))
		s.True(c.BetweenIncludedEnd(c, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.BetweenIncludedEnd(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedEnd(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedEnd(Now(), c))
		s.False(Yesterday().BetweenIncludedEnd(c, c))
		s.False(c.BetweenIncludedEnd(c, Tomorrow()))
		s.False(c.BetweenIncludedEnd(c, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.BetweenIncludedEnd(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedEnd(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedEnd(Now(), c))
		s.False(Yesterday().BetweenIncludedEnd(c, c))
		s.False(c.BetweenIncludedEnd(c, Tomorrow()))
		s.False(c.BetweenIncludedEnd(c, c))
	})

	s.Run("valid carbon", func() {
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedEnd(Parse("2022-08-05 22:00:00"), Parse("2021-08-05 22:00:00")))
	})
}

func (s *ComparerSuite) TestCarbon_BetweenIncludedBoth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.False(c.BetweenIncludedBoth(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedBoth(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedBoth(Now(), c))
		s.False(Yesterday().BetweenIncludedBoth(c, c))
		s.False(c.BetweenIncludedBoth(c, Tomorrow()))
		s.False(c.BetweenIncludedBoth(c, c))

	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.BetweenIncludedBoth(Now(), Tomorrow()))
		s.True(Now().BetweenIncludedBoth(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedBoth(Now(), c))
		s.False(Yesterday().BetweenIncludedBoth(c, c))
		s.True(c.BetweenIncludedBoth(c, Tomorrow()))
		s.True(c.BetweenIncludedBoth(c, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.False(c.BetweenIncludedBoth(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedBoth(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedBoth(Now(), c))
		s.False(Yesterday().BetweenIncludedBoth(c, c))
		s.False(c.BetweenIncludedBoth(c, Tomorrow()))
		s.False(c.BetweenIncludedBoth(c, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.False(c.BetweenIncludedBoth(Now(), Tomorrow()))
		s.False(Now().BetweenIncludedBoth(c, Tomorrow()))
		s.False(Yesterday().BetweenIncludedBoth(Now(), c))
		s.False(Yesterday().BetweenIncludedBoth(c, c))
		s.False(c.BetweenIncludedBoth(c, Tomorrow()))
		s.False(c.BetweenIncludedBoth(c, c))
	})

	s.Run("valid carbon", func() {
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 22:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 23:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 22:00:00"), Parse("2020-08-05 23:00:00")))
		s.True(Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2020-08-05 21:00:00"), Parse("2020-08-05 22:00:00")))
		s.False(Parse("2020-08-05 22:00:00").BetweenIncludedBoth(Parse("2022-08-05 22:00:00"), Parse("2021-08-05 22:00:00")))
	})
}
