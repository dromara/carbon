package carbon

// LayoutTyper defines a LayoutTyper interface
type LayoutTyper interface {
	~string
	Layout() string
}

// FormatTyper defines a FormatTyper interface.
type FormatTyper interface {
	~string
	Format() string
}

// TimestampTyper defines a TimestampTyper interface.
type TimestampTyper interface {
	~int64
	Precision() string
}
