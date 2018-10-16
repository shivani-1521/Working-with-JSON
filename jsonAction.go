package jsonAction

import(
	"fmt"
	"encoding/json"
	"reflect"
)


func doJsonUnmarshal(name []byte) map[string]interface{}{
	if jsb == nil {
		return nil
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(name, &result); err != nil{
		fmt.Println("Failed to Unmarshal")
	}

	return result
}

func isEqualTo(first, second map[string]interface{}) bool{
	if first == nil && second == nil{
		return true
	}

	else if first == nil || second == nil{
		return false
	}

	return reflect.DeepEqual(first, second)

}

func jsonEqual(objectA, objectB map[string]interface{}) bool{
	return isEqualTo(doJsonUnmarshal(objectA), doJsonUnmarshal(objectB))
}