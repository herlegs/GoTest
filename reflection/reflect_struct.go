package reflection

import (
	"reflect"
	"fmt"
)

type book struct{
	Name string
	Content string
	pages int
}

func SetStructWithMap(dto interface{}, vars map[string]string){
	fmt.Println(dto)
	value := reflect.ValueOf(dto).Elem()
	dtoTyp := value.Type()
	if(dtoTyp.Kind() != reflect.Struct){
		return
	}
	n := value.NumField()
	for i := 0; i < n; i++ {
		structField := dtoTyp.Field(i)
		fieldVal := value.Field(i)
		fieldName := structField.Name
		strVal,ok := vars[fieldName]
		if(ok && fieldVal.CanSet()){
			val := reflect.ValueOf(strVal)
			fieldVal.Set(val)
			//newField := fieldVal.Interface()
			//newval := reflect.ValueOf(newField)
			fmt.Println(fieldVal.CanSet())
			//newFieldVal.Set(val)

		}
	}
	fmt.Println(dto)
}

func IsSettable(i interface{}){
	val := reflect.ValueOf(i).Elem().Field(0)
	fmt.Println("can addr:",val.CanAddr(),"can set:",val.CanSet())

} 
