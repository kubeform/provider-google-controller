/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAllUpdatesRule{}).Type1()):        BillingBudgetSpecAllUpdatesRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmount{}).Type1()):                BillingBudgetSpecAmountCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmountSpecifiedAmount{}).Type1()): BillingBudgetSpecAmountSpecifiedAmountCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecBudgetFilter{}).Type1()):          BillingBudgetSpecBudgetFilterCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAllUpdatesRule{}).Type1()):        BillingBudgetSpecAllUpdatesRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmount{}).Type1()):                BillingBudgetSpecAmountCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmountSpecifiedAmount{}).Type1()): BillingBudgetSpecAmountSpecifiedAmountCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecBudgetFilter{}).Type1()):          BillingBudgetSpecBudgetFilterCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type BillingBudgetSpecAllUpdatesRuleCodec struct {
}

func (BillingBudgetSpecAllUpdatesRuleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BillingBudgetSpecAllUpdatesRule)(ptr) == nil
}

func (BillingBudgetSpecAllUpdatesRuleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BillingBudgetSpecAllUpdatesRule)(ptr)
	var objs []BillingBudgetSpecAllUpdatesRule
	if obj != nil {
		objs = []BillingBudgetSpecAllUpdatesRule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAllUpdatesRule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BillingBudgetSpecAllUpdatesRuleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BillingBudgetSpecAllUpdatesRule)(ptr) = BillingBudgetSpecAllUpdatesRule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BillingBudgetSpecAllUpdatesRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAllUpdatesRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BillingBudgetSpecAllUpdatesRule)(ptr) = objs[0]
			} else {
				*(*BillingBudgetSpecAllUpdatesRule)(ptr) = BillingBudgetSpecAllUpdatesRule{}
			}
		} else {
			*(*BillingBudgetSpecAllUpdatesRule)(ptr) = BillingBudgetSpecAllUpdatesRule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BillingBudgetSpecAllUpdatesRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAllUpdatesRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BillingBudgetSpecAllUpdatesRule)(ptr) = obj
		} else {
			*(*BillingBudgetSpecAllUpdatesRule)(ptr) = BillingBudgetSpecAllUpdatesRule{}
		}
	default:
		iter.ReportError("decode BillingBudgetSpecAllUpdatesRule", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BillingBudgetSpecAmountCodec struct {
}

func (BillingBudgetSpecAmountCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BillingBudgetSpecAmount)(ptr) == nil
}

func (BillingBudgetSpecAmountCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BillingBudgetSpecAmount)(ptr)
	var objs []BillingBudgetSpecAmount
	if obj != nil {
		objs = []BillingBudgetSpecAmount{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmount{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BillingBudgetSpecAmountCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BillingBudgetSpecAmount)(ptr) = BillingBudgetSpecAmount{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BillingBudgetSpecAmount

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmount{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BillingBudgetSpecAmount)(ptr) = objs[0]
			} else {
				*(*BillingBudgetSpecAmount)(ptr) = BillingBudgetSpecAmount{}
			}
		} else {
			*(*BillingBudgetSpecAmount)(ptr) = BillingBudgetSpecAmount{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BillingBudgetSpecAmount

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmount{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BillingBudgetSpecAmount)(ptr) = obj
		} else {
			*(*BillingBudgetSpecAmount)(ptr) = BillingBudgetSpecAmount{}
		}
	default:
		iter.ReportError("decode BillingBudgetSpecAmount", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BillingBudgetSpecAmountSpecifiedAmountCodec struct {
}

func (BillingBudgetSpecAmountSpecifiedAmountCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BillingBudgetSpecAmountSpecifiedAmount)(ptr) == nil
}

func (BillingBudgetSpecAmountSpecifiedAmountCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BillingBudgetSpecAmountSpecifiedAmount)(ptr)
	var objs []BillingBudgetSpecAmountSpecifiedAmount
	if obj != nil {
		objs = []BillingBudgetSpecAmountSpecifiedAmount{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmountSpecifiedAmount{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BillingBudgetSpecAmountSpecifiedAmountCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BillingBudgetSpecAmountSpecifiedAmount)(ptr) = BillingBudgetSpecAmountSpecifiedAmount{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BillingBudgetSpecAmountSpecifiedAmount

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmountSpecifiedAmount{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BillingBudgetSpecAmountSpecifiedAmount)(ptr) = objs[0]
			} else {
				*(*BillingBudgetSpecAmountSpecifiedAmount)(ptr) = BillingBudgetSpecAmountSpecifiedAmount{}
			}
		} else {
			*(*BillingBudgetSpecAmountSpecifiedAmount)(ptr) = BillingBudgetSpecAmountSpecifiedAmount{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BillingBudgetSpecAmountSpecifiedAmount

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecAmountSpecifiedAmount{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BillingBudgetSpecAmountSpecifiedAmount)(ptr) = obj
		} else {
			*(*BillingBudgetSpecAmountSpecifiedAmount)(ptr) = BillingBudgetSpecAmountSpecifiedAmount{}
		}
	default:
		iter.ReportError("decode BillingBudgetSpecAmountSpecifiedAmount", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BillingBudgetSpecBudgetFilterCodec struct {
}

func (BillingBudgetSpecBudgetFilterCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BillingBudgetSpecBudgetFilter)(ptr) == nil
}

func (BillingBudgetSpecBudgetFilterCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BillingBudgetSpecBudgetFilter)(ptr)
	var objs []BillingBudgetSpecBudgetFilter
	if obj != nil {
		objs = []BillingBudgetSpecBudgetFilter{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecBudgetFilter{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BillingBudgetSpecBudgetFilterCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BillingBudgetSpecBudgetFilter)(ptr) = BillingBudgetSpecBudgetFilter{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BillingBudgetSpecBudgetFilter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecBudgetFilter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BillingBudgetSpecBudgetFilter)(ptr) = objs[0]
			} else {
				*(*BillingBudgetSpecBudgetFilter)(ptr) = BillingBudgetSpecBudgetFilter{}
			}
		} else {
			*(*BillingBudgetSpecBudgetFilter)(ptr) = BillingBudgetSpecBudgetFilter{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BillingBudgetSpecBudgetFilter

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BillingBudgetSpecBudgetFilter{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BillingBudgetSpecBudgetFilter)(ptr) = obj
		} else {
			*(*BillingBudgetSpecBudgetFilter)(ptr) = BillingBudgetSpecBudgetFilter{}
		}
	default:
		iter.ReportError("decode BillingBudgetSpecBudgetFilter", "unexpected JSON type")
	}
}
