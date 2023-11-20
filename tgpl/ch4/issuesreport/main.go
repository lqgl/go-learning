package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/lqgl/go-learning/tgpl/ch4/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

// $: go run . repo:golang/go is:open json decoder
// Output:
/*
88 issues:
--------------------------------------
Number: 48298
User:   dsnet
Title:  encoding/json: add Decoder.DisallowDuplicateFields
Age:    801 days
--------------------------------------
Number: 11046
User:   kurin
Title:  encoding/json: Decoder internally buffers full input
Age:    3091 days
--------------------------------------
Number: 36225
User:   dsnet
Title:  encoding/json: the Decoder.Decode API lends itself to misuse
Age:    1431 days
--------------------------------------
Number: 56733
User:   rolandshoemaker
Title:  encoding/json: add (*Decoder).SetLimit
Age:    370 days
--------------------------------------
Number: 61627
User:   nabice
Title:  x/tools/gopls: The rename command line may accept identifiers in
Age:    115 days
--------------------------------------
Number: 29035
User:   jaswdr
Title:  proposal: encoding/json: add error var to compare  the returned
Age:    1815 days
--------------------------------------
Number: 42571
User:   dsnet
Title:  encoding/json: clarify Decoder.InputOffset semantics
Age:    1102 days
--------------------------------------
Number: 40128
User:   rogpeppe
Title:  proposal: encoding/json: garbage-free reading of tokens
Age:    1229 days
--------------------------------------
Number: 41144
User:   alvaroaleman
Title:  encoding/json: Unmarshaler breaks DisallowUnknownFields
Age:    1175 days
--------------------------------------
Number: 5901
User:   rsc
Title:  encoding/json: allow per-Encoder/per-Decoder registration of mar
Age:    3777 days
--------------------------------------
Number: 40127
User:   rogpeppe
Title:  encoding/json: add Encoder.EncodeToken method
Age:    1229 days
--------------------------------------
Number: 43716
User:   ggaaooppeenngg
Title:  encoding/json: increment byte counter when using decoder.Token
Age:    1039 days
--------------------------------------
Number: 14750
User:   cyberphone
Title:  encoding/json: parser ignores the case of member names
Age:    2810 days
--------------------------------------
Number: 34543
User:   maxatome
Title:  encoding/json: Unmarshal &amp; json.(*Decoder).Token report differen
Age:    1516 days
--------------------------------------
Number: 59053
User:   joerdav
Title:  proposal: encoding/json: add a generic Decode function
Age:    249 days
--------------------------------------
Number: 32779
User:   rsc
Title:  encoding/json: memoize strings during decode
Age:    1608 days
--------------------------------------
Number: 48950
User:   AlexanderYastrebov
Title:  encoding/json: calculate correct SyntaxError.Offset in the strea
Age:    767 days
--------------------------------------
Number: 31701
User:   lr1980
Title:  encoding/json: second decode after error impossible
Age:    1668 days
--------------------------------------
Number: 6647
User:   btracey
Title:  x/tools/cmd/godoc: display type kind of each named type
Age:    3679 days
--------------------------------------
Number: 40982
User:   Segflow
Title:  encoding/json: use different error type for unknown field if the
Age:    1184 days
--------------------------------------
Number: 16212
User:   josharian
Title:  encoding/json: do all reflect work before decoding
Age:    2699 days
--------------------------------------
Number: 43513
User:   AlexanderYastrebov
Title:  encoding/json: add line number to SyntaxError
Age:    1048 days
--------------------------------------
Number: 56332
User:   gansvv
Title:  encoding/json: clearer error message for boolean like prefix in
Age:    396 days
--------------------------------------
Number: 33714
User:   flimzy
Title:  proposal: encoding/json: Opt-in for true streaming support
Age:    1553 days
--------------------------------------
Number: 7872
User:   extemporalgenome
Title:  encoding/json: Encoder internally buffers full output
Age:    3494 days
--------------------------------------
Number: 34564
User:   mdempsky
Title:  go/internal/gcimporter: single source of truth for decoder logic
Age:    1515 days
--------------------------------------
Number: 33854
User:   Qhesz
Title:  encoding/json: unmarshal option to treat omitted fields as null
Age:    1546 days
--------------------------------------
Number: 22533
User:   ibrt
Title:  proposal: encoding/json: preserve unknown fields
Age:    2209 days
--------------------------------------
Number: 58649
User:   nabokihms
Title:  encoding/json: show nested fields path if DisallowUnknownFields
Age:    270 days
--------------------------------------
Number: 26946
User:   deuill
Title:  encoding/json: clarify what happens when unmarshaling into a non
Age:    1925 days
*/
