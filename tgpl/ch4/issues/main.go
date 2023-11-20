package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lqgl/go-learning/tgpl/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	// $: go run . repo:golang/go is:open json decoder
	// Output:
	/*
		88 issues:
		#48298     dsnet encoding/json: add Decoder.DisallowDuplicateFields
		#11046     kurin encoding/json: Decoder internally buffers full input
		#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
		#56733 rolandsho encoding/json: add (*Decoder).SetLimit
		#61627    nabice x/tools/gopls: The rename command line may accept ident
		#29035    jaswdr proposal: encoding/json: add error var to compare  the
		#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
		#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
		#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
		#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
		#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
		#43716 ggaaooppe encoding/json: increment byte counter when using decode
		#14750 cyberphon encoding/json: parser ignores the case of member names
		#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
		#59053   joerdav proposal: encoding/json: add a generic Decode function
		#32779       rsc encoding/json: memoize strings during decode
		#48950 Alexander encoding/json: calculate correct SyntaxError.Offset in
		#31701    lr1980 encoding/json: second decode after error impossible
		#6647    btracey x/tools/cmd/godoc: display type kind of each named type
		#40982   Segflow encoding/json: use different error type for unknown fie
		#16212 josharian encoding/json: do all reflect work before decoding
		#43513 Alexander encoding/json: add line number to SyntaxError
		#56332    gansvv encoding/json: clearer error message for boolean like p
		#33714    flimzy proposal: encoding/json: Opt-in for true streaming supp
		#7872  extempora encoding/json: Encoder internally buffers full output
		#34564  mdempsky go/internal/gcimporter: single source of truth for deco
		#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
		#22533      ibrt proposal: encoding/json: preserve unknown fields
		#58649 nabokihms encoding/json: show nested fields path if DisallowUnkno
		#26946    deuill encoding/json: clarify what happens when unmarshaling i
	*/
}
