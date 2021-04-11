package repodata

type Package struct {
	Type    string `xml:"type,attr"`
	Name    string `xml:"name"`
	Arch    string `xml:"arch"`
	Version struct {
		Epoch string `xml:"epoch,attr"`
		Ver   string `xml:"ver,attr"`
		Rel   string `xml:"rel,attr"`
	} `xml:"version"`
	Checksum struct {
		Type  string `xml:"type,attr"`
		Pkgid string `xml:"pkgid,attr"`
	} `xml:"checksum"`
	Summary     string `xml:"summary"`
	Description string `xml:"description"`
	Packager    string `xml:"packager"`
	URL         string `xml:"url"`
	Time        struct {
		File  string `xml:"file,attr"`
		Build string `xml:"build,attr"`
	} `xml:"time"`
	Size struct {
		Package   string `xml:"package,attr"`
		Installed string `xml:"installed,attr"`
		Archive   string `xml:"archive,attr"`
	} `xml:"size"`
	Location struct {
		Href string `xml:"href,attr"`
	} `xml:"location"`
}

type Primary struct {
	Count   int       `xml:"packages,attr"`
	Package []Package `xml:"package"`
}
