package main

import (
	"strings"
)

var bannedChirps = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

func cleanChirp(chirp string) string {
	split := strings.Split(chirp, " ")
	cleanedChirps := make([]string, len(split))
	for i, word := range split {
		lower := strings.ToLower(word)
		if _, banned := bannedChirps[lower]; banned {
			cleanedChirps[i] = "****"
		} else {
			cleanedChirps[i] = word
		}
	}
	cleanedChirp := strings.Join(cleanedChirps, " ")
	return cleanedChirp
}
