package since

import (
	"testing"
	"time"
)

func TestBack(t *testing.T) {
	testCases := []struct {
		from      time.Time
		since     string
		want      time.Time
		wantError error
	}{
		{from: stdISO8601("20220101T000000+0900"), since: "1d", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1day", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1days", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1 d", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1 day", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1 days", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		// 7 days does not round up to weeks
		{from: stdISO8601("20220128T000000+0900"), since: "28 days", want: stdISO8601("20211231T000000+0900"), wantError: nil},
		{from: stdISO8601("20220307T000000+0900"), since: "1w", want: stdISO8601("20220228T000000+0900"), wantError: nil},
		{from: stdISO8601("20220307T000000+0900"), since: "1week", want: stdISO8601("20220228T000000+0900"), wantError: nil},
		{from: stdISO8601("20220307T000000+0900"), since: "1weeks", want: stdISO8601("20220228T000000+0900"), wantError: nil},
		{from: stdISO8601("20220307T000000+0900"), since: "1 w", want: stdISO8601("20220228T000000+0900"), wantError: nil},
		{from: stdISO8601("20220307T000000+0900"), since: "1 week", want: stdISO8601("20220228T000000+0900"), wantError: nil},
		{from: stdISO8601("20220307T000000+0900"), since: "1 weeks", want: stdISO8601("20220228T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1mo", want: stdISO8601("20211201T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1month", want: stdISO8601("20211201T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1months", want: stdISO8601("20211201T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1 mo", want: stdISO8601("20211201T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1 month", want: stdISO8601("20211201T000000+0900"), wantError: nil},
		{from: stdISO8601("20220101T000000+0900"), since: "1 months", want: stdISO8601("20211201T000000+0900"), wantError: nil},
		// 4 weeks will be round up to 1 month
		{from: stdISO8601("20220228T000000+0900"), since: "4 weeks", want: stdISO8601("20220128T000000+0900"), wantError: nil},
		{from: stdISO8601("20220228T000000+0900"), since: "8 weeks", want: stdISO8601("20211228T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1y", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1yr", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1yrs", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1year", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1years", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1 y", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1 yr", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1 yrs", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1 year", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		{from: stdISO8601("20000229T000000+0900"), since: "1 years", want: stdISO8601("19990301T000000+0900"), wantError: nil},
		// 12 months will be round up to 1 year
		{from: stdISO8601("20000229T000000+0900"), since: "12mo", want: stdISO8601("19990301T000000+0900"), wantError: nil},
	}

	for idx, tc := range testCases {
		actual, err := Since(tc.since, tc.from)
		if tc.wantError != nil {
			if err == nil || err.Error() != tc.wantError.Error() {
				t.Errorf("%v: wants [%+v] but got [%+v] error", idx, tc.wantError, err)
			}
		} else {
			if err != nil {
				t.Errorf("%v: wants no errors but got [%+v] error", idx, err)
			}
		}
		if actual != tc.want {
			t.Errorf("%v: wants %+v but got %+v", idx, tc.want, actual)
		}
	}
}

func stdISO8601(s string) time.Time {
	t, _ := time.Parse("20060102T150405-0700", s)
	return t
}
