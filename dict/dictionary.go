package dict

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var ErrNotFound = DictionaryErr("key not found")
var ErrKeyExists = DictionaryErr("key exists")
var ErrKeyNotExists = DictionaryErr("key not exists")

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyNotExists
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}
