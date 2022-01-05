package component

import (
	"context"

	"github.com/trallerd/bff-server/graph/model"
)

func Image(ctx context.Context) (*model.ImageContract, error) {
	title := "Fake Cake"
	comment := "Fake Cake for a real party"
	src := "https://scontent.fbfh2-1.fna.fbcdn.net/v/t1.6435-9/36398974_2123678454542623_8852093438628200448_n.jpg?_nc_cat=100&ccb=1-5&_nc_sid=8bfeb9&_nc_eui2=AeHaCG1m4LsNlaNEUSo8HHyD6tkr19zJjYXq2SvX3MmNhdZYOykCZhRxuMqxpT1Y-grS4KGpmi9fVUXnhNxtD7hs&_nc_ohc=qOJU6xog20YAX9oxTXa&_nc_ht=scontent.fbfh2-1.fna&oh=00_AT9t1XI6Bd9LzQj6elbY4ASZf-1vbLCYLkabpM2Xlqhc2Q&oe=61FA20F9"
	return &model.ImageContract{
		Title:   title,
		Comment: &comment,
		Scr:     &src,
	}, nil
}
