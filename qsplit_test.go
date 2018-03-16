package qsplit

import (
	"testing"
)

func Test01(t *testing.T) {

	fields,err := Split("foo bar baz")
	if err != nil {
		t.Error(err)
	}
	
	if len(fields) != 3 ||
		fields[0] != "foo" ||
		fields[1] != "bar" ||
		fields[2] != "baz" {
		for _,v := range(fields) {
			t.Logf("[%v]", v)
		}
		t.Fail()
	}
	
}

func Test02(t *testing.T) {

	fields,err := Split("foo \"barbecue\" baz")
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 3 ||
		fields[0] != "foo" ||
		fields[1] != "barbecue" ||
		fields[2] != "baz" {
		for _,v := range(fields) {
			t.Logf("[%v]", v)
		}
		t.Fail()
	}
	
}

func Test03(t *testing.T) {

	fields,err := Split("\"barbecue\" baz")
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 2 ||
		fields[0] != "barbecue" ||
		fields[1] != "baz" {
		for _,v := range(fields) {
			t.Logf("[%v]", v)
		}
		t.Fail()
	}
	
}

func Test04(t *testing.T) {

	fields,err := Split("foo \"bar be cue\" baz")
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 3 ||
		fields[0] != "foo" ||
		fields[1] != "bar be cue" ||
		fields[2] != "baz" {
		for _,v := range(fields) {
			t.Logf("[%v]", v)
		}
		t.Fail()
	}
	
}

func Test05(t *testing.T) {

	fields,err := Split("foo \"bar be cue baz")


	if err == nil {
		t.Fail()
	}
	if fields != nil {
		t.Fail()		
	}

	if !IsUnterminatedQuote(err) {
		t.Fail()
	}
	
}

func Test06(t *testing.T) {

	fields,err := Split("foo \"bar \\\" cue\" baz")
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 3 ||
		fields[0] != "foo" ||
		fields[1] != "bar \" cue" ||
		fields[2] != "baz" {
		for _,v := range(fields) {
			t.Logf("[%v]", v)
		}
		t.Fail()
	}
	
}

func Test07(t *testing.T) {

	fields,err := Split("foo \"barbecue\\\\\" baz")
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 3 ||
		fields[0] != "foo" ||
		fields[1] != "barbecue\\" ||
		fields[2] != "baz" {
		for _,v := range(fields) {
			t.Logf("[%v]", v)
		}
		t.Fail()
	}
	
}

func Test08(t *testing.T) {

	fields,err := Split("foo \"barbecue\\m\" baz")

	if err == nil {
		t.Fail()
	}
	if fields != nil {
		t.Fail()		
	}

	if !IsInvalidEscapedCharacter(err) {
		t.Fail()
	}
	
}

func Test09(t *testing.T) {

	fields,err := Split("foo \"barbecue\\")

	if err == nil {
		t.Fail()
	}
	if fields != nil {
		t.Fail()		
	}

	if !IsUnterminatedQuote(err) {
		t.Fail()
	}
	
}
