#include <iostream>

int main()
{
	int T = 0;
	int acertos = 0;
	std::cin >> T;

	for (int i = 0; i <= 4; i++) {
		int palpite = 0;
		std::cin >> palpite;
		if (palpite == T) {
			acertos++;
		}
	}

	std::cout << acertos << std::endl;
	
	return 0;
}
