server {
  listen       80;
  server_name  localhost;
  root         /usr/share/nginx/html;

  location / {
    try_files $uri $uri/index.html @app;
  }

  location @app {
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_read_timeout 1d;
    proxy_send_timeout 1d;
    proxy_redirect off;
    proxy_pass http://HOST_IP:APP_PORT;
  }
}
