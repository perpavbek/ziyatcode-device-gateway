package mdns

type ServiceEntry struct {
	Instance  string
	Service   string
	Domain    string
	Addresses []string
	Text      map[string]string
}