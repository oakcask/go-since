package since

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

type since struct {
	Days   int
	Weeks  int
	Months int
	Years  int
}

var periodFormat *regexp.Regexp

func init() {
	periodFormat = regexp.MustCompile(`(?:(\d+)\s*(?:y|yrs?|years?))?(?:(\d+)\s*(?:mo|months?))?(?:(\d+)\s*(?:w|weeks?))?(?:(\d+)\s*(?:d|days?))?`)
}

func parseNumber(s string) int {
	if s == "" {
		return 0
	}

	// won't fail
	n, _ := strconv.Atoi(s)
	return n
}

func normalize(s since) since {
	mo := s.Weeks / 4
	yr := (s.Months + mo) / 12

	return since{
		Days:   s.Days,
		Weeks:  (s.Weeks % 4),
		Months: (s.Months + mo) % 12,
		Years:  s.Years + yr,
	}
}

func parse(period string) (since, error) {
	matches := periodFormat.FindStringSubmatch(period)
	if matches == nil {
		return since{}, errors.New("invalid format")
	}

	s := since{
		Years:  parseNumber(matches[1]),
		Months: parseNumber(matches[2]),
		Weeks:  parseNumber(matches[3]),
		Days:   parseNumber(matches[4]),
	}

	if (s == since{}) {
		return since{}, errors.New("empty period")
	}

	return normalize(s), nil
}

func shift(from time.Time, sign int, s since) time.Time {
	return from.AddDate(sign*s.Years, sign*s.Months, sign*s.Weeks*7).AddDate(0, 0, sign*s.Days)
}

func Since(s string, now time.Time) (time.Time, error) {
	v, e := parse(s)
	if e != nil {
		return time.Time{}, e
	}

	return shift(now, -1, v), nil
}
