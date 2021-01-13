package settings

type Size struct {
    value string
}

func NewSize(value string) *Size {
    return &Size{value: value}
}

func (size *Size) Source() (interface{}, error) {
    return size.value, nil
}

/*func (size Size) Validate() (bool, error) {
    re := regexp.MustCompile(`(?ms)^[0-9]+(mb|gb)$`)

    if !re.Match([]byte(size)) {
        return false, ErrorSizeNotValid(errors.New(string(size + " is not valid interval")))
    }

    return true, nil
}*/
