package ispathinternal

func EmptyArray(paths []string) bool {
	return &paths == nil || paths == nil || len(paths) == 0
}
