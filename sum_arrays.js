// Sum Numbers
function sum (numbers) {
    "use strict";
    let hasil = 0
    for (let i = 0; i < numbers.length;i++) {
      hasil += numbers[i];
    }
    console.log(hasil);
    return hasil;
};