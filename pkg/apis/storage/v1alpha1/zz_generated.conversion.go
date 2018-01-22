// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1alpha1 "k8s.io/api/storage/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	storage "k8s.io/kubernetes/pkg/apis/storage"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment,
		Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment,
		Convert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList,
		Convert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList,
		Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource,
		Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource,
		Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec,
		Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec,
		Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus,
		Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus,
		Convert_v1alpha1_VolumeError_To_storage_VolumeError,
		Convert_storage_VolumeError_To_v1alpha1_VolumeError,
	)
}

func autoConvert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(in *v1alpha1.VolumeAttachment, out *storage.VolumeAttachment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(in *v1alpha1.VolumeAttachment, out *storage.VolumeAttachment, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachment_To_storage_VolumeAttachment(in, out, s)
}

func autoConvert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(in *storage.VolumeAttachment, out *v1alpha1.VolumeAttachment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment is an autogenerated conversion function.
func Convert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(in *storage.VolumeAttachment, out *v1alpha1.VolumeAttachment, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachment_To_v1alpha1_VolumeAttachment(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *v1alpha1.VolumeAttachmentList, out *storage.VolumeAttachmentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.VolumeAttachment)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in *v1alpha1.VolumeAttachmentList, out *storage.VolumeAttachmentList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentList_To_storage_VolumeAttachmentList(in, out, s)
}

func autoConvert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(in *storage.VolumeAttachmentList, out *v1alpha1.VolumeAttachmentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.VolumeAttachment)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(in *storage.VolumeAttachmentList, out *v1alpha1.VolumeAttachmentList, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentList_To_v1alpha1_VolumeAttachmentList(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *v1alpha1.VolumeAttachmentSource, out *storage.VolumeAttachmentSource, s conversion.Scope) error {
	out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
	return nil
}

// Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in *v1alpha1.VolumeAttachmentSource, out *storage.VolumeAttachmentSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(in, out, s)
}

func autoConvert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(in *storage.VolumeAttachmentSource, out *v1alpha1.VolumeAttachmentSource, s conversion.Scope) error {
	out.PersistentVolumeName = (*string)(unsafe.Pointer(in.PersistentVolumeName))
	return nil
}

// Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(in *storage.VolumeAttachmentSource, out *v1alpha1.VolumeAttachmentSource, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *v1alpha1.VolumeAttachmentSpec, out *storage.VolumeAttachmentSpec, s conversion.Scope) error {
	out.Attacher = in.Attacher
	if err := Convert_v1alpha1_VolumeAttachmentSource_To_storage_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	return nil
}

// Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in *v1alpha1.VolumeAttachmentSpec, out *storage.VolumeAttachmentSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentSpec_To_storage_VolumeAttachmentSpec(in, out, s)
}

func autoConvert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(in *storage.VolumeAttachmentSpec, out *v1alpha1.VolumeAttachmentSpec, s conversion.Scope) error {
	out.Attacher = in.Attacher
	if err := Convert_storage_VolumeAttachmentSource_To_v1alpha1_VolumeAttachmentSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	return nil
}

// Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(in *storage.VolumeAttachmentSpec, out *v1alpha1.VolumeAttachmentSpec, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentSpec_To_v1alpha1_VolumeAttachmentSpec(in, out, s)
}

func autoConvert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *v1alpha1.VolumeAttachmentStatus, out *storage.VolumeAttachmentStatus, s conversion.Scope) error {
	out.Attached = in.Attached
	out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
	out.AttachError = (*storage.VolumeError)(unsafe.Pointer(in.AttachError))
	out.DetachError = (*storage.VolumeError)(unsafe.Pointer(in.DetachError))
	return nil
}

// Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in *v1alpha1.VolumeAttachmentStatus, out *storage.VolumeAttachmentStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAttachmentStatus_To_storage_VolumeAttachmentStatus(in, out, s)
}

func autoConvert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(in *storage.VolumeAttachmentStatus, out *v1alpha1.VolumeAttachmentStatus, s conversion.Scope) error {
	out.Attached = in.Attached
	out.AttachmentMetadata = *(*map[string]string)(unsafe.Pointer(&in.AttachmentMetadata))
	out.AttachError = (*v1alpha1.VolumeError)(unsafe.Pointer(in.AttachError))
	out.DetachError = (*v1alpha1.VolumeError)(unsafe.Pointer(in.DetachError))
	return nil
}

// Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus is an autogenerated conversion function.
func Convert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(in *storage.VolumeAttachmentStatus, out *v1alpha1.VolumeAttachmentStatus, s conversion.Scope) error {
	return autoConvert_storage_VolumeAttachmentStatus_To_v1alpha1_VolumeAttachmentStatus(in, out, s)
}

func autoConvert_v1alpha1_VolumeError_To_storage_VolumeError(in *v1alpha1.VolumeError, out *storage.VolumeError, s conversion.Scope) error {
	out.Time = in.Time
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_VolumeError_To_storage_VolumeError is an autogenerated conversion function.
func Convert_v1alpha1_VolumeError_To_storage_VolumeError(in *v1alpha1.VolumeError, out *storage.VolumeError, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeError_To_storage_VolumeError(in, out, s)
}

func autoConvert_storage_VolumeError_To_v1alpha1_VolumeError(in *storage.VolumeError, out *v1alpha1.VolumeError, s conversion.Scope) error {
	out.Time = in.Time
	out.Message = in.Message
	return nil
}

// Convert_storage_VolumeError_To_v1alpha1_VolumeError is an autogenerated conversion function.
func Convert_storage_VolumeError_To_v1alpha1_VolumeError(in *storage.VolumeError, out *v1alpha1.VolumeError, s conversion.Scope) error {
	return autoConvert_storage_VolumeError_To_v1alpha1_VolumeError(in, out, s)
}
