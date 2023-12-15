package translation_test

import (
	"testing"

	"github.com/dmaxim/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "french",
			Translation: "bonjour",
		},
	}

	under_test := translation.NewStaticService()
	for _, test := range tt {
		res := under_test.Translate(test.Word, test.Language)
		if res != test.Translation {
			t.Errorf(`expected %s but received %s`, test.Translation, res)
		}
	}
}
