package partitionbuffer

// startBufferWatchdog starts a watchdog that periodically checks for unbuffered partitions that have buffered jobs in the buffer JobsDB and flushes them.
func (b *jobsDBPartitionBuffer) startBufferWatchdog() { _ = "STUB: not implemented"; return }

// run performs a single iteration of the watchdog logic, returning true if a partition was flushed:
// it gets the first unprocessed job from the buffer JobsDB, checks if its partition is buffered,
// and if not, flushes the partition.
// It returns true if a partition was flushed, false otherwise.

// try to find a job in buffered JobsDB

// no jobs found

// before moving each batch, check if the partition got buffered in the meantime, and if so, stop moving to avoid conflicts with the buffering process

// keep flushing partitions as long as run keeps finding unbuffered partitions with buffered jobs
