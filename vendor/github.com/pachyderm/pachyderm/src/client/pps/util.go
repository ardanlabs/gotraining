package pps

import (
	"sort"

	"github.com/pachyderm/pachyderm/src/client/pfs"
)

// VisitInput visits each input recursively in ascending order (root last)
func VisitInput(input *Input, f func(*Input)) {
	switch {
	case input.Cross != nil:
		for _, input := range input.Cross {
			VisitInput(input, f)
		}
	case input.Union != nil:
		for _, input := range input.Union {
			VisitInput(input, f)
		}
	}
	f(input)
}

// InputName computes the name of an Input.
func InputName(input *Input) string {
	switch {
	case input.Atom != nil:
		return input.Atom.Name
	case input.Cross != nil:
		if len(input.Cross) > 0 {
			return InputName(input.Cross[0])
		}
	case input.Union != nil:
		if len(input.Union) > 0 {
			return InputName(input.Union[0])
		}
	}
	return ""
}

// SortInput sorts an Input.
func SortInput(input *Input) {
	VisitInput(input, func(input *Input) {
		SortInputs := func(inputs []*Input) {
			sort.SliceStable(inputs, func(i, j int) bool { return InputName(inputs[i]) < InputName(inputs[j]) })
		}
		switch {
		case input.Cross != nil:
			SortInputs(input.Cross)
		case input.Union != nil:
			SortInputs(input.Union)
		}
	})
}

// InputCommits returns the commits in an Input.
func InputCommits(input *Input) []*pfs.Commit {
	var result []*pfs.Commit
	VisitInput(input, func(input *Input) {
		if input.Atom != nil {
			result = append(result, &pfs.Commit{
				Repo: &pfs.Repo{input.Atom.Repo},
				ID:   input.Atom.Commit,
			})
		}
		if input.Cron != nil {
			result = append(result, &pfs.Commit{
				Repo: &pfs.Repo{input.Cron.Repo},
				ID:   input.Cron.Commit,
			})
		}
	})
	return result
}
