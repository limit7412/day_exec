package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/limit7412/day_exec/app/repository"
	constants "github.com/limit7412/day_exec/internal"
)

func Run(imputCmd string, year, month, day, period, times int) error {
	from := time.Date(year, time.Month(month), day, 0, 0, 0, 0, constants.JSTLoc)
	to := from.AddDate(0, 0, period*times)

	for target := from; target.Unix() < to.Unix(); target = target.AddDate(0, 0, period) {
		cmd := strings.Replace(strings.Replace(strings.Replace(imputCmd,
			"<yyyy>", fmt.Sprintf("%04d", target.Year()), -1),
			"<mm>", fmt.Sprintf("%02d", int(target.Month())), -1),
			"<dd>", fmt.Sprintf("%02d", target.Day()), -1)

		err := repository.Exec(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}
