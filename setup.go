package main

func setup(funs ...func() error) (err error) {
	for _, fn := range funs {
		if err = fn(); err != nil {
			return err
		}
	}

	return
}
