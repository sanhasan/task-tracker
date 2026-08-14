package repository

func (r *repo) UpdateTaskDescription(id int, desc string) error {
	st, err := r.load()
	if err != nil {
		return err
	}

	pos := st.findIndex(id)
	if pos == -1 {
		return ErrEventNotFound
	}

	st.Tasks[pos].UpdateDescription(desc)
	r.swapTask(pos, st)
	return r.save(st)
}
