package main

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	val, ok := d[key]
	if !ok {
		return "n/a"
	} else {
		return val
	}
}
