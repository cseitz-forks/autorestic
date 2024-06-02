package flags

var (
	CI           bool = false
	VERBOSE      bool = false
	CRON_LEAN    bool = false
	CRON_DRY     bool = false
	RESTIC_BIN   string
	DOCKER_IMAGE string
)
