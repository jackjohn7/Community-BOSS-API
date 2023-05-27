package utilities

import (
  "strings"
)

func ValidateSeason(s string) bool {
  seasons := [4]string{"winter", "spring", "summer", "fall"}
  for _, valid := range seasons {
    if strings.ToLower(s) == valid {
      return true
    }
  }
  return false
}
