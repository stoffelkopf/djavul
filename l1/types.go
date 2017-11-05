package l1

// A Shadow contains the shadows for a 2x2 block of base tile IDs.
//
// PSX def:
//    typedef struct ShadowStruct {
//       unsigned char strig;
//       unsigned char s1;
//       unsigned char s2;
//       unsigned char s3;
//       unsigned char nv1;
//       unsigned char nv2;
//       unsigned char nv3;
//    } ShadowStruct;
type Shadow struct {
	// Shadow trigger base tile ID.
	BottomBase TileID
	TopBase    TileID
	RightBase  TileID
	LeftBase   TileID
	// Replacement tile IDs.
	Top   TileID
	Right TileID
	Left  TileID
}
