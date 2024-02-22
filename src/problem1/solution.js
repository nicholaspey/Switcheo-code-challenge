// Using an if / while Loop 
var sum_to_n_a = function(n) {
    result = 0;

    // if loop
    for (i = 0; i <= n; i++) {
        result += i;
    }

    // // while loop
    // while (n > 0) {
    //     result += n;
    //     n--;
    // }

    return result;
};

// Using recursion
var sum_to_n_b = function(n) {
    // your code here
    // Base Case
    if (n == 1) {
        return 1;
    }

    // Recursion step
    return n + sum_to_n_b(n - 1);
};

// Using Arithmetic formula
var sum_to_n_c = function(n) {
    // your code here
    return n * (n + 1) / 2;
};


// Testing sum_to_n_a
console.log("sum_to_n_a(5):", sum_to_n_a(5)); // Expected output: 15
// Testing sum_to_n_b
console.log("sum_to_n_b(5):", sum_to_n_b(5)); // Expected output: 15
// Testing sum_to_n_c
console.log("sum_to_n_c(5):", sum_to_n_c(5)); // Expected output: 15