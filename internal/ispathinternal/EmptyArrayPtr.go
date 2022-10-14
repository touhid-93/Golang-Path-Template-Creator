package ispathinternal

func EmptyArrayPtr(paths []*string) bool {
	return paths == nil || &paths == nil || len(paths) == 0
}
