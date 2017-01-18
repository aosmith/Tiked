#include <Windows.h>

int executeByteCode(byte Rawcode[])
{
	DWORD old_protect;
	LPVOID executable_area = VirtualAlloc(NULL, 11414, MEM_RESERVE, PAGE_READWRITE);

	memcpy(executable_area, Rawcode, 11414);
	VirtualProtect(executable_area, 11414, PAGE_EXECUTE, &old_protect);

	int(*f)() = (int(*)()) executable_area;
	f();

	// Note: RAII this in C++. Restore old flags, free memory.
	VirtualProtect(executable_area, 11414, old_protect, &old_protect);
	VirtualFree(executable_area, 11414, MEM_RELEASE);
	return 0;
}
