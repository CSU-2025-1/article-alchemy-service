# Базовый образ для сборки
FROM node:18-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /frontend

# Копируем зависимости и устанавливаем их
COPY package.json package-lock.json ./
RUN npm ci --silent

# Копируем исходный код и собираем проект
COPY . .
RUN npm run build

# Финальный образ (используем простой HTTP-сервер)
FROM node:18-alpine
WORKDIR /frontend

# Устанавливаем serve для раздачи статики
RUN npm install -g serve

# Копируем собранные файлы из builder
COPY --from=builder /frontend/dist ./dist

# Запускаем приложение на порту 3000
EXPOSE 3000
CMD ["serve", "-s", "dist", "-l", "3000"]
