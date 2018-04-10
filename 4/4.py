from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer


class HttpProcessor(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('content-type', 'text/html')
        self.end_headers()
        self.wfile.write("Bradbury lab")


serv = HTTPServer(("localhost", 180), HttpProcessor)
serv.serve_forever()
