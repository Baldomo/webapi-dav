package orario

var (
	classi []classe
)

func loadClassi() {
	classi = nil
	var classitemp []classe
	for _, att := range orario.Attivita {
		classitemp = append(classitemp, att.Classe)
	}

	for _, c := range classitemp {
		skip := false
		for _, u := range classi {
			if c == u || c.String() == "" {
				skip = true
				break
			}
		}
		if !skip {
			classi = append(classi, c)
		}
	}
}

func GetAllClassi() *[]classe {
	return &classi
}
