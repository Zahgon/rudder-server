package testhelper

type EventsCountMap map[string]int

func defaultStagingFilesEventsMap() EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

func defaultStagingFilesWithIDResolutionEventsMap() EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

// 32 + 24 (merge events because of ID resolution)

func defaultTableUploadsEventsMap(destType string) EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

func defaultWarehouseEventsMap(destType string) EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

func defaultSourcesStagingFilesEventsMap() EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

func defaultSourcesStagingFilesWithIDResolutionEventsMap() EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

// 8 + 4 (merge events because of ID resolution)

func defaultSourcesTableUploadsEventsMap() EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}

func defaultSourcesWarehouseEventsMap() EventsCountMap {
	_ = "STUB: not implemented"
	return *new(EventsCountMap)
}
