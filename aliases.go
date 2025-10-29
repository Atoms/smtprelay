package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AliasMap map[string]string

func AliasLoadFile(file string) (AliasMap, error) {
	aliasMap := make(AliasMap)
	count := 0
	log.Info().
		Str("file", file).
		Msg("Loading aliases file")

	f, err := os.Open(file)
	if err != nil {
		log.Fatal().
			Str("file", file).
			Err(err).
			Msg("cannot load aliases file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			aliasMap[parts[0]] = parts[1]
			count++
		}
	}
	log.Info().
		Str("file", file).
		Msg(fmt.Sprintf("Loaded %d aliases from file", count))

	if err := scanner.Err(); err != nil {
		log.Fatal().
			Str("file", file).
			Err(err).
			Msg("cannot load aliases file")
	}
	return aliasMap, nil
}
