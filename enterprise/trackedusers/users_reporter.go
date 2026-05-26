package trackedusers

import (
	"context"
	"time"

	"github.com/segmentio/go-hll"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
	"github.com/rudderlabs/rudder-go-kit/stats"

	"github.com/rudderlabs/rudder-server/jobsdb"
	txn "github.com/rudderlabs/rudder-server/utils/tx"
)

const (
	idTypeUserID                = "userID"
	idTypeAnonymousID           = "anonymousID"
	idTypeIdentifiedAnonymousID = "identifiedAnonymousID"

	// changing this will be non backwards compatible
	murmurSeed = 123

	trackUsersTable = "tracked_users_reports"

	eventTypeAlias = "alias"
)

type UsersReport struct {
	WorkspaceID              string
	SourceID                 string
	UserIDHll                *hll.Hll
	AnonymousIDHll           *hll.Hll
	IdentifiedAnonymousIDHll *hll.Hll
}

// UsersReporter is interface to report unique users from reports
type UsersReporter interface {
	ReportUsers(ctx context.Context, reports []*UsersReport, tx *txn.Tx) error
	GenerateReportsFromJobs(jobs []*jobsdb.JobT, sourceIdFilter map[string]bool) []*UsersReport
	MigrateDatabase(dbConn string, conf *config.Config) error
}

type UniqueUsersReporter struct {
	log         logger.Logger
	hllSettings *hll.Settings
	instanceID  string
	now         func() time.Time
	stats       stats.Stats
}

func NewUniqueUsersReporter(log logger.Logger, conf *config.Config, stats stats.Stats) (*UniqueUsersReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *UniqueUsersReporter) MigrateDatabase(dbConn string, conf *config.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *UniqueUsersReporter) GenerateReportsFromJobs(jobs []*jobsdb.JobT, sourceIDtoFilter map[string]bool) []*UsersReport {
	_ = "STUB: not implemented"
	return nil
}

// for alias event we will be adding previousId to identifiedAnonymousID hll,
// so for calculating unique users we do not double count the user
// we also add previousId to userId hll to avoid negative cardinality
// e.g. we receive events
// {type:track, anonymousID: anon1}
// {type:track, userID: user1}
// {type:track, userID: user2}
// {type:identify, userID: user1, anonymousID: anon1}
// {type:alias, previousId: user2, userID: user1}
// userHLL: {user1, user2, anon1}, identifiedAnonHLL: {anon1, user2}
// cardinality: len(userHLL)-len(identifiedAnonHLL): 3-2 = 1

func (u *UniqueUsersReporter) ReportUsers(ctx context.Context, reports []*UsersReport, tx *txn.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// convert hll to hexadecimal encoding
func (u *UniqueUsersReporter) hllToString(hllStruct *hll.Hll) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (u *UniqueUsersReporter) recordIdentifier(idTypeHllMap map[string]*hll.Hll, identifier, identifierType string) map[string]*hll.Hll {
	_ = "STUB: not implemented"
	return nil
}

func (u *UniqueUsersReporter) recordHllSizeStats(report *UsersReport) {
	_ = "STUB: not implemented"
	return
}
