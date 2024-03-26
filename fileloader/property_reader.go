package fileloader

import "github.com/magiconair/properties"

func ReadPropertiesFile(path string) (map[string]string, error) {

	p, err := properties.LoadFile(path, properties.UTF8)

	if err != nil {
		return nil, err
	}

	return p.Map(), nil
}

func ReadPropertiesInto(path string, target interface{}) error {

	p, err := properties.LoadFile(path, properties.UTF8)

	if err != nil {
		return err
	}

	return p.Decode(target)
}
