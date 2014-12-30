package fuelutils_test

import (
	"github.com/fuel-shea/fuel-go-utils/fuelutils"
	"reflect"
	"testing"
)

func TestInterfaceArr2StringArr(t *testing.T) {
	interfaceArr := []interface{}{
		"str1",
		"str2",
		"str3",
	}

	strArr := fuelutils.InterfaceArr2StringArr(interfaceArr)
	if len(interfaceArr) != len(strArr) {
		t.Fatal("The resulting array is", len(strArr), "elements long, but it should be", len(interfaceArr))
	}
	for interfaceIdx, interfaceElem := range interfaceArr {
		strElem := strArr[interfaceIdx]
		if reflect.TypeOf(strElem).Kind() != reflect.String {
			t.Fatal("Element", interfaceIdx, "in converted array should be a string, but it is not")
		}
		if interfaceElem.(string) != strElem {
			t.Error("Element", interfaceIdx, "in converted array should be", interfaceElem, "but is", strElem)
		}
	}
}
