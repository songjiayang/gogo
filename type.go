package gogo

type RunMode string

func (mode RunMode) IsValid() bool {
	switch mode {
	case Development, Test, Production:
		return true
	}

	return false
}

func (mode RunMode) IsDevelopment() bool {
	return mode == Development
}

func (mode RunMode) IsTest() bool {
	return mode == Test
}

func (mode RunMode) IsProduction() bool {
	return mode == Production
}

func (mode RunMode) String() string {
	return string(mode)
}
