package carbon

import (
	"fmt"
	"math"
	"strings"
)

type VLunarFestival struct {
	Day   int
	Month int
	Name  string
}

type LuckyHour struct {
	Chi  string
	From int
	To   int
}

type LuckyHourDetail struct {
	Chi  string
	From int
	To   int
}

type SolarTerm struct {
	Longitude int
	Name      string
}

var (
	PI              = math.Pi
	vLunarTimes     = []string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	vLunarAnimals   = []string{"Thân", "Dậu", "Tuất", "Hợi", "Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi"}
	vLunarFestivals = []VLunarFestival{
		{
			Day:   1,
			Month: 1,
			Name:  "Tết Nguyên Đán",
		},
		{
			Day:   15,
			Month: 1,
			Name:  "Rằm tháng Giêng",
		},
		{
			Day:   10,
			Month: 3,
			Name:  "Giỗ Tổ Hùng Vương",
		},
		{
			Day:   15,
			Month: 4,
			Name:  "Phật Đản",
		},
		{
			Day:   5,
			Month: 5,
			Name:  "Lễ Đoan Ngọ",
		},
		{
			Day:   15,
			Month: 7,
			Name:  "Vu Lan",
		},
		{
			Day:   15,
			Month: 8,
			Name:  "Tết Trung Thu",
		},
		{
			Day:   23,
			Month: 12,
			Name:  "Ông Táo chầu trời",
		},
	}
	vSolarTerms = []SolarTerm{
		{
			Longitude: 0,
			Name:      "Xuân phân",
		},
		{
			Longitude: 15,
			Name:      "Thanh minh",
		},
		{
			Longitude: 30,
			Name:      "Cốc vũ",
		},
		{
			Longitude: 45,
			Name:      "Lập hạ",
		},
		{
			Longitude: 60,
			Name:      "Tiểu mãn",
		},
		{
			Longitude: 75,
			Name:      "Mang chủng",
		},
		{
			Longitude: 90,
			Name:      "Hạ chí",
		},
		{
			Longitude: 105,
			Name:      "Tiểu thử",
		},
		{
			Longitude: 120,
			Name:      "Đại thử",
		},
		{
			Longitude: 135,
			Name:      "Lập thu",
		},
		{
			Longitude: 150,
			Name:      "Xử thử",
		},
		{
			Longitude: 165,
			Name:      "Bạch lộ",
		},
		{
			Longitude: 180,
			Name:      "Thu phân",
		},
		{
			Longitude: 195,
			Name:      "Hàn lộ",
		},
		{
			Longitude: 210,
			Name:      "Sương giáng",
		},
		{
			Longitude: 225,
			Name:      "Lập đông",
		},
		{
			Longitude: 240,
			Name:      "Tiểu tuyết",
		},
		{
			Longitude: 255,
			Name:      "Đại tuyết",
		},
		{
			Longitude: 270,
			Name:      "Đông chí",
		},
		{
			Longitude: 285,
			Name:      "Tiểu hàn",
		},
		{
			Longitude: 300,
			Name:      "Đại hàn",
		},
		{
			Longitude: 315,
			Name:      "Lập xuân",
		},
		{
			Longitude: 330,
			Name:      "Vũ Thủy",
		},
		{
			Longitude: 345,
			Name:      "Kinh trập",
		},
	}
	vLunarLuckyHours = [6]string{"110100101100", "001101001011", "110011010010", "101100110100", "001011001101", "010010110011"}
)

// lunar defines a lunar struct.
type vlunar struct {
	year, month, day, hour, minute, second, offset, timeZone int
	isInvalid                                                bool
	isLeapMonth                                              bool
	Error                                                    error
	jd                                                       int
}

func integerToFloat(i int) float64 {
	return float64(i)
}

func mathFloor(f float64) float64 {
	return math.Floor(f)
}

func floor(f float64) int {
	return int(mathFloor(f))
}

func dateToJuliusDay(dd, mm, yy int) int {
	a := floor((14 - integerToFloat(mm)) / 12)
	y := integerToFloat(yy) + 4800 - integerToFloat(a)
	m := integerToFloat(mm) + 12*integerToFloat(a) - 3
	jd := integerToFloat(dd) + mathFloor((153*m+2)/5) + 365*y + mathFloor(y/4) - math.Floor(y/100) + mathFloor(y/400) - 32045
	if jd < 2299161 {
		jd = integerToFloat(dd) + mathFloor((153*m+2)/5) + 365*y + mathFloor(y/4) - 32083
	}

	return floor(jd)
}

func newMoon(k int) float64 {
	var kf, t, t2, t3, dr, jd1, m, mpr, f, c1, deltat, jdNew float64
	kf = integerToFloat(k)
	t = kf / 1236.85 // Time in Julian centuries from 1900 January 0.5
	t2 = t * t
	t3 = t2 * t
	dr = PI / 180
	jd1 = 2415020.75933 + 29.53058868*kf + 0.0001178*t2 - 0.000000155*t3
	jd1 = jd1 + 0.00033*math.Sin((166.56+132.87*t-0.009173*t2)*dr)  // Mean new moon
	m = 359.2242 + 29.10535608*kf - 0.0000333*t2 - 0.00000347*t3    // Sun's mean anomaly
	mpr = 306.0253 + 385.81691806*kf + 0.0107306*t2 + 0.00001236*t3 // Moon's mean anomaly
	f = 21.2964 + 390.67050646*kf - 0.0016528*t2 - 0.00000239*t3    // Moon's argument of latitude
	c1 = (0.1734-0.000393*t)*math.Sin(m*dr) + 0.0021*math.Sin(2*dr*m)
	c1 = c1 - 0.4068*math.Sin(mpr*dr) + 0.0161*math.Sin(dr*2*mpr)
	c1 = c1 - 0.0004*math.Sin(dr*3*mpr)
	c1 = c1 + 0.0104*math.Sin(dr*2*f) - 0.0051*math.Sin(dr*(m+mpr))
	c1 = c1 - 0.0074*math.Sin(dr*(m-mpr)) + 0.0004*math.Sin(dr*(2*f+m))
	c1 = c1 - 0.0004*math.Sin(dr*(2*f-m)) - 0.0006*math.Sin(dr*(2*f+mpr))
	c1 = c1 + 0.0010*math.Sin(dr*(2*f-mpr)) + 0.0005*math.Sin(dr*(2*mpr+m))
	if t < -11 {
		deltat = 0.001 + 0.000839*t + 0.0002261*t2 - 0.00000845*t3 - 0.000000081*t*t3
	} else {
		deltat = -0.000278 + 0.000265*t + 0.000262*t2
	}
	jdNew = jd1 + c1 - deltat

	return jdNew
}

func getNewMoonDay(k, tz int) int {
	return floor(newMoon(k) + 0.5 + integerToFloat(tz)/24)
}

func sunLongitude(jdn float64) float64 {
	var t, t2, dr, m, l0, dl, l float64
	t = (jdn - 2451545.0) / 36525 // Time in Julian centuries from 2000-01-01 12:00:00 GMT
	t2 = t * t
	dr = PI / 180                                                  // degree to radian
	m = 357.52910 + 35999.05030*t - 0.0001559*t2 - 0.00000048*t*t2 // mean anomaly, degree
	l0 = 280.46645 + 36000.76983*t + 0.0003032*t2                  // mean longitude, degree
	dl = (1.914600 - 0.004817*t - 0.000014*t2) * math.Sin(dr*m)
	dl = dl + (0.019993-0.000101*t)*math.Sin(dr*2*m) + 0.00029*math.Sin(dr*3*m)
	l = l0 + dl // true longitude, degree
	l = l * dr
	l = l - PI*2*(math.Floor(l/(PI*2))) // Normalize to (0, 2*PI)
	return l
}

func getSunLongitude(d, tz int) int {
	return floor((sunLongitude(integerToFloat(d)-0.5-integerToFloat(tz)/24) / PI) * 6)
}

func getSolarTerm(dayNumber int, timeZone int) int {
	return floor(sunLongitude(integerToFloat(dayNumber)-0.5-integerToFloat(timeZone)/24.0) / PI * 12)
}

/* Find the day that starts the lunar month 11 of the given year for the given time zone */
func getLunarStartNovember(yy, tz int) int {
	var off, k, nm int

	off = dateToJuliusDay(31, 12, yy) - 2415021
	k = floor(integerToFloat(off) / 29.530588853)
	nm = getNewMoonDay(k, tz)
	sunLong := getSunLongitude(nm, tz) // sun longitude at local midnight
	if sunLong >= 9 {
		nm = getNewMoonDay(k-1, tz)
	}
	return nm
}

func getLeapMonthOffset(a11 float64, tz int) int {
	k := floor((a11-2415021.076998695)/29.530588853 + 0.5)
	last := 0
	i := 1 // We start with the month following lunar month 11
	arc := getSunLongitude(getNewMoonDay(k+i, tz), tz)
	for {
		last = arc
		i++
		newmoon := getNewMoonDay(k+i, tz)
		arc = getSunLongitude(newmoon, tz)

		if arc == last || i >= 14 {
			break
		}
	}
	return i - 1
}

// The function SolarToLunar converts a given Solar date (dd, mm, yy) into
// the corresponding Lunar date (lunarD, lunarM, lunarY) in the Vietnamese Lunar calendar.
// It also determines if the given Lunar month is a leap month or not (lunarLeap).
// The time zone (tz) is used to calculate the date in the Lunar calendar.
func solarToLunar(dd, mm, yy, tz int) (lunarD, lunarM, lunarY, lunarLeap int) {
	dayNumber := dateToJuliusDay(dd, mm, yy)
	k := floor((integerToFloat(dayNumber) - 2415021.076998695) / 29.530588853)
	monthStart := getNewMoonDay(k+1, tz)
	if monthStart > dayNumber {
		monthStart = getNewMoonDay(k, tz)
	}
	a11 := getLunarStartNovember(yy, tz)
	b11 := a11
	if a11 >= monthStart {
		lunarY = yy
		a11 = getLunarStartNovember(yy-1, tz)
	} else {
		lunarY = yy + 1
		b11 = getLunarStartNovember(yy+1, tz)
	}
	lunarD = dayNumber - monthStart + 1
	diff := floor((integerToFloat(monthStart) - integerToFloat(a11)) / 29)

	lunarLeap = 0
	lunarM = diff + 11
	if b11-a11 > 365 {
		leapMonthDiff := getLeapMonthOffset(integerToFloat(a11), tz)
		if diff >= leapMonthDiff {
			lunarM = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = 1
			}
		}
	}
	if lunarM > 12 {
		lunarM = lunarM - 12
	}
	if lunarM >= 11 && diff < 4 {
		lunarY -= 1
	}

	return
}

// Lunar converts the gregorian calendar to the vietnam lunar calendar.
func (c Carbon) VLunar() (l vlunar) {
	l.isInvalid, l.isLeapMonth = false, false
	if c.IsInvalid() {
		l.Error = c.Error
		l.isInvalid = true
		return
	}
	l.offset = c.Offset()
	l.timeZone = l.offset / 3600
	lD, lM, lY, leap := solarToLunar(c.Day(), c.Month(), c.Year(), l.timeZone)
	l.day = lD
	l.month = lM
	l.year = lY
	l.isLeapMonth = false
	if leap == 1 {
		l.isLeapMonth = true
	}

	l.hour, l.minute, l.second = c.Time()
	l.jd = dateToJuliusDay(c.Day(), c.Month(), c.Year())
	return
}

// Animal gets lunar animal name like "猴".
func (l vlunar) Animal() string {
	if l.isInvalid {
		return ""
	}
	return vLunarAnimals[l.year%MonthsPerYear]
}

// Festival gets lunar festival name like "Trung thu".
func (l vlunar) Festivals() (festivals []VLunarFestival) {
	festivals = []VLunarFestival{}
	if l.isInvalid {
		return
	}
	month, day := l.month, l.day

	for i := 0; i < len(vLunarFestivals); i++ {
		event := vLunarFestivals[i]
		if event.Day == day && event.Month == month {
			festivals = append(festivals, event)
		}
	}

	return
}

func (l vlunar) SolarTerm() SolarTerm {
	if l.isInvalid {
		return SolarTerm{}
	}

	solarTerm := getSolarTerm(l.jd+1, 7.0)
	return vSolarTerms[solarTerm]
}

func (l vlunar) LuckyHour() string {
	if l.isInvalid {
		return ""
	}

	chiOfDay := (l.jd + 1) % 12
	luckyHour := vLunarLuckyHours[chiOfDay%6]

	return luckyHour
}

func (l vlunar) LuckyHourDetails() []LuckyHourDetail {
	ret := []LuckyHourDetail{}
	luckyHour := l.LuckyHour()
	if len(luckyHour) <= 0 {
		return ret
	}
	for i := 0; i < 12; i++ {
		index := luckyHour[i]
		if index == '1' {
			detail := LuckyHourDetail{
				Chi:  vLunarTimes[i],
				From: (i*2 + 23) % 24,
				To:   (i*2 + 1) % 24,
			}
			ret = append(ret, detail)
		}
	}
	return ret
}

// DateTime gets lunar year, month, day, hour, minute, and second like 2020, 8, 5, 13, 14, 15.
func (l vlunar) DateTime() (year, month, day, hour, minute, second int) {
	if l.isInvalid {
		return
	}
	return l.year, l.month, l.day, l.hour, l.minute, l.second
}

// Date gets lunar year, month and day like 2020, 8, 5.
func (l vlunar) Date() (year, month, day int) {
	if l.isInvalid {
		return
	}
	return l.year, l.month, l.day
}

// Time gets lunar hour, minute, and second like 13, 14, 15.
func (l vlunar) Time() (hour, minute, second int) {
	if l.isInvalid {
		return
	}
	return l.hour, l.minute, l.second
}

// Year gets lunar year like 2020.
func (l vlunar) Year() int {
	if l.isInvalid {
		return 0
	}
	return l.year
}

// Month gets lunar month like 8.
func (l vlunar) Month() int {
	if l.isInvalid {
		return 0
	}
	return l.month
}

// LeapMonth gets lunar leap month like 8.
func (l vlunar) LeapMonth() int {
	if l.isInvalid {
		return 0
	}
	month := l.month
	leapMonth := 0
	a11 := getLunarStartNovember(l.year-1, l.timeZone)
	b11 := getLunarStartNovember(l.year, l.timeZone)
	if month >= 11 {
		a11 = getLunarStartNovember(l.year, l.timeZone)
		b11 = getLunarStartNovember(l.year+1, l.timeZone)
	}
	if b11-a11 > 365 {
		leapOffset := getLeapMonthOffset(float64(a11), l.timeZone)
		leapMonth = leapOffset - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
	}
	return leapMonth
}

// Day gets lunar day like 5.
func (l vlunar) Day() int {
	if l.isInvalid {
		return 0
	}
	return l.day
}

// ToYearString outputs a string in lunar year format like "Giáp Tý".
func (l vlunar) ToYearString() string {
	if l.isInvalid {
		return ""
	}
	year := fmt.Sprintf("%d", l.year)
	return year
}

// ToMonthString outputs a string in lunar month format like "Tháng Một".
func (l vlunar) ToMonthString() string {
	if l.isInvalid {
		return ""
	}
	month := fmt.Sprintf("%d", l.month)
	return month
}

// ToDayString outputs a string in lunar day format like "Mùng Một".
func (l vlunar) ToDayString() (day string) {
	if l.isInvalid {
		return
	}
	day = fmt.Sprintf("%d", l.day)
	return
}

// ToDateString outputs a string in lunar date format like "Ngày mười sáu tháng chín năm 2020".
func (l vlunar) ToDateString() string {
	if l.isInvalid {
		return ""
	}
	return "Ngày " + strings.ToLower(l.ToDayString()) + " tháng " + strings.ToLower(l.ToMonthString()) + " năm " + strings.ToLower(l.ToYearString())
}

// String outputs a string in YYYY-MM-DD HH::ii::ss format, implement Stringer interface.
func (l vlunar) String() string {
	if l.isInvalid {
		return ""
	}
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", l.year, l.month, l.day, l.hour, l.minute, l.second)
}

// IsLeapYear reports whether is leap year.
func (l vlunar) IsLeapYear() bool {
	if l.isInvalid {
		return false
	}
	return l.LeapMonth() != 0
}

// IsLeapMonth reports whether is leap month.
func (l vlunar) IsLeapMonth() bool {
	if l.isInvalid {
		return false
	}
	return l.month == l.LeapMonth()
}

// IsRatYear reports whether is year of Rat.
func (l vlunar) IsRatYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 4 {
		return true
	}
	return false
}

// IsOxYear reports whether is year of Ox.
func (l vlunar) IsOxYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 5 {
		return true
	}
	return false
}

// IsTigerYear reports whether is year of Tiger.
func (l vlunar) IsTigerYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 6 {
		return true
	}
	return false
}

// IsRabbitYear reports whether is year of Rabbit.
func (l vlunar) IsCatYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 7 {
		return true
	}
	return false
}

// IsDragonYear reports whether is year of Dragon.
func (l vlunar) IsDragonYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 8 {
		return true
	}
	return false
}

// IsSnakeYear reports whether is year of Snake.
func (l vlunar) IsSnakeYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 9 {
		return true
	}
	return false
}

// IsHorseYear reports whether is year of Horse.
func (l vlunar) IsHorseYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 10 {
		return true
	}
	return false
}

// IsGoatYear reports whether is year of Goat.
func (l vlunar) IsGoatYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 11 {
		return true
	}
	return false
}

// IsMonkeyYear reports whether is year of Monkey.
func (l vlunar) IsMonkeyYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 0 {
		return true
	}
	return false
}

// IsRoosterYear reports whether is year of Rooster.
func (l vlunar) IsRoosterYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 1 {
		return true
	}
	return false
}

// IsDogYear reports whether is year of Dog.
func (l vlunar) IsDogYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 2 {
		return true
	}
	return false
}

// IsPigYear reports whether is year of Pig.
func (l vlunar) IsPigYear() bool {
	if l.isInvalid {
		return false
	}
	if l.year%MonthsPerYear == 3 {
		return true
	}
	return false
}

// DoubleHour gets double-hour name like "子时".
func (l vlunar) DoubleHour() (dh string) {
	if l.isInvalid {
		return ""
	}
	hour, minute := l.hour, l.minute
	switch {
	case hour >= 23, hour == 0 && minute <= 59:
		dh = vLunarTimes[0] // FirstDoubleHour
	case hour >= 1 && hour < 3:
		dh = vLunarTimes[1] // SecondDoubleHour
	case hour >= 3 && hour < 5:
		dh = vLunarTimes[2] // ThirdDoubleHour
	case hour >= 5 && hour < 7:
		dh = vLunarTimes[3] // FourthDoubleHour
	case hour >= 7 && hour < 9:
		dh = vLunarTimes[4] // FifthDoubleHour
	case hour >= 9 && hour < 11:
		dh = vLunarTimes[5] // SixthDoubleHour
	case hour >= 11 && hour < 13:
		dh = vLunarTimes[6] // SeventhDoubleHour
	case hour >= 13 && hour < 15:
		dh = vLunarTimes[7] // EighthDoubleHour
	case hour >= 15 && hour < 17:
		dh = vLunarTimes[8] // NinthDoubleHour
	case hour >= 17 && hour < 19:
		dh = vLunarTimes[9] // TenthDoubleHour
	case hour >= 19 && hour < 21:
		dh = vLunarTimes[10] // EleventhDoubleHour
	case hour >= 21 && hour < 23:
		dh = vLunarTimes[11] // TwelfthDoubleHour
	}
	return
}

// IsFirstDoubleHour reports whether is FirstDoubleHour.
func (l vlunar) IsFirstDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour, minute := l.hour, l.minute
	if hour >= 23 {
		return true
	}
	if hour == 0 && minute <= 59 {
		return true
	}
	return false
}

// IsSecondDoubleHour reports whether is SecondDoubleHour.
func (l vlunar) IsSecondDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 1 && hour < 3 {
		return true
	}
	return false
}

// IsThirdDoubleHour reports whether is ThirdDoubleHour.
func (l vlunar) IsThirdDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 3 && hour < 5 {
		return true
	}
	return false
}

// IsFourthDoubleHour reports whether is FourthDoubleHour.
func (l vlunar) IsFourthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 5 && hour < 7 {
		return true
	}
	return false
}

// IsFifthDoubleHour reports whether is FifthDoubleHour.
func (l vlunar) IsFifthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 7 && hour < 9 {
		return true
	}
	return false
}

// IsSixthDoubleHour reports whether is SixthDoubleHour.
func (l vlunar) IsSixthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 9 && hour < 11 {
		return true
	}
	return false
}

// IsSeventhDoubleHour reports whether is SeventhDoubleHour.
func (l vlunar) IsSeventhDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 11 && hour < 13 {
		return true
	}
	return false
}

// IsEighthDoubleHour reports whether is EighthDoubleHour.
func (l vlunar) IsEighthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 13 && hour < 15 {
		return true
	}
	return false
}

// IsNinthDoubleHour reports whether is NinthDoubleHour.
func (l vlunar) IsNinthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 15 && hour < 17 {
		return true
	}
	return false
}

// IsTenthDoubleHour reports whether is TenthDoubleHour.
func (l vlunar) IsTenthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 17 && hour < 19 {
		return true
	}
	return false
}

// IsEleventhDoubleHour reports whether is EleventhDoubleHour.
func (l vlunar) IsEleventhDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 19 && hour < 21 {
		return true
	}
	return false
}

// IsTwelfthDoubleHour reports whether is TwelfthDoubleHour.
func (l vlunar) IsTwelfthDoubleHour() bool {
	if l.isInvalid {
		return false
	}
	hour := l.hour
	if hour >= 21 && hour < 23 {
		return true
	}
	return false
}
