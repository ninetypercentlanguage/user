package getflags

import (
	"flag"
	"fmt"
)

type FlagValues struct {
	DB    string
	Words string
}

func GetFlags() *FlagValues {
	dbConnectionStrPtr := flag.String("db", "", "the connection string for the user knowledge database")
	wordsPtr := flag.String("words", "", "the address of the words server")
	flag.Parse()

	dbConnectionStr := *dbConnectionStrPtr
	wordsHost := *wordsPtr

	if dbConnectionStr == "" || wordsHost == "" {
		fmt.Println("Must provide the following flags:")
		flag.PrintDefaults()
		panic("missing flags")
	}
	return &FlagValues{
		DB:    dbConnectionStr,
		Words: wordsHost,
	}
}
