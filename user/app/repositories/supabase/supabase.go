package supabase

import (
	supa "github.com/nedpals/supabase-go"
)

type supabase struct {
	supaClient *supa.Client
}

func NewSupabase(
	supaClient *supa.Client,
) *supabase {
	return &supabase{
		supaClient: supaClient,
	}
}
