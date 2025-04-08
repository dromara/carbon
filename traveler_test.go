package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Now(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.Empty(t, Now("").ToString())
		assert.Empty(t, Now("0").ToString())
		assert.Empty(t, Now("xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, time.Now().Format(DateLayout), Now().Layout(DateLayout, Local))
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, time.Now().In(time.UTC).Format(DateLayout), Now(UTC).Layout(DateLayout))
	})

	t.Run("frozen time", func(t *testing.T) {
		SetTestNow(Parse("2020-08-05"))
		assert.Equal(t, "2020-08-05", Now(UTC).Layout(DateLayout))
		CleanTestNow()
	})
}

func TestCarbon_Tomorrow(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.Empty(t, Tomorrow("").ToString())
		assert.Empty(t, Tomorrow("0").ToString())
		assert.Empty(t, Tomorrow("xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, time.Now().Add(time.Hour*24).Format(DateLayout), Tomorrow().Layout(DateLayout, Local))
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, time.Now().Add(time.Hour*24).In(time.UTC).Format(DateLayout), Tomorrow(UTC).Layout(DateLayout))
	})

	t.Run("frozen time", func(t *testing.T) {
		SetTestNow(Parse("2020-08-05"))
		assert.Equal(t, "2020-08-06", Tomorrow(UTC).Layout(DateLayout))
		CleanTestNow()
	})
}

func TestCarbon_Yesterday(t *testing.T) {
	t.Run("invalid timezone", func(t *testing.T) {
		assert.Empty(t, Yesterday("").ToString())
		assert.Empty(t, Yesterday("0").ToString())
		assert.Empty(t, Yesterday("xxx").ToString())
	})

	t.Run("without timezone", func(t *testing.T) {
		assert.Equal(t, time.Now().Add(time.Hour*-24).Format(DateLayout), Yesterday().Layout(DateLayout, Local))
	})

	t.Run("with timezone", func(t *testing.T) {
		assert.Equal(t, time.Now().Add(time.Hour*-24).In(time.UTC).Format(DateLayout), Yesterday(UTC).Layout(DateLayout))
	})

	t.Run("frozen time", func(t *testing.T) {
		SetTestNow(Parse("2020-08-05"))
		assert.Equal(t, "2020-08-04", Yesterday(UTC).Layout(DateLayout))
		CleanTestNow()
	})
}

func TestCarbon_AddDuration(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 10:00:00 +0000 UTC", NewCarbon().AddDuration("10h").ToString())
	})

	t.Run("invalid duration", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").AddDuration("").ToString())
		assert.Empty(t, Parse("2020-08-05").AddDuration("xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDuration("10h").ToString())
		assert.Empty(t, Parse("xxx").AddDuration("10h").ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 23:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10h").ToString())
		assert.Equal(t, "2020-01-01 23:44:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10.5h").ToString())
		assert.Equal(t, "2020-01-01 13:24:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10m").ToString())
		assert.Equal(t, "2020-01-01 13:24:45 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10.5m").ToString())
	})
}

func TestCarbon_SubDuration(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 14:00:00 +0000 UTC", NewCarbon().SubDuration("10h").ToString())
	})

	t.Run("invalid duration", func(t *testing.T) {
		assert.Empty(t, Parse("2020-08-05").SubDuration("").ToString())
		assert.Empty(t, Parse("2020-08-05").SubDuration("xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDuration("10h").ToString())
		assert.Empty(t, Parse("0").SubDuration("10h").ToString())
		assert.Empty(t, Parse("xxx").SubDuration("10h").ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 03:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10h").ToString())
		assert.Equal(t, "2020-01-01 02:44:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10.5h").ToString())
		assert.Equal(t, "2020-01-01 13:04:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10m").ToString())
		assert.Equal(t, "2020-01-01 13:03:45 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10.5m").ToString())
	})
}

func TestCarbon_AddCenturies(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0201-01-01 00:00:00 +0000 UTC", NewCarbon().AddCenturies(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddCenturies(2).ToString())
		assert.Empty(t, Parse("0").AddCenturies(2).ToString())
		assert.Empty(t, Parse("xxx").AddCenturies(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(0).ToString())
		assert.Equal(t, "2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(1).ToString())
		assert.Equal(t, "2220-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(2).ToString())
	})
}

func TestCarbon_AddCenturiesNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0201-01-01 00:00:00 +0000 UTC", NewCarbon().AddCenturiesNoOverflow(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddCenturiesNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").AddCenturiesNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").AddCenturiesNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(0).ToString())
		assert.Equal(t, "2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(1).ToString())
		assert.Equal(t, "2220-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(2).ToString())
	})
}

func TestCarbon_AddCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0101-01-01 00:00:00 +0000 UTC", NewCarbon().AddCentury().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddCentury().ToString())
		assert.Empty(t, Parse("0").AddCentury().ToString())
		assert.Empty(t, Parse("xxx").AddCentury().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCentury().ToString())
	})
}

func TestCarbon_AddCenturyNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0101-01-01 00:00:00 +0000 UTC", NewCarbon().AddCenturyNoOverflow().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddCenturyNoOverflow().ToString())
		assert.Empty(t, Parse("0").AddCenturyNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").AddCenturyNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturyNoOverflow().ToString())
	})
}

func TestCarbon_SubCenturies(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0199-01-01 00:00:00 +0000 UTC", NewCarbon().SubCenturies(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubCenturies(2).ToString())
		assert.Empty(t, Parse("0").SubCenturies(2).ToString())
		assert.Empty(t, Parse("xxx").SubCenturies(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(0).ToString())
		assert.Equal(t, "1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(1).ToString())
		assert.Equal(t, "1820-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(2).ToString())
	})
}

func TestCarbon_SubCenturiesNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0199-01-01 00:00:00 +0000 UTC", NewCarbon().SubCenturiesNoOverflow(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubCenturiesNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SubCenturiesNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SubCenturiesNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(0).ToString())
		assert.Equal(t, "1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(1).ToString())
		assert.Equal(t, "1820-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(2).ToString())
	})
}

func TestCarbon_SubCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0099-01-01 00:00:00 +0000 UTC", NewCarbon().SubCentury().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubCentury().ToString())
		assert.Empty(t, Parse("0").SubCentury().ToString())
		assert.Empty(t, Parse("xxx").SubCentury().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCentury().ToString())
	})
}

func TestCarbon_SubCenturyNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0099-01-01 00:00:00 +0000 UTC", NewCarbon().SubCenturyNoOverflow().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubCenturyNoOverflow().ToString())
		assert.Empty(t, Parse("0").SubCenturyNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").SubCenturyNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturyNoOverflow().ToString())
	})
}

func TestCarbon_AddDecades(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0021-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecades(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDecades(2).ToString())
		assert.Empty(t, Parse("0").AddDecades(2).ToString())
		assert.Empty(t, Parse("xxx").AddDecades(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(0).ToString())
		assert.Equal(t, "2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(1).ToString())
		assert.Equal(t, "2040-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(2).ToString())
	})
}

func TestCarbon_AddDecadesNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0021-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecadesNoOverflow(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDecadesNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").AddDecadesNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").AddDecadesNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(0).ToString())
		assert.Equal(t, "2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(1).ToString())
		assert.Equal(t, "2040-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(2).ToString())
	})
}

func TestCarbon_AddDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0011-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecade().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDecade().ToString())
		assert.Empty(t, Parse("0").AddDecade().ToString())
		assert.Empty(t, Parse("xxx").AddDecade().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecade().ToString())
	})
}

func TestCarbon_AddDecadeNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0011-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecadeNoOverflow().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDecadeNoOverflow().ToString())
		assert.Empty(t, Parse("0").AddDecadeNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").AddDecadeNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadeNoOverflow().ToString())
	})
}

func TestCarbon_SubDecades(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0019-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecades(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDecades(2).ToString())
		assert.Empty(t, Parse("0").SubDecades(2).ToString())
		assert.Empty(t, Parse("xxx").SubDecades(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(0).ToString())
		assert.Equal(t, "2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(1).ToString())
		assert.Equal(t, "2000-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(2).ToString())
	})
}

func TestCarbon_SubDecadesNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0019-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecadesNoOverflow(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDecadesNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SubDecadesNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SubDecadesNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(0).ToString())
		assert.Equal(t, "2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(1).ToString())
		assert.Equal(t, "2000-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(2).ToString())
	})
}

func TestCarbon_SubDecade(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0009-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecade().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDecade().ToString())
		assert.Empty(t, Parse("0").SubDecade().ToString())
		assert.Empty(t, Parse("xxx").SubDecade().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecade().ToString())
	})
}

func TestCarbon_SubDecadeNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0009-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecadeNoOverflow().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDecadeNoOverflow().ToString())
		assert.Empty(t, Parse("0").SubDecadeNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").SubDecadeNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2010-01-01", Parse("2020-01-01").SubDecadeNoOverflow().ToDateString())
	})
}

func TestCarbon_AddYears(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0003-01-01", NewCarbon().AddYears(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddYears(2).ToString())
		assert.Empty(t, Parse("0").AddYears(2).ToString())
		assert.Empty(t, Parse("xxx").AddYears(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddYears(0).ToDateString())
		assert.Equal(t, "2021-01-01", Parse("2020-01-01").AddYears(1).ToDateString())
		assert.Equal(t, "2022-01-01", Parse("2020-01-01").AddYears(2).ToDateString())
		assert.Equal(t, "2023-03-01", Parse("2020-02-29").AddYears(3).ToDateString())
	})
}

func TestCarbon_AddYearsNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0003-01-01", NewCarbon().AddYearsNoOverflow(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddYearsNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").AddYearsNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").AddYearsNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddYearsNoOverflow(0).ToDateString())
		assert.Equal(t, "2021-01-01", Parse("2020-01-01").AddYearsNoOverflow(1).ToDateString())
		assert.Equal(t, "2022-01-01", Parse("2020-01-01").AddYearsNoOverflow(2).ToDateString())
		assert.Equal(t, "2023-02-28", Parse("2020-02-29").AddYearsNoOverflow(3).ToDateString())
	})
}

func TestCarbon_AddYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0002-01-01", NewCarbon().AddYear().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddYear().ToString())
		assert.Empty(t, Parse("0").AddYear().ToString())
		assert.Empty(t, Parse("xxx").AddYear().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2021-01-01", Parse("2020-01-01").AddYear().ToDateString())
		assert.Equal(t, "2021-02-28", Parse("2020-02-28").AddYear().ToDateString())
		assert.Equal(t, "2021-03-01", Parse("2020-02-29").AddYear().ToDateString())
	})
}

func TestCarbon_AddYearNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0002-01-01", NewCarbon().AddYearNoOverflow().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddYearNoOverflow().ToString())
		assert.Empty(t, Parse("0").AddYearNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").AddYearNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2021-01-01", Parse("2020-01-01").AddYearNoOverflow().ToDateString())
		assert.Equal(t, "2021-02-28", Parse("2020-02-28").AddYearNoOverflow().ToDateString())
		assert.Equal(t, "2021-02-28", Parse("2020-02-29").AddYearNoOverflow().ToDateString())
	})
}

func TestCarbon_SubYears(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0001-01-01", NewCarbon().SubYears(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubYears(2).ToString())
		assert.Empty(t, Parse("0").SubYears(2).ToString())
		assert.Empty(t, Parse("xxx").SubYears(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubYears(0).ToDateString())
		assert.Equal(t, "2019-01-01", Parse("2020-01-01").SubYears(1).ToDateString())
		assert.Equal(t, "2018-01-01", Parse("2020-01-01").SubYears(2).ToDateString())
		assert.Equal(t, "2017-03-01", Parse("2020-02-29").SubYears(3).ToDateString())
	})
}

func TestCarbon_SubYearsNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "-0001-01-01", NewCarbon().SubYearsNoOverflow(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubYearsNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SubYearsNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SubYearsNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubYearsNoOverflow(0).ToDateString())
		assert.Equal(t, "2019-01-01", Parse("2020-01-01").SubYearsNoOverflow(1).ToDateString())
		assert.Equal(t, "2018-01-01", Parse("2020-01-01").SubYearsNoOverflow(2).ToDateString())
		assert.Equal(t, "2017-02-28", Parse("2020-02-29").SubYearsNoOverflow(3).ToDateString())
	})
}

func TestCarbon_SubYear(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-01-01", NewCarbon().SubYear().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubYear().ToString())
		assert.Empty(t, Parse("0").SubYear().ToString())
		assert.Empty(t, Parse("xxx").SubYear().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-01-01", Parse("2020-01-01").SubYear().ToDateString())
		assert.Equal(t, "2019-02-28", Parse("2020-02-28").SubYear().ToDateString())
		assert.Equal(t, "2019-03-01", Parse("2020-02-29").SubYear().ToDateString())
	})
}

func TestCarbon_SubYearNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-01-01", NewCarbon().SubYearNoOverflow().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubYearNoOverflow().ToString())
		assert.Empty(t, Parse("0").SubYearNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").SubYearNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-01-01", Parse("2020-01-01").SubYearNoOverflow().ToDateString())
		assert.Equal(t, "2019-02-28", Parse("2020-02-28").SubYearNoOverflow().ToDateString())
		assert.Equal(t, "2019-02-28", Parse("2020-02-29").SubYearNoOverflow().ToDateString())
	})
}

func TestCarbon_AddQuarters(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-07-01", NewCarbon().AddQuarters(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddQuarters(2).ToString())
		assert.Empty(t, Parse("0").AddQuarters(2).ToString())
		assert.Empty(t, Parse("xxx").AddQuarters(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddQuarters(0).ToDateString())
		assert.Equal(t, "2020-04-01", Parse("2020-01-01").AddQuarters(1).ToDateString())
		assert.Equal(t, "2020-07-01", Parse("2020-01-01").AddQuarters(2).ToDateString())
		assert.Equal(t, "2020-11-29", Parse("2020-02-29").AddQuarters(3).ToDateString())
		assert.Equal(t, "2021-03-03", Parse("2020-08-31").AddQuarters(2).ToDateString())
	})
}

func TestCarbon_AddQuartersNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-07-01", NewCarbon().AddQuartersNoOverflow(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddQuartersNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").AddQuartersNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").AddQuartersNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddQuartersNoOverflow(0).ToDateString())
		assert.Equal(t, "2020-04-01", Parse("2020-01-01").AddQuartersNoOverflow(1).ToDateString())
		assert.Equal(t, "2020-07-01", Parse("2020-01-01").AddQuartersNoOverflow(2).ToDateString())
		assert.Equal(t, "2020-11-29", Parse("2020-02-29").AddQuartersNoOverflow(3).ToDateString())
		assert.Equal(t, "2021-02-28", Parse("2020-08-31").AddQuartersNoOverflow(2).ToDateString())
	})
}

func TestCarbon_AddQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-04-01", NewCarbon().AddQuarter().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddQuarter().ToString())
		assert.Empty(t, Parse("0").AddQuarter().ToString())
		assert.Empty(t, Parse("xxx").AddQuarter().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-04-01", Parse("2020-01-01").AddQuarter().ToDateString())
		assert.Equal(t, "2020-05-28", Parse("2020-02-28").AddQuarter().ToDateString())
		assert.Equal(t, "2020-05-29", Parse("2020-02-29").AddQuarter().ToDateString())
		assert.Equal(t, "2021-03-02", Parse("2020-11-30").AddQuarter().ToDateString())
	})
}

func TestCarbon_AddQuarterNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-04-01", NewCarbon().AddQuarterNoOverflow().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddQuarterNoOverflow().ToString())
		assert.Empty(t, Parse("0").AddQuarterNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").AddQuarterNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-04-01", Parse("2020-01-01").AddQuarterNoOverflow().ToDateString())
		assert.Equal(t, "2020-05-28", Parse("2020-02-28").AddQuarterNoOverflow().ToDateString())
		assert.Equal(t, "2020-05-29", Parse("2020-02-29").AddQuarterNoOverflow().ToDateString())
		assert.Equal(t, "2021-02-28", Parse("2020-11-30").AddQuarterNoOverflow().ToDateString())
	})
}

func TestCarbon_SubQuarters(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-07-01", NewCarbon().SubQuarters(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubQuarters(2).ToString())
		assert.Empty(t, Parse("0").SubQuarters(2).ToString())
		assert.Empty(t, Parse("xxx").SubQuarters(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubQuarters(0).ToDateString())
		assert.Equal(t, "2019-10-01", Parse("2020-01-01").SubQuarters(1).ToDateString())
		assert.Equal(t, "2019-07-01", Parse("2020-01-01").SubQuarters(2).ToDateString())
		assert.Equal(t, "2019-05-29", Parse("2020-02-29").SubQuarters(3).ToDateString())
		assert.Equal(t, "2020-03-02", Parse("2020-08-31").SubQuarters(2).ToDateString())
	})
}

func TestCarbon_SubQuartersNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-07-01", NewCarbon().SubQuartersNoOverflow(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubQuartersNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SubQuartersNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SubQuartersNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubQuartersNoOverflow(0).ToDateString())
		assert.Equal(t, "2019-10-01", Parse("2020-01-01").SubQuartersNoOverflow(1).ToDateString())
		assert.Equal(t, "2019-07-01", Parse("2020-01-01").SubQuartersNoOverflow(2).ToDateString())
		assert.Equal(t, "2019-05-29", Parse("2020-02-29").SubQuartersNoOverflow(3).ToDateString())
		assert.Equal(t, "2020-02-29", Parse("2020-08-31").SubQuartersNoOverflow(2).ToDateString())
	})
}

func TestCarbon_SubQuarter(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-10-01", NewCarbon().SubQuarter().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubQuarter().ToString())
		assert.Empty(t, Parse("0").SubQuarter().ToString())
		assert.Empty(t, Parse("xxx").SubQuarter().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-10-01", Parse("2020-01-01").SubQuarter().ToDateString())
		assert.Equal(t, "2019-11-28", Parse("2020-02-28").SubQuarter().ToDateString())
		assert.Equal(t, "2019-11-29", Parse("2020-02-29").SubQuarter().ToDateString())
		assert.Equal(t, "2020-08-30", Parse("2020-11-30").SubQuarter().ToDateString())
	})
}

func TestCarbon_SubQuarterNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-10-01", NewCarbon().SubQuarterNoOverflow().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubQuarterNoOverflow().ToString())
		assert.Empty(t, Parse("0").SubQuarterNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").SubQuarterNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-10-01", Parse("2020-01-01").SubQuarterNoOverflow().ToDateString())
		assert.Equal(t, "2019-11-28", Parse("2020-02-28").SubQuarterNoOverflow().ToDateString())
		assert.Equal(t, "2019-11-29", Parse("2020-02-29").SubQuarterNoOverflow().ToDateString())
		assert.Equal(t, "2020-08-30", Parse("2020-11-30").SubQuarterNoOverflow().ToDateString())
	})
}

func TestCarbon_AddMonths(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-03-01", NewCarbon().AddMonths(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMonths(2).ToString())
		assert.Empty(t, Parse("0").AddMonths(2).ToString())
		assert.Empty(t, Parse("xxx").AddMonths(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddMonths(0).ToDateString())
		assert.Equal(t, "2020-02-01", Parse("2020-01-01").AddMonths(1).ToDateString())
		assert.Equal(t, "2020-03-01", Parse("2020-01-01").AddMonths(2).ToDateString())
		assert.Equal(t, "2020-05-29", Parse("2020-02-29").AddMonths(3).ToDateString())
		assert.Equal(t, "2020-10-31", Parse("2020-08-31").AddMonths(2).ToDateString())
	})
}

func TestCarbon_AddMonthsNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-03-01", NewCarbon().AddMonthsNoOverflow(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMonthsNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").AddMonthsNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").AddMonthsNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddMonthsNoOverflow(0).ToDateString())
		assert.Equal(t, "2020-02-01", Parse("2020-01-01").AddMonthsNoOverflow(1).ToDateString())
		assert.Equal(t, "2020-03-01", Parse("2020-01-01").AddMonthsNoOverflow(2).ToDateString())
		assert.Equal(t, "2020-05-29", Parse("2020-02-29").AddMonthsNoOverflow(3).ToDateString())
		assert.Equal(t, "2020-10-31", Parse("2020-08-31").AddMonthsNoOverflow(2).ToDateString())
	})
}

func TestCarbon_AddMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-01", NewCarbon().AddMonth().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMonth().ToString())
		assert.Empty(t, Parse("0").AddMonth().ToString())
		assert.Empty(t, Parse("xxx").AddMonth().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01", Parse("2020-01-01").AddMonth().ToDateString())
		assert.Equal(t, "2020-03-28", Parse("2020-02-28").AddMonth().ToDateString())
		assert.Equal(t, "2020-03-29", Parse("2020-02-29").AddMonth().ToDateString())
		assert.Equal(t, "2020-12-30", Parse("2020-11-30").AddMonth().ToDateString())
	})
}

func TestCarbon_AddMonthNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-02-01", NewCarbon().AddMonthNoOverflow().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMonthNoOverflow().ToString())
		assert.Empty(t, Parse("0").AddMonthNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").AddMonthNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-02-01", Parse("2020-01-01").AddMonthNoOverflow().ToDateString())
		assert.Equal(t, "2020-03-28", Parse("2020-02-28").AddMonthNoOverflow().ToDateString())
		assert.Equal(t, "2020-03-29", Parse("2020-02-29").AddMonthNoOverflow().ToDateString())
		assert.Equal(t, "2020-12-30", Parse("2020-11-30").AddMonthNoOverflow().ToDateString())
	})
}

func TestCarbon_SubMonths(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-11-01", NewCarbon().SubMonths(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMonths(2).ToString())
		assert.Empty(t, Parse("0").SubMonths(2).ToString())
		assert.Empty(t, Parse("xxx").SubMonths(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubMonths(0).ToDateString())
		assert.Equal(t, "2019-12-01", Parse("2020-01-01").SubMonths(1).ToDateString())
		assert.Equal(t, "2019-11-01", Parse("2020-01-01").SubMonths(2).ToDateString())
		assert.Equal(t, "2019-11-29", Parse("2020-02-29").SubMonths(3).ToDateString())
		assert.Equal(t, "2020-07-01", Parse("2020-08-31").SubMonths(2).ToDateString())
	})
}

func TestCarbon_SubMonthsNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-11-01", NewCarbon().SubMonthsNoOverflow(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMonthsNoOverflow(2).ToString())
		assert.Empty(t, Parse("0").SubMonthsNoOverflow(2).ToString())
		assert.Empty(t, Parse("xxx").SubMonthsNoOverflow(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubMonthsNoOverflow(0).ToDateString())
		assert.Equal(t, "2019-12-01", Parse("2020-01-01").SubMonthsNoOverflow(1).ToDateString())
		assert.Equal(t, "2019-11-01", Parse("2020-01-01").SubMonthsNoOverflow(2).ToDateString())
		assert.Equal(t, "2019-11-29", Parse("2020-02-29").SubMonthsNoOverflow(3).ToDateString())
		assert.Equal(t, "2020-06-30", Parse("2020-08-31").SubMonthsNoOverflow(2).ToDateString())
	})
}

func TestCarbon_SubMonth(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-01", NewCarbon().SubMonth().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMonth().ToString())
		assert.Empty(t, Parse("0").SubMonth().ToString())
		assert.Empty(t, Parse("xxx").SubMonth().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-01", Parse("2020-01-01").SubMonth().ToDateString())
		assert.Equal(t, "2020-01-28", Parse("2020-02-28").SubMonth().ToDateString())
		assert.Equal(t, "2020-01-29", Parse("2020-02-29").SubMonth().ToDateString())
		assert.Equal(t, "2020-10-30", Parse("2020-11-30").SubMonth().ToDateString())
	})
}

func TestCarbon_SubMonthNoOverflow(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-01", NewCarbon().SubMonthNoOverflow().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMonthNoOverflow().ToString())
		assert.Empty(t, Parse("0").SubMonthNoOverflow().ToString())
		assert.Empty(t, Parse("xxx").SubMonthNoOverflow().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-01", Parse("2020-01-01").SubMonthNoOverflow().ToDateString())
		assert.Equal(t, "2020-01-28", Parse("2020-02-28").SubMonthNoOverflow().ToDateString())
		assert.Equal(t, "2020-01-29", Parse("2020-02-29").SubMonthNoOverflow().ToDateString())
		assert.Equal(t, "2020-10-30", Parse("2020-11-30").SubMonthNoOverflow().ToDateString())
	})
}

func TestCarbon_AddWeeks(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-15", NewCarbon().AddWeeks(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddWeeks(2).ToString())
		assert.Empty(t, Parse("0").AddWeeks(2).ToString())
		assert.Empty(t, Parse("xxx").AddWeeks(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddWeeks(0).ToDateString())
		assert.Equal(t, "2020-01-08", Parse("2020-01-01").AddWeeks(1).ToDateString())
		assert.Equal(t, "2020-01-15", Parse("2020-01-01").AddWeeks(2).ToDateString())
		assert.Equal(t, "2020-03-21", Parse("2020-02-29").AddWeeks(3).ToDateString())
		assert.Equal(t, "2020-09-14", Parse("2020-08-31").AddWeeks(2).ToDateString())
	})
}

func TestCarbon_AddWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-08", NewCarbon().AddWeek().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddWeek().ToString())
		assert.Empty(t, Parse("0").AddWeek().ToString())
		assert.Empty(t, Parse("xxx").AddWeek().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-08", Parse("2020-01-01").AddWeek().ToDateString())
		assert.Equal(t, "2020-03-06", Parse("2020-02-28").AddWeek().ToDateString())
		assert.Equal(t, "2020-03-07", Parse("2020-02-29").AddWeek().ToDateString())
		assert.Equal(t, "2020-12-07", Parse("2020-11-30").AddWeek().ToDateString())
	})
}

func TestCarbon_SubWeeks(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-18", NewCarbon().SubWeeks(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubWeeks(2).ToString())
		assert.Empty(t, Parse("0").SubWeeks(2).ToString())
		assert.Empty(t, Parse("xxx").SubWeeks(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubWeeks(0).ToDateString())
		assert.Equal(t, "2019-12-25", Parse("2020-01-01").SubWeeks(1).ToDateString())
		assert.Equal(t, "2019-12-18", Parse("2020-01-01").SubWeeks(2).ToDateString())
		assert.Equal(t, "2020-02-08", Parse("2020-02-29").SubWeeks(3).ToDateString())
		assert.Equal(t, "2020-08-17", Parse("2020-08-31").SubWeeks(2).ToDateString())
	})
}

func TestCarbon_SubWeek(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-25", NewCarbon().SubWeek().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubWeek().ToString())
		assert.Empty(t, Parse("0").SubWeek().ToString())
		assert.Empty(t, Parse("xxx").SubWeek().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-25", Parse("2020-01-01").SubWeek().ToDateString())
		assert.Equal(t, "2020-02-21", Parse("2020-02-28").SubWeek().ToDateString())
		assert.Equal(t, "2020-02-22", Parse("2020-02-29").SubWeek().ToDateString())
		assert.Equal(t, "2020-11-23", Parse("2020-11-30").SubWeek().ToDateString())
	})
}

func TestCarbon_AddDays(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-03", NewCarbon().AddDays(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDays(2).ToString())
		assert.Empty(t, Parse("0").AddDays(2).ToString())
		assert.Empty(t, Parse("xxx").AddDays(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").AddDays(0).ToDateString())
		assert.Equal(t, "2020-01-02", Parse("2020-01-01").AddDays(1).ToDateString())
		assert.Equal(t, "2020-01-03", Parse("2020-01-01").AddDays(2).ToDateString())
		assert.Equal(t, "2020-03-03", Parse("2020-02-29").AddDays(3).ToDateString())
		assert.Equal(t, "2020-09-02", Parse("2020-08-31").AddDays(2).ToDateString())
	})
}

func TestCarbon_AddDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-02", NewCarbon().AddDay().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddDay().ToString())
		assert.Empty(t, Parse("0").AddDay().ToString())
		assert.Empty(t, Parse("xxx").AddDay().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-02", Parse("2020-01-01").AddDay().ToDateString())
		assert.Equal(t, "2020-02-29", Parse("2020-02-28").AddDay().ToDateString())
		assert.Equal(t, "2020-03-01", Parse("2020-02-29").AddDay().ToDateString())
		assert.Equal(t, "2020-12-01", Parse("2020-11-30").AddDay().ToDateString())
	})
}

func TestCarbon_SubDays(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-30", NewCarbon().SubDays(2).ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDays(2).ToString())
		assert.Empty(t, Parse("0").SubDays(2).ToString())
		assert.Empty(t, Parse("xxx").SubDays(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01", Parse("2020-01-01").SubDays(0).ToDateString())
		assert.Equal(t, "2019-12-31", Parse("2020-01-01").SubDays(1).ToDateString())
		assert.Equal(t, "2019-12-30", Parse("2020-01-01").SubDays(2).ToDateString())
		assert.Equal(t, "2020-02-26", Parse("2020-02-29").SubDays(3).ToDateString())
		assert.Equal(t, "2020-08-29", Parse("2020-08-31").SubDays(2).ToDateString())
	})
}

func TestCarbon_SubDay(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31", NewCarbon().SubDay().ToDateString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubDay().ToString())
		assert.Empty(t, Parse("0").SubDay().ToString())
		assert.Empty(t, Parse("xxx").SubDay().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2019-12-31", Parse("2020-01-01").SubDay().ToDateString())
		assert.Equal(t, "2020-02-27", Parse("2020-02-28").SubDay().ToDateString())
		assert.Equal(t, "2020-02-28", Parse("2020-02-29").SubDay().ToDateString())
		assert.Equal(t, "2020-11-29", Parse("2020-11-30").SubDay().ToDateString())
	})
}

func TestCarbon_AddHours(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 02:00:00 +0000 UTC", NewCarbon().AddHours(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddHours(2).ToString())
		assert.Empty(t, Parse("0").AddHours(2).ToString())
		assert.Empty(t, Parse("xxx").AddHours(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(0).ToString())
		assert.Equal(t, "2020-01-01 14:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(1).ToString())
		assert.Equal(t, "2020-01-01 15:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(2).ToString())
	})
}

func TestCarbon_AddHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 01:00:00 +0000 UTC", NewCarbon().AddHour().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddHour().ToString())
		assert.Empty(t, Parse("0").AddHour().ToString())
		assert.Empty(t, Parse("xxx").AddHour().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 14:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").AddHour().ToString())
	})
}

func TestCarbon_SubHours(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 22:00:00 +0000 UTC", NewCarbon().SubHours(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubHours(2).ToString())
		assert.Empty(t, Parse("0").SubHours(2).ToString())
		assert.Empty(t, Parse("xxx").SubHours(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(0).ToString())
		assert.Equal(t, "2020-01-01 12:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(1).ToString())
		assert.Equal(t, "2020-01-01 11:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(2).ToString())
	})
}

func TestCarbon_SubHour(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:00:00 +0000 UTC", NewCarbon().SubHour().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubHour().ToString())
		assert.Empty(t, Parse("0").SubHour().ToString())
		assert.Empty(t, Parse("xxx").SubHour().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 12:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").SubHour().ToString())
	})
}

func TestCarbon_AddMinutes(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:02:00 +0000 UTC", NewCarbon().AddMinutes(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMinutes(2).ToString())
		assert.Empty(t, Parse("0").AddMinutes(2).ToString())
		assert.Empty(t, Parse("xxx").AddMinutes(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(0).ToString())
		assert.Equal(t, "2020-01-01 13:15:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(1).ToString())
		assert.Equal(t, "2020-01-01 13:16:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(2).ToString())
	})
}

func TestCarbon_AddMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:01:00 +0000 UTC", NewCarbon().AddMinute().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMinute().ToString())
		assert.Empty(t, Parse("0").AddMinute().ToString())
		assert.Empty(t, Parse("xxx").AddMinute().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:15:15 +0000 UTC", Parse("2020-08-05 13:14:15").AddMinute().ToString())
	})
}

func TestCarbon_SubMinutes(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:58:00 +0000 UTC", NewCarbon().SubMinutes(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMinutes(2).ToString())
		assert.Empty(t, Parse("0").SubMinutes(2).ToString())
		assert.Empty(t, Parse("xxx").SubMinutes(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(0).ToString())
		assert.Equal(t, "2020-01-01 13:13:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(1).ToString())
		assert.Equal(t, "2020-01-01 13:12:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(2).ToString())
	})
}

func TestCarbon_SubMinute(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:00 +0000 UTC", NewCarbon().SubMinute().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMinute().ToString())
		assert.Empty(t, Parse("0").SubMinute().ToString())
		assert.Empty(t, Parse("xxx").SubMinute().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:13:15 +0000 UTC", Parse("2020-08-05 13:14:15").SubMinute().ToString())
	})
}

func TestCarbon_AddSeconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:02 +0000 UTC", NewCarbon().AddSeconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddSeconds(2).ToString())
		assert.Empty(t, Parse("0").AddSeconds(2).ToString())
		assert.Empty(t, Parse("xxx").AddSeconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:16 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:17 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(2).ToString())
	})
}

func TestCarbon_AddSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:01 +0000 UTC", NewCarbon().AddSecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddSecond().ToString())
		assert.Empty(t, Parse("0").AddSecond().ToString())
		assert.Empty(t, Parse("xxx").AddSecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:16 +0000 UTC", Parse("2020-08-05 13:14:15").AddSecond().ToString())
	})
}

func TestCarbon_SubSeconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:58 +0000 UTC", NewCarbon().SubSeconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubSeconds(2).ToString())
		assert.Empty(t, Parse("0").SubSeconds(2).ToString())
		assert.Empty(t, Parse("xxx").SubSeconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:14 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:13 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(2).ToString())
	})
}

func TestCarbon_SubSecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59 +0000 UTC", NewCarbon().SubSecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubSecond().ToString())
		assert.Empty(t, Parse("0").SubSecond().ToString())
		assert.Empty(t, Parse("xxx").SubSecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:14 +0000 UTC", Parse("2020-08-05 13:14:15").SubSecond().ToString())
	})
}

func TestCarbon_AddMilliseconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.002 +0000 UTC", NewCarbon().AddMilliseconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMilliseconds(2).ToString())
		assert.Empty(t, Parse("0").AddMilliseconds(2).ToString())
		assert.Empty(t, Parse("xxx").AddMilliseconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:15.001 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:15.002 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(2).ToString())
	})
}

func TestCarbon_AddMillisecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.001 +0000 UTC", NewCarbon().AddMillisecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMillisecond().ToString())
		assert.Empty(t, Parse("0").AddMillisecond().ToString())
		assert.Empty(t, Parse("xxx").AddMillisecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.001 +0000 UTC", Parse("2020-08-05 13:14:15").AddMillisecond().ToString())
	})
}

func TestCarbon_SubMilliseconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59.998 +0000 UTC", NewCarbon().SubMilliseconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMilliseconds(2).ToString())
		assert.Empty(t, Parse("0").SubMilliseconds(2).ToString())
		assert.Empty(t, Parse("xxx").SubMilliseconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:14.999 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:14.998 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(2).ToString())
	})
}

func TestCarbon_SubMillisecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59.999 +0000 UTC", NewCarbon().SubMillisecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMillisecond().ToString())
		assert.Empty(t, Parse("0").SubMillisecond().ToString())
		assert.Empty(t, Parse("xxx").SubMillisecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:14.999 +0000 UTC", Parse("2020-08-05 13:14:15").SubMillisecond().ToString())
	})
}

func TestCarbon_AddMicroseconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.000002 +0000 UTC", NewCarbon().AddMicroseconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMicroseconds(2).ToString())
		assert.Empty(t, Parse("0").AddMicroseconds(2).ToString())
		assert.Empty(t, Parse("xxx").AddMicroseconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:15.000001 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:15.000002 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(2).ToString())
	})
}

func TestCarbon_AddMicrosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.000001 +0000 UTC", NewCarbon().AddMicrosecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddMicrosecond().ToString())
		assert.Empty(t, Parse("0").AddMicrosecond().ToString())
		assert.Empty(t, Parse("xxx").AddMicrosecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.000001 +0000 UTC", Parse("2020-08-05 13:14:15").AddMicrosecond().ToString())
	})
}

func TestCarbon_SubMicroseconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59.999998 +0000 UTC", NewCarbon().SubMicroseconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMicroseconds(2).ToString())
		assert.Empty(t, Parse("0").SubMicroseconds(2).ToString())
		assert.Empty(t, Parse("xxx").SubMicroseconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:14.999999 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:14.999998 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(2).ToString())
	})
}

func TestCarbon_SubMicrosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59.999999 +0000 UTC", NewCarbon().SubMicrosecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubMicrosecond().ToString())
		assert.Empty(t, Parse("0").SubMicrosecond().ToString())
		assert.Empty(t, Parse("xxx").SubMicrosecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:14.999999 +0000 UTC", Parse("2020-08-05 13:14:15").SubMicrosecond().ToString())
	})
}

func TestCarbon_AddNanoseconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.000000002 +0000 UTC", NewCarbon().AddNanoseconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddNanoseconds(2).ToString())
		assert.Empty(t, Parse("0").AddNanoseconds(2).ToString())
		assert.Empty(t, Parse("xxx").AddNanoseconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:15.000000001 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:15.000000002 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(2).ToString())
	})
}

func TestCarbon_AddNanosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0001-01-01 00:00:00.000000001 +0000 UTC", NewCarbon().AddNanosecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").AddNanosecond().ToString())
		assert.Empty(t, Parse("0").AddNanosecond().ToString())
		assert.Empty(t, Parse("xxx").AddNanosecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15.000000001 +0000 UTC", Parse("2020-08-05 13:14:15").AddNanosecond().ToString())
	})
}

func TestCarbon_SubNanoseconds(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59.999999998 +0000 UTC", NewCarbon().SubNanoseconds(2).ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubNanoseconds(2).ToString())
		assert.Empty(t, Parse("0").SubNanoseconds(2).ToString())
		assert.Empty(t, Parse("xxx").SubNanoseconds(2).ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(0).ToString())
		assert.Equal(t, "2020-01-01 13:14:14.999999999 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(1).ToString())
		assert.Equal(t, "2020-01-01 13:14:14.999999998 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(2).ToString())
	})
}

func TestCarbon_SubNanosecond(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().SubNanosecond().ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Empty(t, Parse("").SubNanosecond().ToString())
		assert.Empty(t, Parse("0").SubNanosecond().ToString())
		assert.Empty(t, Parse("xxx").SubNanosecond().ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:14.999999999 +0000 UTC", Parse("2020-08-05 13:14:15").SubNanosecond().ToString())
	})
}
