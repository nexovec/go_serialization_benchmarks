package cbor

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/fxamacker/cbor/v2"
)

// type CBORToArraySerializer struct{}

// type AnnotatedSmallStruct struct {
// 	_ struct{} `cbor:",toarray"`
// 	goserbench.SmallStruct
// }

// func NewCBORToArraySerializer() goserbench.Serializer {
// 	return CBORToArraySerializer{}
// }

// func (serializer CBORToArraySerializer) Marshal(value any) ([]byte, error) {
// 	thing := value.(*goserbench.SmallStruct)
// 	arrayizedThing := AnnotatedSmallStruct{
// 		SmallStruct: *thing,
// 	}
// 	return cbor.Marshal(arrayizedThing)
// }
// func (serializer CBORToArraySerializer) Unmarshal(buffer []byte, output any) error {
// 	thing := output.(*goserbench.SmallStruct)
// 	optimized := AnnotatedSmallStruct{}
// 	err := cbor.Unmarshal(buffer, &optimized)
// 	*thing = optimized.SmallStruct
// 	return err
// }

// type CBORPrecastSerializer struct{}

// func NewCBORPrecastSerializer() goserbench.Serializer {
// 	return CBORPrecastSerializer{}
// }

// func (serializer CBORPrecastSerializer) Marshal(value any) ([]byte, error) {
// 	thing := value.(*goserbench.SmallStruct)
// 	if thing == nil {
// 		return []byte{}, fmt.Errorf("nil pointer")
// 	}
// 	return cbor.Marshal(thing)
// }
// func (serializer CBORPrecastSerializer) Unmarshal(buffer []byte, output any) error {
// 	thing := output.(*goserbench.SmallStruct)
// 	if thing == nil {
// 		return fmt.Errorf("nil pointer")
// 	}
// 	return cbor.Unmarshal(buffer, thing)
// }

type CBORSerializer struct{}

func NewCBORSerializer() goserbench.Serializer {
	return CBORSerializer{}
}

func (serializer CBORSerializer) Marshal(value any) ([]byte, error) {
	return cbor.Marshal(value)
}
func (serializer CBORSerializer) Unmarshal(buffer []byte, output any) error {
	return cbor.Unmarshal(buffer, output)
}
