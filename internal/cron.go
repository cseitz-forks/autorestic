package internal

import (
	"github.com/cseitz-forks/autorestic/internal/colors"
	"github.com/cseitz-forks/autorestic/internal/flags"
)

func RunCron() error {
	if flags.CRON_DRY == true {
		return DryRunCron()
	}

	c := GetConfig()
	for name, l := range c.Locations {
		l.name = name
		if err := l.RunCron(); err != nil {
			return err
		}
	}
	return nil
}

func DryRunCron() error {
	c := GetConfig()
	dueCount := 0
	for name, l := range c.Locations {
		l.name = name
		shouldRun, err := l.IsCronDue()
		if err != nil {
			return err
		} else if shouldRun == true {
			colors.Body.Printf("Location \"%s\" is due\n", l.name)
			dueCount++
		}
	}
	if dueCount > 0 {
		colors.Body.Printf("There are \"%d\" locations are due\n", dueCount)
		colors.Body.Print("Cron is ready\n")
	} else {
		colors.Body.Printf("No cron jobs are due\n")
	}
	return nil
}
