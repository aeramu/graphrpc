package main

type Resolver struct {}

type getFilmReq struct {
	ID string
}

func (r *Resolver) GetFilm(req getFilmReq) (*filmResolver, error) {
	return &filmResolver{}, nil
}