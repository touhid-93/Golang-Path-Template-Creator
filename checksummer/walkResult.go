package checksummer

// Code from: https://blog.golang.org/pipelines#TOC_9.
// Still not huge performance boost, because of single threaded directory traversal
// If we use go routine, the performance boost is visible
// With larger files this version should be faster, because the hasher will be the
// dominating factor, hot the directory traversal
// With larger files sizes, it's 3~4 times faster
// A walkResult is the product of reading and summing a file using MD5.
type walkResult struct {
	path string
	sum  []byte
	err  error
}
