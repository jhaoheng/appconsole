package module

import "embed"

var Resource *embed.FS

type ResourceOP struct {
	Resource *embed.FS
}

func NewResourceOP(resource *embed.FS) *ResourceOP {
	return &ResourceOP{
		Resource: resource,
	}
}

func (r *ResourceOP) GetImage(file string) []byte {
	data, err := r.Resource.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return data
}
