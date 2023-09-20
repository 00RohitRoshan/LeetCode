#include <iostream>
#include <string>
using namespace std;

int main() {
    string S1, S2;
    cin >> S1 >> S2;

    int L1 = S1.length();
    int L2 = S2.length();
    char smallest_char = *min_element(S1.begin(), S1.end());
    string result;

    for (int i = 0; i < L2; i++) {
        if (S2[i] == 'T') {
            // If S2[i] is 'T', append S1 to the result.
            result += S1;
        } else {
            // If S2[i] is 'F', append the lexicographically smallest character of S1.
            result += smallest_char;
        }
    }
    for (int a = result.size();a< L1 + L2 - 1; a++){
            result+=smallest_char;
    }
    
    // Check if the length of the result is (L1 + L2 - 1).
    if (result.length() != (L1 + L2 - 1)) {
        cout << "-1" << endl;
    } else {
        cout << result << endl;
    }

    return 0;
}
