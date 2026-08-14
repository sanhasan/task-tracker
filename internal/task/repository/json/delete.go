package repository

import (
	"fmt"
)

func (r *repo) DeleteEvent(id int) error {
	st, err := r.load()
	if err != nil {
		return err
	}

	ind := st.findIndex(id)
	if ind == -1 {
		return fmt.Errorf("%w: task with id %d", ErrEventNotFound, id)
	}

	st.Tasks = append(st.Tasks[:ind], st.Tasks[ind+1:]...)

	return r.save(st)
}
