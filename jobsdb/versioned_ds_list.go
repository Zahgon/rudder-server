package jobsdb

import (
	"errors"
	"sync"
)

var errInvalidDSListDrainVersion = errors.New("drain version must be older than current version")

func newVersionedDSList(list dataSetTList, rangeList dataSetRangeTList) *versionedDSList {
	_ = "STUB: not implemented"
	return nil
}

type versionedDSList struct {
	mu sync.Mutex

	// version is the version assigned to readers reading the currently published snapshots.
	version uint64

	// list and rangeList are the currently published immutable snapshots.
	// They are owned by versionedDSList so read and update cannot drift apart.
	list      dataSetTList
	rangeList dataSetRangeTList

	// readers counts readers per version. A missing key means no reader is using that version.
	readers map[uint64]int

	// waiters are operations waiting for all readers at or before a version to drain.
	waiters []dsListDrainWaiter
}

type dsListDrainWaiter struct {
	through uint64        // through is inclusive: a drop for version N must wait for readers in versions 0..N.
	drained chan struct{} // closed when there are no more readers at or before through
}

// get returns the currently published snapshots and a release function that must be called when the reader is done with the snapshots.
func (v *versionedDSList) get() (list dataSetTList, ranges dataSetRangeTList, version uint64, release func()) {
	_ = "STUB: not implemented"
	return *new(dataSetTList), *new(dataSetRangeTList), 0, nil
}

func (v *versionedDSList) snapshot() (dataSetTList, dataSetRangeTList) {
	_ = "STUB: not implemented"
	return *new(dataSetTList), *new(dataSetRangeTList)
}

func (v *versionedDSList) currentVersion() uint64 { _ = "STUB: not implemented"; return 0 }

func (v *versionedDSList) set(list dataSetTList, ranges dataSetRangeTList) {
	_ = "STUB: not implemented"
	return
}

// wait returns a channel that will be closed when there are no more readers at or before through. through must be less than the current version.
func (v *versionedDSList) wait(through uint64) (<-chan struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// closeDrainedLocked closes the drained channels of all waiters that are waiting for versions that have no more readers. It must be called with v.mu held.
func (v *versionedDSList) closeDrainedLocked() { _ = "STUB: not implemented"; return }
