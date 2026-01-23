const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

readline.question('', x => {
    readline.question('', y => {
        console.log(`${Number(x) + Number(y)}`);
        readline.close();
    });
});