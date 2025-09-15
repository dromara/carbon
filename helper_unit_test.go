package carbon

import (
	"testing"
	"time"
)

// Test format2layout function
func TestFormat2layout(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "Empty format",
			format:   "",
			expected: "",
		},
		{
			name:     "Simple date format",
			format:   "Y-m-d",
			expected: "2006-01-02",
		},
		{
			name:     "Date and time format",
			format:   "Y-m-d H:i:s",
			expected: "2006-01-02 15:04:05",
		},
		{
			name:     "Full datetime format",
			format:   "Y-m-d H:i:s.u",
			expected: "2006-01-02 15:04:05.999",
		},
		{
			name:     "Format with escape characters",
			format:   "Y-m-d \\T H:i:s",
			expected: "2006-01-02 T 15:04:05",
		},
		{
			name:     "Format with mixed characters",
			format:   "Y-m-d H:i:s [UTC]",
			expected: "2006-01-02 15:04:05 [unixMilliTC]",
		},
		{
			name:     "All format characters",
			format:   "Y-y-m-n-d-j D-l H-h-g a-A i-s u-v-x",
			expected: "2006-06-01-1-02-2 Mon-Monday 15-03-3 pm-PM 04-05 999-999999-999999999",
		},
		{
			name:     "Timezone formats",
			format:   "O-P-Q-R-Z",
			expected: "-0700--07:00-Z0700-Z07:00-MST",
		},
		{
			name:     "Timestamp formats",
			format:   "S-U-V-X",
			expected: "unix-unixMilli-unixMicro-unixNano",
		},
		{
			name:     "Multiple escape characters",
			format:   "Y\\-m\\-d H\\:i\\:s",
			expected: "2006-01-02 15:04:05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := format2layout(tt.format)
			if result != tt.expected {
				t.Errorf("format2layout(%q) = %q, want %q", tt.format, result, tt.expected)
			}
		})
	}
}

// Test format2layout caching
func TestFormat2layout_Caching(t *testing.T) {
	format := "Y-m-d H:i:s"
	expected := "2006-01-02 15:04:05"

	// First call should convert and cache
	result1 := format2layout(format)
	if result1 != expected {
		t.Fatalf("First format2layout call returned %q, want %q", result1, expected)
	}

	// Second call should return cached result
	result2 := format2layout(format)
	if result2 != expected {
		t.Fatalf("Second format2layout call returned %q, want %q", result2, expected)
	}

	// Cache should contain the format
	if cached, exists := layoutCache.Load(format); !exists {
		t.Errorf("format2layout cache should contain %q", format)
	} else if cached.(string) != expected {
		t.Errorf("format2layout cache contains %q, want %q", cached, expected)
	}
}

// Test format2layout cache limit
func TestFormat2layout_CacheLimit(t *testing.T) {
	// Test short format (should be cached)
	shortFormat := "Y-m-d"
	result := format2layout(shortFormat)
	if result == "" {
		t.Fatalf("format2layout(%q) failed", shortFormat)
	}
	if _, exists := layoutCache.Load(shortFormat); !exists {
		t.Errorf("Short format %q should be cached", shortFormat)
	}

	// Test long format (should not be cached)
	longFormat := "verylongformatstringthatislongerthan50charactersandshouldnotbecached"
	result = format2layout(longFormat)
	// This should succeed, but not be cached
	if result == "" {
		t.Fatalf("format2layout(%q) should have succeeded", longFormat)
	}
	if _, exists := layoutCache.Load(longFormat); exists {
		t.Errorf("Long format %q should not be cached", longFormat)
	}
}

// Test parseTimezone function
func TestParseTimezone(t *testing.T) {
	tests := []struct {
		name        string
		timezone    string
		expectError bool
	}{
		{
			name:        "Empty timezone",
			timezone:    "",
			expectError: true,
		},
		{
			name:        "Valid timezone UTC",
			timezone:    "UTC",
			expectError: false,
		},
		{
			name:        "Valid timezone Local",
			timezone:    "Local",
			expectError: false,
		},
		{
			name:        "Valid timezone America/New_York",
			timezone:    "America/New_York",
			expectError: false,
		},
		{
			name:        "Valid timezone Europe/London",
			timezone:    "Europe/London",
			expectError: false,
		},
		{
			name:        "Valid timezone Asia/Shanghai",
			timezone:    "Asia/Shanghai",
			expectError: false,
		},
		{
			name:        "Invalid timezone",
			timezone:    "Invalid/Timezone",
			expectError: true,
		},
		{
			name:        "Invalid timezone format",
			timezone:    "InvalidTimezone",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc, err := parseTimezone(tt.timezone)

			if tt.expectError {
				if err == nil {
					t.Errorf("parseTimezone(%q) expected error, got nil", tt.timezone)
				}
				if loc != nil {
					t.Errorf("parseTimezone(%q) expected nil location, got %v", tt.timezone, loc)
				}
			} else {
				if err != nil {
					t.Errorf("parseTimezone(%q) unexpected error: %v", tt.timezone, err)
				}
				if loc == nil {
					t.Errorf("parseTimezone(%q) expected location, got nil", tt.timezone)
				}
			}
		})
	}
}

// Test parseTimezone caching
func TestParseTimezone_Caching(t *testing.T) {
	timezone := "UTC"

	// First call should parse and cache
	loc1, err1 := parseTimezone(timezone)
	if err1 != nil {
		t.Fatalf("First parseTimezone call failed: %v", err1)
	}

	// Second call should return cached result
	loc2, err2 := parseTimezone(timezone)
	if err2 != nil {
		t.Fatalf("Second parseTimezone call failed: %v", err2)
	}

	// Should be the same location instance (cached)
	if loc1 != loc2 {
		t.Errorf("parseTimezone caching failed: got different instances")
	}

	// Cache should contain the timezone
	if _, exists := timezoneCache.Load(timezone); !exists {
		t.Errorf("parseTimezone cache should contain %q", timezone)
	}
}

// Test parseDuration function
func TestParseDuration(t *testing.T) {
	tests := []struct {
		name        string
		duration    string
		expectError bool
		expectedDur time.Duration
	}{
		{
			name:        "Empty duration",
			duration:    "",
			expectError: true,
		},
		{
			name:        "Valid duration 1s",
			duration:    "1s",
			expectError: false,
			expectedDur: time.Second,
		},
		{
			name:        "Valid duration 1m",
			duration:    "1m",
			expectError: false,
			expectedDur: time.Minute,
		},
		{
			name:        "Valid duration 1h",
			duration:    "1h",
			expectError: false,
			expectedDur: time.Hour,
		},
		{
			name:        "Valid duration 24h",
			duration:    "24h",
			expectError: false,
			expectedDur: 24 * time.Hour,
		},
		{
			name:        "Valid duration 0.5s",
			duration:    "0.5s",
			expectError: false,
			expectedDur: 500 * time.Millisecond,
		},
		{
			name:        "Valid duration 1.5h",
			duration:    "1.5h",
			expectError: false,
			expectedDur: 90 * time.Minute,
		},
		{
			name:        "Valid duration 2h30m",
			duration:    "2h30m",
			expectError: false,
			expectedDur: 2*time.Hour + 30*time.Minute,
		},
		{
			name:        "Valid duration 1h2m3s",
			duration:    "1h2m3s",
			expectError: false,
			expectedDur: time.Hour + 2*time.Minute + 3*time.Second,
		},
		{
			name:        "Valid duration 1ms",
			duration:    "1ms",
			expectError: false,
			expectedDur: time.Millisecond,
		},
		{
			name:        "Valid duration 1us",
			duration:    "1us",
			expectError: false,
			expectedDur: time.Microsecond,
		},
		{
			name:        "Valid duration 1ns",
			duration:    "1ns",
			expectError: false,
			expectedDur: time.Nanosecond,
		},
		{
			name:        "Invalid duration xyz",
			duration:    "xyz",
			expectError: true,
		},
		{
			name:        "Invalid duration 1x",
			duration:    "1x",
			expectError: true,
		},
		{
			name:        "Invalid duration 1h2x",
			duration:    "1h2x",
			expectError: true,
		},
		{
			name:        "Invalid duration verylongdurationstring",
			duration:    "verylongdurationstring",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur, err := parseDuration(tt.duration)

			if tt.expectError {
				if err == nil {
					t.Errorf("parseDuration(%q) expected error, got nil", tt.duration)
				}
				if dur != 0 {
					t.Errorf("parseDuration(%q) expected 0 duration, got %v", tt.duration, dur)
				}
			} else {
				if err != nil {
					t.Errorf("parseDuration(%q) unexpected error: %v", tt.duration, err)
				}
				if dur != tt.expectedDur {
					t.Errorf("parseDuration(%q) = %v, want %v", tt.duration, dur, tt.expectedDur)
				}
			}
		})
	}
}

// Test parseDuration caching
func TestParseDuration_Caching(t *testing.T) {
	duration := "1s"
	expectedDur := time.Second

	// First call should parse and cache
	dur1, err1 := parseDuration(duration)
	if err1 != nil {
		t.Fatalf("First parseDuration call failed: %v", err1)
	}
	if dur1 != expectedDur {
		t.Fatalf("First parseDuration call returned %v, want %v", dur1, expectedDur)
	}

	// Second call should return cached result
	dur2, err2 := parseDuration(duration)
	if err2 != nil {
		t.Fatalf("Second parseDuration call failed: %v", err2)
	}
	if dur2 != expectedDur {
		t.Fatalf("Second parseDuration call returned %v, want %v", dur2, expectedDur)
	}

	// Cache should contain the duration
	if cached, exists := durationCache.Load(duration); !exists {
		t.Errorf("parseDuration cache should contain %q", duration)
	} else if cached.(Duration) != expectedDur {
		t.Errorf("parseDuration cache contains %v, want %v", cached, expectedDur)
	}
}

// Test parseDuration cache limit
func TestParseDuration_CacheLimit(t *testing.T) {
	// Test short duration (should be cached)
	shortDur := "1s"
	_, err := parseDuration(shortDur)
	if err != nil {
		t.Fatalf("parseDuration(%q) failed: %v", shortDur, err)
	}
	if _, exists := durationCache.Load(shortDur); !exists {
		t.Errorf("Short duration %q should be cached", shortDur)
	}

	// Test long duration (should not be cached)
	longDur := "verylongdurationstringthatislongerthan10characters"
	_, err = parseDuration(longDur)
	// This should fail, so we expect an error
	if err == nil {
		t.Fatalf("parseDuration(%q) should have failed", longDur)
	}
	if _, exists := durationCache.Load(longDur); exists {
		t.Errorf("Long duration %q should not be cached", longDur)
	}
}

// Test getAbsValue function
func TestGetAbsValue(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected int64
	}{
		{
			name:     "Positive value 0",
			value:    0,
			expected: 0,
		},
		{
			name:     "Positive value 1",
			value:    1,
			expected: 1,
		},
		{
			name:     "Positive value 123456",
			value:    123456,
			expected: 123456,
		},
		{
			name:     "Positive value max int64",
			value:    9223372036854775807,
			expected: 9223372036854775807,
		},
		{
			name:     "Negative value -1",
			value:    -1,
			expected: 1,
		},
		{
			name:     "Negative value -123456",
			value:    -123456,
			expected: 123456,
		},
		{
			name:     "Negative value -9223372036854775807",
			value:    -9223372036854775807,
			expected: 9223372036854775807,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getAbsValue(tt.value)
			if result != tt.expected {
				t.Errorf("getAbsValue(%d) = %d, want %d", tt.value, result, tt.expected)
			}
		})
	}
}
