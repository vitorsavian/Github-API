package value

import "fmt"

type CustomError struct {
	key         string
	title       string
	description string
}

func NewCustomError(key, title, description string) *CustomError {
	return &CustomError{
		key:         key,
		title:       title,
		description: description,
	}
}

func (ref *CustomError) Error() string {
	return "key: " + ref.Key() + " | title: " + ref.Title() + " | description: " + ref.Description()
}

func (ref *CustomError) Key() string {
	return ref.key
}

func (ref *CustomError) Title() string {
	return ref.title
}

func (ref *CustomError) Description() string {
	return ref.description
}

func (ref *CustomError) Wrap(key, title, description string) {
	if key != "" {
		ref.key = fmt.Sprintf("%s. %s", key, ref.key)
	}
	if title != "" {
		ref.title = fmt.Sprintf("%s. %s", title, ref.title)
	}
	if description != "" {
		ref.description = fmt.Sprintf("%s. %s", description, ref.description)
	}
}
