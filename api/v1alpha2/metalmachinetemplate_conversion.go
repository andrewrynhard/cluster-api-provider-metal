package v1alpha2

import (
	infrav1alpha3 "github.com/talos-systems/cluster-api-provider-metal/api/v1alpha3"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this MetalMachineTemplate to the Hub version (v1alpha3).
func (src *MetalMachineTemplate) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha3.MetalMachineTemplate)
	if err := Convert_v1alpha2_MetalMachineTemplate_To_v1alpha3_MetalMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data from annotations
	restored := &infrav1alpha3.MetalMachineTemplate{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *MetalMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha3.MetalMachineTemplate)
	if err := Convert_v1alpha3_MetalMachineTemplate_To_v1alpha2_MetalMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts this MetalMachineTemplateList to the Hub version (v1alpha3).
func (src *MetalMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha3.MetalMachineTemplateList)
	return Convert_v1alpha2_MetalMachineTemplateList_To_v1alpha3_MetalMachineTemplateList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *MetalMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha3.MetalMachineTemplateList)
	return Convert_v1alpha3_MetalMachineTemplateList_To_v1alpha2_MetalMachineTemplateList(src, dst, nil)
}
