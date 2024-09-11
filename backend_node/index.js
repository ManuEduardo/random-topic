import { listen } from 'soap';
import { createServer } from 'http';
import { readFileSync } from 'fs';

// Definimos las funciones que implementará el servicio
const myService = {
  MyService: {
    MyServicePort: {
      MyFunction: function(args) {
        const response = `Received: ${args.input}`;
        return { output: response };
      }
    }
  }
};

// Cargar el archivo WSDL
const wsdl = readFileSync('service.wsdl', 'utf8');

// Crear el servidor HTTP y asociar el servicio SOAP
const server = createServer(function(request, response) {
  response.end('404: Not Found');
});

server.listen(8000, function() {
  console.log('Server listening on port 8000');
});

// Crear el servidor SOAP
listen(server, '/wsdl', myService, wsdl);
