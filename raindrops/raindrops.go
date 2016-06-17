package raindrops
import "strconv"

const testVersion = 2

func Convert(i int) string {

  var s string

  if (i % 3 == 0) {
    s += "Pling"
  } 
  if (i % 5 == 0) {
    s += "Plang"
  }
  if (i % 7 == 0) {
    s += "Plong"
  }
  if (s == "") {
    s = strconv.Itoa(i)
  }
  
  return s

}