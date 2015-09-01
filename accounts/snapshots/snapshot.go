package snapshots

import "github.com/victor-castellanos/compz/dollars"

type Snapshot struct {
	Year int32
	TotalGrowth dollars.Dollar
	CurrentGrowth dollars.Dollar
	TotalInvested dollars.Dollar
	CurrentInvested dollars.Dollar
}
