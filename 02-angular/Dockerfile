FROM node:20 AS build

WORKDIR /build

# intentionally done separately from copying sources
COPY package*.json .
RUN npm ci --verbose && npm cache clean --force

COPY . .
RUN npm run build --prod

FROM nginx:latest

RUN chown -R nginx /etc/nginx

USER nginx:nginx

COPY --from=build /build/dist/02-angular/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

# don't have to specify command, nginx already provides that
