package kata

//import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
  // Implement me
  if dadYearsOld > (sonYearsOld*2) {
    return dadYearsOld - (sonYearsOld*2)
  } else if dadYearsOld < sonYearsOld*2 {
    return (sonYearsOld*2) - dadYearsOld
  } else {
    return 0
  }
}