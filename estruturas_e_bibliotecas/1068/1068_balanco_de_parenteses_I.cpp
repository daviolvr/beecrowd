#include <iostream>
#include <string>

int main()
{
	std::string expressao;

	while (std::getline(std::cin, expressao)) {
		int contador = 0;
		bool correto = true;

		for (char c: expressao) {
			if (c == '(') {
				contador++;
			} else if (c == ')') {
				contador--;

				if (contador < 0) {
					correto = false;
					break;
				}
			}
		}

		if (contador != 0) {
			correto = false;
		}

		if (correto) {
			std::cout << "correct" << std::endl;
		} else {
			std::cout << "incorrect" << std::endl;
		}
	}

	return 0;
}
	

	
