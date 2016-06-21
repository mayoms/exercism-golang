package bob
import "strings"


const testVersion = 2

const question = "Sure."
const yells = "Whoa, chill out!"
const silence = "Fine. Be that way!"
const other = "Whatever."

func Hey(prompt string) string {

  prompt = strings.TrimSpace(prompt)

  if len(prompt) == 0 {
    return silence
  } else if prompt == strings.ToUpper(prompt) && prompt != strings.ToLower(prompt) {
    return yells
  } else if prompt[len(prompt)-1] == '?' {
    return question
  }

  return other
}


