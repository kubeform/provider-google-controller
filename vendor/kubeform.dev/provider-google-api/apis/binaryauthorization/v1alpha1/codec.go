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
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNote{}).Type1()):                        AttestorSpecAttestationAuthorityNoteCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}).Type1()): AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamBindingSpecCondition{}).Type1()):                             AttestorIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamMemberSpecCondition{}).Type1()):                              AttestorIamMemberSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDefaultAdmissionRule{}).Type1()):                              PolicySpecDefaultAdmissionRuleCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNote{}).Type1()):                        AttestorSpecAttestationAuthorityNoteCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}).Type1()): AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamBindingSpecCondition{}).Type1()):                             AttestorIamBindingSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamMemberSpecCondition{}).Type1()):                              AttestorIamMemberSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDefaultAdmissionRule{}).Type1()):                              PolicySpecDefaultAdmissionRuleCodec{},
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
type AttestorSpecAttestationAuthorityNoteCodec struct {
}

func (AttestorSpecAttestationAuthorityNoteCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AttestorSpecAttestationAuthorityNote)(ptr) == nil
}

func (AttestorSpecAttestationAuthorityNoteCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AttestorSpecAttestationAuthorityNote)(ptr)
	var objs []AttestorSpecAttestationAuthorityNote
	if obj != nil {
		objs = []AttestorSpecAttestationAuthorityNote{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNote{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AttestorSpecAttestationAuthorityNoteCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AttestorSpecAttestationAuthorityNote)(ptr) = AttestorSpecAttestationAuthorityNote{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AttestorSpecAttestationAuthorityNote

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNote{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AttestorSpecAttestationAuthorityNote)(ptr) = objs[0]
			} else {
				*(*AttestorSpecAttestationAuthorityNote)(ptr) = AttestorSpecAttestationAuthorityNote{}
			}
		} else {
			*(*AttestorSpecAttestationAuthorityNote)(ptr) = AttestorSpecAttestationAuthorityNote{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AttestorSpecAttestationAuthorityNote

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNote{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AttestorSpecAttestationAuthorityNote)(ptr) = obj
		} else {
			*(*AttestorSpecAttestationAuthorityNote)(ptr) = AttestorSpecAttestationAuthorityNote{}
		}
	default:
		iter.ReportError("decode AttestorSpecAttestationAuthorityNote", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKeyCodec struct {
}

func (AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKeyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) == nil
}

func (AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKeyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr)
	var objs []AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey
	if obj != nil {
		objs = []AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKeyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) = AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) = objs[0]
			} else {
				*(*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) = AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}
			}
		} else {
			*(*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) = AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) = obj
		} else {
			*(*AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey)(ptr) = AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey{}
		}
	default:
		iter.ReportError("decode AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AttestorIamBindingSpecConditionCodec struct {
}

func (AttestorIamBindingSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AttestorIamBindingSpecCondition)(ptr) == nil
}

func (AttestorIamBindingSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AttestorIamBindingSpecCondition)(ptr)
	var objs []AttestorIamBindingSpecCondition
	if obj != nil {
		objs = []AttestorIamBindingSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamBindingSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AttestorIamBindingSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AttestorIamBindingSpecCondition)(ptr) = AttestorIamBindingSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AttestorIamBindingSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamBindingSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AttestorIamBindingSpecCondition)(ptr) = objs[0]
			} else {
				*(*AttestorIamBindingSpecCondition)(ptr) = AttestorIamBindingSpecCondition{}
			}
		} else {
			*(*AttestorIamBindingSpecCondition)(ptr) = AttestorIamBindingSpecCondition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AttestorIamBindingSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamBindingSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AttestorIamBindingSpecCondition)(ptr) = obj
		} else {
			*(*AttestorIamBindingSpecCondition)(ptr) = AttestorIamBindingSpecCondition{}
		}
	default:
		iter.ReportError("decode AttestorIamBindingSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AttestorIamMemberSpecConditionCodec struct {
}

func (AttestorIamMemberSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AttestorIamMemberSpecCondition)(ptr) == nil
}

func (AttestorIamMemberSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AttestorIamMemberSpecCondition)(ptr)
	var objs []AttestorIamMemberSpecCondition
	if obj != nil {
		objs = []AttestorIamMemberSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamMemberSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AttestorIamMemberSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AttestorIamMemberSpecCondition)(ptr) = AttestorIamMemberSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AttestorIamMemberSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamMemberSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AttestorIamMemberSpecCondition)(ptr) = objs[0]
			} else {
				*(*AttestorIamMemberSpecCondition)(ptr) = AttestorIamMemberSpecCondition{}
			}
		} else {
			*(*AttestorIamMemberSpecCondition)(ptr) = AttestorIamMemberSpecCondition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AttestorIamMemberSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AttestorIamMemberSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AttestorIamMemberSpecCondition)(ptr) = obj
		} else {
			*(*AttestorIamMemberSpecCondition)(ptr) = AttestorIamMemberSpecCondition{}
		}
	default:
		iter.ReportError("decode AttestorIamMemberSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PolicySpecDefaultAdmissionRuleCodec struct {
}

func (PolicySpecDefaultAdmissionRuleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PolicySpecDefaultAdmissionRule)(ptr) == nil
}

func (PolicySpecDefaultAdmissionRuleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PolicySpecDefaultAdmissionRule)(ptr)
	var objs []PolicySpecDefaultAdmissionRule
	if obj != nil {
		objs = []PolicySpecDefaultAdmissionRule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDefaultAdmissionRule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PolicySpecDefaultAdmissionRuleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PolicySpecDefaultAdmissionRule)(ptr) = PolicySpecDefaultAdmissionRule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PolicySpecDefaultAdmissionRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDefaultAdmissionRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PolicySpecDefaultAdmissionRule)(ptr) = objs[0]
			} else {
				*(*PolicySpecDefaultAdmissionRule)(ptr) = PolicySpecDefaultAdmissionRule{}
			}
		} else {
			*(*PolicySpecDefaultAdmissionRule)(ptr) = PolicySpecDefaultAdmissionRule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PolicySpecDefaultAdmissionRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PolicySpecDefaultAdmissionRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PolicySpecDefaultAdmissionRule)(ptr) = obj
		} else {
			*(*PolicySpecDefaultAdmissionRule)(ptr) = PolicySpecDefaultAdmissionRule{}
		}
	default:
		iter.ReportError("decode PolicySpecDefaultAdmissionRule", "unexpected JSON type")
	}
}
