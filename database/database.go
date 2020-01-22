package database

type DatabaseX interface {
	Close() error
	ListDrafts() ([]Draft, error)
}

type Draft struct {
	// Client ID and secret used to identify the client.
	ID     string `json:"id" yaml:"id"`
	Secret string `json:"secret" yaml:"secret"`

	// A registered set of redirect URIs. When redirecting from dex to the client, the URI
	// requested to redirect to MUST match one of these values, unless the client is "public".
	RedirectURIs []string `json:"redirectURIs" yaml:"redirectURIs"`

	// TrustedPeers are a list of peers which can issue tokens on this client's behalf using
	// the dynamic "oauth2:server:client_id:(client_id)" scope. If a peer makes such a request,
	// this client's ID will appear as the ID Token's audience.
	//
	// Clients inherently trust themselves.
	TrustedPeers []string `json:"trustedPeers" yaml:"trustedPeers"`

	// Public clients must use either use a redirectURL 127.0.0.1:X or "urn:ietf:wg:oauth:2.0:oob"
	Public bool `json:"public" yaml:"public"`

	// Name and LogoURL used when displaying this client to the end user.
	Name    string `json:"name" yaml:"name"`
	LogoURL string `json:"logoURL" yaml:"logoURL"`
}
