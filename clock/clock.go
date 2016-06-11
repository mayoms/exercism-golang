package clock
import "fmt"


// as an aside, the stub code indicated testVersion 3, not sure why
// as the test suite appears to be 4.
const testVersion = 4

// Go doesn't have a real modulus operator: '%' is remainder, 
// so it doesn't work backwords - this is to provide that 
// ability.

func mod(x, y int) int {

  mod := x % y; 
  if mod < 0 { 
    mod += y 
  }
  return mod
}


// struct with hours and minutes seems appropriate.

type Clock struct { h, m int}

// returns a new Clock

func New(h, m int) Clock {

  // if minutes are negative, we need to subtract 60
  // to adjust for the impact on hours. Otherwise
  // we end up being an hour off.
  if m < 0 {
    m -= 60
  }
  h += m / 60
  m = mod(m, 60)
  h = mod(h, 24)
  return Clock { h, m }
}

// returns string representation of Clock

func (c Clock) String() string {

  return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Adds Minutes to Clock

func (c Clock) Add(minutes int) Clock {

  return New( c.h, c.m+minutes )
}
