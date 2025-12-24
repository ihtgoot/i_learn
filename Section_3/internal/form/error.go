package form

type error map[string][]string

// add will add an error message to specify form field
func (e error) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get return the firdt error messssage from teh specifuc form filed
func (e error) Get(field string) string {
	errorString := e[field]
	if len(errorString) == 0 {
		return ""
	}
	return errorString[0]
}
