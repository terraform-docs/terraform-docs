package doc

// Module represents a Terraform module.
type Module struct {
	Name     string
	Source   string
	basepath string
}

func (o *Module) GetBasepath() string {
	return o.basepath
}
