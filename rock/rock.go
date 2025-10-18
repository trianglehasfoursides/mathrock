package rock

import "github.com/dgraph-io/ristretto/v2"

var Rock *ristretto.Cache[string, string]

func Setup() (err error) {
	Rock, err = ristretto.
		NewCache(
			&ristretto.
				Config[string, string]{},
		)

	return
}
