package hda

var VerbMap = map[string]string{
	"0xa00": "AC_VERB_GET_STREAM_FORMAT",
	"0xb00": "AC_VERB_GET_AMP_GAIN_MUTE",
	"0xc00": "AC_VERB_GET_PROC_COEF",
	"0xd00": "AC_VERB_GET_COEF_INDEX",
	"0xf00": "AC_VERB_PARAMETERS",
	"0xf01": "AC_VERB_GET_CONNECT_SEL",
	"0xf02": "AC_VERB_GET_CONNECT_LIST",
	"0xf03": "AC_VERB_GET_PROC_STATE",
	"0xf04": "AC_VERB_GET_SDI_SELECT",
	"0xf05": "AC_VERB_GET_POWER_STATE",
	"0xf06": "AC_VERB_GET_CONV",
	"0xf07": "AC_VERB_GET_PIN_WIDGET_CONTROL",
	"0xf08": "AC_VERB_GET_UNSOLICITED_RESPONSE",
	"0xf09": "AC_VERB_GET_PIN_SENSE",
	"0xf0a": "AC_VERB_GET_BEEP_CONTROL",
	"0xf0c": "AC_VERB_GET_EAPD_BTLENABLE",
	"0xf0d": "AC_VERB_GET_DIGI_CONVERT_1",
	"0xf0e": "AC_VERB_GET_DIGI_CONVERT_2",
	"0xf0f": "AC_VERB_GET_VOLUME_KNOB_CONTROL",
	"0xf15": "AC_VERB_GET_GPIO_DATA",
	"0xf16": "AC_VERB_GET_GPIO_MASK",
	"0xf17": "AC_VERB_GET_GPIO_DIRECTION",
	"0xf18": "AC_VERB_GET_GPIO_WAKE_MASK",
	"0xf19": "AC_VERB_GET_GPIO_UNSOLICITED_RSP_MASK",
	"0xf1a": "AC_VERB_GET_GPIO_STICKY_MASK",
	"0xf1c": "AC_VERB_GET_CONFIG_DEFAULT",
	"0xf20": "AC_VERB_GET_SUBSYSTEM_ID",
	"0x200": "AC_VERB_SET_STREAM_FORMAT",
	"0x300": "AC_VERB_SET_AMP_GAIN_MUTE",
	"0x400": "AC_VERB_SET_PROC_COEF",
	"0x500": "AC_VERB_SET_COEF_INDEX",
	"0x701": "AC_VERB_SET_CONNECT_SEL",
	"0x703": "AC_VERB_SET_PROC_STATE",
	"0x704": "AC_VERB_SET_SDI_SELECT",
	"0x705": "AC_VERB_SET_POWER_STATE",
	"0x706": "AC_VERB_SET_CHANNEL_STREAMID",
	"0x707": "AC_VERB_SET_PIN_WIDGET_CONTROL",
	"0x708": "AC_VERB_SET_UNSOLICITED_ENABLE",
	"0x709": "AC_VERB_SET_PIN_SENSE",
	"0x70a": "AC_VERB_SET_BEEP_CONTROL",
	"0x70c": "AC_VERB_SET_EAPD_BTLENABLE",
	"0x70d": "AC_VERB_SET_DIGI_CONVERT_1",
	"0x70e": "AC_VERB_SET_DIGI_CONVERT_2",
	"0x70f": "AC_VERB_SET_VOLUME_KNOB_CONTROL",
	"0x715": "AC_VERB_SET_GPIO_DATA",
	"0x716": "AC_VERB_SET_GPIO_MASK",
	"0x717": "AC_VERB_SET_GPIO_DIRECTION",
	"0x718": "AC_VERB_SET_GPIO_WAKE_MASK",
	"0x719": "AC_VERB_SET_GPIO_UNSOLICITED_RSP_MASK",
	"0x71a": "AC_VERB_SET_GPIO_STICKY_MASK",
	"0x71c": "AC_VERB_SET_CONFIG_DEFAULT_BYTES_0",
	"0x71d": "AC_VERB_SET_CONFIG_DEFAULT_BYTES_1",
	"0x71e": "AC_VERB_SET_CONFIG_DEFAULT_BYTES_2",
	"0x71f": "AC_VERB_SET_CONFIG_DEFAULT_BYTES_3",
	"0x7ff": "AC_VERB_SET_CODEC_RESET",
}

var ParamMap = map[string]string{
	"0x0":  "AC_PAR_VENDOR_ID",
	"0x1":  "AC_PAR_SUBSYSTEM_ID",
	"0x2":  "AC_PAR_REV_ID",
	"0x4":  "AC_PAR_NODE_COUNT",
	"0x5":  "AC_PAR_FUNCTION_TYPE",
	"0x8":  "AC_PAR_AUDIO_FG_CAP",
	"0x9":  "AC_PAR_AUDIO_WIDGET_CAP",
	"0xa":  "AC_PAR_PCM",
	"0xb":  "AC_PAR_STREAM",
	"0xc":  "AC_PAR_PIN_CAP",
	"0xd":  "AC_PAR_AMP_IN_CAP",
	"0xe":  "AC_PAR_CONNLIST_LEN",
	"0xf":  "AC_PAR_POWER_STATE",
	"0x10": "AC_PAR_PROC_CAP",
	"0x11": "AC_PAR_GPIO_CAP",
	"0x12": "AC_PAR_AMP_OUT_CAP",
	"0x13": "AC_PAR_VOL_KNB_CAP",
}
