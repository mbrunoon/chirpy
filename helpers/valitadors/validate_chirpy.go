package validators

import (
	"fmt"
	"strings"

	"github.com/mbrunoon/chirpy/app/models"
)

func ValidateChirpy(chirpy *models.CreateChirpyParams) error {
	const maxChirpLength = 140

	if len(chirpy.Body) > maxChirpLength {
		return fmt.Errorf("body chirp max length is 140")
	}

	chirpy.Body = cleanProfane(chirpy.Body)

	return nil
}

func cleanProfane(text string) string {

	profaneWords := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	replacer := "****"

	separator := " "
	splitedText := strings.Split(text, separator)

	for i, word := range splitedText {
		if _, exists := profaneWords[strings.ToLower(word)]; exists {
			splitedText[i] = replacer
		}
	}

	return strings.Join(splitedText, separator)
}
