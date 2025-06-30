# Stage 1: Build
FROM node:18 AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

# Stage 2: Run (development server with Vite)
FROM node:18

WORKDIR /app

COPY --from=builder /app /app

# Install only runtime dependencies
RUN npm install --only=production

CMD ["npm", "run", "dev"]
