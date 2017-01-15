package reflect

import "testing"

func TestSetStructWithMap(t *testing.T) {
	var dto = &book{}
	bookName := "harry potter"
	bookContent := "harry has long hair"
	var vars = map[string]string{
		"name": bookName,
		"content": bookContent,
	}
	SetStructWithMap(dto, vars)
	if(dto.name != bookName || dto.content != bookContent){
		t.Fail()
	}
}