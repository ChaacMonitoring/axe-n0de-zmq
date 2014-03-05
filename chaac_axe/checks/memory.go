package chaac_axe_checks

import "github.com/abhishekkr/gol/golbin"

func MemoryCheck() map[string]string{
  return map[string]string{
    "mem-total": golbin.MemInfo("MemTotal"),
    "mem-free": golbin.MemInfo("MemFree"),
  }
}
