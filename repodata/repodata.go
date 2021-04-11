package repodata

type RepoMDItem struct {
	Type     string `xml:"type,attr"`
	Checksum struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"checksum"`
	OpenChecksum struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"open-checksum"`
	Location struct {
		Href string `xml:"href,attr"`
	} `xml:"location"`
	Timestamp      string `xml:"timestamp"`
	Size           string `xml:"size"`
	OpenSize       string `xml:"open-size"`
	HeaderChecksum struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"header-checksum"`
	HeaderSize string `xml:"header-size"`
}

type RepoMD struct {
	Revision uint         `xml:"revision"`
	Data     []RepoMDItem `xml:"data"`
}
