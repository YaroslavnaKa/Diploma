package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	d, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}
	if repeat == "" {
		return "", errors.New("repeat is required")
	}
	r := strings.Split(repeat, " ")
	sr := r[0]

	if sr == "d" {
		if len(r) < 2 {
			return "", errors.New("wrong format for d")
		}
		days, er := strconv.Atoi(r[1])
		if er != nil {
			return "", er
		}
		if days > 400 {
			return "", errors.New("day out of range")
		}
		d = d.AddDate(0, 0, days)
		for !d.After(now) {
			d = d.AddDate(0, 0, days)
		}
		return d.Format(DateFormat), nil
	}
	if sr == "y" {
		d = d.AddDate(1, 0, 0)
		for !d.After(now) {
			d = d.AddDate(1, 0, 0)
		}
		return d.Format(DateFormat), nil
	}
	if sr == "w" {
		if len(r) < 2 {
			return "", errors.New("day out of range")
		}
		myDays := make(map[int]bool)
		dayStrings := strings.Split(r[1], ",")
		for _, v := range dayStrings {
			day, e := strconv.Atoi(v)
			if e != nil {
				return "", e
			}
			if day < 1 || day > 7 {
				return "", errors.New("day out of range")
			}

			myDays[day] = true
		}
		for {
			d = d.AddDate(0, 0, 1)
			wkDay := int(d.Weekday())
			if wkDay == 0 {
				wkDay = 7
			}
			if myDays[wkDay] && d.After(now) {
				break
			}
		}
		return d.Format(DateFormat), nil

	}
	if sr == "m" {
		days := make(map[int]bool)
		month := make(map[int]bool)
		if len(r) < 2 {
			return "", errors.New("wrong format")

		}
		daysStrings := strings.Split(r[1], ",")
		for _, v := range daysStrings {
			dni, e := strconv.Atoi(v)
			if e != nil {
				return "", e
			}
			if dni < -2 || dni == 0 || dni > 31 {
				return "", errors.New("wrong format")
			}
			days[dni] = true
		}

		if len(r) == 3 {
			monthStrings := strings.Split(r[2], ",")
			for _, v := range monthStrings {
				mon, e := strconv.Atoi(v)
				if e != nil {
					return "", e
				}
				if mon < 0 || mon > 12 {
					return "", errors.New("month out of range")
				}
				month[mon] = true
			}
		}
		for {
			d = d.AddDate(0, 0, 1)
			if !d.After(now) {
				continue
			}
			if len(month) > 0 {
				mNow := int(d.Month())
				if !month[mNow] {
					continue
				}
			}
			if days[d.Day()] {
				return d.Format(DateFormat), nil
			}
			nextDay := d.AddDate(0, 0, 1)
			isLastDay := nextDay.Month() != d.Month()
			if isLastDay {
				if days[-1] {
					return d.Format(DateFormat), nil
				}
			}

			dayAfterNext := d.AddDate(0, 0, 2)
			isPenultimateDay := dayAfterNext.Month() != d.Month()

			if isPenultimateDay && !isLastDay {
				if days[-2] {
					return d.Format(DateFormat), nil
				}
			}
		}

	}
	return "", errors.New("wrong format")
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(DateFormat, nowStr)
		if err != nil {
			http.Error(w, "wrong format date now", http.StatusBadRequest)
			return
		}
	}

	nextDate, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write([]byte(nextDate))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
