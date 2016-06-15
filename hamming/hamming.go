package hamming
import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {

  hamming := 0
  if (len(a) != len(b)) {
    return -1, errors.New("Strings not the same length.")
  }
  if (a == b) {
    return hamming, nil
  }
  for index := range a {
    if (a[index] != b[index]) {
      hamming += 1
    }
  }
  return hamming, nil
}
