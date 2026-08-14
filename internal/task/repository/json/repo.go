package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"task-tracker/internal/task/entity"
)

var ErrEventNotFound = errors.New("event not found")

type repo struct {
	path string
}

func New(path string) (*repo, error) {
	ext := strings.ToLower(filepath.Ext(path))
	if ext == ".json" {
		return &repo{path: path}, nil
	}
	return nil, fmt.Errorf("invalid repository path, check configuration")
}

type storage struct {
	NextID int           `json:"next_id"`
	Tasks  []entity.Task `json:"tasks"`
}

func (r *repo) load() (storage, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		if os.IsNotExist(err) {
			return storage{NextID: 1}, nil
		}
		return storage{}, fmt.Errorf("read storage: %w", err)
	}

	if len(data) == 0 {
		return storage{NextID: 1}, nil
	}

	var st storage
	if err := json.Unmarshal(data, &st); err != nil {
		return storage{}, fmt.Errorf("unmarshal storage: %w", err)
	}

	return st, nil
}

func (r *repo) save(st storage) error {
	if err := os.MkdirAll(filepath.Dir(r.path), 0o755); err != nil {
		return fmt.Errorf("create storage dir: %w", err)
	}

	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal storage: %w", err)
	}

	tmp := r.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return fmt.Errorf("write storage: %w", err)
	}

	return os.Rename(tmp, r.path)
}

func (s *storage) findIndex(id int) int {
	for i, v := range s.Tasks {
		if v.Id() == id {
			return i
		}
	}
	return -1
}

func (r *repo) swapTask(pos int, st storage) {
	for i := pos; i < len(st.Tasks)-1; i++ {
		st.Tasks[pos], st.Tasks[pos+1] = st.Tasks[pos+1], st.Tasks[pos]
	}
}
