package spec

type Data struct {
	Location         string `yaml:"location"`
	ReadOnly         bool   `yaml:"readOnly,omitempty"`
	KMSDecryptPolicy string `yaml:"kmsDecryptPolicy,omitempty"`
}
