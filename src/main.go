package main

import (
	"fmt"

	"github.com/mateusVedoy/go-my-songs.git/src/application/converter"
	"github.com/mateusVedoy/go-my-songs.git/src/application/dto"
)

func main() {
	dto := dto.NewMusicDTO("One", "...And justice for all", "Metallica", "06:24")

	music, businessErr := converter.Convert(*dto)

	if businessErr != nil {
		errors := businessErr.GetErrors()
		fmt.Println("error")
		fmt.Println(errors)
	} else {
		fmt.Println("success")
		fmt.Println(*music)
		fmt.Println(music.GetName())
	}
}
