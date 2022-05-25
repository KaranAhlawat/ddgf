package service

import repo "github.com/KaranAhlawat/ddgf/internal/repo/postgresql"

type Advice struct {
	r *repo.Advice
}

func NewAdvice(repository *repo.Advice) *Advice {
	return &Advice{
		repository,
	}
}

// Add a new advice with no tag

// Delete an advice

// Select an advice with all it's tags

// Select all advices with all their tags
