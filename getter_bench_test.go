package carbon

import "testing"

func BenchmarkCarbon_StdTime(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.StdTime()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.StdTime()
	}
}

func BenchmarkCarbon_DaysInYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DaysInYear()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DaysInYear()
	}
}

func BenchmarkCarbon_DaysInMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DaysInMonth()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DaysInMonth()
	}
}

func BenchmarkCarbon_MonthOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.MonthOfYear()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.MonthOfYear()
	}
}

func BenchmarkCarbon_DayOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DayOfYear()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DayOfYear()
	}
}

func BenchmarkCarbon_DayOfMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DayOfMonth()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DayOfMonth()
	}
}

func BenchmarkCarbon_DayOfWeek(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DayOfWeek()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DayOfWeek()
	}
}

func BenchmarkCarbon_WeekOfYear(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.WeekOfYear()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.WeekOfYear()
	}
}

func BenchmarkCarbon_WeekOfMonth(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.WeekOfMonth()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.WeekOfMonth()
	}
}

func BenchmarkCarbon_DateTime(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTime()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateTime()
	}
}

func BenchmarkCarbon_DateTimeMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTimeMilli()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateTimeMilli()
	}
}

func BenchmarkCarbon_DateTimeMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTimeMilli()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateTimeMilli()
	}
}

func BenchmarkCarbon_DateTimeNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateTimeNano()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateTimeNano()
	}
}

func BenchmarkCarbon_Date(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Date()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Date()
	}
}

func BenchmarkCarbon_DateMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateMilli()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateMilli()
	}
}

func BenchmarkCarbon_DateMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateMicro()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateMicro()
	}
}

func BenchmarkCarbon_DateNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.DateNano()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.DateNano()
	}
}

func BenchmarkCarbon_Time(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Time()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Time()
	}
}

func BenchmarkCarbon_TimeMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimeMilli()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.TimeMilli()
	}
}

func BenchmarkCarbon_TimeMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimeMicro()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.TimeMicro()
	}
}

func BenchmarkCarbon_TimeNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimeNano()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.TimeNano()
	}
}

func BenchmarkCarbon_Century(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Century()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Century()
	}
}

func BenchmarkCarbon_Decade(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Decade()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Decade()
	}
}

func BenchmarkCarbon_Year(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Year()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Year()
	}
}

func BenchmarkCarbon_Quarter(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Quarter()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Quarter()
	}
}

func BenchmarkCarbon_Month(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Month()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Month()
	}
}

func BenchmarkCarbon_Week(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Week()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Week()
	}
}

func BenchmarkCarbon_Day(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Day()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Day()
	}
}

func BenchmarkCarbon_Hour(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Hour()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Hour()
	}
}

func BenchmarkCarbon_Minute(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Minute()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Minute()
	}
}

func BenchmarkCarbon_Second(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Second()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Second()
	}
}

func BenchmarkCarbon_Millisecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Millisecond()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Millisecond()
	}
}

func BenchmarkCarbon_Microsecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Microsecond()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Microsecond()
	}
}

func BenchmarkCarbon_Nanosecond(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Nanosecond()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Nanosecond()
	}
}

func BenchmarkCarbon_Timestamp(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Timestamp()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Timestamp()
	}
}

func BenchmarkCarbon_TimestampMilli(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimestampMilli()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.TimestampMilli()
	}
}

func BenchmarkCarbon_TimestampMicro(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimestampMicro()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.TimestampMicro()
	}
}

func BenchmarkCarbon_TimestampNano(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.TimestampNano()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.TimestampNano()
	}
}

func BenchmarkCarbon_Location(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Location()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Location()
	}
}

func BenchmarkCarbon_Timezone(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Timezone()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Timezone()
	}
}

func BenchmarkCarbon_Offset(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Offset()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Offset()
	}
}

func BenchmarkCarbon_Locale(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Locale()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Locale()
	}
}

func BenchmarkCarbon_Age(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Age()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.Age()
	}
}

func BenchmarkCarbon_CurrentLayout(b *testing.B) {
	now := Now()
	for n := 0; n < b.N; n++ {
		now.Age()
	}

	c := NewCarbon()
	for n := 0; n < b.N; n++ {
		c.CurrentLayout()
	}
}
