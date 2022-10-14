package pathhelpercore

type CompiledPathWithErrorWrapper struct {
	*CompiledPath
	*BaseErrorWrapper
}
