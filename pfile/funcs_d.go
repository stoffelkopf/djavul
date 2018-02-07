//+build djavul

package pfile

// #include <stdint.h>
// typedef int32_t bool32_t;
//
//    // HeroInfo specifies information about the hero to be presented in the user
//    // interface.
//    //
//    // PSX def:
//    //    typedef struct _uiheroinfo {
//    //       struct _uiheroinfo* next;
//    //       char name[16];
//    //       unsigned short level;
//    //       unsigned char heroclass;
//    //       unsigned char herorank;
//    //       unsigned short strength;
//    //       unsigned short magic;
//    //       unsigned short dexterity;
//    //       unsigned short vitality;
//    //       unsigned long gold;
//    //       unsigned char hassaved;
//    //       unsigned char spawned;
//    //    } TUIHEROINFO, _uiheroinfo;
//    typedef struct HeroInfo { // size = 0x2C
//       // offset: 0000 (4 bytes)
//       struct HeroInfo *next;
//       // offset: 0004 (16 bytes)
//       char name[16];
//       // offset: 0014 (2 bytes)
//       int16_t clvl;
//       // offset: 0016 (1 bytes)
//       uint8_t player_class; // enum player_class
//       // offset: 0017 (1 bytes)
//       int8_t difficulty; // TODO: use difficulty enum
//       // offset: 0018 (2 bytes)
//       int16_t str_cur;
//       // offset: 001A (2 bytes)
//       int16_t mag_cur;
//       // offset: 001C (2 bytes)
//       int16_t dex_cur;
//       // offset: 001E (2 bytes)
//       int16_t vit_cur;
//       // offset: 0020 (4 bytes)
//       int32_t gold_total;
//       // offset: 0024 (4 bytes)
//       bool32_t has_save;
//       // offset: 0028 (4 bytes)
//       bool32_t spawned;
//    } HeroInfo;
//
// static bool32_t __stdcall pfile_ui_create_save(HeroInfo *hero_info) {
//    bool32_t (__stdcall *f)(HeroInfo*) = (void *)0x44A220;
//    return f(hero_info);
// }
import "C"
import "unsafe"

const useGo = true

// UICreateSave creates a save file based on the given hero information.
//
// ref: 0x44A220
func UICreateSave(heroInfo *HeroInfo) bool {
	if useGo {
		_ = uiCreateSave(heroInfo)
		// TODO: return result directly; now fallthrough to C implementation.
	}
	return C.pfile_ui_create_save((*C.HeroInfo)(unsafe.Pointer(heroInfo))) == 1
}
