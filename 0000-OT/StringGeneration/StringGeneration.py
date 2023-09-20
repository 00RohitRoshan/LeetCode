def generate_lexicographically_smallest_string(S1, S2):
    L1 = len(S1)
    L2 = len(S2)
    result = []

    for i in range(L2):
        if S2[i] == 'T':
            result.extend(S1)
        else:
            smallest_char = min(S1)
            result.append(smallest_char)

    while len(result) < L1 + L2 - 1:
        smallest_char = min(S1)
        result.append(smallest_char)

    if len(result) != L1 + L2 - 1:
        print("-1")
    else:
        print("".join(result))

# Example usage:
S1 = "ABCA"
S2 = "TFTF"
generate_lexicographically_smallest_string(S1, S2)
