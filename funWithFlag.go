/* package main

import (
	//"errors"
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetName() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names more than once")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)

	}

	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "bir integer")
	minusO := flag.String("o", "Mihalis", "bir adAd")
	flag.Var(&manyNames, "names", "vergul ile ayrilan list")

	flag.Parse()
	fmt.Println("-k", *minusK)
	fmt.Println("-o", *minusO)

	for i, item := range manyNames.GetName() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining cmd arguments")

	for i, v := range flag.Arg(1) {
		fmt.Println(i, ":", v)
	}

}
 */