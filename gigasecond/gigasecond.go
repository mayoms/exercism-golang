// Package clause.
package gigasecond
import "time"

const testVersion = 4 

func AddGigasecond(t time.Time) time.Time {

  oneBSec := time.Duration(time.Second) * 1e9
  t = t.Add(oneBSec)
  return t

}

