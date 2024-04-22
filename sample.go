package wave

import (
	"github.com/bbars/wave/internal/binary"
)

//goland:noinspection ALL
var (
	SampleFormat_S8 = SampleFormat{"S8", 1, 8, binary.LittleEndian, uint8(0)}
	//SampleFormat_U8                 = SampleFormat{"U8"}
	SampleFormat_S16_LE = SampleFormat{"S16_LE", 1, 16, binary.LittleEndian, int16(0)}
	//SampleFormat_S16_BE             = SampleFormat{"S16_BE"}
	//SampleFormat_U16_LE             = SampleFormat{"U16_LE"}
	//SampleFormat_U16_BE             = SampleFormat{"U16_BE"}
	SampleFormat_S24_LE = SampleFormat{"S24_LE", 1, 24, binary.LittleEndian, binary.Int24(0)}
	//SampleFormat_S24_BE             = SampleFormat{"S24_BE"}
	//SampleFormat_U24_LE             = SampleFormat{"U24_LE"}
	//SampleFormat_U24_BE             = SampleFormat{"U24_BE"}
	SampleFormat_S32_LE = SampleFormat{"S32_LE", 1, 32, binary.LittleEndian, int32(0)}
	//SampleFormat_S32_BE             = SampleFormat{"S32_BE"}
	//SampleFormat_U32_LE             = SampleFormat{"U32_LE"}
	//SampleFormat_U32_BE             = SampleFormat{"U32_BE"}
	SampleFormat_FLOAT_LE = SampleFormat{"FLOAT_LE", 3, 32, binary.LittleEndian, float32(0)}
	//SampleFormat_FLOAT_BE           = SampleFormat{"FLOAT_BE"}
	//SampleFormat_FLOAT64_LE         = SampleFormat{"FLOAT64_LE"}
	//SampleFormat_FLOAT64_BE         = SampleFormat{"FLOAT64_BE"}
	//SampleFormat_IEC958_SUBFRAME_LE = SampleFormat{"IEC958_SUBFRAME_LE"}
	//SampleFormat_IEC958_SUBFRAME_BE = SampleFormat{"IEC958_SUBFRAME_BE"}
	//SampleFormat_MU_LAW             = SampleFormat{"MU_LAW"}
	//SampleFormat_A_LAW              = SampleFormat{"A_LAW"}
	//SampleFormat_IMA_ADPCM          = SampleFormat{"IMA_ADPCM"}
	//SampleFormat_MPEG               = SampleFormat{"MPEG"}
	//SampleFormat_GSM                = SampleFormat{"GSM"}
	//SampleFormat_S20_LE             = SampleFormat{"S20_LE"}
	//SampleFormat_S20_BE             = SampleFormat{"S20_BE"}
	//SampleFormat_U20_LE             = SampleFormat{"U20_LE"}
	//SampleFormat_U20_BE             = SampleFormat{"U20_BE"}
	//SampleFormat_SPECIAL            = SampleFormat{"SPECIAL"}
	SampleFormat_S24_3LE = SampleFormat{"S24_3LE", 1, 24, binary.LittleEndian, binary.Int24(0)}
	//SampleFormat_S24_3BE            = SampleFormat{"S24_3BE"}
	//SampleFormat_U24_3LE            = SampleFormat{"U24_3LE"}
	//SampleFormat_U24_3BE            = SampleFormat{"U24_3BE"}
	//SampleFormat_S20_3LE            = SampleFormat{"S20_3LE"}
	//SampleFormat_S20_3BE            = SampleFormat{"S20_3BE"}
	//SampleFormat_U20_3LE            = SampleFormat{"U20_3LE"}
	//SampleFormat_U20_3BE            = SampleFormat{"U20_3BE"}
	//SampleFormat_S18_3LE            = SampleFormat{"S18_3LE"}
	//SampleFormat_S18_3BE            = SampleFormat{"S18_3BE"}
	//SampleFormat_U18_3LE            = SampleFormat{"U18_3LE"}
	//SampleFormat_U18_3BE            = SampleFormat{"U18_3BE"}
	//SampleFormat_G723_24            = SampleFormat{"G723_24"}
	//SampleFormat_G723_24_1B         = SampleFormat{"G723_24_1B"}
	//SampleFormat_G723_40            = SampleFormat{"G723_40"}
	//SampleFormat_G723_40_1B         = SampleFormat{"G723_40_1B"}
	//SampleFormat_DSD_U8             = SampleFormat{"DSD_U8"}
	//SampleFormat_DSD_U16_LE         = SampleFormat{"DSD_U16_LE"}
	//SampleFormat_DSD_U32_LE         = SampleFormat{"DSD_U32_LE"}
	//SampleFormat_DSD_U16_BE         = SampleFormat{"DSD_U16_BE"}
	//SampleFormat_DSD_U32_BE         = SampleFormat{"DSD_U32_BE"}

	KnownSampleFormats = map[string]SampleFormat{
		SampleFormat_S8.Name: SampleFormat_S8,
		//SampleFormat_U8.Name: SampleFormat_U8,
		SampleFormat_S16_LE.Name: SampleFormat_S16_LE,
		//SampleFormat_S16_BE.Name: SampleFormat_S16_BE,
		//SampleFormat_U16_LE.Name: SampleFormat_U16_LE,
		//SampleFormat_U16_BE.Name: SampleFormat_U16_BE,
		SampleFormat_S24_LE.Name: SampleFormat_S24_LE,
		//SampleFormat_S24_BE.Name: SampleFormat_S24_BE,
		//SampleFormat_U24_LE.Name: SampleFormat_U24_LE,
		//SampleFormat_U24_BE.Name: SampleFormat_U24_BE,
		SampleFormat_S32_LE.Name: SampleFormat_S32_LE,
		//SampleFormat_S32_BE.Name: SampleFormat_S32_BE,
		//SampleFormat_U32_LE.Name: SampleFormat_U32_LE,
		//SampleFormat_U32_BE.Name: SampleFormat_U32_BE,
		SampleFormat_FLOAT_LE.Name: SampleFormat_FLOAT_LE,
		//SampleFormat_FLOAT_BE.Name: SampleFormat_FLOAT_BE,
		//SampleFormat_FLOAT64_LE.Name: SampleFormat_FLOAT64_LE,
		//SampleFormat_FLOAT64_BE.Name: SampleFormat_FLOAT64_BE,
		//SampleFormat_IEC958_SUBFRAME_LE.Name: SampleFormat_IEC958_SUBFRAME_LE,
		//SampleFormat_IEC958_SUBFRAME_BE.Name: SampleFormat_IEC958_SUBFRAME_BE,
		//SampleFormat_MU_LAW.Name: SampleFormat_MU_LAW,
		//SampleFormat_A_LAW.Name: SampleFormat_A_LAW,
		//SampleFormat_IMA_ADPCM.Name: SampleFormat_IMA_ADPCM,
		//SampleFormat_MPEG.Name: SampleFormat_MPEG,
		//SampleFormat_GSM.Name: SampleFormat_GSM,
		//SampleFormat_S20_LE.Name: SampleFormat_S20_LE,
		//SampleFormat_S20_BE.Name: SampleFormat_S20_BE,
		//SampleFormat_U20_LE.Name: SampleFormat_U20_LE,
		//SampleFormat_U20_BE.Name: SampleFormat_U20_BE,
		//SampleFormat_SPECIAL.Name: SampleFormat_SPECIAL,
		SampleFormat_S24_3LE.Name: SampleFormat_S24_3LE,
		//SampleFormat_S24_3BE.Name: SampleFormat_S24_3BE,
		//SampleFormat_U24_3LE.Name: SampleFormat_U24_3LE,
		//SampleFormat_U24_3BE.Name: SampleFormat_U24_3BE,
		//SampleFormat_S20_3LE.Name: SampleFormat_S20_3LE,
		//SampleFormat_S20_3BE.Name: SampleFormat_S20_3BE,
		//SampleFormat_U20_3LE.Name: SampleFormat_U20_3LE,
		//SampleFormat_U20_3BE.Name: SampleFormat_U20_3BE,
		//SampleFormat_S18_3LE.Name: SampleFormat_S18_3LE,
		//SampleFormat_S18_3BE.Name: SampleFormat_S18_3BE,
		//SampleFormat_U18_3LE.Name: SampleFormat_U18_3LE,
		//SampleFormat_U18_3BE.Name: SampleFormat_U18_3BE,
		//SampleFormat_G723_24.Name: SampleFormat_G723_24,
		//SampleFormat_G723_24_1B.Name: SampleFormat_G723_24_1B,
		//SampleFormat_G723_40.Name: SampleFormat_G723_40,
		//SampleFormat_G723_40_1B.Name: SampleFormat_G723_40_1B,
		//SampleFormat_DSD_U8.Name: SampleFormat_DSD_U8,
		//SampleFormat_DSD_U16_LE.Name: SampleFormat_DSD_U16_LE,
		//SampleFormat_DSD_U32_LE.Name: SampleFormat_DSD_U32_LE,
		//SampleFormat_DSD_U16_BE.Name: SampleFormat_DSD_U16_BE,
		//SampleFormat_DSD_U32_BE.Name: SampleFormat_DSD_U32_BE,
	}
)

type SampleFormat struct {
	Name          string
	FormatTag     FormatTag
	BitsPerSample uint16
	ByteOrder     binary.ByteOrder
	valueSample   any
}

func (f SampleFormat) String() string {
	return f.Name
}

type Sample struct {
	N      uint64
	Values []float64
}
