const http = require('http');

const port = 3000;

const requestHandler = (req, res) => {
    res.end('Microservicio sesión 6 funcionando');
};

const server = http.createServer(requestHandler);

server.listen(port, () => {
    console.log(`Servidor Node corriendo en puerto ${port}`);
});