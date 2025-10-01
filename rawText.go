package main

type rawTextRenderer struct { }

func (rt rawTextRenderer) Render(in string) (string, error) {
	return in, nil
}
