package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TravelerSuite struct {
	suite.Suite
}

func TestTravelerSuite(t *testing.T) {
	suite.Run(t, new(TravelerSuite))
}

func (s *TravelerSuite) TearDownTest() {
	ClearTestNow()
}

func (s *TravelerSuite) TestNow() {
	s.Run("without timezone", func() {
		c := Now()
		s.False(c.HasError())
		s.Equal(time.Now().Format(DateLayout), c.Layout(DateLayout, Local))
	})

	s.Run("empty timezone", func() {
		c := Now("")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error timezone", func() {
		c := Now("xxx")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid timezone", func() {
		c := Now(UTC)
		s.False(c.HasError())
		s.Equal(time.Now().In(time.UTC).Format(DateLayout), c.Layout(DateLayout))
	})

	s.Run("frozen time", func() {
		SetTestNow(Parse("2020-08-05"))
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Now().ToString())
		s.Equal("2020-08-05 08:00:00 +0800 CST", Now(PRC).ToString())
	})
}

func (s *TravelerSuite) TestTomorrow() {
	s.Run("without timezone", func() {
		c := Tomorrow()
		s.False(c.HasError())
		s.Equal(time.Now().Add(time.Hour*24).Format(DateLayout), c.Layout(DateLayout, Local))
	})

	s.Run("empty timezone", func() {
		c := Tomorrow("")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error timezone", func() {
		c := Tomorrow("xxx")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid timezone", func() {
		c := Tomorrow(UTC)
		s.False(c.HasError())
		s.Equal(time.Now().Add(time.Hour*24).In(time.UTC).Format(DateLayout), c.Layout(DateLayout))
	})

	s.Run("frozen time", func() {
		SetTestNow(Parse("2020-08-05"))
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Tomorrow().ToString())
		s.Equal("2020-08-06 08:00:00 +0800 CST", Tomorrow(PRC).ToString())
	})
}

func (s *TravelerSuite) TestYesterday() {
	s.Run("without timezone", func() {
		c := Yesterday()
		s.False(c.HasError())
		s.Equal(time.Now().Add(time.Hour*-24).Format(DateLayout), c.Layout(DateLayout, Local))
	})

	s.Run("empty timezone", func() {
		c := Yesterday("")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error timezone", func() {
		c := Yesterday("xxx")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid timezone", func() {
		c := Yesterday(UTC)
		s.False(c.HasError())
		s.Equal(time.Now().Add(time.Hour*-24).In(time.UTC).Format(DateLayout), c.Layout(DateLayout))
	})

	s.Run("frozen time", func() {
		SetTestNow(Parse("2020-08-05"))
		s.Equal("2020-08-04 00:00:00 +0000 UTC", Yesterday().ToString())
		s.Equal("2020-08-04 08:00:00 +0800 CST", Yesterday(PRC).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDuration() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDuration("10h")
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDuration("10h")
		s.False(c.HasError())
		s.Equal("0001-01-01 10:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDuration("10h")
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDuration("10h")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("empty duration", func() {
		c := Parse("2020-08-05").AddDuration("")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error duration", func() {
		c := Parse("2020-08-05").AddDuration("xxx")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15")
		s.Equal("2020-01-01 23:14:15 +0000 UTC", c.Copy().AddDuration("10h").ToString())
		s.Equal("2020-01-01 23:44:15 +0000 UTC", c.Copy().AddDuration("10.5h").ToString())
		s.Equal("2020-01-01 13:24:15 +0000 UTC", c.Copy().AddDuration("10m").ToString())
		s.Equal("2020-01-01 13:24:45 +0000 UTC", c.Copy().AddDuration("10.5m").ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDuration() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDuration("10h")
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDuration("10h")
		s.False(c.HasError())
		s.Equal("0000-12-31 14:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDuration("10h")
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDuration("10h")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error duration", func() {
		c := Parse("2020-08-05").SubDuration("xxx")
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 03:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10h").ToString())
		s.Equal("2020-01-01 02:44:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10.5h").ToString())
		s.Equal("2020-01-01 13:04:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10m").ToString())
		s.Equal("2020-01-01 13:03:45 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10.5m").ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCenturies() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddCenturies(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddCenturies(2)
		s.False(c.HasError())
		s.Equal("0201-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddCenturies(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddCenturies(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(0).ToString())
		s.Equal("2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(1).ToString())
		s.Equal("2220-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCenturiesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddCenturiesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddCenturiesNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0201-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddCenturiesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddCenturiesNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(0).ToString())
		s.Equal("2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(1).ToString())
		s.Equal("2220-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddCentury()
		s.False(c.HasError())
		s.Equal("0101-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddCentury()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").AddCentury()
		s.False(c.HasError())
		s.Equal("2120-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCenturyNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddCenturyNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddCenturyNoOverflow()
		s.False(c.HasError())
		s.Equal("0101-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddCenturyNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddCenturyNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").AddCenturyNoOverflow()
		s.False(c.HasError())
		s.Equal("2120-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCenturies() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubCenturies(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubCenturies(2)
		s.False(c.HasError())
		s.Equal("-0199-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubCenturies(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubCenturies(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(0).ToString())
		s.Equal("1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(1).ToString())
		s.Equal("1820-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCenturiesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubCenturiesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubCenturiesNoOverflow(2)
		s.False(c.HasError())
		s.Equal("-0199-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubCenturiesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubCenturiesNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(0).ToString())
		s.Equal("1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(1).ToString())
		s.Equal("1820-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubCentury()
		s.False(c.HasError())
		s.Equal("-0099-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubCentury()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubCentury()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").SubCentury()
		s.False(c.HasError())
		s.Equal("1920-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCenturyNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubCenturyNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubCenturyNoOverflow()
		s.False(c.HasError())
		s.Equal("-0099-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubCenturyNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubCenturyNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").SubCenturyNoOverflow()
		s.False(c.HasError())
		s.Equal("1920-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecades() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDecades(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDecades(2)
		s.False(c.HasError())
		s.Equal("0021-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDecades(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDecades(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(0).ToString())
		s.Equal("2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(1).ToString())
		s.Equal("2040-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecadesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDecadesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDecadesNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0021-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDecadesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDecadesNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(0).ToString())
		s.Equal("2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(1).ToString())
		s.Equal("2040-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDecade()
		s.False(c.HasError())
		s.Equal("0011-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDecade()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").AddDecade()
		s.False(c.HasError())
		s.Equal("2030-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecadeNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDecadeNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDecadeNoOverflow()
		s.False(c.HasError())
		s.Equal("0011-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDecadeNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDecadeNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").AddDecadeNoOverflow()
		s.False(c.HasError())
		s.Equal("2030-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecades() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDecades(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDecades(2)
		s.False(c.HasError())
		s.Equal("-0019-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDecades(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDecades(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(0).ToString())
		s.Equal("2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(1).ToString())
		s.Equal("2000-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecadesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDecadesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDecadesNoOverflow(2)
		s.False(c.HasError())
		s.Equal("-0019-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDecadesNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDecadesNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(0).ToString())
		s.Equal("2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(1).ToString())
		s.Equal("2000-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDecade()
		s.False(c.HasError())
		s.Equal("-0009-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDecade()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDecade()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01 13:14:15").SubDecade()
		s.False(c.HasError())
		s.Equal("2010-01-01 13:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecadeNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDecadeNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDecadeNoOverflow()
		s.False(c.HasError())
		s.Equal("-0009-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDecadeNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDecadeNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-01-01").SubDecadeNoOverflow()
		s.False(c.HasError())
		s.Equal("2010-01-01", c.ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYears() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddYears(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddYears(2)
		s.False(c.HasError())
		s.Equal("0003-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddYears(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddYears(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddYears(0).ToDateString())
		s.Equal("2021-01-01", Parse("2020-01-01").AddYears(1).ToDateString())
		s.Equal("2022-01-01", Parse("2020-01-01").AddYears(2).ToDateString())
		s.Equal("2023-03-01", Parse("2020-02-29").AddYears(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYearsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddYearsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddYearsNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0003-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddYearsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddYearsNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddYearsNoOverflow(0).ToDateString())
		s.Equal("2021-01-01", Parse("2020-01-01").AddYearsNoOverflow(1).ToDateString())
		s.Equal("2022-01-01", Parse("2020-01-01").AddYearsNoOverflow(2).ToDateString())
		s.Equal("2023-02-28", Parse("2020-02-29").AddYearsNoOverflow(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddYear()
		s.False(c.HasError())
		s.Equal("0002-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddYear()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2021-01-01", Parse("2020-01-01").AddYear().ToDateString())
		s.Equal("2021-02-28", Parse("2020-02-28").AddYear().ToDateString())
		s.Equal("2021-03-01", Parse("2020-02-29").AddYear().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYearNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddYearNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddYearNoOverflow()
		s.False(c.HasError())
		s.Equal("0002-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddYearNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddYearNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2021-01-01", Parse("2020-01-01").AddYearNoOverflow().ToDateString())
		s.Equal("2021-02-28", Parse("2020-02-28").AddYearNoOverflow().ToDateString())
		s.Equal("2021-02-28", Parse("2020-02-29").AddYearNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYears() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubYears(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubYears(2)
		s.False(c.HasError())
		s.Equal("-0001-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubYears(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubYears(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubYears(0).ToDateString())
		s.Equal("2019-01-01", Parse("2020-01-01").SubYears(1).ToDateString())
		s.Equal("2018-01-01", Parse("2020-01-01").SubYears(2).ToDateString())
		s.Equal("2017-03-01", Parse("2020-02-29").SubYears(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYearsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubYearsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubYearsNoOverflow(2)
		s.False(c.HasError())
		s.Equal("-0001-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubYearsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubYearsNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubYearsNoOverflow(0).ToDateString())
		s.Equal("2019-01-01", Parse("2020-01-01").SubYearsNoOverflow(1).ToDateString())
		s.Equal("2018-01-01", Parse("2020-01-01").SubYearsNoOverflow(2).ToDateString())
		s.Equal("2017-02-28", Parse("2020-02-29").SubYearsNoOverflow(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubYear()
		s.False(c.HasError())
		s.Equal("0000-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubYear()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubYear()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-01-01", Parse("2020-01-01").SubYear().ToDateString())
		s.Equal("2019-02-28", Parse("2020-02-28").SubYear().ToDateString())
		s.Equal("2019-03-01", Parse("2020-02-29").SubYear().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYearNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubYearNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubYearNoOverflow()
		s.False(c.HasError())
		s.Equal("0000-01-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubYearNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubYearNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-01-01", Parse("2020-01-01").SubYearNoOverflow().ToDateString())
		s.Equal("2019-02-28", Parse("2020-02-28").SubYearNoOverflow().ToDateString())
		s.Equal("2019-02-28", Parse("2020-02-29").SubYearNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuarters() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddQuarters(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddQuarters(2)
		s.False(c.HasError())
		s.Equal("0001-07-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddQuarters(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddQuarters(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddQuarters(0).ToDateString())
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuarters(1).ToDateString())
		s.Equal("2020-07-01", Parse("2020-01-01").AddQuarters(2).ToDateString())
		s.Equal("2020-11-29", Parse("2020-02-29").AddQuarters(3).ToDateString())
		s.Equal("2021-03-03", Parse("2020-08-31").AddQuarters(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuartersNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddQuartersNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddQuartersNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0001-07-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddQuartersNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddQuartersNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddQuartersNoOverflow(0).ToDateString())
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuartersNoOverflow(1).ToDateString())
		s.Equal("2020-07-01", Parse("2020-01-01").AddQuartersNoOverflow(2).ToDateString())
		s.Equal("2020-11-29", Parse("2020-02-29").AddQuartersNoOverflow(3).ToDateString())
		s.Equal("2021-02-28", Parse("2020-08-31").AddQuartersNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddQuarter()
		s.False(c.HasError())
		s.Equal("0001-04-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddQuarter()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuarter().ToDateString())
		s.Equal("2020-05-28", Parse("2020-02-28").AddQuarter().ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddQuarter().ToDateString())
		s.Equal("2021-03-02", Parse("2020-11-30").AddQuarter().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuarterNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddQuarterNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddQuarterNoOverflow()
		s.False(c.HasError())
		s.Equal("0001-04-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddQuarterNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddQuarterNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuarterNoOverflow().ToDateString())
		s.Equal("2020-05-28", Parse("2020-02-28").AddQuarterNoOverflow().ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddQuarterNoOverflow().ToDateString())
		s.Equal("2021-02-28", Parse("2020-11-30").AddQuarterNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuarters() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubQuarters(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubQuarters(2)
		s.False(c.HasError())
		s.Equal("0000-07-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubQuarters(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubQuarters(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubQuarters(0).ToDateString())
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuarters(1).ToDateString())
		s.Equal("2019-07-01", Parse("2020-01-01").SubQuarters(2).ToDateString())
		s.Equal("2019-05-29", Parse("2020-02-29").SubQuarters(3).ToDateString())
		s.Equal("2020-03-02", Parse("2020-08-31").SubQuarters(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuartersNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubQuartersNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubQuartersNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0000-07-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubQuartersNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubQuartersNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubQuartersNoOverflow(0).ToDateString())
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuartersNoOverflow(1).ToDateString())
		s.Equal("2019-07-01", Parse("2020-01-01").SubQuartersNoOverflow(2).ToDateString())
		s.Equal("2019-05-29", Parse("2020-02-29").SubQuartersNoOverflow(3).ToDateString())
		s.Equal("2020-02-29", Parse("2020-08-31").SubQuartersNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubQuarter()
		s.False(c.HasError())
		s.Equal("0000-10-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubQuarter()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubQuarter()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuarter().ToDateString())
		s.Equal("2019-11-28", Parse("2020-02-28").SubQuarter().ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubQuarter().ToDateString())
		s.Equal("2020-08-30", Parse("2020-11-30").SubQuarter().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuarterNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubQuarterNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubQuarterNoOverflow()
		s.False(c.HasError())
		s.Equal("0000-10-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubQuarterNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubQuarterNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuarterNoOverflow().ToDateString())
		s.Equal("2019-11-28", Parse("2020-02-28").SubQuarterNoOverflow().ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubQuarterNoOverflow().ToDateString())
		s.Equal("2020-08-30", Parse("2020-11-30").SubQuarterNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMonths(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMonths(2)
		s.False(c.HasError())
		s.Equal("0001-03-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMonths(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMonths(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddMonths(0).ToDateString())
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonths(1).ToDateString())
		s.Equal("2020-03-01", Parse("2020-01-01").AddMonths(2).ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddMonths(3).ToDateString())
		s.Equal("2020-10-31", Parse("2020-08-31").AddMonths(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonthsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMonthsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMonthsNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0001-03-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMonthsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMonthsNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddMonthsNoOverflow(0).ToDateString())
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonthsNoOverflow(1).ToDateString())
		s.Equal("2020-03-01", Parse("2020-01-01").AddMonthsNoOverflow(2).ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddMonthsNoOverflow(3).ToDateString())
		s.Equal("2020-10-31", Parse("2020-08-31").AddMonthsNoOverflow(2).ToDateString())
	})

	// https://github.com/dromara/carbon/issues/303
	s.Run("issue303", func() {
		c := CreateFromDate(2025, 6, 11, "Asia/Shanghai")
		c = c.SetWeekStartsAt(time.Sunday)
		c = c.SubMonthNoOverflow()
		c = c.StartOfWeek()
		s.Equal("2025-05-11", c.ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMonth()
		s.False(c.HasError())
		s.Equal("0001-02-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMonth()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonth().ToDateString())
		s.Equal("2020-03-28", Parse("2020-02-28").AddMonth().ToDateString())
		s.Equal("2020-03-29", Parse("2020-02-29").AddMonth().ToDateString())
		s.Equal("2020-12-30", Parse("2020-11-30").AddMonth().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonthNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMonthNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMonthNoOverflow()
		s.False(c.HasError())
		s.Equal("0001-02-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMonthNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMonthNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonthNoOverflow().ToDateString())
		s.Equal("2020-03-28", Parse("2020-02-28").AddMonthNoOverflow().ToDateString())
		s.Equal("2020-03-29", Parse("2020-02-29").AddMonthNoOverflow().ToDateString())
		s.Equal("2020-12-30", Parse("2020-11-30").AddMonthNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMonths(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMonths(2)
		s.False(c.HasError())
		s.Equal("0000-11-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMonths(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMonths(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubMonths(0).ToDateString())
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonths(1).ToDateString())
		s.Equal("2019-11-01", Parse("2020-01-01").SubMonths(2).ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubMonths(3).ToDateString())
		s.Equal("2020-07-01", Parse("2020-08-31").SubMonths(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonthsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMonthsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMonthsNoOverflow(2)
		s.False(c.HasError())
		s.Equal("0000-11-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMonthsNoOverflow(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMonthsNoOverflow(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubMonthsNoOverflow(0).ToDateString())
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonthsNoOverflow(1).ToDateString())
		s.Equal("2019-11-01", Parse("2020-01-01").SubMonthsNoOverflow(2).ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubMonthsNoOverflow(3).ToDateString())
		s.Equal("2020-06-30", Parse("2020-08-31").SubMonthsNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMonth()
		s.False(c.HasError())
		s.Equal("0000-12-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMonth()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMonth()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonth().ToDateString())
		s.Equal("2020-01-28", Parse("2020-02-28").SubMonth().ToDateString())
		s.Equal("2020-01-29", Parse("2020-02-29").SubMonth().ToDateString())
		s.Equal("2020-10-30", Parse("2020-11-30").SubMonth().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonthNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMonthNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMonthNoOverflow()
		s.False(c.HasError())
		s.Equal("0000-12-01", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMonthNoOverflow()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMonthNoOverflow()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonthNoOverflow().ToDateString())
		s.Equal("2020-01-28", Parse("2020-02-28").SubMonthNoOverflow().ToDateString())
		s.Equal("2020-01-29", Parse("2020-02-29").SubMonthNoOverflow().ToDateString())
		s.Equal("2020-10-30", Parse("2020-11-30").SubMonthNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddWeeks() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddWeeks(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddWeeks(2)
		s.False(c.HasError())
		s.Equal("0001-01-15", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddWeeks(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddWeeks(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddWeeks(0).ToDateString())
		s.Equal("2020-01-08", Parse("2020-01-01").AddWeeks(1).ToDateString())
		s.Equal("2020-01-15", Parse("2020-01-01").AddWeeks(2).ToDateString())
		s.Equal("2020-03-21", Parse("2020-02-29").AddWeeks(3).ToDateString())
		s.Equal("2020-09-14", Parse("2020-08-31").AddWeeks(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddWeek()
		s.False(c.HasError())
		s.Equal("0001-01-08", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddWeek()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-08", Parse("2020-01-01").AddWeek().ToDateString())
		s.Equal("2020-03-06", Parse("2020-02-28").AddWeek().ToDateString())
		s.Equal("2020-03-07", Parse("2020-02-29").AddWeek().ToDateString())
		s.Equal("2020-12-07", Parse("2020-11-30").AddWeek().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubWeeks() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubWeeks(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubWeeks(2)
		s.False(c.HasError())
		s.Equal("0000-12-18", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubWeeks(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubWeeks(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubWeeks(0).ToDateString())
		s.Equal("2019-12-25", Parse("2020-01-01").SubWeeks(1).ToDateString())
		s.Equal("2019-12-18", Parse("2020-01-01").SubWeeks(2).ToDateString())
		s.Equal("2020-02-08", Parse("2020-02-29").SubWeeks(3).ToDateString())
		s.Equal("2020-08-17", Parse("2020-08-31").SubWeeks(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubWeek()
		s.False(c.HasError())
		s.Equal("0000-12-25", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubWeek()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubWeek()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-25", Parse("2020-01-01").SubWeek().ToDateString())
		s.Equal("2020-02-21", Parse("2020-02-28").SubWeek().ToDateString())
		s.Equal("2020-02-22", Parse("2020-02-29").SubWeek().ToDateString())
		s.Equal("2020-11-23", Parse("2020-11-30").SubWeek().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDays(2)
		s.False(c.HasError())
		s.Equal("0001-01-03", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDays(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddDays(0).ToDateString())
		s.Equal("2020-01-02", Parse("2020-01-01").AddDays(1).ToDateString())
		s.Equal("2020-01-03", Parse("2020-01-01").AddDays(2).ToDateString())
		s.Equal("2020-03-03", Parse("2020-02-29").AddDays(3).ToDateString())
		s.Equal("2020-09-02", Parse("2020-08-31").AddDays(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddDay()
		s.False(c.HasError())
		s.Equal("0001-01-02", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddDay()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-02", Parse("2020-01-01").AddDay().ToDateString())
		s.Equal("2020-02-29", Parse("2020-02-28").AddDay().ToDateString())
		s.Equal("2020-03-01", Parse("2020-02-29").AddDay().ToDateString())
		s.Equal("2020-12-01", Parse("2020-11-30").AddDay().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDays(2)
		s.False(c.HasError())
		s.Equal("0000-12-30", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDays(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubDays(0).ToDateString())
		s.Equal("2019-12-31", Parse("2020-01-01").SubDays(1).ToDateString())
		s.Equal("2019-12-30", Parse("2020-01-01").SubDays(2).ToDateString())
		s.Equal("2020-02-26", Parse("2020-02-29").SubDays(3).ToDateString())
		s.Equal("2020-08-29", Parse("2020-08-31").SubDays(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubDay()
		s.False(c.HasError())
		s.Equal("0000-12-31", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubDay()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubDay()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-31", Parse("2020-01-01").SubDay().ToDateString())
		s.Equal("2020-02-27", Parse("2020-02-28").SubDay().ToDateString())
		s.Equal("2020-02-28", Parse("2020-02-29").SubDay().ToDateString())
		s.Equal("2020-11-29", Parse("2020-11-30").SubDay().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddWeekdays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddWeekdays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddWeekdays(2)
		s.False(c.HasError())
		s.Equal("0001-01-03", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddWeekdays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddWeekdays(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		// 2026-09-21 is Monday, 2026-09-26 is Saturday
		s.Equal("2026-09-22", Parse("2026-09-21").AddWeekdays(1).ToDateString())
		s.Equal("2026-09-25", Parse("2026-09-21").AddWeekdays(4).ToDateString())
		s.Equal("2026-09-28", Parse("2026-09-21").AddWeekdays(5).ToDateString())
	})

	s.Run("negative weekdays", func() {
		s.Equal("2026-09-18", Parse("2026-09-21").AddWeekdays(-1).ToDateString())
		s.Equal("2026-09-16", Parse("2026-09-21").AddWeekdays(-3).ToDateString())
		s.Equal("2026-09-14", Parse("2026-09-21").AddWeekdays(-5).ToDateString())
	})

	s.Run("custom weekend days", func() {
		fridayAndSaturday := []Weekday{Friday, Saturday}
		s.Equal("2026-09-22", Parse("2026-09-21").SetWeekendDays(fridayAndSaturday).AddWeekdays(1).ToDateString())
		s.Equal("2026-09-27", Parse("2026-09-24").SetWeekendDays(fridayAndSaturday).AddWeekdays(1).ToDateString())
		s.Equal("2026-09-27", Parse("2026-09-25").SetWeekendDays(fridayAndSaturday).AddWeekdays(1).ToDateString())
		s.Equal("2026-09-28", Parse("2026-09-21").SetWeekendDays(fridayAndSaturday).AddWeekdays(5).ToDateString())
		s.Equal("2026-09-20", Parse("2026-09-21").SetWeekendDays(fridayAndSaturday).AddWeekdays(-1).ToDateString())
		s.Equal("2026-09-14", Parse("2026-09-21").SetWeekendDays(fridayAndSaturday).AddWeekdays(-5).ToDateString())

		sundayOnly := []Weekday{Sunday}
		s.Equal("2026-09-22", Parse("2026-09-21").SetWeekendDays(sundayOnly).AddWeekdays(1).ToDateString())
		s.Equal("2026-09-26", Parse("2026-09-25").SetWeekendDays(sundayOnly).AddWeekdays(1).ToDateString())
		s.Equal("2026-09-26", Parse("2026-09-21").SetWeekendDays(sundayOnly).AddWeekdays(5).ToDateString())
		s.Equal("2026-09-19", Parse("2026-09-21").SetWeekendDays(sundayOnly).AddWeekdays(-1).ToDateString())
		s.Equal("2026-09-15", Parse("2026-09-21").SetWeekendDays(sundayOnly).AddWeekdays(-5).ToDateString())
	})

	s.Run("no weekend days", func() {
		none := []Weekday{}
		s.Equal("2026-09-26", Parse("2026-09-21").SetWeekendDays(none).AddWeekdays(5).ToDateString())
		s.Equal("2026-09-16", Parse("2026-09-21").SetWeekendDays(none).AddWeekdays(-5).ToDateString())
	})

	s.Run("starting on a weekend day never lands on one", func() {
		fridayAndSaturday := []Weekday{Friday, Saturday}
		c := Parse("2026-09-25").SetWeekendDays(fridayAndSaturday).AddWeekdays(5)
		s.Equal("2026-10-01", c.ToDateString())
		s.False(c.IsWeekend())
	})

	s.Run("the time of day survives a daylight saving transition", func() {
		// in America/New_York the clocks jump from 02:00 to 03:00 on Sunday 2026-03-08,
		// which a step from Friday to Monday passes over
		tz := "America/New_York"
		s.Equal("2026-03-09 02:30:00", Parse("2026-03-06 02:30:00", tz).AddWeekdays(1).ToDateTimeString())
		s.Equal("2026-03-09 10:00:00", Parse("2026-03-06 10:00:00", tz).AddWeekdays(1).ToDateTimeString())
		// and on Sunday 2026-11-01 they go back from 02:00 to 01:00
		s.Equal("2026-11-02 01:30:00", Parse("2026-10-30 01:30:00", tz).AddWeekdays(1).ToDateTimeString())
	})

	s.Run("weekend days outside the week are ignored", func() {
		// Saturday is the only day left as a weekday, so it must stay reachable
		weekendDays := []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Weekday(42)}
		s.Equal("2026-09-26", Parse("2026-09-21").SetWeekendDays(weekendDays).AddWeekdays(1).ToDateString())
	})

	s.Run("every day is a weekend day", func() {
		weekendDays := []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}
		s.Equal("2026-09-28", Parse("2026-09-21").SetWeekendDays(weekendDays).AddWeekdays(1).ToDateString())
		s.Equal("2026-10-12", Parse("2026-09-21").SetWeekendDays(weekendDays).AddWeekdays(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddWeekday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddWeekday()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddWeekday()
		s.False(c.HasError())
		s.Equal("0001-01-02", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddWeekday()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddWeekday()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2026-09-22", Parse("2026-09-21").AddWeekday().ToDateString())
		s.Equal("2026-09-28", Parse("2026-09-25").AddWeekday().ToDateString())
		s.Equal("2026-09-28", Parse("2026-09-26").AddWeekday().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubWeekdays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubWeekdays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubWeekdays(2)
		s.False(c.HasError())
		s.Equal("0000-12-28", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubWeekdays(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubWeekdays(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2026-09-18", Parse("2026-09-21").SubWeekdays(1).ToDateString())
		s.Equal("2026-09-16", Parse("2026-09-21").SubWeekdays(3).ToDateString())
		s.Equal("2026-09-14", Parse("2026-09-21").SubWeekdays(5).ToDateString())
		s.Equal("2026-09-25", Parse("2026-09-26").SubWeekdays(1).ToDateString())
		s.Equal("2026-09-21", Parse("2026-09-26").SubWeekdays(5).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubWeekday() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubWeekday()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubWeekday()
		s.False(c.HasError())
		s.Equal("0000-12-29", c.ToDateString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubWeekday()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubWeekday()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2026-09-18", Parse("2026-09-21").SubWeekday().ToDateString())
		s.Equal("2026-09-25", Parse("2026-09-26").SubWeekday().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddHours() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddHours(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddHours(2)
		s.False(c.HasError())
		s.Equal("0001-01-01 02:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddHours(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddHours(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(0).ToString())
		s.Equal("2020-01-01 14:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(1).ToString())
		s.Equal("2020-01-01 15:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddHour()
		s.False(c.HasError())
		s.Equal("0001-01-01 01:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddHour()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").AddHour()
		s.False(c.HasError())
		s.Equal("2020-08-05 14:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubHours() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubHours(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubHours(2)
		s.False(c.HasError())
		s.Equal("0000-12-31 22:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubHours(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubHours(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(0).ToString())
		s.Equal("2020-01-01 12:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(1).ToString())
		s.Equal("2020-01-01 11:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubHour()
		s.False(c.HasError())
		s.Equal("0000-12-31 23:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubHour()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubHour()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").SubHour()
		s.False(c.HasError())
		s.Equal("2020-08-05 12:14:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMinutes() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMinutes(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMinutes(2)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:02:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMinutes(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMinutes(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(0).ToString())
		s.Equal("2020-01-01 13:15:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(1).ToString())
		s.Equal("2020-01-01 13:16:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMinute()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:01:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMinute()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").AddMinute()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:15:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMinutes() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMinutes(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMinutes(2)
		s.False(c.HasError())
		s.Equal("0000-12-31 23:58:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMinutes(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMinutes(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(0).ToString())
		s.Equal("2020-01-01 13:13:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(1).ToString())
		s.Equal("2020-01-01 13:12:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMinute()
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMinute()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMinute()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").SubMinute()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:13:15 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddSeconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddSeconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddSeconds(2)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:02 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddSeconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddSeconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(0).ToString())
		s.Equal("2020-01-01 13:14:16 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(1).ToString())
		s.Equal("2020-01-01 13:14:17 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddSecond()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:01 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddSecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").AddSecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:16 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubSeconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubSeconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubSeconds(2)
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:58 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubSeconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubSeconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(0).ToString())
		s.Equal("2020-01-01 13:14:14 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(1).ToString())
		s.Equal("2020-01-01 13:14:13 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubSecond()
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:59 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubSecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubSecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").SubSecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:14 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMilliseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMilliseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMilliseconds(2)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.002 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMilliseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMilliseconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(0).ToString())
		s.Equal("2020-01-01 13:14:15.001 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(1).ToString())
		s.Equal("2020-01-01 13:14:15.002 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMillisecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMillisecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMillisecond()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.001 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMillisecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMillisecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").AddMillisecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.001 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMilliseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMilliseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMilliseconds(2)
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:59.998 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMilliseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMilliseconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(0).ToString())
		s.Equal("2020-01-01 13:14:14.999 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(1).ToString())
		s.Equal("2020-01-01 13:14:14.998 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMillisecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMillisecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMillisecond()
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:59.999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMillisecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMillisecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").SubMillisecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:14.999 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMicroseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMicroseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMicroseconds(2)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.000002 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMicroseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMicroseconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(0).ToString())
		s.Equal("2020-01-01 13:14:15.000001 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(1).ToString())
		s.Equal("2020-01-01 13:14:15.000002 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMicrosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddMicrosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddMicrosecond()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.000001 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddMicrosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddMicrosecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").AddMicrosecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.000001 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMicroseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMicroseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.999998 +0000 UTC", NewCarbon().SubMicroseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMicroseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMicroseconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(0).ToString())
		s.Equal("2020-01-01 13:14:14.999999 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(1).ToString())
		s.Equal("2020-01-01 13:14:14.999998 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMicrosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubMicrosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubMicrosecond()
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:59.999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubMicrosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubMicrosecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").SubMicrosecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:14.999999 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddNanoseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddNanoseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddNanoseconds(2)
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.000000002 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddNanoseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddNanoseconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(0).ToString())
		s.Equal("2020-01-01 13:14:15.000000001 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(1).ToString())
		s.Equal("2020-01-01 13:14:15.000000002 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddNanosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.AddNanosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().AddNanosecond()
		s.False(c.HasError())
		s.Equal("0001-01-01 00:00:00.000000001 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").AddNanosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").AddNanosecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").AddNanosecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:15.000000001 +0000 UTC", c.ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubNanoseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubNanoseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubNanoseconds(2)
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:59.999999998 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubNanoseconds(2)
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubNanoseconds(2)
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(0).ToString())
		s.Equal("2020-01-01 13:14:14.999999999 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(1).ToString())
		s.Equal("2020-01-01 13:14:14.999999998 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubNanosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		c = c.SubNanosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().SubNanosecond()
		s.False(c.HasError())
		s.Equal("0000-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").SubNanosecond()
		s.False(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").SubNanosecond()
		s.True(c.HasError())
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05 13:14:15").SubNanosecond()
		s.False(c.HasError())
		s.Equal("2020-08-05 13:14:14.999999999 +0000 UTC", c.ToString())
	})
}
