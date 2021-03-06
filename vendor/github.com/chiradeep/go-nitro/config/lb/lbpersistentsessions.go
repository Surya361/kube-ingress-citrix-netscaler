package lb

type Lbpersistentsessions struct {
	Destip               string `json:"destip,omitempty"`
	Destipv6             string `json:"destipv6,omitempty"`
	Destport             int    `json:"destport,omitempty"`
	Flags                bool   `json:"flags,omitempty"`
	Persistenceparam     string `json:"persistenceparam,omitempty"`
	Persistenceparameter string `json:"persistenceparameter,omitempty"`
	Referencecount       int    `json:"referencecount,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srcipv6              string `json:"srcipv6,omitempty"`
	Timeout              int    `json:"timeout,omitempty"`
	Type                 int    `json:"type,omitempty"`
	Vserver              string `json:"vserver,omitempty"`
	Vservername          string `json:"vservername,omitempty"`
}
