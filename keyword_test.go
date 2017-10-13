package keyword

import (
	"reflect"
	"testing"
)

var keyword = New("k∑¥, qwerty", "   exc, LUde,, ", false)

func TestNew(t *testing.T) {
	if keyword == nil {
		t.Error()
	}

	if !reflect.DeepEqual(keyword.keyword, [][]string{{"k∑¥", "qwerty"}}) {
		t.Error()
	}

	if !reflect.DeepEqual(keyword.exclude, [][]string{{"exc", "lude"}}) {
		t.Error()
	}
}

func TestAndNew(t *testing.T) {
	checkerAnd := New("k∑¥, qwerty && abc, 123,", "   exc, LUde,, && , , 478", false)

	if checkerAnd == nil {
		t.Error()
	}

	if !reflect.DeepEqual(checkerAnd.keyword, [][]string{
		{"k∑¥", "qwerty"},
		{"abc", "123"},
	}) {
		t.Error()
	}

	if !reflect.DeepEqual(checkerAnd.exclude, [][]string{
		{"exc", "lude"},
		{"478"},
	}) {
		t.Error()
	}
}

func TestAndNewEmpty(t *testing.T) {
	checkerAnd := New(" , , && ", "&&&&", false)

	if checkerAnd == nil {
		t.Error()
	}

	if checkerAnd.keyword != nil {
		t.Error()
	}

	if checkerAnd.exclude != nil {
		t.Error()
	}
}

func TestCheckKeyword(t *testing.T) {
	if !keyword.CheckKeyword("k∑¥word") {
		t.Error()
	}

	if !keyword.CheckKeyword("       k∑¥word        ") {
		t.Error()
	}

	if !keyword.CheckKeyword("Lorem ipsum dolor sit amet, k∑¥word consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestAndCheckKeyword(t *testing.T) {
	checkerAnd := New("k∑¥ && qwerty", "", false)

	if !checkerAnd.CheckKeyword("k∑¥qwerty") {
		t.Error()
	}

	if !checkerAnd.CheckKeyword("Lorem ipsum dolor sit amet, k∑¥word consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore qwerty et dolore magna aliqua") {
		t.Error()
	}
}

func TestCheckKeywordFail(t *testing.T) {
	if keyword.CheckKeyword("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestAndCheckKeywordFail(t *testing.T) {
	checkerAnd := New("k∑¥ && qwerty", "", false)

	if checkerAnd.CheckKeyword("Lorem ipsum dolor sit amet, k∑¥ consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestCheckExclude(t *testing.T) {
	if !keyword.CheckExclude("exclue") {
		t.Error()
	}

	if !keyword.CheckExclude("       exclue        ") {
		t.Error()
	}

	if !keyword.CheckExclude("Lorem ipsum dolor sit amet, exclue consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestAndCheckExclude(t *testing.T) {
	checkerAnd := New("k∑¥ && qwerty", "exclue && 123", false)

	if !checkerAnd.CheckExclude("exclue 123") {
		t.Error()
	}

	if !checkerAnd.CheckExclude("       exclue123        ") {
		t.Error()
	}

	if !checkerAnd.CheckExclude("Lorem ipsum dolor sit amet, exclue consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore 123 magna aliqua") {
		t.Error()
	}
}

func TestCheckExcludeFail(t *testing.T) {
	if keyword.CheckExclude("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestAndCheckExcludeFail(t *testing.T) {
	checkerAnd := New("k∑¥ && qwerty", "exclue && 123", false)

	if checkerAnd.CheckExclude("Lorem ipsum dolor sit amet, exclue consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestCheck(t *testing.T) {
	if !keyword.Check("Lorem ipsum dolor sit amet, k∑¥word consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}

	if !keyword.Check("Lorem ipsum dolor sit amet, k∑¥ consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestAndCheck(t *testing.T) {
	checkerAnd := New("k∑¥ && word", "exclue && 123", false)

	if !checkerAnd.Check("Lorem ipsum dolor sit amet, k∑¥ consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua word") {
		t.Error()
	}

	if !checkerAnd.Check("Lorem ipsum dolor sit amet, k∑¥word consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestCheckFail(t *testing.T) {
	if keyword.Check("Lorem ipsum dolor sit amet, k∑¥word exclude consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}

	if keyword.Check("Lorem ipsum dolor sit amet, k∑¥ EXCLUDE consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}

func TestAndCheckFail(t *testing.T) {
	checkerAnd := New("k∑¥ && word", "exclude && 123", false)

	if checkerAnd.Check("Lorem ipsum dolor sit amet, k∑¥word exclude consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua 123") {
		t.Error()
	}

	if checkerAnd.Check("Lorem ipsum dolor sit amet, k∑¥ EXCLUDE consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua 123") {
		t.Error()
	}
}

func TestCheckEmptyKeywordExclude(t *testing.T) {
	keyword := New("", "", false)

	if !keyword.Check("Lorem ipsum dolor sit amet") {
		t.Error()
	}

	if !keyword.Check("") {
		t.Error()
	}
}

func TestAndCheckEmptyKeywordExclude(t *testing.T) {
	keyword := New(",,,&&,,", "&&,,,", false)

	if !keyword.Check("Lorem ipsum dolor sit amet") {
		t.Error()
	}

	if !keyword.Check("") {
		t.Error()
	}
}

func TestCheckEmptyKeyword(t *testing.T) {
	keyword := New("", "exclude", false)

	if !keyword.Check("Lorem ipsum dolor sit amet") {
		t.Error()
	}

	if !keyword.Check("") {
		t.Error()
	}

	if keyword.Check("exclude") {
		t.Error()
	}
}

func TestAndCheckEmptyKeyword(t *testing.T) {
	keyword := New("", "exclude && 123", false)

	if !keyword.Check("Lorem ipsum dolor sit amet") {
		t.Error()
	}

	if !keyword.Check("") {
		t.Error()
	}

	if keyword.Check("exclude 123") {
		t.Error()
	}
}

func TestCheckEmptyExclude(t *testing.T) {
	keyword := New("keyword", "", false)

	if keyword.Check("Lorem ipsum dolor sit amet") {
		t.Error()
	}

	if keyword.Check("") {
		t.Error()
	}

	if !keyword.Check("keyword") {
		t.Error()
	}
}

func TestAndCheckEmptyExclude(t *testing.T) {
	keyword := New("keyword && 123", "", false)

	if keyword.Check("Lorem ipsum dolor sit amet") {
		t.Error()
	}

	if keyword.Check("") {
		t.Error()
	}

	if !keyword.Check("123 keyword") {
		t.Error()
	}
}
