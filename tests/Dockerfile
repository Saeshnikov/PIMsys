# Используем более полный базовый образ
FROM python:3.9

# Первый шаг: только обновление пакетов
RUN apt-get update -yq --fix-missing && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    gnupg2

# Второй шаг: добавление репозитория Chromium
RUN echo "deb http://deb.debian.org/debian stable main" > /etc/apt/sources.list.d/debian-stable.list && \
    curl -sSL https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list

# Третий шаг: установка зависимостей
RUN apt-get update -yq && \
    apt-get install -y --no-install-recommends \
    unzip \
    chromium \
    google-chrome-stable \
    fonts-liberation \
    libxss1 \
    libxtst6 \
    xdg-utils \
    && rm -rf /var/lib/apt/lists/*

# Переменные окружения
ENV CHROME_BIN=/usr/bin/google-chrome-stable \
    CHROMIUM_BIN=/usr/bin/chromium

# Копируем зависимости и устанавливаем их
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Копируем тестовый скрипт
COPY autotests.py .

# Команда запуска
CMD ["python", "autotests.py"]