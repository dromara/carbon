package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DifferenceSuite struct {
	suite.Suite
}

func TestDifferenceSuite(t *testing.T) {
	suite.Run(t, new(DifferenceSuite))
}

func (s *DifferenceSuite) SetupTest() {
	SetTestNow(Parse("2020-08-05 13:14:15"))
}

func (s *DifferenceSuite) TearDownTest() {
	ClearTestNow()
}

func (s *DifferenceSuite) TestCarbon_DiffInYears() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInYears())
		s.Zero(Now().DiffInYears(c))
		s.Zero(c.DiffInYears(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(2019), c.DiffInYears())
		s.Equal(int64(-2019), Now().DiffInYears(c))
		s.Equal(int64(0), c.DiffInYears(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInYears())
		s.Zero(Now().DiffInYears(c))
		s.Zero(c.DiffInYears(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInYears())
		s.Zero(Now().DiffInYears(c))
		s.Zero(c.DiffInYears(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInYears())
		s.Equal(int64(0), Parse("2021-01-01 13:14:15").DiffInYears())
		s.Equal(int64(-1), Parse("2021-08-28 13:14:59").DiffInYears())
		s.Equal(int64(1), Parse("2018-08-28 13:14:59").DiffInYears())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-12-31 13:14:15").DiffInYears(Parse("2021-01-01 13:14:15")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2021-08-28 13:14:59")))
		s.Equal(int64(-1), Parse("2020-08-05 13:14:15").DiffInYears(Parse("2018-08-28 13:14:59")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInYears() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInYears())
		s.Zero(Now().DiffAbsInYears(c))
		s.Zero(c.DiffAbsInYears(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(2019), c.DiffAbsInYears())
		s.Equal(int64(2019), Now().DiffAbsInYears(c))
		s.Equal(int64(0), c.DiffAbsInYears(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInYears())
		s.Zero(Now().DiffAbsInYears(c))
		s.Zero(c.DiffAbsInYears(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInYears())
		s.Zero(Now().DiffAbsInYears(c))
		s.Zero(c.DiffAbsInYears(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInYears())
		s.Equal(int64(0), Parse("2021-01-01 13:14:15").DiffAbsInYears())
		s.Equal(int64(1), Parse("2021-08-28 13:14:59").DiffAbsInYears())
		s.Equal(int64(1), Parse("2018-08-28 13:14:59").DiffAbsInYears())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-12-31 13:14:15").DiffAbsInYears(Parse("2021-01-01 13:14:15")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2021-08-28 13:14:59")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInYears(Parse("2018-08-28 13:14:59")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInMonths())
		s.Zero(Now().DiffInMonths(c))
		s.Zero(c.DiffInMonths(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(24235), c.DiffInMonths())
		s.Equal(int64(-24235), Now().DiffInMonths(c))
		s.Equal(int64(0), c.DiffInMonths(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInMonths())
		s.Zero(Now().DiffInMonths(c))
		s.Zero(c.DiffInMonths(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInMonths())
		s.Zero(Now().DiffInMonths(c))
		s.Zero(c.DiffInMonths(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInMonths())
		s.Equal(int64(0), Parse("2020-07-28 13:14:00").DiffInMonths())
		s.Equal(int64(-1), Parse("2020-09-06 13:14:59").DiffInMonths())
		s.Equal(int64(23), Parse("2018-08-28 13:14:59").DiffInMonths())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-07-28 13:14:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2020-09-06 13:14:59")))
		s.Equal(int64(-23), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2018-08-28 13:14:59")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInMonths())
		s.Zero(Now().DiffAbsInMonths(c))
		s.Zero(c.DiffAbsInMonths(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(24235), c.DiffAbsInMonths())
		s.Equal(int64(24235), Now().DiffAbsInMonths(c))
		s.Equal(int64(0), c.DiffAbsInMonths(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInMonths())
		s.Zero(Now().DiffAbsInMonths(c))
		s.Zero(c.DiffAbsInMonths(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInMonths())
		s.Zero(Now().DiffAbsInMonths(c))
		s.Zero(c.DiffAbsInMonths(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths())
		s.Equal(int64(0), Parse("2020-07-28 13:14:00").DiffAbsInMonths())
		s.Equal(int64(1), Parse("2020-09-06 13:14:59").DiffAbsInMonths())
		s.Equal(int64(23), Parse("2018-08-28 13:14:59").DiffAbsInMonths())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-07-28 13:14:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2020-09-06 13:14:59")))
		s.Equal(int64(23), Parse("2020-08-05 13:14:15").DiffAbsInMonths(Parse("2018-08-28 13:14:59")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInWeeks() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInWeeks())
		s.Zero(Now().DiffInWeeks(c))
		s.Zero(c.DiffInWeeks(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(105377), c.DiffInWeeks())
		s.Equal(int64(-105377), Now().DiffInWeeks(c))
		s.Equal(int64(0), c.DiffInWeeks(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInWeeks())
		s.Zero(Now().DiffInWeeks(c))
		s.Zero(c.DiffInWeeks(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInWeeks())
		s.Zero(Now().DiffInWeeks(c))
		s.Zero(c.DiffInWeeks(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInWeeks())
		s.Equal(int64(1), Parse("2020-07-28 13:14:00").DiffInWeeks())
		s.Equal(int64(-1), Parse("2020-08-12 13:14:15").DiffInWeeks())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(-1), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-07-28 13:14:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffInWeeks(Parse("2020-08-12 13:14:15")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInWeeks() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInWeeks())
		s.Zero(Now().DiffAbsInWeeks(c))
		s.Zero(c.DiffAbsInWeeks(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(105377), c.DiffAbsInWeeks())
		s.Equal(int64(105377), Now().DiffAbsInWeeks(c))
		s.Equal(int64(0), c.DiffAbsInWeeks(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInWeeks())
		s.Zero(Now().DiffAbsInWeeks(c))
		s.Zero(c.DiffAbsInWeeks(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInWeeks())
		s.Zero(Now().DiffAbsInWeeks(c))
		s.Zero(c.DiffAbsInWeeks(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInWeeks())
		s.Equal(int64(1), Parse("2020-07-28 13:14:00").DiffAbsInWeeks())
		s.Equal(int64(1), Parse("2020-08-12 13:14:15").DiffAbsInWeeks())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-07-28 13:14:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInWeeks(Parse("2020-08-12 13:14:15")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInDays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInDays())
		s.Zero(Now().DiffInDays(c))
		s.Zero(c.DiffInDays(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(737641), c.DiffInDays())
		s.Equal(int64(-737641), Now().DiffInDays(c))
		s.Equal(int64(0), c.DiffInDays(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInDays())
		s.Zero(Now().DiffInDays(c))
		s.Zero(c.DiffInDays(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInDays())
		s.Zero(Now().DiffInDays(c))
		s.Zero(c.DiffInDays(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInDays())
		s.Equal(int64(0), Parse("2020-08-04 13:14:59").DiffInDays())
		s.Equal(int64(-1), Parse("2020-08-06 13:14:15").DiffInDays())
		s.Equal(int64(1), Parse("2020-08-04 13:00:00").DiffInDays())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-04 13:14:59")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-06 13:14:15")))
		s.Equal(int64(-1), Parse("2020-08-05 13:14:15").DiffInDays(Parse("2020-08-04 13:00:00")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInDays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInDays())
		s.Zero(Now().DiffAbsInDays(c))
		s.Zero(c.DiffAbsInDays(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(737641), c.DiffAbsInDays())
		s.Equal(int64(737641), Now().DiffAbsInDays(c))
		s.Equal(int64(0), c.DiffAbsInDays(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInDays())
		s.Zero(Now().DiffAbsInDays(c))
		s.Zero(c.DiffAbsInDays(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInDays())
		s.Zero(Now().DiffAbsInDays(c))
		s.Zero(c.DiffAbsInDays(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays())
		s.Equal(int64(0), Parse("2020-08-04 13:14:59").DiffAbsInDays())
		s.Equal(int64(1), Parse("2020-08-06 13:14:15").DiffAbsInDays())
		s.Equal(int64(1), Parse("2020-08-04 13:00:00").DiffAbsInDays())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-04 13:14:59")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-06 13:14:15")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInDays(Parse("2020-08-04 13:00:00")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInHours() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInHours())
		s.Zero(Now().DiffInHours(c))
		s.Zero(c.DiffInHours(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(17703397), c.DiffInHours())
		s.Equal(int64(-17703397), Now().DiffInHours(c))
		s.Equal(int64(0), c.DiffInHours(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInHours())
		s.Zero(Now().DiffInHours(c))
		s.Zero(c.DiffInHours(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInHours())
		s.Zero(Now().DiffInHours(c))
		s.Zero(c.DiffInHours(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInHours())
		s.Equal(int64(0), Parse("2020-08-05 14:13:15").DiffInHours())
		s.Equal(int64(1), Parse("2020-08-05 12:14:00").DiffInHours())
		s.Equal(int64(-1), Parse("2020-08-05 14:14:15").DiffInHours())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 14:13:15")))
		s.Equal(int64(-1), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 12:14:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffInHours(Parse("2020-08-05 14:14:15")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInHours() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInHours())
		s.Zero(Now().DiffAbsInHours(c))
		s.Zero(c.DiffAbsInHours(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(17703397), c.DiffAbsInHours())
		s.Equal(int64(17703397), Now().DiffAbsInHours(c))
		s.Equal(int64(0), c.DiffAbsInHours(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInHours())
		s.Zero(Now().DiffAbsInHours(c))
		s.Zero(c.DiffAbsInHours(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInHours())
		s.Zero(Now().DiffAbsInHours(c))
		s.Zero(c.DiffAbsInHours(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours())
		s.Equal(int64(0), Parse("2020-08-05 14:13:15").DiffAbsInHours())
		s.Equal(int64(1), Parse("2020-08-05 12:14:00").DiffAbsInHours())
		s.Equal(int64(1), Parse("2020-08-05 14:14:15").DiffAbsInHours())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 14:13:15")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 12:14:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInHours(Parse("2020-08-05 14:14:15")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInMinutes() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInHours())
		s.Zero(Now().DiffAbsInHours(c))
		s.Zero(c.DiffAbsInHours(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(1062203834), c.DiffInMinutes())
		s.Equal(int64(-1062203834), Now().DiffInMinutes(c))
		s.Equal(int64(0), c.DiffInMinutes(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInMinutes())
		s.Zero(Now().DiffInMinutes(c))
		s.Zero(c.DiffInMinutes(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInMinutes())
		s.Zero(Now().DiffInMinutes(c))
		s.Zero(c.DiffInMinutes(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes())
		s.Equal(int64(0), Parse("2020-08-05 13:15:10").DiffInMinutes())
		s.Equal(int64(1), Parse("2020-08-05 13:13:00").DiffInMinutes())
		s.Equal(int64(-1), Parse("2020-08-05 13:15:15").DiffInMinutes())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:15:10")))
		s.Equal(int64(-1), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:13:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffInMinutes(Parse("2020-08-05 13:15:15")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInMinutes() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInMinutes())
		s.Zero(Now().DiffAbsInMinutes(c))
		s.Zero(c.DiffAbsInMinutes(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(1062203834), c.DiffAbsInMinutes())
		s.Equal(int64(1062203834), Now().DiffAbsInMinutes(c))
		s.Equal(int64(0), c.DiffAbsInMinutes(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInMinutes())
		s.Zero(Now().DiffAbsInMinutes(c))
		s.Zero(c.DiffAbsInMinutes(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInMinutes())
		s.Zero(Now().DiffAbsInMinutes(c))
		s.Zero(c.DiffAbsInMinutes(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes())
		s.Equal(int64(0), Parse("2020-08-05 13:15:10").DiffAbsInMinutes())
		s.Equal(int64(1), Parse("2020-08-05 13:13:00").DiffAbsInMinutes())
		s.Equal(int64(1), Parse("2020-08-05 13:15:15").DiffAbsInMinutes())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:15:10")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:13:00")))
		s.Equal(int64(1), Parse("2020-08-05 13:14:15").DiffAbsInMinutes(Parse("2020-08-05 13:15:15")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInSeconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInSeconds())
		s.Zero(Now().DiffInSeconds(c))
		s.Zero(c.DiffInSeconds(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(63732230055), c.DiffInSeconds())
		s.Equal(int64(-63732230055), Now().DiffInSeconds(c))
		s.Equal(int64(0), c.DiffInSeconds(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInSeconds())
		s.Zero(Now().DiffInSeconds(c))
		s.Zero(c.DiffInSeconds(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInSeconds())
		s.Zero(Now().DiffInSeconds(c))
		s.Zero(c.DiffInSeconds(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds())
		s.Equal(int64(0), Parse("2020-08-05 13:14:15.999999").DiffInSeconds())
		s.Equal(int64(-5), Parse("2020-08-05 13:14:20").DiffInSeconds())
		s.Equal(int64(5), Parse("2020-08-05 13:14:10").DiffInSeconds())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:15.999999")))
		s.Equal(int64(5), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:20")))
		s.Equal(int64(-5), Parse("2020-08-05 13:14:15").DiffInSeconds(Parse("2020-08-05 13:14:10")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInSeconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInSeconds())
		s.Zero(Now().DiffAbsInSeconds(c))
		s.Zero(c.DiffAbsInSeconds(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(63732230055), c.DiffAbsInSeconds())
		s.Equal(int64(63732230055), Now().DiffAbsInSeconds(c))
		s.Equal(int64(0), c.DiffAbsInSeconds(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInSeconds())
		s.Zero(Now().DiffAbsInSeconds(c))
		s.Zero(c.DiffAbsInSeconds(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInSeconds())
		s.Zero(Now().DiffAbsInSeconds(c))
		s.Zero(c.DiffAbsInSeconds(c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds())
		s.Equal(int64(0), Parse("2020-08-05 13:14:15.999999").DiffAbsInSeconds())
		s.Equal(int64(5), Parse("2020-08-05 13:14:20").DiffAbsInSeconds())
		s.Equal(int64(5), Parse("2020-08-05 13:14:10").DiffAbsInSeconds())

		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:15")))
		s.Equal(int64(0), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:15.999999")))
		s.Equal(int64(5), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:20")))
		s.Equal(int64(5), Parse("2020-08-05 13:14:15").DiffAbsInSeconds(Parse("2020-08-05 13:14:10")))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInString())
		s.Zero(Now().DiffInString(c))
		s.Zero(c.DiffInString(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2019 years", c.DiffInString())
		s.Equal("-2019 years", Now().DiffInString(c))
		s.Equal("just now", c.DiffInString(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInString())
		s.Zero(Now().DiffInString(c))
		s.Zero(c.DiffInString(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInString())
		s.Zero(Now().DiffInString(c))
		s.Zero(c.DiffInString(c))
	})

	s.Run("valid carbon", func() {
		s.Equal("just now", Now().DiffInString())
		s.Equal("-1 year", Now().AddYearsNoOverflow(1).DiffInString())
		s.Equal("1 year", Now().SubYearsNoOverflow(1).DiffInString())
		s.Equal("-1 month", Now().AddMonthsNoOverflow(1).DiffInString())
		s.Equal("1 month", Now().SubMonthsNoOverflow(1).DiffInString())
		s.Equal("-1 week", Now().AddWeeks(1).DiffInString())
		s.Equal("1 week", Now().SubWeeks(1).DiffInString())
		s.Equal("-1 day", Now().AddDays(1).DiffInString())
		s.Equal("1 day", Now().SubDays(1).DiffInString())
		s.Equal("-1 hour", Now().AddHours(1).DiffInString())
		s.Equal("1 hour", Now().SubHours(1).DiffInString())
		s.Equal("-1 minute", Now().AddMinutes(1).DiffInString())
		s.Equal("1 minute", Now().SubMinutes(1).DiffInString())
		s.Equal("-1 second", Now().AddSeconds(1).DiffInString())
		s.Equal("1 second", Now().SubSeconds(1).DiffInString())

		s.Equal("just now", Now().DiffInString(Now()))
		s.Equal("-1 year", Now().AddYearsNoOverflow(1).DiffInString(Now()))
		s.Equal("1 year", Now().SubYearsNoOverflow(1).DiffInString(Now()))
		s.Equal("-1 month", Now().AddMonthsNoOverflow(1).DiffInString(Now()))
		s.Equal("1 month", Now().SubMonthsNoOverflow(1).DiffInString(Now()))
		s.Equal("-1 week", Now().AddWeeks(1).DiffInString(Now()))
		s.Equal("1 week", Now().SubWeeks(1).DiffInString(Now()))
		s.Equal("-1 day", Now().AddDays(1).DiffInString(Now()))
		s.Equal("1 day", Now().SubDays(1).DiffInString(Now()))
		s.Equal("-1 hour", Now().AddHours(1).DiffInString(Now()))
		s.Equal("1 hour", Now().SubHours(1).DiffInString(Now()))
		s.Equal("-1 minute", Now().AddMinutes(1).DiffInString(Now()))
		s.Equal("1 minute", Now().SubMinutes(1).DiffInString(Now()))
		s.Equal("-1 second", Now().AddSeconds(1).DiffInString(Now()))
		s.Equal("1 second", Now().SubSeconds(1).DiffInString(Now()))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInString() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInString())
		s.Zero(Now().DiffAbsInString(c))
		s.Zero(c.DiffAbsInString(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2019 years", c.DiffAbsInString())
		s.Equal("2019 years", Now().DiffAbsInString(c))
		s.Equal("just now", c.DiffAbsInString(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInString())
		s.Zero(Now().DiffAbsInString(c))
		s.Zero(c.DiffAbsInString(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInString())
		s.Zero(Now().DiffAbsInString(c))
		s.Zero(c.DiffAbsInString(c))
	})

	s.Run("valid carbon", func() {
		s.Equal("just now", Now().DiffAbsInString())
		s.Equal("1 year", Now().AddYearsNoOverflow(1).DiffAbsInString())
		s.Equal("1 year", Now().SubYearsNoOverflow(1).DiffAbsInString())
		s.Equal("1 month", Now().AddMonthsNoOverflow(1).DiffAbsInString())
		s.Equal("1 month", Now().SubMonthsNoOverflow(1).DiffAbsInString())
		s.Equal("1 day", Now().AddDays(1).DiffAbsInString())
		s.Equal("1 day", Now().SubDays(1).DiffAbsInString())
		s.Equal("1 hour", Now().AddHours(1).DiffAbsInString())
		s.Equal("1 hour", Now().SubHours(1).DiffAbsInString())
		s.Equal("1 minute", Now().AddMinutes(1).DiffAbsInString())
		s.Equal("1 minute", Now().SubMinutes(1).DiffAbsInString())
		s.Equal("1 second", Now().AddSeconds(1).DiffAbsInString())
		s.Equal("1 second", Now().SubSeconds(1).DiffAbsInString())

		s.Equal("just now", Parse("2020-08-05 13:14:15").DiffAbsInString(Now()))
		s.Equal("1 year", Now().AddYearsNoOverflow(1).DiffAbsInString(Now()))
		s.Equal("1 year", Now().SubYearsNoOverflow(1).DiffAbsInString(Now()))
		s.Equal("1 month", Now().AddMonthsNoOverflow(1).DiffAbsInString(Now()))
		s.Equal("1 month", Now().SubMonthsNoOverflow(1).DiffAbsInString(Now()))
		s.Equal("1 day", Now().AddDays(1).DiffAbsInString(Now()))
		s.Equal("1 day", Now().SubDays(1).DiffAbsInString(Now()))
		s.Equal("1 hour", Now().AddHours(1).DiffAbsInString(Now()))
		s.Equal("1 hour", Now().SubHours(1).DiffAbsInString(Now()))
		s.Equal("1 minute", Now().AddMinutes(1).DiffAbsInString(Now()))
		s.Equal("1 minute", Now().SubMinutes(1).DiffAbsInString(Now()))
		s.Equal("1 second", Now().AddSeconds(1).DiffAbsInString(Now()))
		s.Equal("1 second", Now().SubSeconds(1).DiffAbsInString(Now()))
	})
}

func (s *DifferenceSuite) TestCarbon_DiffInDuration() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffInDuration())
		s.Zero(Now().DiffInDuration(c))
		s.Zero(c.DiffInDuration(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2562047h47m16.854775807s", c.DiffInDuration().String())
		s.Equal("-2562047h47m16.854775808s", Now().DiffInDuration(c).String())
		s.Equal("0s", c.DiffInDuration(c).String())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffInDuration())
		s.Zero(Now().DiffInDuration(c))
		s.Zero(c.DiffInDuration(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffInDuration())
		s.Zero(Now().DiffInDuration(c))
		s.Zero(c.DiffInDuration(c))
	})

	s.Run("valid carbon", func() {
		s.Equal("0s", Now().DiffInDuration().String())
		s.Equal("-8760h0m0s", Now().AddYearsNoOverflow(1).DiffInDuration().String())
		s.Equal("8784h0m0s", Now().SubYearsNoOverflow(1).DiffInDuration().String())
		s.Equal("-744h0m0s", Now().AddMonthsNoOverflow(1).DiffInDuration().String())
		s.Equal("744h0m0s", Now().SubMonthsNoOverflow(1).DiffInDuration().String())
		s.Equal("-24h0m0s", Now().AddDays(1).DiffInDuration().String())
		s.Equal("24h0m0s", Now().SubDays(1).DiffInDuration().String())

		s.Equal("0s", Now().DiffInDuration(Now()).String())
		s.Equal("-8760h0m0s", Now().AddYearsNoOverflow(1).DiffInDuration(Now()).String())
		s.Equal("8784h0m0s", Now().SubYearsNoOverflow(1).DiffInDuration(Now()).String())
		s.Equal("-744h0m0s", Now().AddMonthsNoOverflow(1).DiffInDuration(Now()).String())
		s.Equal("744h0m0s", Now().SubMonthsNoOverflow(1).DiffInDuration(Now()).String())
		s.Equal("-24h0m0s", Now().AddDays(1).DiffInDuration(Now()).String())
		s.Equal("24h0m0s", Now().SubDays(1).DiffInDuration(Now()).String())
	})
}

func (s *DifferenceSuite) TestCarbon_DiffAbsInDuration() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(c.DiffAbsInDuration())
		s.Zero(Now().DiffAbsInDuration(c))
		s.Zero(c.DiffAbsInDuration(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2562047h47m16.854775807s", c.DiffAbsInDuration().String())
		s.Equal("2562047h47m16.854775807s", Now().DiffAbsInDuration(c).String())
		s.Equal("0s", c.DiffAbsInDuration(c).String())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.DiffAbsInDuration())
		s.Zero(Now().DiffAbsInDuration(c))
		s.Zero(c.DiffAbsInDuration(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.DiffAbsInDuration())
		s.Zero(Now().DiffAbsInDuration(c))
		s.Zero(c.DiffAbsInDuration(c))
	})

	s.Run("valid carbon", func() {
		s.Equal("0s", Now().DiffAbsInDuration().String())
		s.Equal("8760h0m0s", Now().AddYearsNoOverflow(1).DiffAbsInDuration().String())
		s.Equal("8784h0m0s", Now().SubYearsNoOverflow(1).DiffAbsInDuration().String())
		s.Equal("744h0m0s", Now().AddMonthsNoOverflow(1).DiffAbsInDuration().String())
		s.Equal("744h0m0s", Now().SubMonthsNoOverflow(1).DiffAbsInDuration().String())
		s.Equal("24h0m0s", Now().AddDays(1).DiffAbsInDuration().String())
		s.Equal("24h0m0s", Now().SubDays(1).DiffAbsInDuration().String())

		s.Equal("0s", Now().DiffAbsInDuration(Now()).String())
		s.Equal("8760h0m0s", Now().AddYearsNoOverflow(1).DiffAbsInDuration(Now()).String())
		s.Equal("8784h0m0s", Now().SubYearsNoOverflow(1).DiffAbsInDuration(Now()).String())
		s.Equal("744h0m0s", Now().AddMonthsNoOverflow(1).DiffAbsInDuration(Now()).String())
		s.Equal("744h0m0s", Now().SubMonthsNoOverflow(1).DiffAbsInDuration(Now()).String())
		s.Equal("24h0m0s", Now().AddDays(1).DiffAbsInDuration(Now()).String())
		s.Equal("24h0m0s", Now().SubDays(1).DiffAbsInDuration(Now()).String())
	})
}

func (s *DifferenceSuite) TestCarbon_DiffForHumans() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.DiffForHumans())
		s.Empty(Now().DiffForHumans(c))
		s.Empty(c.DiffForHumans(c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2019 years ago", c.DiffForHumans())
		s.Equal("2019 years after", Now().DiffForHumans(c))
		s.Equal("just now", c.DiffForHumans(c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Empty(c.DiffForHumans())
		s.Empty(Now().DiffForHumans(c))
		s.Empty(c.DiffForHumans(c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Empty(c.DiffForHumans())
		s.Empty(Now().DiffForHumans(c))
		s.Empty(c.DiffForHumans(c))
	})

	s.Run("nil lang", func() {
		c := Now()
		c.lang = nil
		s.Empty(c.DiffForHumans())
	})

	s.Run("valid carbon", func() {
		s.Equal("just now", Parse("2020-08-05 13:14:15").DiffForHumans())
		s.Equal("2 days ago", Parse("2020-08-03 13:14:15").DiffForHumans())
		s.Equal("2 days from now", Parse("2020-08-07 13:14:15").DiffForHumans())
		s.Equal("1 year from now", Now().AddYearsNoOverflow(1).DiffForHumans())
		s.Equal("1 year ago", Now().SubYearsNoOverflow(1).DiffForHumans())
		s.Equal("1 month from now", Now().AddMonthsNoOverflow(1).DiffForHumans())
		s.Equal("1 month ago", Now().SubMonthsNoOverflow(1).DiffForHumans())
		s.Equal("1 day from now", Now().AddDays(1).DiffForHumans())
		s.Equal("1 day ago", Now().SubDays(1).DiffForHumans())

		s.Equal("just now", Parse("2020-08-05 13:14:15").DiffForHumans(Now()))
		s.Equal("2 days before", Parse("2020-08-03 13:14:15").DiffForHumans(Now()))
		s.Equal("2 days after", Parse("2020-08-07 13:14:15").DiffForHumans(Now()))
		s.Equal("1 year after", Now().AddYearsNoOverflow(1).DiffForHumans(Now()))
		s.Equal("1 year before", Now().SubYearsNoOverflow(1).DiffForHumans(Now()))
		s.Equal("1 month after", Now().AddMonthsNoOverflow(1).DiffForHumans(Now()))
		s.Equal("1 month before", Now().SubMonthsNoOverflow(1).DiffForHumans(Now()))
		s.Equal("1 day after", Now().AddDays(1).DiffForHumans(Now()))
		s.Equal("1 day before", Now().SubDays(1).DiffForHumans(Now()))
	})
}

func (s *DifferenceSuite) TestCarbon_getDiffInMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Zero(getDiffInMonths(Now(), c))
		s.Zero(getDiffInMonths(c, Now()))
		s.Zero(getDiffInMonths(c, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(int64(-24236), getDiffInMonths(Now(), c))
		s.Equal(int64(24235), getDiffInMonths(c, Now()))
		s.Equal(int64(0), getDiffInMonths(c, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(getDiffInMonths(Now(), c))
		s.Zero(getDiffInMonths(c, Now()))
		s.Zero(getDiffInMonths(c, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(getDiffInMonths(Now(), c))
		s.Zero(getDiffInMonths(c, Now()))
		s.Zero(getDiffInMonths(c, c))
	})

	s.Run("valid carbon", func() {
		s.Equal(int64(0), getDiffInMonths(Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15")))
		s.Equal(int64(1), getDiffInMonths(Parse("2020-07-05 13:14:15"), Parse("2020-08-05 13:14:15")))
		s.Equal(int64(-1), getDiffInMonths(Parse("2020-09-05 13:14:15"), Parse("2020-08-05 13:14:15")))
	})
}

// https://github.com/dromara/carbon/issues/255
func (s *DifferenceSuite) TestCarbon_Issue255() {
	s.Equal(int64(0), Parse("2024-10-11").DiffInMonths(Parse("2024-11-10")))
	s.Equal(int64(0), Parse("2024-11-10").DiffInMonths(Parse("2024-10-11")))
	s.Equal(int64(1), Parse("2024-10-11").DiffInMonths(Parse("2024-11-11")))
	s.Equal(int64(-1), Parse("2024-11-11").DiffInMonths(Parse("2024-10-11")))
	s.Equal(int64(0), Parse("2024-10-11 23:59:00").DiffInMonths(Parse("2024-11-11 00:00:00")))
	s.Equal(int64(-23), Parse("2020-08-05 13:14:15").DiffInMonths(Parse("2018-08-28 13:14:59")))
	s.Equal(int64(23), Parse("2018-08-28 13:14:59").DiffInMonths(Parse("2020-08-05 13:14:15")))
	s.Equal(int64(11999), Parse("1024-12-25 13:14:20").DiffInMonths(Parse("2024-12-25 13:14:19")))
}

// TestCarbon_diff tests the internal diff function for 100% coverage
func (s *DifferenceSuite) TestCarbon_diff() {
	s.Run("same time", func() {
		now := Parse("2020-08-05 13:14:15")
		unit, value := now.diff(now)
		s.Equal("now", unit)
		s.Equal(int64(0), value)
	})

	s.Run("years difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2022-08-05 13:14:15")
		unit, value := start.diff(end)
		s.Equal("year", unit)
		s.Equal(int64(2), value)
	})

	s.Run("months difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-10-05 13:14:15")
		unit, value := start.diff(end)
		s.Equal("month", unit)
		s.Equal(int64(2), value)
	})

	s.Run("weeks difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-08-19 13:14:15")
		unit, value := start.diff(end)
		s.Equal("week", unit)
		s.Equal(int64(2), value)
	})

	s.Run("days difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-08-07 13:14:15")
		unit, value := start.diff(end)
		s.Equal("day", unit)
		s.Equal(int64(2), value)
	})

	s.Run("hours difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-08-05 15:14:15")
		unit, value := start.diff(end)
		s.Equal("hour", unit)
		s.Equal(int64(2), value)
	})

	s.Run("minutes difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-08-05 13:16:15")
		unit, value := start.diff(end)
		s.Equal("minute", unit)
		s.Equal(int64(2), value)
	})

	s.Run("seconds difference", func() {
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-08-05 13:14:17")
		unit, value := start.diff(end)
		s.Equal("second", unit)
		s.Equal(int64(2), value)
	})

	s.Run("edge case - all units zero", func() {
		// This should trigger the final return "now", 0
		start := Parse("2020-08-05 13:14:15")
		end := Parse("2020-08-05 13:14:15")
		unit, value := start.diff(end)
		s.Equal("now", unit)
		s.Equal(int64(0), value)
	})

	s.Run("edge case - very small difference", func() {
		// Test case where all absolute differences are 0 but secondsDiff is not 0
		// This should trigger the final return "now", 0
		start := Parse("2020-08-05 13:14:15.000000001")
		end := Parse("2020-08-05 13:14:15.000000002")
		unit, value := start.diff(end)
		s.Equal("now", unit)
		s.Equal(int64(0), value)
	})
}
