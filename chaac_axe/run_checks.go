package chaac_axe

import (
  golhashmap "github.com/abhishekkr/gol/golhashmap"

  chaac_axe_checks "github.com/ChaacMonitoring/axe-n0de-zmq/chaac_axe/checks"
)


func CheckResult() string{
  var result_hmap golhashmap.HashMap
  result_hmap = make(golhashmap.HashMap)

  all_checks := [...]golhashmap.HashMap{
                  chaac_axe_checks.BasicCheck(),
                  chaac_axe_checks.MemoryCheck(),
                }
  for _, check_n_result := range all_checks {
    for check, result := range check_n_result {
      result_hmap[check] = result
    }
  }

  return golhashmap.Hashmap_to_csv(result_hmap)
}
