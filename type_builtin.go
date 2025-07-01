package carbon

type (
	timestampType      int64
	timestampMicroType int64
	timestampMilliType int64
	timestampNanoType  int64

	datetimeType      string
	datetimeMicroType string
	datetimeMilliType string
	datetimeNanoType  string

	dateType      string
	dateMilliType string
	dateMicroType string
	dateNanoType  string

	timeType      string
	timeMilliType string
	timeMicroType string
	timeNanoType  string
)

type (
	Timestamp      = TimestampType[timestampType]
	TimestampMilli = TimestampType[timestampMilliType]
	TimestampMicro = TimestampType[timestampMicroType]
	TimestampNano  = TimestampType[timestampNanoType]

	DateTime      = LayoutType[datetimeType]
	DateTimeMicro = LayoutType[datetimeMicroType]
	DateTimeMilli = LayoutType[datetimeMilliType]
	DateTimeNano  = LayoutType[datetimeNanoType]

	Date      = LayoutType[dateType]
	DateMilli = LayoutType[dateMilliType]
	DateMicro = LayoutType[dateMicroType]
	DateNano  = LayoutType[dateNanoType]

	Time      = LayoutType[timeType]
	TimeMilli = LayoutType[timeMilliType]
	TimeMicro = LayoutType[timeMicroType]
	TimeNano  = LayoutType[timeNanoType]
)

func NewTimestamp(c Carbon) Timestamp {
	return NewTimestampType[timestampType](c)
}
func NewTimestampMilli(c Carbon) TimestampMilli {
	return NewTimestampType[timestampMilliType](c)
}
func NewTimestampMicro(c Carbon) TimestampMicro {
	return NewTimestampType[timestampMicroType](c)
}
func NewTimestampNano(c Carbon) TimestampNano {
	return NewTimestampType[timestampNanoType](c)
}

func NewDateTime(c Carbon) DateTime {
	return NewLayoutType[datetimeType](c)
}
func NewDateTimeMilli(c Carbon) DateTimeMilli {
	return NewLayoutType[datetimeMilliType](c)
}
func NewDateTimeMicro(c Carbon) DateTimeMicro {
	return NewLayoutType[datetimeMicroType](c)
}
func NewDateTimeNano(c Carbon) DateTimeNano {
	return NewLayoutType[datetimeNanoType](c)
}

func NewDate(c Carbon) Date {
	return NewLayoutType[dateType](c)
}
func NewDateMilli(c Carbon) DateMilli {
	return NewLayoutType[dateMilliType](c)
}
func NewDateMicro(c Carbon) DateMicro {
	return NewLayoutType[dateMicroType](c)
}
func NewDateNano(c Carbon) DateNano {
	return NewLayoutType[dateNanoType](c)
}

func NewTime(c Carbon) Time {
	return NewLayoutType[timeType](c)
}
func NewTimeMilli(c Carbon) TimeMilli {
	return NewLayoutType[timeMilliType](c)
}
func NewTimeMicro(c Carbon) TimeMicro {
	return NewLayoutType[timeMicroType](c)
}
func NewTimeNano(c Carbon) TimeNano {
	return NewLayoutType[timeNanoType](c)
}

func (t timestampType) Precision() string {
	return PrecisionSecond
}

func (t timestampMilliType) Precision() string {
	return PrecisionMillisecond
}

func (t timestampMicroType) Precision() string {
	return PrecisionMicrosecond
}

func (t timestampNanoType) Precision() string {
	return PrecisionNanosecond
}

func (t datetimeType) Layout() string {
	return DateTimeLayout
}

func (t datetimeMilliType) Layout() string {
	return DateTimeMilliLayout
}

func (t datetimeMicroType) Layout() string {
	return DateTimeMicroLayout
}

func (t datetimeNanoType) Layout() string {
	return DateTimeNanoLayout
}

func (t dateType) Layout() string {
	return DateLayout
}

func (t dateMilliType) Layout() string {
	return DateMilliLayout
}

func (t dateMicroType) Layout() string {
	return DateMicroLayout
}

func (t dateNanoType) Layout() string {
	return DateNanoLayout
}

func (t timeType) Layout() string {
	return TimeLayout
}

func (t timeMilliType) Layout() string {
	return TimeMilliLayout
}

func (t timeMicroType) Layout() string {
	return TimeMicroLayout
}

func (t timeNanoType) Layout() string {
	return TimeNanoLayout
}
