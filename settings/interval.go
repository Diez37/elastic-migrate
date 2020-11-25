package settings

import (
    "errors"
    "regexp"
)

type Interval string

func (interval Interval) Validate() (bool, error) {
    re := regexp.MustCompile(`(?ms)^[0-9]+(s|m|h)$`)

    if !re.Match([]byte(interval)) {
        return false, ErrorIntervalNotValid(errors.New(string(interval + " is not valid interval")))
    }

    return true, nil
}
