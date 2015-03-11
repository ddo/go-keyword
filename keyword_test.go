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

	if !reflect.DeepEqual(keyword.keyword, []string{"k∑¥", "qwerty"}) {
		t.Error()
	}

	if !reflect.DeepEqual(keyword.exclude, []string{"exc", "lude"}) {
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

func TestCheckKeywordFail(t *testing.T) {
	if keyword.CheckKeyword("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
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

func TestCheckExcludeFail(t *testing.T) {
	if keyword.CheckExclude("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
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

func TestCheckFail(t *testing.T) {
	if keyword.Check("Lorem ipsum dolor sit amet, k∑¥word exclude consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}

	if keyword.Check("Lorem ipsum dolor sit amet, k∑¥ EXCLUDE consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua") {
		t.Error()
	}
}
