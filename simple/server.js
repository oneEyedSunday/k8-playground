'use strict';

const http = require('http');

const testRegExp = new RegExp(/^\/compute\/\d+\/\d+\/?$/, 'i');

function compute(a, b) {
    if (!(a && b)) return 0;
    if (isNaN(Number(a) || isNaN(Number(b)))) return 0;
    return Number(a) * Number(b);
}

function handleRoute(route = '/') {
    switch (true) {
        case testRegExp.test(route):
            const [a, b] = route.split('/').slice(2);
            return compute(a, b);
        default:
            console.log('Failed to match route: ', route);
            return 'Ok';
    }
}

const server = http.createServer((req, res) => {
    console.log('Ok we have a hit: ', req.url);
    res.setHeader('ContentType', 'application/json');
    res.setHeader('Content-Type', 'application/json');
    const message = handleRoute(req.url);
    res.write(JSON.stringify({ message }));
    return res.end();
});

server.listen(3000, () => {
    console.log('Listening on port: 3000');
})
