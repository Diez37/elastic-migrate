package settings

type Interval struct {
    value string
}

func NewInterval(value string) *Interval {
    return &Interval{value: value}
}

func (interval *Interval) Source() (interface{}, error) {
    return interval.value, nil
}

/*func (interval Interval) Validate() (bool, error) {
   re := regexp.MustCompile(`(?ms)^[0-9]+(s|m|h)$`)

   if !re.Match([]byte(interval)) {
       return false, ErrorIntervalNotValid(errors.New(string(interval + " is not valid interval")))
   }

   return true, nil
}*/
