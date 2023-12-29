package carbon

import "testing"

func BenchmarkCarbon_VLunar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Now().VLunar()
	}
}

func BenchmarkVLunar_Animal(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Animal()
	}
}

func BenchmarkVLunar_Festival(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Festivals()
	}
}

func BenchmarkVLunar_DateTime(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.DateTime()
	}
}

func BenchmarkVLunar_Date(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Date()
	}
}

func BenchmarkVLunar_Time(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Time()
	}
}

func BenchmarkVLunar_Year(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Year()
	}
}

func BenchmarkVLunar_Month(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Month()
	}
}

func BenchmarkVLunar_LeapMonth(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.LeapMonth()
	}
}

func BenchmarkVLunar_SolarTerm(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.SolarTerm()
	}
}

func BenchmarkVLunar_LuckyHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.LuckyHour()
	}
}

func BenchmarkVLunar_LuckyHourDetails(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.LuckyHourDetails()
	}
}

func BenchmarkVLunar_Day(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.Day()
	}
}

func BenchmarkVLunar_ToYearString(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.ToYearString()
	}
}

func BenchmarkVLunar_ToMonthString(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.ToMonthString()
	}
}

func BenchmarkVLunar_ToDayString(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.ToDayString()
	}
}

func BenchmarkVLunar_ToDateString(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.ToDateString()
	}
}

func BenchmarkVLunar_String(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.String()
	}
}

func BenchmarkVLunar_IsLeapYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsLeapYear()
	}
}

func BenchmarkVLunar_IsLeapMonth(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsLeapMonth()
	}
}

func BenchmarkVLunar_IsRatYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsRatYear()
	}
}

func BenchmarkVLunar_IsOxYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsOxYear()
	}
}

func BenchmarkVLunar_IsTigerYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsTigerYear()
	}
}

func BenchmarkVLunar_IsRabbitYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsCatYear()
	}
}

func BenchmarkVLunar_IsDragonYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsDragonYear()
	}
}

func BenchmarkVLunar_IsSnakeYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsSnakeYear()
	}
}

func BenchmarkVLunar_IsHorseYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsHorseYear()
	}
}

func BenchmarkVLunar_IsGoatYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsGoatYear()
	}
}

func BenchmarkVLunar_IsMonkeyYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsMonkeyYear()
	}
}

func BenchmarkVLunar_IsRoosterYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsRoosterYear()
	}
}

func BenchmarkVLunar_IsDogYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsDogYear()
	}
}

func BenchmarkVLunar_IsPigYear(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsPigYear()
	}
}

func BenchmarkVLunar_DoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.DoubleHour()
	}
}

func BenchmarkVLunar_IsFirstDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsFirstDoubleHour()
	}
}

func BenchmarkVLunar_IsSecondDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsSecondDoubleHour()
	}
}

func BenchmarkVLunar_IsThirdDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsThirdDoubleHour()
	}
}

func BenchmarkVLunar_IsFourthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsFourthDoubleHour()
	}
}

func BenchmarkVLunar_IsFifthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsFifthDoubleHour()
	}
}

func BenchmarkVLunar_IsSixthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsSixthDoubleHour()
	}
}

func BenchmarkVLunar_IsSeventhDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsSeventhDoubleHour()
	}
}

func BenchmarkVLunar_IsEighthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsEighthDoubleHour()
	}
}

func BenchmarkVLunar_IsNinthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsNinthDoubleHour()
	}
}

func BenchmarkVLunar_IsTenthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsTenthDoubleHour()
	}
}

func BenchmarkVLunar_IsEleventhDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsEleventhDoubleHour()
	}
}

func BenchmarkVLunar_IsTwelfthDoubleHour(b *testing.B) {
	l := Now().VLunar()
	for n := 0; n < b.N; n++ {
		l.IsTwelfthDoubleHour()
	}
}
