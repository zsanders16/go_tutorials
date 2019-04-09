package form

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
}

func fields(strct interface{}) field {

	return field{
		Label:       "Name",
		Name:        "Name",
		Type:        "text",
		Placeholder: "Name",
	}
}
