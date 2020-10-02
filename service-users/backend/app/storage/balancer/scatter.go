package balancer

import (
	"fmt"
	"github.com/pkg/errors"
)

func scatter(n int, fn func(i int) error) error {
	errs := make(chan error, n)

	var i int
	for i = 0; i < n; i++ {
		go func(i int) {
			errs <- fn(i)
		}(i)
	}

	var err error
	var innerErr error
	for i = 0; i < cap(errs); i++ {
		if innerErr = <-errs; innerErr != nil {
			err = errors.Errorf(fmt.Sprintf("; An error in scatter with db ", i), innerErr, err)
		}
	}

	return err
}
