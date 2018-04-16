FROM scratch

COPY CustomRedis CustomRedis
COPY config/config.toml config/config.toml

CMD ["./CustomRedis"]