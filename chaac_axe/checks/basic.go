package chaac_axe_checks

import (
  golbin "github.com/abhishekkr/gol/golbin"
)

func BasicCheck() map[string]string{
  return map[string]string{
    "uptime": golbin.Uptime(),
  }
}
