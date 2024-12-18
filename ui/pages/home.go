package pages

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type HomeProps struct {
	Excerpts []HomeExcerpt
	Loggedin bool
}

type HomeExcerpt struct {
	Name       string
	Link       string
	DeleteLink string
}

func HomePage(props HomeProps) g.Node {
	// TODO(Amr Ojjeh): Change to PageBase
	return QuizBase(
		QuizBaseProps{
			Title: "NahwApp",
			Body: []g.Node{ID("home-page"),
				Div(ID("home-page--grid"),
					Div(ID("home-page__create"),
						A(Href("https://tags.nahw.app/"),
							Button(Class("button button--primary"),
								g.Text("Create excerpt"))),
					),
					g.Group(g.Map(props.Excerpts, homePageExcerpt)),
				),
			},
			Loggedin: props.Loggedin,
		},
	)
}

func homePageExcerpt(h HomeExcerpt) g.Node {
	return A(Class("button excerpt"), Href(h.Link),
		g.Text(h.Name),
	)
}
