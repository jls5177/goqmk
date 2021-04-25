package via

type ViaCommandId = uint8

const (
	IdGetProtocolVersion              ViaCommandId = 0x01 // always 0x01
	IdGetKeyboardValue                ViaCommandId = 0x02
	IdSetKeyboardValue                ViaCommandId = 0x03
	IdDynamicKeymapGetKeycode         ViaCommandId = 0x04
	IdDynamicKeymapSetKeycode         ViaCommandId = 0x05
	IdDynamicKeymapReset              ViaCommandId = 0x06
	IdLightingSetValue                ViaCommandId = 0x07
	IdLightingGetValue                ViaCommandId = 0x08
	IdLightingSave                    ViaCommandId = 0x09
	IdEepromReset                     ViaCommandId = 0x0A
	IdBootloaderJump                  ViaCommandId = 0x0B
	IdDynamicKeymapMacroGetCount      ViaCommandId = 0x0C
	IdDynamicKeymapMacroGetBufferSize ViaCommandId = 0x0D
	IdDynamicKeymapMacroGetBuffer     ViaCommandId = 0x0E
	IdDynamicKeymapMacroSetBuffer     ViaCommandId = 0x0F
	IdDynamicKeymapMacroReset         ViaCommandId = 0x10
	IdDynamicKeymapGetLayerCount      ViaCommandId = 0x11
	IdDynamicKeymapGetBuffer          ViaCommandId = 0x12
	IdDynamicKeymapSetBuffer          ViaCommandId = 0x13
	IdDebug                           ViaCommandId = 0xFE
	IdUnhandled                       ViaCommandId = 0xFF
)

type ViaKeyboardValueId = uint8

const (
	IdUptime            ViaKeyboardValueId = 0x01
	IdLayoutOptions     ViaKeyboardValueId = 0x01
	IdSwitchMatrixState ViaKeyboardValueId = 0x01
)

type ViaLightingSubCmdId = uint8

const (
	// QMK BACKLIGHT
	IdQmkBacklightBrightness ViaLightingSubCmdId = 0x09
	IdQmkBacklightEffect     ViaLightingSubCmdId = 0x0A

	// QMK RGBLIGHT
	IdQmkRgblightBrightness  ViaLightingSubCmdId = 0x80
	IdQmkRgblightEffect      ViaLightingSubCmdId = 0x81
	IdQmkRgblightEffectSpeed ViaLightingSubCmdId = 0x82
	IdQmkRgblightColor       ViaLightingSubCmdId = 0x83
)
