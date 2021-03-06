package kinc 


type KeyCode int32 
const (
	KEY_UNKNOWN KeyCode = 0
	KEY_BACK  KeyCode = 1 // Android
	KEY_CANCEL KeyCode = 3
	KEY_HELP KeyCode = 6
	KEY_BACKSPACE KeyCode = 8
	KEY_TAB KeyCode = 9
	KEY_CLEAR KeyCode = 12
	KEY_RETURN KeyCode = 13
	KEY_SHIFT KeyCode = 16
	KEY_CONTROL KeyCode = 17
	KEY_ALT KeyCode = 18
	KEY_PAUSE KeyCode = 19
	KEY_CAPS_LOCK KeyCode = 20
	KEY_KANA KeyCode = 21
	KEY_HANGUL KeyCode = 21
	KEY_EISU KeyCode = 22
	KEY_JUNJA KeyCode = 23
	KEY_FINAL KeyCode = 24
	KEY_HANJA KeyCode = 25
	KEY_KANJI KeyCode = 25
	KEY_ESCAPE KeyCode = 27
	KEY_CONVERT KeyCode = 28
	KEY_NON_CONVERT KeyCode = 29
	KEY_ACCEPT KeyCode = 30
	KEY_MODE_CHANGE KeyCode = 31
	KEY_SPACE KeyCode = 32
	KEY_PAGE_UP KeyCode = 33
	KEY_PAGE_DOWN KeyCode = 34
	KEY_END KeyCode = 35
	KEY_HOME KeyCode = 36
	KEY_LEFT KeyCode = 37
	KEY_UP KeyCode = 38
	KEY_RIGHT KeyCode = 39
	KEY_DOWN KeyCode = 40
	KEY_SELECT KeyCode = 41
	KEY_PRINT KeyCode = 42
	KEY_EXECUTE KeyCode = 43
	KEY_PRINT_SCREEN KeyCode = 44
	KEY_INSERT KeyCode = 45
	KEY_DELETE KeyCode = 46
	KEY_0 KeyCode = 48
	KEY_1 KeyCode = 49
	KEY_2 KeyCode = 50
	KEY_3 KeyCode = 51
	KEY_4 KeyCode = 52
	KEY_5 KeyCode = 53
	KEY_6 KeyCode = 54
	KEY_7 KeyCode = 55
	KEY_8 KeyCode = 56
	KEY_9 KeyCode = 57
	KEY_COLON KeyCode = 58
	KEY_SEMICOLON KeyCode = 59
	KEY_LESS_THAN KeyCode = 60
	KEY_EQUALS KeyCode = 61
	KEY_GREATER_THAN KeyCode = 62
	KEY_QUESTIONMARK KeyCode = 63
	KEY_AT KeyCode = 64
	KEY_A KeyCode = 65
	KEY_B KeyCode = 66
	KEY_C KeyCode = 67
	KEY_D KeyCode = 68
	KEY_E KeyCode = 69
	KEY_F KeyCode = 70
	KEY_G KeyCode = 71
	KEY_H KeyCode = 72
	KEY_I KeyCode = 73
	KEY_J KeyCode = 74
	KEY_K KeyCode = 75
	KEY_L KeyCode = 76
	KEY_M KeyCode = 77
	KEY_N KeyCode = 78
	KEY_O KeyCode = 79
	KEY_P KeyCode = 80
	KEY_Q KeyCode = 81
	KEY_R KeyCode = 82
	KEY_S KeyCode = 83
	KEY_T KeyCode = 84
	KEY_U KeyCode = 85
	KEY_V KeyCode = 86
	KEY_W KeyCode = 87
	KEY_X KeyCode = 88
	KEY_Y KeyCode = 89
	KEY_Z KeyCode = 90
	KEY_WIN KeyCode = 91
	KEY_CONTEXT_MENU KeyCode = 93
	KEY_SLEEP KeyCode = 95
	KEY_NUMPAD_0 KeyCode = 96
	KEY_NUMPAD_1 KeyCode = 97
	KEY_NUMPAD_2 KeyCode = 98
	KEY_NUMPAD_3 KeyCode = 99
	KEY_NUMPAD_4 KeyCode = 100
	KEY_NUMPAD_5 KeyCode = 101
	KEY_NUMPAD_6 KeyCode = 102
	KEY_NUMPAD_7 KeyCode = 103
	KEY_NUMPAD_8 KeyCode = 104
	KEY_NUMPAD_9 KeyCode = 105
	KEY_MULTIPLY KeyCode = 106
	KEY_ADD KeyCode = 107
	KEY_SEPARATOR KeyCode = 108
	KEY_SUBTRACT KeyCode = 109
	KEY_DECIMAL KeyCode = 110
	KEY_DIVIDE KeyCode = 111
	KEY_F1 KeyCode = 112
	KEY_F2 KeyCode = 113
	KEY_F3 KeyCode = 114
	KEY_F4 KeyCode = 115
	KEY_F5 KeyCode = 116
	KEY_F6 KeyCode = 117
	KEY_F7 KeyCode = 118
	KEY_F8 KeyCode = 119
	KEY_F9 KeyCode = 120
	KEY_F10 KeyCode = 121
	KEY_F11 KeyCode = 122
	KEY_F12 KeyCode = 123
	KEY_F13 KeyCode = 124
	KEY_F14 KeyCode = 125
	KEY_F15 KeyCode = 126
	KEY_F16 KeyCode = 127
	KEY_F17 KeyCode = 128
	KEY_F18 KeyCode = 129
	KEY_F19 KeyCode = 130
	KEY_F20 KeyCode = 131
	KEY_F21 KeyCode = 132
	KEY_F22 KeyCode = 133
	KEY_F23 KeyCode = 134
	KEY_F24 KeyCode = 135
	KEY_NUM_LOCK KeyCode = 144
	KEY_SCROLL_LOCK KeyCode = 145
	KEY_WIN_OEM_FJ_JISHO KeyCode = 146
	KEY_WIN_OEM_FJ_MASSHOU KeyCode = 147
	KEY_WIN_OEM_FJ_TOUROKU KeyCode = 148
	KEY_WIN_OEM_FJ_LOYA KeyCode = 149
	KEY_WIN_OEM_FJ_ROYA KeyCode = 150
	KEY_CIRCUMFLEX KeyCode = 160
	KEY_EXCLAMATION KeyCode = 161
	KEY_DOUBLE_QUOTE KeyCode = 162
	KEY_HASH KeyCode = 163
	KEY_DOLLAR KeyCode = 164
	KEY_PERCENT KeyCode = 165
	KEY_AMPERSAND KeyCode = 166
	KEY_UNDERSCORE KeyCode = 167
	KEY_OPEN_PAREN KeyCode = 168
	KEY_CLOSE_PAREN KeyCode = 169
	KEY_ASTERISK KeyCode = 170
	KEY_PLUS KeyCode = 171
	KEY_PIPE KeyCode = 172
	KEY_HYPHEN_MINUS KeyCode = 173
	KEY_OPEN_CURLY_BRACKET KeyCode = 174
	KEY_CLOSE_CURLY_BRACKET KeyCode =  175
	KEY_TILDE KeyCode = 176
	KEY_VOLUME_MUTE KeyCode = 181
	KEY_VOLUME_DOWN KeyCode = 182
	KEY_VOLUME_UP KeyCode = 183
	KEY_COMMA KeyCode = 188
	KEY_PERIOD KeyCode = 190
	KEY_SLASH KeyCode = 191
	KEY_BACK_QUOTE KeyCode = 192
	KEY_OPEN_BRACKET KeyCode = 219
	KEY_BACK_SLASH KeyCode = 220
	KEY_CLOSE_BRACKET KeyCode = 221
	KEY_QUOTE KeyCode = 222
	KEY_META KeyCode = 224
	KEY_ALT_GR KeyCode = 225
	KEY_WIN_ICO_HELP KeyCode = 227
	KEY_WIN_ICO_00 KeyCode = 228
	KEY_WIN_ICO_CLEAR KeyCode = 230
	KEY_WIN_OEM_RESET KeyCode = 233
	KEY_WIN_OEM_JUMP KeyCode = 234
	KEY_WIN_OEM_PA1 KeyCode = 235
	KEY_WIN_OEM_PA2 KeyCode = 236
	KEY_WIN_OEM_PA3 KeyCode = 237
	KEY_WIN_OEM_WSCTRL KeyCode =  238
	KEY_WIN_OEM_CUSEL KeyCode = 239
	KEY_WIN_OEM_ATTN KeyCode = 240
	KEY_WIN_OEM_FINISH KeyCode = 241
	KEY_WIN_OEM_COPY KeyCode = 242
	KEY_WIN_OEM_AUTO KeyCode = 243
	KEY_WIN_OEM_ENLW KeyCode = 244
	KEY_WIN_OEM_BACK_TAB KeyCode = 245
	KEY_ATTN KeyCode = 246
	KEY_CRSEL KeyCode = 247
	KEY_EXSEL KeyCode = 248
	KEY_EREOF KeyCode = 249
	KEY_PLAY KeyCode = 250
	KEY_ZOOM KeyCode = 251
	KEY_PA1 KeyCode = 253
	KEY_WIN_OEM_CLEAR KeyCode = 254
)