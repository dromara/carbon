---
head:
  - - meta
    - name: description
      content: Detailed introduction to time and date format symbols supported by Format and ParseByFormat methods, including symbol description, length, range and examples
  - - meta
    - name: keywords
      content: format signs, Format, ParseByFormat, time format, date format, golang, carbon, time processing
---

# Format signs
`Format` and `ParseByFormat` methods are not completely reversible. `Format` method supports all of the following signs, while `ParseByFormat` method does not support `K`, `W`, `N`, `L`, `w`, `t`, `o`, `q`, `c` because these signs are custom and do not have corresponding signs in the standard time package.

| sign |                                                  desc                                                  | length |      range       |       example       |
|:----:|:------------------------------------------------------------------------------------------------------:|:------:|:----------------:|:-------------------:|
|  d   |                                     Day of the month, padded to 2                                      |   2    |      01-31       |         02          |
|  D   |                           Day of the week, as an abbreviate localized string                           |   3    |     Mon-Sun      |         Mon         |
|  j   |                                      Day of the month, no padding                                      |   -    |       1-31       |          2          |
|  K   | English ordinal suffix for the day of the month, 2 characters. Eg: st, nd, rd or th. Works well with j |   2    |   st/nd/rd/th    |         th          |
|  l   |                         Day of the week, as an unabbreviated localized string                          |   -    |  Monday-Sunday   |       Monday        |
|  F   |                               Month as an unabbreviated localized string                               |   -    | January-December |       January       |
|  m   |                                           Month, padded to 2                                           |   2    |      01-12       |         01          |
|  M   |                                Month as an abbreviated localized string                                |   3    |     Jan-Dec      |         Jan         |
|  n   |                                           Month, no padding                                            |   -    |       1-12       |          1          |
|  Y   |                                            Four-digit year                                             |   4    |    0000-9999     |        2006         |
|  y   |                                             Two-digit year                                             |   2    |      00-99       |         06          |
|  a   |                                  Lowercase morning or afternoon sign                                   |   2    |      am/pm       |         pm          |
|  A   |                                  Uppercase morning or afternoon sign                                   |   2    |      AM/PM       |         PM          |
|  g   |                                   Hour in 12-hour format, no padding                                   |   -    |       1-12       |          3          |
|  G   |                                   Hour in 24-hour format, no padding                                   |   -    |       0-23       |         15          |
|  h   |                                  Hour in 12-hour format, padded to 2                                   |   2    |      00-11       |         03          |
|  H   |                                  Hour in 24-hour format, padded to 2                                   |   2    |      00-23       |         15          |
|  i   |                                          Minute, padded to 2                                           |   2    |      01-59       |         04          |
|  s   |                                          Second, padded to 2                                           |   2    |      01-59       |         05          |
|  O   |               Difference to Greenwich time (GMT) without colon between hours and minutes               |   -    |        -         |        -0700        |
|  P   |                Difference to Greenwich time (GMT) with colon between hours and minutes                 |   -    |        -         |       -07:00        |
|  Z   |                                               Zone name                                                |   -    |        -         |         MST         |
|  W   |                                            week of the year                                            |   -    |       1-52       |          1          |
|  N   |                                            day of the week                                             |   1    |       1-7        |          2          |
|  L   |                                        Whether it's a leap year                                        |   1    |       0-1        |          0          |
|  S   |                                       Unix timestamp with second                                       |   -    |        -         |     1596604455      |
|  U   |                               Unix timestamp with millisecond precision                                |   -    |        -         |    1596604455666    |
|  V   |                               Unix timestamp with microsecond precision                                |   -    |        -         |  1596604455666666   |
|  X   |                                Unix timestamp with nanosecond precision                                |   -    |        -         | 1596604455666666666 |
|  u   |                                              Millisecond                                               |   -    |      1-999       |         999         |
|  v   |                                              Microsecond                                               |   -    |     1-999999     |       999999        |
|  x   |                                               Nanosecond                                               |   -    |   1-999999999    |      999999999      |
|  w   |                                            Day of the week                                             |   1    |       0-6        |          1          |
|  t   |                                        Total days of the month                                         |   2    |      28-31       |         31          |
|  z   |                                               Time zone                                                |   -    |        -         |    Asia/Shanghai    |
|  o   |                                              Time offset                                               |   -    |        -         |        28800        |
|  q   |                                                Quarter                                                 |   1    |       1-4        |          1          |
|  c   |                                                Century                                                 |   -    |       0-99       |         21          |
