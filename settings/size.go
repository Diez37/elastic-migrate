package settings

import (
    "errors"
    "regexp"
)

type Size string

func (size Size) Validate() (bool, error) {
    re := regexp.MustCompile(`(?ms)^[0-9]+(mb|gb)$`)

    if !re.Match([]byte(size)) {
        return false, ErrorSizeNotValid(errors.New(string(size + " is not valid interval")))
    }

    return true, nil
}
