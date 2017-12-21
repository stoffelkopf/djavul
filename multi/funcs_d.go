//+build djavul

package multi

// static void multi_process_network_packets(void) {
//    void (*f)(void) = (void*)0x440153;
//    f();
// }
import "C"

// useGo specifies whether to use the Go implementation.
const useGo = true

// ProcessNetworkPackets processes network incoming packets.
//
// ref: 0x440153
func ProcessNetworkPackets() {
	if useGo {
		processNetworkPackets()
	}
	C.multi_process_network_packets()
}
