'use strict';

const http = require('http');

const server = http.createServer((req, res) => {
    console.log('Ok we have a hit: ', req.url);
    res.setHeader('ContentType', 'application/json');
    res.setHeader('Content-Type', 'application/json');
    res.write(JSON.stringify({ message: `We have a hit`, time: new Date().toDateString(), method: req.method, url: req.url }));
    return res.end();
});

server.listen(3000, () => {
    console.log('Listening on port: 3000');
})
