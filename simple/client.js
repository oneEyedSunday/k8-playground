'use strict';

const http = require('http');

function getNumber(limit = 10) {
    return Math.floor((Math.random() * limit) + 1);
}

const server = http.createServer((req, res) => {
    console.info('Ok, we have a hit: ', req.url);

    http.get({ path: `/compute/${getNumber(100)}/${getNumber(200)}`, host: 'server-service', port: 3000, timeout: 200 }, (response) => {
        if (response.statusMessage === 'OK') {
            let result = '';

            response.on('data', (chunk) => {
                result += Buffer.from(chunk).toString();
            });

            response.on('end', () => {
                res.write(result);
                return res.end();
            });
        } else {
            res.setHeader('ContentType', 'application/json');
            res.setHeader('Content-Type', 'application/json');
            res.statusCode = 400;
            res.write(JSON.stringify({ message: 'Failed to do something' }));
            return res.end();
        }
    })
});

server.listen(3000, () => {
    console.log('Listening on port: 3000');
})
