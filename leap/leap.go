package leap

const testVersion = 2

func IsLeapYear(year int) bool {
  // If year is divisible by 4, then maybe..
  if year % 4 == 0 {
    // if year is divisible by 100 no, unless
    // also divisible by 400
    if year % 100 == 0 {
      if year % 400 == 0 {
        return true
      } else {
        return false
      }
    }
    // all other cases divisble by 4 are leap years
    return true
  }
  // if not divisible by 4, always no
  return false
}
