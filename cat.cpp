#include <string>
#include <cstring>
#include <cstdlib>

extern "C" char *cat(char *a, char *b)
{
	std::string x = a;
	std::string y = b;
	std::string z = x + y;
	free(a);
	free(b);
	return strdup(z.c_str());
}

