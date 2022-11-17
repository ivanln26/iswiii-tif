FROM node:19-alpine3.15 AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY package.json package-lock.json /app
RUN npm ci

FROM node:19-alpine3.15 AS build
WORKDIR /app
COPY --from=deps /app/node_modules /app/node_modules
COPY . /app
ENV NEXT_TELEMETRY_DISABLED=1
RUN npm run build

FROM node:19-alpine3.15 AS prod
WORKDIR /app
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1 \
    PORT=3000
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs
COPY --from=build /app/public /app/public
COPY --from=build --chown=nextjs:nodejs /app/.next/standalone /app
COPY --from=build --chown=nextjs:nodejs /app/.next/static /app/.next/static
USER nextjs
EXPOSE 3000
CMD ["node", "server.js"]
