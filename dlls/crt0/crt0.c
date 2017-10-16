#include <stdio.h>
#include <windows.h>

void crt0_cinit(void);

int (*djavul_WinMain)(HINSTANCE, HINSTANCE, char *, int) = (void *) 0x408B4A;
void (**djavul_cpp_init_funcs)(void) = (void *) 0x483000;

__declspec(dllexport) void crt0_start(void) {
	crt0_cinit();
	HINSTANCE hInstance = GetModuleHandleA(NULL);
	// Skip program name in command line arguments.
	char *szCmdLine = GetCommandLineA();
	if (*szCmdLine == '"') {
		while (*++szCmdLine != '"' && *szCmdLine) {
			if (*szCmdLine) {
				++szCmdLine;
			}
		}
		if (*szCmdLine == '"') {
			++szCmdLine;
		}
	} else {
		while (*szCmdLine > ' ') {
			++szCmdLine;
		}
	}
	while (*szCmdLine && *szCmdLine <= ' ') {
		++szCmdLine;
	}
	int status;
	status = djavul_WinMain(hInstance, NULL, szCmdLine, SW_SHOWDEFAULT);
	exit(status);
}

void crt0_cinit(void) {
	for (int i = 1; i < 34; i++) {
		void (*p)(void) = djavul_cpp_init_funcs[i];
		if (p == NULL) {
			break;
		}
		printf("crt0_init: %p\n", p);
		p();
	}
}
