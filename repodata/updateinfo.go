package repodata


type Reference struct {
	Text      string `xml:",chardata"`
	Reference struct {
		Text  string `xml:",chardata"`
		Href  string `xml:"href,attr"`
		ID    string `xml:"id,attr"`
		Type  string `xml:"type,attr"`
		Title string `xml:"title,attr"`
	} `xml:"reference"`
}

type Update struct {
	From    string `xml:"from,attr"`
	Status  string `xml:"status,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Issued  struct {
		Date string `xml:"date,attr"`
	} `xml:"issued"`
	Updated struct {
		Date string `xml:"date,attr"`
	} `xml:"updated"`
	Rights      string      `xml:"rights"`
	Release     string      `xml:"release"`
	Severity    string      `xml:"severity"`
	Summary     string      `xml:"summary"`
	Description string      `xml:"description"`
	References  []Reference `xml:"references"`
	Pkglist     struct {
		Collection struct {
			Short    string `xml:"short,attr"`
			Name     string `xml:"name"`
			Packages []struct {
				Name     string `xml:"name,attr"`
				Version  string `xml:"version,attr"`
				Release  string `xml:"release,attr"`
				Epoch    string `xml:"epoch,attr"`
				Arch     string `xml:"arch,attr"`
				Src      string `xml:"src,attr"`
				Filename string `xml:"filename"`
			} `xml:"package"`
		} `xml:"collection"`
	} `xml:"pkglist"`
}

type Updateinfo struct {
	Update []Update `xml:"update"`
}

