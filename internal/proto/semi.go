package proto

// === [ Semi-reliable connection ] ============================================
//
// Commands sent from the client to the game server, that are valid only as long
// as no new command has been issued (such as mouse click). Resend and track
// packet delivery only until a new command has replaced an ond one.
//
// Semi-reliable packet loss is nto good: resend until new command is issued.

// TODO: figure out what type of network communication fits this senario.
// Probably light weight wrapper around UDP.
