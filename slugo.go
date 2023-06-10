package slugo

import (
	"math/rand"
	"reflect"
)

func FilterStructSlices(sliceRef interface{}, property string, operator string, value interface{}) {
	v := reflect.ValueOf(sliceRef)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		panic("First Argument Should Be A Slice Pointer")
	}

	slice := v.Elem()
	filteredSlice := reflect.MakeSlice(slice.Type(), 0, 0)

	for i := 0; i < slice.Len(); i++ {
		element := slice.Index(i)
		prop := element.FieldByName(property)

		if prop.IsValid() && compareValues(prop.Interface(), value, operator) {
			filteredSlice = reflect.Append(filteredSlice, element)
		}
	}

	slice.Set(filteredSlice)
}

func compareValues(a interface{}, b interface{}, operator string) bool {
	switch operator {
	case "==":
		return a == b
	case "<":
		return lessThan(a, b)
	case ">":
		return greaterThan(a, b)
	case "<=":
		return lessThanOrEqual(a, b)
	case ">=":
		return greaterThanOrEqual(a, b)
	default:
		return false
	}
}

func lessThan(a interface{}, b interface{}) bool {
	switch a := a.(type) {
	case int:
		return a < b.(int)
	case int8:
		return a < b.(int8)
	case int16:
		return a < b.(int16)
	case int32:
		return a < b.(int32)
	case int64:
		return a < b.(int64)
	case uint:
		return a < b.(uint)
	case uint8:
		return a < b.(uint8)
	case uint16:
		return a < b.(uint16)
	case uint32:
		return a < b.(uint32)
	case uint64:
		return a < b.(uint64)
	case float32:
		return a < b.(float32)
	case float64:
		return a < b.(float64)
	case string:
		return false
	case bool:
		return false
	default:
		return false
	}
}

func greaterThan(a interface{}, b interface{}) bool {
	switch a := a.(type) {
	case int:
		return a > b.(int)
	case int8:
		return a > b.(int8)
	case int16:
		return a > b.(int16)
	case int32:
		return a > b.(int32)
	case int64:
		return a > b.(int64)
	case uint:
		return a > b.(uint)
	case uint8:
		return a > b.(uint8)
	case uint16:
		return a > b.(uint16)
	case uint32:
		return a > b.(uint32)
	case uint64:
		return a > b.(uint64)
	case float32:
		return a > b.(float32)
	case float64:
		return a > b.(float64)
	case string:
		return false
	case bool:
		return false
	default:
		return false
	}
}

func lessThanOrEqual(a interface{}, b interface{}) bool {
	return lessThan(a, b) || equal(a, b)
}

func greaterThanOrEqual(a interface{}, b interface{}) bool {
	return greaterThan(a, b) || equal(a, b)
}

func equal(a interface{}, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// PopSlice Function:

func PopSlice(sliceRef interface{}) {
	v := reflect.ValueOf(sliceRef)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return
	}

	slice := v.Elem()
	if slice.Len() == 0 {
		return
	}

	newSlice := slice.Slice(0, slice.Len()-1)
	slice.Set(newSlice)
}

// ReverseSlice Function:

func ReverseSlice(slice interface{}) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	length := sliceValue.Len()
	swap := reflect.Swapper(slice)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// ShuffleSlice Function:

func ShuffleSlice(slice interface{}) {
	reflectValue := reflect.ValueOf(slice)
	if reflectValue.Kind() != reflect.Slice {
		panic("Argument Should Be Slice")
	}

	length := reflectValue.Len()
	randomIndex := rand.Perm(length)

	reflectValueCopy := reflect.MakeSlice(reflectValue.Type(), length, length)
	reflect.Copy(reflectValueCopy, reflectValue)

	for i := 0; i < length; i++ {
		reflectValue.Index(i).Set(reflectValueCopy.Index(randomIndex[i]))
	}
}

// ReduceSlice functions. This function uses individual functions for
// converting various int, uint and float types, but it works only for
// int, uint and float typed slices, not struct slices.:

func ReduceSlice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return nil
	}

	switch v.Index(0).Interface().(type) {
	case uint:
		return ReduceUintSlice(slice, operator)
	case uint8:
		return ReduceUint8Slice(slice, operator)
	case uint16:
		return ReduceUint16Slice(slice, operator)
	case uint32:
		return ReduceUint32Slice(slice, operator)
	case uint64:
		return ReduceUint64Slice(slice, operator)
	case int:
		return ReduceIntSlice(slice, operator)
	case int8:
		return ReduceInt8Slice(slice, operator)
	case int16:
		return ReduceInt16Slice(slice, operator)
	case int32:
		return ReduceInt32Slice(slice, operator)
	case int64:
		return ReduceInt64Slice(slice, operator)
	case float32:
		return ReduceFloat32Slice(slice, operator)
	case float64:
		return ReduceFloat64Slice(slice, operator)
	default:
		return nil
	}
}

func ReduceUintSlice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(uint)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(uint)
		case "-":
			result -= v.Index(i).Interface().(uint)
		case "*":
			result *= v.Index(i).Interface().(uint)
		case "/":
			result /= v.Index(i).Interface().(uint)
		}
	}
	return result
}

func ReduceUint8Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(uint8)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(uint8)
		case "-":
			result -= v.Index(i).Interface().(uint8)
		case "*":
			result *= v.Index(i).Interface().(uint8)
		case "/":
			result /= v.Index(i).Interface().(uint8)
		}
	}
	return result
}

func ReduceUint16Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(uint16)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(uint16)
		case "-":
			result -= v.Index(i).Interface().(uint16)
		case "*":
			result *= v.Index(i).Interface().(uint16)
		case "/":
			result /= v.Index(i).Interface().(uint16)
		}
	}
	return result
}

func ReduceUint32Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(uint32)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(uint32)
		case "-":
			result -= v.Index(i).Interface().(uint32)
		case "*":
			result *= v.Index(i).Interface().(uint32)
		case "/":
			result /= v.Index(i).Interface().(uint32)
		}
	}
	return result
}

func ReduceUint64Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(uint64)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(uint64)
		case "-":
			result -= v.Index(i).Interface().(uint64)
		case "*":
			result *= v.Index(i).Interface().(uint64)
		case "/":
			result /= v.Index(i).Interface().(uint64)
		}
	}
	return result
}

func ReduceIntSlice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(int)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(int)
		case "-":
			result -= v.Index(i).Interface().(int)
		case "*":
			result *= v.Index(i).Interface().(int)
		case "/":
			result /= v.Index(i).Interface().(int)
		}
	}
	return result
}

func ReduceInt8Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(int8)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(int8)
		case "-":
			result -= v.Index(i).Interface().(int8)
		case "*":
			result *= v.Index(i).Interface().(int8)
		case "/":
			result /= v.Index(i).Interface().(int8)
		}
	}
	return result
}

func ReduceInt16Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(int16)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(int16)
		case "-":
			result -= v.Index(i).Interface().(int16)
		case "*":
			result *= v.Index(i).Interface().(int16)
		case "/":
			result /= v.Index(i).Interface().(int16)
		}
	}
	return result
}

func ReduceInt32Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(int32)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(int32)
		case "-":
			result -= v.Index(i).Interface().(int32)
		case "*":
			result *= v.Index(i).Interface().(int32)
		case "/":
			result /= v.Index(i).Interface().(int32)
		}
	}
	return result
}

func ReduceInt64Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).Interface().(int64)
	for i := 1; i < v.Len(); i++ {
		switch operator {
		case "+":
			result += v.Index(i).Interface().(int64)
		case "-":
			result -= v.Index(i).Interface().(int64)
		case "*":
			result *= v.Index(i).Interface().(int64)
		case "/":
			result /= v.Index(i).Interface().(int64)
		}
	}
	return result
}

func ReduceFloat32Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)

	var result float32

	if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Float32 {
		result = v.Index(0).Interface().(float32)
		for i := 1; i < v.Len(); i++ {
			switch operator {
			case "+":
				result += v.Index(i).Interface().(float32)
			case "-":
				result -= v.Index(i).Interface().(float32)
			case "*":
				result *= v.Index(i).Interface().(float32)
			case "/":
				result /= v.Index(i).Interface().(float32)
			}
		}
	}

	return result
}

func ReduceFloat64Slice(slice interface{}, operator string) interface{} {
	v := reflect.ValueOf(slice)

	var result float64

	if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Float64 {
		result = v.Index(0).Interface().(float64)
		for i := 1; i < v.Len(); i++ {
			switch operator {
			case "+":
				result += v.Index(i).Interface().(float64)
			case "-":
				result -= v.Index(i).Interface().(float64)
			case "*":
				result *= v.Index(i).Interface().(float64)
			case "/":
				result /= v.Index(i).Interface().(float64)
			}
		}
	}

	return result
}

// If you want you could do reducing action on a specific field
// of slice of structs, with this functions. If you want to
// maximize your speed i advice to use specific functions, like
// "ReduceInt8StructSlice" rather than "ReduceStructSlice".

func ReduceStructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return nil
	}

	firstElement := v.Index(0)
	fieldValue := getFieldValue(firstElement, field)

	switch fieldValue.Kind() {
	case reflect.Uint:
		return ReduceUintStructSlice(slice, field, operator)
	case reflect.Uint8:
		return ReduceUint8StructSlice(slice, field, operator)
	case reflect.Uint16:
		return ReduceUint16StructSlice(slice, field, operator)
	case reflect.Uint32:
		return ReduceUint32StructSlice(slice, field, operator)
	case reflect.Uint64:
		return ReduceUint64StructSlice(slice, field, operator)
	case reflect.Int:
		return ReduceIntStructSlice(slice, field, operator)
	case reflect.Int8:
		return ReduceInt8StructSlice(slice, field, operator)
	case reflect.Int16:
		return ReduceInt16StructSlice(slice, field, operator)
	case reflect.Int32:
		return ReduceInt32StructSlice(slice, field, operator)
	case reflect.Int64:
		return ReduceInt64StructSlice(slice, field, operator)
	case reflect.Float32:
		return ReduceFloat32StructSlice(slice, field, operator)
	case reflect.Float64:
		return ReduceFloat64StructSlice(slice, field, operator)
	default:
		return nil
	}
}

func getFieldValue(element reflect.Value, field string) reflect.Value {
	fieldValue := element.FieldByName(field)
	if !fieldValue.IsValid() {
		panic("That Field Not Exist")
	}
	return fieldValue
}

func ReduceUintStructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(uint)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(uint)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceUint8StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(uint8)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(uint8)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceUint16StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(uint16)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(uint16)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceUint32StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(uint32)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(uint32)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceUint64StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(uint64)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(uint64)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceIntStructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(int)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(int)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceInt8StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(int8)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(int8)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceInt16StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(int16)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(int16)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceInt32StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(int32)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(int32)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceInt64StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(int64)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(int64)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceFloat32StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(float32)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(float32)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}

func ReduceFloat64StructSlice(slice interface{}, field string, operator string) interface{} {
	v := reflect.ValueOf(slice)
	result := v.Index(0).FieldByName(field).Interface().(float64)
	for i := 1; i < v.Len(); i++ {
		element := v.Index(i)
		fieldValue := getFieldValue(element, field).Interface().(float64)
		switch operator {
		case "+":
			result += fieldValue
		case "-":
			result -= fieldValue
		case "*":
			result *= fieldValue
		case "/":
			result /= fieldValue
		}
	}
	return result
}
